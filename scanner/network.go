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

func getHostname(ip string) string {

	names, err := net.LookupAddr(ip)

	if err != nil || len(names) == 0 {
		return "Unknown"
	}

	return names[0]
}
