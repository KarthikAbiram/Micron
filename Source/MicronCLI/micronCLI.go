package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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
		return "", fmt.Errorf("Error reading file:%v", err)
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

func micronCLI(Args []string) (string, error) {
	// fmt.Println(Args)

	help := `
	Usage:
		micronCLI <network> register <service_name> <connection_string> : Register a service along with its connection string
		micronCLI <network> query <service_name> : Returns the connection string of the service, if present.
		micronCLI <network> unregister <service_name> <connection_string> : Unregister/remove a service
		micronCLI <network> clear : Clear/remove all services from network
	`

	response := ""
	var err error = nil

	if len(Args) < 3 {
		fmt.Println(help)
		return response, fmt.Errorf("Invalid number of args. %v", help)
	}

	network := strings.ToLower(Args[1])
	cmd := strings.ToLower(Args[2])

	switch cmd {
	case "register":
		{
			if len(Args) < 5 {
				return response, fmt.Errorf("Invalid number of args. %v", help)
			}
			service := strings.ToLower(Args[3])
			connection_string := Args[4]
			err = RegisterService(network, service, connection_string)
			return "", err
		}
	case "query":
		{
			if len(Args) < 4 {
				return response, fmt.Errorf("Invalid number of args. %v", help)
			}
			service := strings.ToLower(Args[3])
			response, err = QueryService(network, service)
			return response, err
		}
	case "unregister":
		{
			if len(Args) < 4 {
				return response, fmt.Errorf("Invalid number of args. %v", help)
			}
			service := strings.ToLower(Args[3])
			err = UnregisterService(network, service)
			return "", err
		}
	case "clear":
		{
			err = Clear(network)
			return "", err
		}
	default:
		{
			return "", fmt.Errorf("Command not supported: %v", cmd)
		}
	}
}

func example() {
	err := RegisterService("ProdNet", "DatabaseService1", "Server=127.0.0.1;Database=mydb;User Id=admin;Password=secret;")
	if err != nil {
		fmt.Println("Error:", err)
	}

	connection, err := QueryService("ProdNet", "databaseservice1")
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(connection)

	err = UnregisterService("ProdNet", "DatabaseService1")
	if err != nil {
		fmt.Println("Error:", err)
	}

	err = Clear("ProdNet")
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func main() {
	response, err := micronCLI(os.Args)
	if err == nil {
		if response != "" {
			fmt.Println(response)
		}
	} else {
		fmt.Println("Error:", err)
	}
	// example()
}
