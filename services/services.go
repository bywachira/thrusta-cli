package services

import (
	"github.com/tesh254/thrusta-cli/config"
)

// FormatAPIUrl handles setting the api url based on protocol and endpoint
func FormatAPIUrl(endpoint string) string {
	cliConfig := config.ReadConfig()

	if len(cliConfig.URL) > 1 {
		if cliConfig.SSL {
			return "https://" + cliConfig.URL + "/api/v1" + endpoint
		} else {
			return "http://" + cliConfig.URL + "/api/v1" + endpoint
		}
	} else {
		return "https://usethrusta.com/api/v1" + endpoint
	}
}
