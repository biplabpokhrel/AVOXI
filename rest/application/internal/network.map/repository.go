package network

// Repository interface contains method and their signature
// It gauranteed us that any service data implement this Repository interface
// must have the function CheckIPStatus which should accept one string type value
// and return boolean and error value
type Repository interface {
	CheckIPStatus(ipAddress string, whiteListedContries *[]string) (bool, error)
}
