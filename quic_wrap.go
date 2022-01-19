package main

import (
	"github.com/lucas-clemente/quic-go"
	"net"
	"time"
)

type quicConn struct {
	quic.Session
	quic.Stream
}

func (q quicConn) Read(b []byte) (n int, err error) {
	return q.Stream.Read(b)
}

func (q quicConn) Write(b []byte) (n int, err error) {
	return q.Stream.Write(b)
}

func (q quicConn) Close() error {
	return q.Stream.Close()
}

func (q quicConn) LocalAddr() net.Addr {
	return q.Session.LocalAddr()
}

func (q quicConn) RemoteAddr() net.Addr {
	return q.Session.RemoteAddr()
}

func (q quicConn) SetDeadline(t time.Time) error {
	return q.Stream.SetDeadline(t)
}

func (q quicConn) SetReadDeadline(t time.Time) error {
	return q.Stream.SetReadDeadline(t)
}

func (q quicConn) SetWriteDeadline(t time.Time) error {
	return q.Stream.SetWriteDeadline(t)
}
