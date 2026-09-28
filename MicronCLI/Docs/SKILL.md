# MicronCLI skill

Use this skill when working with the Micron CLI tool for service registration, querying, clearing, unregistering, and messaging services in a network.

## Quick reference

```powershell
microncli help
microncli version
microncli list
microncli list --network mynetwork
microncli message <network> <service> help
microncli message <network> <service> help <command name>
microncli message <network> <service> help <command name> <payload_json>
```

## Command patterns


### List networks and services

```powershell
microncli list
microncli list --network mynetwork
microncli list mynetwork
```

## Message usage

```powershell
microncli message <network> <service> help
microncli message <network> <service> help <command name>
microncli message <network> <service> <command name> <payload_json>
```

Example:

```powershell
microncli message default app3service Add '{"Numeric1":10,"Numeric2":20}'
```

For `cmd.exe`, escape the JSON quotes:

```cmd
microncli message default app3service Add "{\"Numeric1\":10,\"Numeric2\":20}"
```

## JSON rules

When sending JSON payloads:

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