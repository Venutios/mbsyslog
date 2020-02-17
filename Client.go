package mbsyslog

import (
	"errors"
	"net"
	"sync"
)

//Client is a syslog client to send messages to syslog servers
type Client struct {
	syncSend   bool
	wg         sync.WaitGroup
	asyncError error
	mutex      *sync.Mutex
}

//NewClient prepares a client to send messages
func NewClient(syncSend bool) *Client {
	result := new(Client)
	result.syncSend = syncSend
	result.asyncError = nil
	result.mutex = &sync.Mutex{}
	return result
}

//SendData sends raw data messages to remote IP or hostname address
func (c *Client) SendData(addr string, data []byte) error {
	if c.syncSend {
		return c.syncSendData(addr, data)
	}

	c.asyncSendData(addr, data)
	return nil
}

//AsyncError returns the last error from an asynchronous operation
func (c *Client) AsyncError() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.asyncError
}

//Wait for all current asynchronous operations to complete
func (c *Client) Wait() {
	c.wg.Wait()
}

func (c *Client) syncSendData(addr string, data []byte) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr+":514")
	if err != nil {
		return err
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		return err
	}
	defer conn.Close()

	count, err := conn.Write(data)
	if err != nil {
		return err
	}

	if count != len(data) {
		return errors.New("Wrong number of bytes written")
	}
	return nil
}

func (c *Client) asyncSendData(addr string, data []byte) {
	c.wg.Add(1)
	go func(a string, d []byte) {
		defer c.wg.Done()
		c.setAsyncError(c.syncSendData(a, d))
	}(addr, data)
}

func (c *Client) setAsyncError(err error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.asyncError = err
}
