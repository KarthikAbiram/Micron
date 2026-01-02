package library

import (
	"fmt"
	"io/fs"
	"microncli/library/grpcclient"
	"os"
	"path/filepath"
	"strings"
)

type ConnectionInfo struct {
	Network          string
	Service          string
	ConnectionString string
}

func GetMicronNetworksDirectory() (string, error) {
	// Step 1: Get the AppData folder directory (platform-independent approach)
	appDataDir, err := os.UserConfigDir() // Use UserConfigDir for application-specific config
	if err != nil {
		return "", fmt.Errorf("failed to get AppData directory: %v", err)
	}

	// Step 2: Create the Networks directory and network-specific file path
	micronNetworksDir := filepath.Join(appDataDir, "Micron", "Networks")
	if err := os.MkdirAll(micronNetworksDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create Network directory: %v", err)
	}

	return micronNetworksDir, err
}

func GetNetworkFolderPath(network string) (string, error) {
	// Step 1: Get Micron Networks directory
	micronNetworksDir, err := GetMicronNetworksDirectory()
	if err != nil {
		return "", fmt.Errorf("failed to get Micron networks directory: %v", err)
	}

	// Step 2: Create the Network folder and network-specific file path
	networkDir := filepath.Join(micronNetworksDir, network)
	if err := os.MkdirAll(networkDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("failed to create Network directory: %v", err)
	}

	return networkDir, err
}

func GetServiceFilePath(network, service string) (string, error) {
	networkDir, err := GetNetworkFolderPath(network)

	serviceFilePath := filepath.Join(networkDir, service+".txt")
	return serviceFilePath, err
}

// RegisterService registers the service and its connection string
func RegisterService(network, service, connectionString string) error {
	serviceFilePath, err := GetServiceFilePath(network, service)
	if err != nil {
		return fmt.Errorf("failed to get service file path: %v", err)
	}

	// Write connection string to service file
	if err := os.WriteFile(serviceFilePath, []byte(connectionString), 0644); err != nil {
		return fmt.Errorf("failed to write to network file: %v", err)
	}

	// fmt.Println(serviceFilePath)

	return nil
}

// RegisterService registers the service and its connection string
func UnregisterService(network, service string) error {
	serviceFilePath, err := GetServiceFilePath(network, service)
	if err != nil {
		return fmt.Errorf("failed to get service file path: %v", err)
	}

	err = os.Remove(serviceFilePath)

	// fmt.Println(serviceFilePath)

	return err
}

// QueryService queries the service for its connection string
func QueryService(network, service string) (string, error) {
	serviceFilePath, err := GetServiceFilePath(network, service)
	if err != nil {
		return "", fmt.Errorf("failed to get service file path: %v", err)
	}

	data, err := os.ReadFile(serviceFilePath)
	if err != nil {
		return "", fmt.Errorf("error reading file:%v", err)
	}

	//Convert bytes to string
	connectionString := string(data)

	return connectionString, err
}

func ListNetworkAndServices(network string) ([]ConnectionInfo, error) {
	var connections []ConnectionInfo
	var connectionString string
	// Step 1: Get Micron Networks directory
	micronNetworksDir, err := GetMicronNetworksDirectory()
	if err != nil {
		return connections, fmt.Errorf("failed to get Micron networks directory: %v", err)
	}

	//Step 2: Update query path, if a network is passed as an argument
	queryPath := micronNetworksDir
	if network != "" {
		queryPath = filepath.Join(micronNetworksDir, network)
	}

	//Step 3: Get list of all service files under the query path
	var serviceFilePaths []string
	err = filepath.WalkDir(queryPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err // stop walking if there's an error
		}

		if !d.IsDir() && filepath.Ext(d.Name()) == ".txt" {
			serviceFilePaths = append(serviceFilePaths, path)
		}
		return nil
	})

	//Step 4: Iterate over each and return network, service and connection string
	for _, serviceFilePath := range serviceFilePaths {
		network := filepath.Base(filepath.Dir(serviceFilePath))
		service := strings.TrimSuffix(
			filepath.Base(serviceFilePath),
			filepath.Ext(serviceFilePath),
		)
		connectionString, err = QueryService(network, service)

		connections = append(connections, ConnectionInfo{
			Network:          network,
			Service:          service,
			ConnectionString: connectionString,
		})
	}

	return connections, err
}

func MessageService(network, service, command, payload string) (string, error) {
	// fmt.Println(network, service, command, payload)
	//Step 1 : Get service connection string
	connectionString, err := QueryService(network, service)
	if err != nil {
		return "", fmt.Errorf("failed to get connection string for %s/%s: %w", network, service, err)
	}

	// Step 2: Create gRPC client
	client, err := grpcclient.New(connectionString)
	if err != nil {
		return "", fmt.Errorf("failed to connect to gRPC service: %w", err)
	}
	defer client.Close()

	// Step 3: Send the message
	resp, err := client.SendMessage(command, payload)
	if err != nil {
		return "", fmt.Errorf("error sending message via gRPC: %w", err)
	}

	// Step 4: Return the payload (or any other info)
	return resp.GetPayload(), err
}

// Clear/reset a network
func Clear(network string) error {
	networkDir, err := GetNetworkFolderPath(network)
	if err != nil {
		return fmt.Errorf("failed to get network folder path: %v", err)
	}

	err = os.RemoveAll(networkDir)
	return err
}
