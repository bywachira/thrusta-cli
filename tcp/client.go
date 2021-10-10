package tcp

import (
	"encoding/json"
	"net"
	"os"
	Stats "thrusta/status"
	Utils "thrusta/utils"
)

type ServerStatus struct {
	NodeCredentials Utils.CredentialsConfiguration `json:"node_credentials"`
	SystemStats     Stats.Statistics               `json:"system_stats"`
}

func ConnectToBackend() {
	stats := Stats.GetSystemStats()
	creds := Utils.ReadCredentialsFile()

	var serverStatus ServerStatus = ServerStatus{
		NodeCredentials: creds,
		SystemStats:     stats,
	}

	statsData, err := json.Marshal(&serverStatus)
	strEcho := string(statsData)
	servAddr := "localhost:6000"
	tcpAddr, err := net.ResolveTCPAddr("tcp", servAddr)
	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte(strEcho))
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	// println("write to server = ", strEcho)

	reply := make([]byte, 1024)

	_, err = conn.Read(reply)
	if err != nil {
		println("Write to server failed:", err.Error())
		os.Exit(1)
	}

	println("reply from server=", string(reply))

	conn.Close()
}
