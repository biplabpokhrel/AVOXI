package repository

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/biplabpokhrel/avoxi/internal/network.map"
)

func getCountryName(URL string, IP string) string {
	resp, err := http.PostForm(URL, url.Values{"ip": {IP}})
	if err != nil {
		print(err)
	}

	defer resp.Body.Close()

	var data map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		print(err)
	}
	return data["country_name"].(string)
}

type locationNetworkRepository struct {
	httpURL string
}

// LocationNetworkRepository return repository instance of Network
func LocationNetworkRepository(httpURL string) network.Repository {
	return &locationNetworkRepository{httpURL}
}

// CheckIPStatus accept IPv4 or IP4 and if it matched found then it return true otherwise false
func (repo *locationNetworkRepository) CheckIPStatus(IP string, whiteListedCountries *[]string) (bool, error) {
	_countryName := getCountryName(repo.httpURL, IP)
	for _, country := range *whiteListedCountries {
		if country == _countryName {
			return true, nil
		}
	}
	return false, nil
}
