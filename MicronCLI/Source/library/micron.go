package library

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetNetworkFolderPath(network string) (string, error) {
	// Step 1: Get the AppData folder directory (platform-independent approach)
	appDataDir, err := os.UserConfigDir() // Use UserConfigDir for application-specific config
	if err != nil {
		return "", fmt.Errorf("failed to get AppData directory: %v", err)
	}

	// Step 2: Create the Networks directory and network-specific file path
	networkDir := filepath.Join(appDataDir, "Micron", "Networks", network)
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

// Clear/reset a network
func Clear(network string) error {
	networkDir, err := GetNetworkFolderPath(network)
	if err != nil {
		return fmt.Errorf("failed to get network folder path: %v", err)
	}

	err = os.RemoveAll(networkDir)
	return err
}
