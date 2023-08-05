package getuserlocation

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type IpInfoResponse struct {
	IP       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func GetUserLocation(ipAddress string) (*IpInfoResponse, error) {
	ipAddress = strings.Trim(ipAddress, `" `)
	formated_url := strings.Trim(fmt.Sprintf("https://ipinfo.io/%v", ipAddress), "\n")

	response, err := http.Get(formated_url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	parsedBody := &IpInfoResponse{}
	json.NewDecoder(response.Body).Decode(parsedBody)

	return parsedBody, nil
}
