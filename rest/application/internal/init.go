package internal

import "github.com/biplabpokhrel/avoxi/internal/network.map"

// Internal ...
type Internal struct {
	NetworkRepo network.Repository
}

// New function initialize network repository
func New(networkRepo network.Repository) *Internal {
	return &Internal{NetworkRepo: networkRepo}
}
