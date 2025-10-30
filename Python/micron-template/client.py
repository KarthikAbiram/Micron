import logging

import grpc
import generated.micron_pb2 as micron_pb2
import generated.micron_pb2_grpc as micron_pb2_grpc

def run():
    print("Starting communication to server...")
    with grpc.insecure_channel("localhost:50051") as channel:
        micron = micron_pb2_grpc.MicronGRPCStub(channel)

        # Send a ping command
        input = micron_pb2.MessageRequest(command="Ping", payload="Micron Client")
        response = micron.Message(input)
        print(f"{input.command}:{input.payload} -> {response.payload}")

        # Send a stop command to stop the server
        input.command = "Stop"
        response = micron.Message(input)
        print(f"{input.command}:{input.payload} -> {response.payload}")


if __name__ == "__main__":
    logging.basicConfig()
    run()