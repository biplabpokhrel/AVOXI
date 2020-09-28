package utility

import (
	"log"
	"net"
)

// CheckIPRange function accept network address with subnet mask as string "192.0.0.1/24"
// and compare whether the checkAgainstIP fall between the range or not
func CheckIPRange(networkAddress string, checkAgainstIP string) bool {
	log.Println("CheckIPRange")
	_, subnet, err := net.ParseCIDR(networkAddress)
	if err != nil {
		return false
	}
	ip := net.ParseIP(checkAgainstIP)
	if subnet.Contains(ip) {
		return true
	}
	return false
}
