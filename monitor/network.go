package monitor

import (
	"fmt"
	"os"

	"github.com/mackerelio/go-osstat/network"
)

func getNetwork() []NetworkPayload {
	network, err := network.Get()

	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}

	var networkData []NetworkPayload

	for i := 0; i < len(network); i++ {
		networkData = append(networkData, NetworkPayload{
			Name:     network[i].Name,
			Receive:  network[i].RxBytes,
			Transmit: network[i].TxBytes,
		})
	}

	return networkData
}
