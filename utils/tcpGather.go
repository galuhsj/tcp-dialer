package utils

import (
	"net"
	"time"
)

// TCPGather ...
func TCPGather(ip string, ports []string, dialTimeoutInMinute time.Duration) map[string]string {
	results := make(map[string]string)

	for _, port := range ports {
		address := net.JoinHostPort(ip, port)
		// 3 second timeout
		conn, err := net.DialTimeout("tcp", address, dialTimeoutInMinute*time.Second)
		if err != nil {
			results[port] = "failed"
			// todo log handler
		} else {
			if conn != nil {
				results[port] = "success"
				_ = conn.Close()
			} else {
				results[port] = "failed"
			}
		}
	}

	return results
}
