package mbsyslog_test

import (
	"strings"
	"testing"
	"time"

	"github.com/venutios/mbsyslog"
)

func TestServer(t *testing.T) {
	clientMsgs := []struct {
		data     []byte
		received bool
	}{
		{[]byte("<34>1 2003-10-11T22:14:15.003Z mymachine.example.com su - ID47 - BOM'su root' failed for lonvick on /dev/pts/8"), false},
		{[]byte("<165>1 2003-08-24T05:14:15.000003-07:00 192.0.2.1 myproc 8710 - - %% It's time to make the do-nuts."), false},
	}
	client := mbsyslog.NewClient(false)

	messages := make(chan mbsyslog.Message, 5)
	s := mbsyslog.NewServer(messages)
	startTime := time.Now()
	maxWait := float64(10)

	//run the server
	go func() {
		if err := s.Listen(); err != nil {
			t.Errorf("Server failed to start listening: %s", err.Error())
		}
	}()

	//wait for the server to start
	for !s.Running() {
		if time.Now().Sub(startTime).Seconds() >= maxWait {
			t.Error("Server failed to start running")
		}
		time.Sleep(1 * time.Second)
	}

	//send the messages from the client
	for _, cm := range clientMsgs {
		client.SendData("127.0.0.1", cm.data)
	}

	//wait for the messages to be recieved from the server, verify they match
	testTime := time.Now()
	for {
		if time.Now().Sub(testTime).Seconds() >= maxWait {
			break
		}

		done := true
		for _, cm := range clientMsgs {
			if !cm.received {
				done = false
				break
			}
		}

		if done {
			break
		}

		select {
		case m := <-messages:
			found := false
			s := m.String()[strings.Index(m.String(), " ")+1:]
			for x, _ := range clientMsgs {
				if string(clientMsgs[x].data) == s {
					if clientMsgs[x].received {
						t.Errorf("Server.TestServer(), duplicate message received: %s", s)
					} else {
						clientMsgs[x].received = true
					}
					found = true
				}
			}
			if !found {
				t.Errorf("Server.TestServer(), message received with no send match: %s", s)
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	for _, cm := range clientMsgs {
		if !cm.received {
			t.Errorf("Server.TestServer(), message never received: %s", string(cm.data))
		}
	}

	//Cleanup the client async operations
	client.Wait()
	if err := client.AsyncError(); err != nil {
		t.Errorf("Server.TestServer(), client error: %s", err)
	}

	//stop the server and wait for it to stop
	s.Stop()
	stopTime := time.Now()
	for s.Running() {
		if time.Now().Sub(stopTime).Seconds() >= maxWait {
			t.Error("Server failed to stop running")
		}
		time.Sleep(1 * time.Second)
	}
}
