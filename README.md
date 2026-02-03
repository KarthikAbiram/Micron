# Micron's Vision
- Enable easy exchange and swapping of microservices built with different programming languages easily.
- Create a docker compose type system for service executables, instead of containers.
- Provide tooling, libraries and batteries missing in default gRPC, like a client starting and stopping services.
- Reduce the barrier of entry in creating microservices in different programming languages.

# Features
- MicronCLI provides ability for registering/deregistering, when the service is started/stopped
- Smart Start - If already running, connect to the already running instance. If not, start new instance.
- Supports running multiple instances of a service in parallel, each with their own unique IDs & ability to send message to each instance
- Dynamic port allocation. No need to hard code port numbers for services and avoids port conflicts.
- Automatic generation of help with ability to list available APIs.
- Ability to show help for each API and their default values
- Ability to send a message to the service from commandline
- Automatic retrying sending of messages
- Show/hide service UI
- Dynamic loading of API's UI into service UI for showing progress/debugging
- LabVIEW Micron Template to start a new service with a click of a button with build spec ready to build & deploy

# Languages
Currently supported languages: LabVIEW, Python

## Other Supported Languages
Micron leverages and uses gRPC. So, any programming language supported by gRPC can be used - C#, Go, Java & more.

# Installation & Usage
Please refer to the corresponding programming languages:
1. LabVIEW

# Author & License
Micron is an open source project developed by Karthik Abiram and is licensed under MIT License.