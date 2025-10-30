# Command to generate from proto file
Execute parallel to 'generated' folder.
```
uv run python -m grpc_tools.protoc -I. --proto_path=. --python_out=. --pyi_out=. --grpc_python_out=. generated/micron.proto
```

# Credits
Template instantiated from [Micron](https://github.com/KarthikAbiram/Micron)