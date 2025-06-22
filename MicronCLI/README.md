# Developer Guide
https://github.com/spf13/cobra-cli/blob/main/README.md
https://github.com/spf13/cobra/blob/main/site/content/user_guide.md#using-the-cobra-library

# Commands
Run below commands from the folder which contains micronCLI.go:

## Run Source
go run main.go
go run main.go register --network mynetwork --service myservice --connection localhost:50051

## Build
go build -o ..\..\Builds\MicronCLI\micronCLI.exe main.go

# Usage
  micronCLI register --network mynetwork --service myservice --connection localhost:50051
  micronCLI register mynetwork myservice localhost:50051
  
  micronCLI query --network mynetwork --service myservice
  micronCLI query mynetwork myservice

  micronCLI clear --network mynetwork
  micronCLI clear mynetwork
  
  micronCLI unregister --network mynetwork --service myservice
  micronCLI unregister mynetwork myservice