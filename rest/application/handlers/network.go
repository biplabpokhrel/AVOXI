package handlers

import (
	"net"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// GetNetwork method read GET HTTP method's data and respond with status data
func (route *Route) GetNetwork(_context *gin.Context) {
	IPAddress := _context.Query("ip")

	//convert string to list of string
	whiteListedCountries := strings.Split(_context.Query("whitelist"), ",")

	if net.ParseIP(IPAddress) == nil {
		_context.JSON(http.StatusBadRequest, gin.H{"IP": "Invalid IP address"})
		return
	}

	checkIPStatus, _error := route.Internal.NetworkRepo.CheckIPStatus(IPAddress, &whiteListedCountries)
	if _error != nil {
		_context.JSON(http.StatusInternalServerError, gin.H{"IpStatus": false})
		return
	}

	status := "Not Allowed"

	if checkIPStatus {
		status = "Allowed"
	}

	_context.JSON(http.StatusOK, gin.H{"IP": status})
}
