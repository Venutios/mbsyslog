package mbsyslog_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/venutios/mbsyslog"
)

func TestServer(t *testing.T) {
	messages := make(chan mbsyslog.Message, 5)
	s := mbsyslog.NewServer(messages)
	startTime := time.Now()
	secondsToWait := float64(10)

	//run the server
	go func() {
		if err := s.Listen(); err != nil {
			t.Errorf("Server failed to start listening: %s", err.Error())
		}
	}()

	//wait for the server to start
	for !s.Running() {
		if time.Now().Sub(startTime).Seconds() >= secondsToWait {
			t.Error("Server failed to start running")
		}
		time.Sleep(1 * time.Second)
	}

	testTime := time.Now()
	for time.Now().Sub(testTime).Seconds() < secondsToWait {
		select {
		case m := <-messages:
			fmt.Println(m)
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}

	//stop the server and wait for it to stop
	s.Stop()
	stopTime := time.Now()
	for s.Running() {
		if time.Now().Sub(stopTime).Seconds() >= secondsToWait {
			t.Error("Server failed to stop running")
		}
		time.Sleep(1 * time.Second)
	}
}
