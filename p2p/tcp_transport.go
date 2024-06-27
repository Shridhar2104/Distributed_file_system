package p2p

import (
	"bytes"
	"fmt"
	"net"
	"sync"
)

//tcp peer represents remote node over a tcp established connection

type TCPPeer struct{

	conn net.Conn
	//if we dial a connetion
	outbound bool
}
func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer{
	return &TCPPeer{
		conn: conn,
		outbound: outbound,
	}
}

type TCPTransport struct{
	listenAddress string
	listener net.Listener
	handshakeFunc HandshakeFunc
	decoder Decoder

	mu sync.RWMutex
	peers map[net.Addr]Peer
}

func NewTCPTransport(listenAddress string) *TCPTransport{
	return &TCPTransport{
		listenAddress: listenAddress,
	}
}

func (t *TCPTransport) ListenAndAccept() error{
	var err error

	t.listener, err = net.Listen("tcp", t.listenAddress)

	if err!=nil{
		return err
	}
	go t.startAcceptLoop()

	return nil
}

func (t *TCPTransport) startAcceptLoop() error{
	for{
		conn, err :=t.listener.Accept()
		if err!=nil{
			fmt.Printf("TCP transport accept error: %s\n",err )
		}
		go t.handleConn(conn)
	}
}
func (t *TCPTransport) handleConn(conn net.Conn){
	peer := NewTCPPeer(conn, false)

	if err := t.shakeHands(conn); err!=nil{
		
	}
	buf:=new(bytes.Buffer)
	for{
		n, _ := conn.read(buf)
	}

	fmt.Printf("new incoming connection %v\n", peer)
}