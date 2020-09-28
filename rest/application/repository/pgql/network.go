package repository

import (
	"log"

	"github.com/biplabpokhrel/avoxi/internal/network.map"
	"github.com/biplabpokhrel/avoxi/models"
	"github.com/biplabpokhrel/avoxi/utility"
	"github.com/go-pg/pg/v10"
)

type pgqlNetworkRepository struct {
	DB *pg.DB
}

// PgqlNetworkRepository return repository instance of Network
func PgqlNetworkRepository(DB *pg.DB) network.Repository {
	return &pgqlNetworkRepository{DB}
}

// CheckIPStatus accept IPv4 or IP4 and if it matched found then it return true otherwise false
func (repo *pgqlNetworkRepository) CheckIPStatus(IP string, whiteListedCountries *[]string) (bool, error) {
	var _networks []*models.NetworkMap
	err := repo.DB.Model(&_networks).Select()
	if err != nil {
		log.Print("Database related issue")
		log.Print(err)
		return false, nil
	}

	// TODO we can improvise this logic, it's super slow but working for now
	// We must have to come back
	for _, _network := range _networks {
		log.Printf("%+v\n", _network)
		if status := utility.CheckIPRange(_network.Network, IP); status == true {
			for _, country := range *whiteListedCountries {
				if country == _network.AutonomousSystemOrganization {
					return status, nil
				}
			}
		}
	}

	return false, nil
}
