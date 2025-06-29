package gosocket

import (
	"bufio"
	"encoding/binary"
	"errors"
	"io"
	"net"
	"os"
)

// Message mirrors the wire format used by socket.cpp.
// It begins with three uldat values: total length, request id and code.
// Data then contains the remaining payload bytes.
type Message struct {
	Len  uint32
	Req  uint32
	Code uint32
	Data []byte
}

// Encode serializes the Message into a byte slice.
func (m Message) Encode() []byte {
	buf := make([]byte, 12+len(m.Data))
	binary.BigEndian.PutUint32(buf[0:4], uint32(12+len(m.Data)))
	binary.BigEndian.PutUint32(buf[4:8], m.Req)
	binary.BigEndian.PutUint32(buf[8:12], m.Code)
	copy(buf[12:], m.Data)
	return buf
}

// DecodeMessage reads one message from r.
func DecodeMessage(r io.Reader) (Message, error) {
	var hdr [12]byte
	if _, err := io.ReadFull(r, hdr[:]); err != nil {
		return Message{}, err
	}
	length := binary.BigEndian.Uint32(hdr[0:4])
	if length < 12 {
		return Message{}, errors.New("invalid length")
	}
	msg := Message{
		Len:  length,
		Req:  binary.BigEndian.Uint32(hdr[4:8]),
		Code: binary.BigEndian.Uint32(hdr[8:12]),
	}
	payload := make([]byte, length-12)
	if _, err := io.ReadFull(r, payload); err != nil {
		return Message{}, err
	}
	msg.Data = payload
	return msg, nil
}

// WriteMessage writes the encoded message to w.
func WriteMessage(w io.Writer, m Message) error {
	_, err := w.Write(m.Encode())
	return err
}

// Server listens on both a Unix domain socket and a TCP address.
type Server struct {
	UnixPath string
	TCPAddr  string
	unixLn   net.Listener
	tcpLn    net.Listener
}

// Start opens the listeners.
func (s *Server) Start() error {
	if s.UnixPath != "" {
		if err := os.RemoveAll(s.UnixPath); err != nil {
			return err
		}
		ln, err := net.Listen("unix", s.UnixPath)
		if err != nil {
			return err
		}
		s.unixLn = ln
	}
	if s.TCPAddr != "" {
		ln, err := net.Listen("tcp", s.TCPAddr)
		if err != nil {
			if s.unixLn != nil {
				s.unixLn.Close()
			}
			return err
		}
		s.tcpLn = ln
	}
	return nil
}

// Serve accepts incoming connections.
func (s *Server) Serve(handler func(Message, net.Conn)) error {
	if s.unixLn == nil && s.tcpLn == nil {
		return errors.New("server not started")
	}
	serveOne := func(ln net.Listener) {
		for {
			c, err := ln.Accept()
			if err != nil {
				return
			}
			go s.handleConn(c, handler)
		}
	}
	if s.unixLn != nil {
		go serveOne(s.unixLn)
	}
	if s.tcpLn != nil {
		go serveOne(s.tcpLn)
	}
	select {}
}

// Stop closes the listeners.
func (s *Server) Stop() {
	if s.unixLn != nil {
		s.unixLn.Close()
	}
	if s.tcpLn != nil {
		s.tcpLn.Close()
	}
}

func (s *Server) handleConn(c net.Conn, handler func(Message, net.Conn)) {
	defer c.Close()
	// very small handshake to mimic the C++ implementation
	// send protocol/version info
	c.Write([]byte("Twin-go\n"))

	r := bufio.NewReader(c)
	for {
		msg, err := DecodeMessage(r)
		if err != nil {
			return
		}
		if handler != nil {
			handler(msg, c)
		}
	}
}
