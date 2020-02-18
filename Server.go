package mbsyslog

import (
	"net"
	"sync"
	"sync/atomic"
	"time"
)

//Server is a syslog server that forms the basis for a relay or collector in
//the syslog system. Data is received and parsed into syslog messages.
type Server struct {
	maxMessageSize int
	port           int
	stopChan       chan struct{}
	running        int32
	messagesOut    chan<- Message
}

//NewServer prepares the server to listen for messages. The server will listen
//on port 514 and have an 8KB maximum message size. The message channel will
//receive all messages received. The channel should not be closed until the
//server is not running by calling Running().
func NewServer(messageChan chan<- Message) *Server {
	result := new(Server)
	result.port = 514
	result.maxMessageSize = 8192
	result.messagesOut = messageChan
	result.stopChan = make(chan struct{}, 1)
	result.running = 0
	return result
}

//Listen starts the server accepting syslog messages. The server will not stop
//until the Stop() method is called, and all outstanding parsers have chance to
//finish processing and write their message to the output channel.
func (s *Server) Listen() error {
	var wg sync.WaitGroup

	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: s.port, Zone: ""})
	if err != nil {
		return err
	}

	s.setRunning(true)
	//defer evalaute as a stack, when stopping close the socket, wait for
	//all gorountines to finish parsing, and then signal the server is stopped
	defer func() { s.setRunning(false) }()
	defer wg.Wait()
	defer conn.Close()

	buffer := make([]byte, s.maxMessageSize)
	for {
		select {
		case <-s.stopChan: //supposed to stop, everything is deferred above
			return nil
		default:
			conn.SetDeadline(time.Now().Add(1 * time.Second))
			count, addr, err := conn.ReadFromUDP(buffer)
			if err == nil {
				//copy of the buffer, so it doesn't change while the goroutine runs
				data := make([]byte, count)
				copy(data, buffer)
				wg.Add(1)
				go func(data []byte) {
					defer wg.Done()
					s.messagesOut <- *NewMessage(addr, data)
				}(data)
			}
		}
	}
}

//Port that the server is configured to listen on
func (s Server) Port() int {
	return s.port
}

//MaximumMessageSize is the largest syslog message that can be accepted.
func (s Server) MaximumMessageSize() int {
	return s.maxMessageSize
}

//Stop signals the server to shutdown, but doesn't stop immediately
func (s *Server) Stop() {
	s.stopChan <- *new(struct{})
}

//Running returns if the server is currently accepting syslog messages or not
func (s *Server) Running() bool {
	if atomic.LoadInt32(&s.running) == 1 {
		return true
	}
	return false
}

func (s *Server) setRunning(val bool) {
	if val {
		atomic.StoreInt32(&s.running, 1)
	} else {
		atomic.StoreInt32(&s.running, 0)
	}
}
