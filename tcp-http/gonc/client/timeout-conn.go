package client

import (
	"net"
	"time"
)

type timeoutConn struct {
	net.Conn
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewTimeoutConn(conn net.Conn, readTimeout, writeTimeout time.Duration) net.Conn {
	return &timeoutConn{
		Conn:         conn,
		ReadTimeout:  readTimeout,
		WriteTimeout: writeTimeout,
	}
}

func (self *timeoutConn) Read(p []byte) (n int, err error) {
	var zero time.Duration
	if self.ReadTimeout != zero {
		self.Conn.SetReadDeadline(time.Now().Add(self.ReadTimeout))
	}
	return self.Conn.Read(p)
}

func (self *timeoutConn) Write(p []byte) (n int, err error) {
	var zero time.Duration
	if self.ReadTimeout != zero {
		self.Conn.SetReadDeadline(time.Now().Add(self.ReadTimeout))
	}
	return self.Conn.Write(p)
}
