package tcpserver

import (
	"fmt"
	"github.com/mcLyu/tcp-proof-of-work/internal/proof-of-work/hashcash/challenge"
	"github.com/mcLyu/tcp-proof-of-work/internal/proof-of-work/hashcash/solution"
	"github.com/mcLyu/tcp-proof-of-work/internal/quotes"
	"log"
	"net"
	"strings"
)

const TcpNetwork = "tcp"

type TcpServer struct {
	listener net.Listener
}

func New() *TcpServer {
	return &TcpServer{}
}

func (s *TcpServer) Listen(port string) {
	listener, err := net.Listen(TcpNetwork, ":"+port)
	if err != nil {
		fmt.Printf("Error starting server: %s", err.Error())
		return
	}

	fmt.Println("Server successfully started, waiting for connections..")
	defer func(listener net.Listener) {
		err = listener.Close()
		if err != nil {

		}
	}(listener)

	connCounter := &ConnectionsCounter{}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection: ", err)
			continue
		}

		connCounter.Increment()
		go s.handleConnection(conn, connCounter)
	}
}

func (s *TcpServer) handleConnection(conn net.Conn, connCounter *ConnectionsCounter) {
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			fmt.Printf("Error during closing server connection : %v\n", err.Error())
		}

		connCounter.Decrement()
	}(conn)

	remoteAddr := conn.RemoteAddr().String()

	difficulty := s.resolveChallengeDifficulty(connCounter)
	fmt.Printf("Accepted new connection : %s, connection num : %d, challenge difficulty : %d\n", remoteAddr, connCounter.Count(), difficulty)

	err := s.handlePowChallenge(remoteAddr, conn, s.resolveChallengeDifficulty(connCounter))
	if err != nil {
		fmt.Printf("Error during pow challenge : %s", err.Error())
		return
	}

	fmt.Printf("PoW challenge passed by client %s, sending quote from 'Book of Wisdom'..\n", remoteAddr)
	s.sendQuote(remoteAddr, conn)

}

func (s *TcpServer) handlePowChallenge(remoteAddr string, conn net.Conn, difficulty int) error {
	readBuffer := make([]byte, 4096)

	addrWithoutPort := strings.Split(remoteAddr, ":")[0]
	defaultChallenge := challenge.NewDefaultChallenge(addrWithoutPort, difficulty)
	challengeRequest := defaultChallenge.String()

	_, err := conn.Write([]byte(challengeRequest))
	if err != nil {
		fmt.Printf("Fail sending challenge to client %s\n", remoteAddr)
		return err
	}

	fmt.Printf("PoW challenge '%s' sent to %s, waiting for solution..\n", challengeRequest, remoteAddr)
	numberOfBytes, err := conn.Read(readBuffer)
	if err != nil {
		fmt.Println("Fail reading from connection")
		return err
	}

	challengeResponse := string(readBuffer[:numberOfBytes])
	fmt.Printf("PoW solution received from client %s, validating..\n", remoteAddr)

	if !strings.Contains(challengeResponse, challengeRequest) {
		err = fmt.Errorf("challenge response '%s' from client doesn't contain original string", challengeResponse)
		return err
	}

	challengeSolution, err := solution.FromString(challengeResponse)
	if err != nil {
		fmt.Println("Fail converting client solution")
		return err
	}

	hasher, _ := challengeSolution.Algorithm.GetHasher()
	valid := challengeSolution.IsValidSolution(challengeResponse, hasher)
	if !valid {
		err = fmt.Errorf("PoW solution '%s' from client %s is not valid\n", challengeResponse, remoteAddr)
		return err
	}

	return nil
}

func (s *TcpServer) sendQuote(remoteAddr string, conn net.Conn) {
	quote := quotes.GetRandomQuote()

	_, err := conn.Write([]byte(quote))
	if err != nil {
		fmt.Printf("Fail sending quote '%s' to client %s : %s\n", quote, remoteAddr, err.Error())
	}
}

// resolveChallengeDifficulty is the simple implementation example of increasing PoW difficulty
// This difficulty is increasing linear based on number of active connections
func (s *TcpServer) resolveChallengeDifficulty(counter *ConnectionsCounter) int {
	return int(challenge.DefaultDifficulty + counter.Count())
}
