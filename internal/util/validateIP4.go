package util

import (
	"net"
)

func ValidateIP4(ip string) bool {
	ipAddress := net.ParseIP(ip)

	if ipAddress == nil {
		return false
	}
	if ipAddress.To4() == nil {
		return false
	}

	return true
}
