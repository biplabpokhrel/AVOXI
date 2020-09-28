package handlers

import "github.com/biplabpokhrel/avoxi/internal"

// Route provide a way to access the repository to the routeHandler
type Route struct {
	Internal *internal.Internal
}
