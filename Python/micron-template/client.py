import logging
from datetime import datetime
import grpc
import generated.micron_pb2 as micron_pb2
import generated.micron_pb2_grpc as micron_pb2_grpc

def run():
    print("Starting communication to server...")
    with grpc.insecure_channel("localhost:50052") as channel:
        micron = micron_pb2_grpc.MicronGRPCStub(channel)

        # Create micron message
        message = micron_pb2.MessageRequest(command="Ping", payload="Micron Client")

        # Loop for sending multiple commands to server based on user choice
        stop_flag = False
        while not stop_flag:
            choice = input("\nChoose command: 1. Ping 2. Custom Message 3. Stop\n")

            if choice == "1":
                message.command = "Ping"
                message.payload = "Micron Client"
            elif choice == "2":
                message.command = input("Enter custom command:")
                message.payload = input("Enter payload:")
            elif choice == "3":
                message.command = "Stop"
                message.payload = "Micron Client"
                stop_flag = True
            else:
                print("Invalid choice. Pinging instead.")
                message.command = "Ping"
                message.payload = "Micron Client"

            response = micron.Message(message)
            print(f"{datetime.now()} {message.command}:{message.payload} -> {response.payload}")


if __name__ == "__main__":
    logging.basicConfig()
    run()