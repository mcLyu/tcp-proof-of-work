package tcpclient

import (
	"fmt"
	"github.com/mcLyu/tcp-proof-of-work/internal/proof-of-work/hashcash/challenge"
	"github.com/mcLyu/tcp-proof-of-work/internal/tcpserver"
	"net"
)

type TcpClient struct{}

func New() *TcpClient {
	return &TcpClient{}
}

func (c *TcpClient) Connect(port string) {
	conn, err := net.Dial(tcpserver.TcpNetwork, ":"+port)
	if err != nil {
		fmt.Printf("Error during connecting to server : %v\n", err.Error())
		return
	}

	defer func(conn net.Conn) {
		err = conn.Close()
		if err != nil {
			fmt.Printf("Error during closing client connection : %v\n", err.Error())
		}
	}(conn)

	err = c.handlePowChallenge(conn)
	if err != nil {
		fmt.Printf("Error during processing PoW challenge : %v\n", err.Error())
		return
	}

	err = c.receiveQuote(conn)
	if err != nil {
		fmt.Printf("Error during receiving quote : %v\n", err.Error())
		return
	}
}

func (c *TcpClient) handlePowChallenge(conn net.Conn) error {
	readBuf := make([]byte, 4096)

	fmt.Printf("Connected to server : %s, waiting for PoW challenge..\n", conn.RemoteAddr().String())
	numberOfBytes, err := conn.Read(readBuf)
	if err != nil {
		fmt.Printf("Error during reading hashcash challenge")
		return err
	}

	hashcashChallengeRequest := string(readBuf[:numberOfBytes])
	fmt.Printf("Received PoW challenge request: %s\n", hashcashChallengeRequest)

	hashcashChallenge, err := challenge.FromString(hashcashChallengeRequest)
	if err != nil {
		fmt.Println("Error during parsing hashcash string")
		return err
	}

	solution, err := hashcashChallenge.FindSolution()
	if err != nil {
		fmt.Println("Error during finding PoW solution")
		return err
	}

	fmt.Printf("Found PoW solution : %s\n", *solution)
	_, err = conn.Write([]byte(*solution))
	if err != nil {
		fmt.Println("Error writing hashcash solution")
		return err
	}

	return nil
}

func (c *TcpClient) receiveQuote(conn net.Conn) error {
	readBuf := make([]byte, 4096)

	numberOfBytes, err := conn.Read(readBuf)
	if err != nil {
		fmt.Printf("Error during reading data\n")
		return err
	}

	fmt.Printf("Received quote from server: '%s'\n", string(readBuf[:numberOfBytes]))

	return nil
}
