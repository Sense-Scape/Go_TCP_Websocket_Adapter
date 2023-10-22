package main

import (
	"encoding/json"
	"net"
	"os"

	"github.com/Sense-Scape/Go_TCP_Websocket_Adapter/v2/Routines"
)

func main() {

	// Create a decoder to read JSON data from the file
	// Open the JSON file for reading
	configFile, err := os.Open("Config.json")

	if err != nil {
		os.Exit(1)
		return
	}
	defer configFile.Close()

	decoder := json.NewDecoder(configFile)

	// Create a variable to store the decoded JSON data
	var serverConfigMap map[string]interface{}
	// Decode the JSON data into a map
	if err := decoder.Decode(&serverConfigMap); err != nil {
		return
	}

	// Define the TCP port to listen on
	port := "10100"
	// Create a channel to pass time chunk json docs around
	TimeChunkDataChannel := make(chan string) // Create an integer channel

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal().Msg("Error:" + err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	logger.Info().Msg("TCP server is listening on port:" + port)

	// Accept incoming TCP connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Error().Msg("Error:" + err.Error())
			continue
		}
		// Start up TCP and websocket threads
		go Routines.HandleTCPReceivals(TimeChunkDataChannel, conn)
		//go Routines.HandleWebSocketTimeChunkTransmissions(TimeChunkDataChannel, "/DataTypes/TimeChunk")
	}

}
