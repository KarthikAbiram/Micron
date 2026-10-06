# MicronCLI skill

Use this skill when working with the Micron CLI tool for registering, listing, querying, starting, stopping, messaging, clearing, unregistering, and managing services in a network.

## Quick reference

Prefer flags for every argument supported by the command. The command implementation also accepts positional arguments as fallbacks, but flag-based syntax is the preferred form.

```powershell
microncli list --network mynetwork
microncli start --network mynetwork --service-id myservice --path "C:\Services\MyService\MyService.exe" --timeout 30
microncli message --network mynetwork --service-id myservice --command Add --payload '{"Numeric1":10,"Numeric2":20}'
microncli query --network mynetwork --service-id myservice
microncli stop --network mynetwork --service-id myservice --timeout 30
microncli register --network mynetwork --service-id myservice --connection localhost:50051
microncli unregister --network mynetwork --service-id myservice
microncli logs --limit 100
microncli purge --keep 1000
microncli help
microncli version
```

## Commands and flag-based syntax

### List
Lists active networks and services.

```powershell
microncli list
microncli list --network mynetwork
```

### Start a service
Starts a service instance from the specified path and registers it under the user provided network.

```powershell
microncli start --network mynetwork --service-id myservice --path C:\Services\MyService\MyService.exe --timeout 30
```

- `--network`: network name
- `--service-id`: service identifier
- `--path`: service path
- `--timeout`: timeout in seconds; defaults to `30`

The service instance is expected to register itself with the network after startup using the 'microncli register' command. The command waits for the service to register and returns a nonzero exit code if the service fails to register within the timeout period.

The command prints the connection string returned by the service startup operation.

### Message a service
Send a message to a service instance through the commandline.

```powershell
microncli message --network mynetwork --service-id myservice --command Add --payload '{"Numeric1":10,"Numeric2":20}'
microncli message --network mynetwork --service-id myservice --command help
microncli message --network mynetwork --service-id myservice --command help --payload '{"command":"Add"}'
```

- `--network`: network name
- `--service-id`: service identifier
- `--command`: command to invoke
- `--payload`: command payload as JSON text
- `--timeout`: timeout in seconds; defaults to `30`

For `cmd.exe`, escape the JSON quotes when the payload is supplied through a quoted command line:

```cmd
microncli message --network default --service-id app3service --command Add --payload "{\"Numeric1\":10,\"Numeric2\":20}"
```

### Query a service
The query returns the registered service connection string and reports a nonzero service status as an error.

```powershell
microncli query --network mynetwork --service-id myservice
```

### Stop a service
Send a stop message to the service and waits for it to unregister itself.

```powershell
microncli stop --network mynetwork --service-id myservice --timeout 30
```

- `--network`: network name
- `--service-id`: service identifier
- `--timeout`: timeout in seconds; defaults to `30`

### Register a service
When a microservice starts, it should register itself with micronCLI specifying its connection string. If there are any errors during the microservice startup, it can report a nonzero status code and optional additional information.
```powershell
microncli register --network mynetwork --service-id myservice --connection localhost:50051 --status 0 --info "Sample Info"
```

- `--network`: network name
- `--service-id`: service identifier
- `--connection`: service connection string
- `--status`: service status; defaults to `0`
- `--info`: optional additional information


### Unregister a service
When a microservice instance stops, it should unregister itself from micronCLI. 
```powershell
microncli unregister --network mynetwork --service-id myservice
```

### Clear a network
Clears all the available services in the specified network. This is useful for cleaning up a network before starting a new set of services.
```powershell
microncli clear --network mynetwork
```

### Find a free port
A helpful tool to check if a preferred port is available or to find a free port for a service to use. The command returns the port number that is available for use.
```powershell
microncli freeport --prefer 50051
```

- `--prefer`: preferred TCP port number; defaults to `0`, which requests any available port

### List recent logs
Lists the recent N log entries from the MicronCLI related to registration and unregistration of services along with timestamp for debugging.
```powershell
microncli logs --limit 100
```

- `--limit`: number of log entries to return; defaults to `100`

### Purge old logs
Deletes old log entries from the MicronCLI log file, retaining only the most recent N entries.
```powershell
microncli purge --keep 1000
```

- `--keep`: number of recent log entries to retain; defaults to `1000`
- Negative values are rejected

## Positional argument compatibility

The source still accepts positional arguments for commands but the flag-based syntax is preferred for clarity and maintainability.

```powershell
microncli list mynetwork
microncli start mynetwork myservice C:\Services\MyService\MyService.exe 30
microncli message mynetwork myservice Add '{"Numeric1":10,"Numeric2":20}'
microncli query mynetwork myservice
microncli stop mynetwork myservice 30
microncli register mynetwork myservice localhost:50051 0 "Sample Info"
microncli unregister mynetwork myservice
microncli clear mynetwork
microncli freeport 50051
microncli logs 100
microncli purge 1000
```

Prefer the flag forms above because they are clearer and remain aligned with the command definitions.

## JSON rules

When sending a JSON payload:

- use valid JSON
- keys must use double quotes
- strings must use double quotes
- do not use single quotes inside the JSON object itself

Valid JSON:

```json
{"Numeric1":10,"Numeric2":20}
```

Invalid JSON:

```json
{'Numeric1':10,'Numeric2':20}
{Numeric1:10,Numeric2:20}
```

For `cmd.exe`, escape the JSON quotes when the payload is supplied through a quoted command line:

```cmd
microncli message --network default --service-id app3service --command Add --payload "{\"Numeric1\":10,\"Numeric2\":20}"
```