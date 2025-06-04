micronCLI <network> register <service_name> <connection_string> : Register a service along with its connection string
micronCLI <network> query <service_name> : Returns the connection string of the service, if present.
micronCLI <network> unregister <service_name> <connection_string> : Unregister/remove a service
micronCLI <network> list : List all services in the network
micronCLI <network> clear : Clear/remove all services from network


Commands:
Run below commands from the folder which contains micronCLI.go:

go run micronCLI.go
go build -o ..\..\Builds\MicronCLI\micronCLI.exe micronCLI.go