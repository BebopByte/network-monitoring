package main

import (
	"net"
	"time"
)

func isAlive(ip string) bool {

	timeout := 300 * time.Millisecond
	ports := []string{"22", "80", "443"}

	for _, port := range ports {

		conn, err := net.DialTimeout("tcp", ip+":"+port, timeout)

		if err == nil {
			// for testing connectivity, close immediately after it connects
			conn.Close()
			return true
		}

	}

	return false
}
