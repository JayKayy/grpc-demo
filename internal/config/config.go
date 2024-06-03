package config

import "net"

const (
	host = "localhost"
	port = "8080"
)

var SrvAddr = net.JoinHostPort(host, port)
