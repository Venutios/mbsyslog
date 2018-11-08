package mbsyslog

import (
	"fmt"
	"net"
)

//Server is a syslog server that forms the basis for a relay or collector in
//the syslog system. Data is received and parsed into syslog messages.
type Server struct {
	maxMessageSize int
	port           int
}

//NewServer prepares the server to listen for messages. The server will listen
//on port 514 and have an 8KB buffer.
func NewServer() *Server {
	result := new(Server)
	result.port = 514
	result.maxMessageSize = 8192
	return result
}

//Listen starts the server accepting syslog messages
func (s Server) Listen() error {
	conn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: s.port, Zone: ""})
	if err != nil {
		return err
	}
	defer conn.Close()

	buffer := make([]byte, s.maxMessageSize)
	for {
		count, addr, err := conn.ReadFromUDP(buffer)
		if err != nil {
			return err
		}
		fmt.Println("Message from " + addr.String())
		go func(data []byte) {
			message := NewMessage(data)
			fmt.Println(message.String())
		}(buffer[0:count])
	}
}

/**
 * Starts the server accepting messages until the server is requested to
 * stop. Run doesn't return until the server has been stopped. Events will
 * be triggered in multiple separate threads.
 */
// @Override
// public void run() {
//     stop = false;
//     running = true;

//     ExecutorService taskExec = Executors.newCachedThreadPool();

//     while (!stop) {
//         //Create a new packet to receive the next message
//         DatagramPacket packet = new DatagramPacket(buffer, buffer.length);

//         try {
//             //listen until the next message is received
//             socket.receive(packet);

//             //Decode the packet into a message
//             Message m = new Message(packet);

//             //Call each listener passing it to the thread pool for execution
//             synchronized (messageMonitor) {
//                 taskExec.submit(() -> {
//                     messageListeners.forEach(listener -> listener.onMessageEvent(m));
//                 });
//             }
//         } catch (SocketTimeoutException ex) {
//             //ignore, the exception, lets the loop poll for the stop flag
//         } catch (IOException ex) {
//             //Call each listener passing it to the thread pool for execution
//             synchronized (errorMonitor) {
//                 taskExec.submit(() -> {
//                     errorListeners.forEach(listener -> listener.onErrorEvent(ex,
//                             Thread.currentThread().getStackTrace()[1].getClassName() + "."
//                             + Thread.currentThread().getStackTrace()[1].getMethodName() + " on line "
//                             + Thread.currentThread().getStackTrace()[1].getLineNumber()));
//                 });
//             }
//         }
//     }
//     //stop accepting messages
//     socket.close();

//     //stop accepting new tasks and wait for all tasks to complete
//     taskExec.shutdown();
//     while (!taskExec.isTerminated()) {
//         try {
//             taskExec.awaitTermination(1, TimeUnit.MINUTES);
//         } catch (InterruptedException ex) {
//             //ignore the timeout, and keep waiting if necessary
//         }
//     }

//     running = false;
// }

//Port that the server is configured to listen on
func (s Server) Port() int {
	return s.port
}

//MaximumMessageSize is the largest syslog message that can be accepted.
func (s Server) MaximumMessageSize() int {
	return s.maxMessageSize
}
