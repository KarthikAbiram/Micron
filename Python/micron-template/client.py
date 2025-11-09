import logging
from datetime import datetime
import grpc
import generated.micron_pb2 as micron_pb2
import generated.micron_pb2_grpc as micron_pb2_grpc
import subprocess

def run():
    print("Starting communication to server...")
    default_connection = "localhost:50052" # Default connection string
    network = "default"
    service_id = "pymicron" 
    connection = query_service(network, service_id)
    if not connection:
        connection = default_connection
    with grpc.insecure_channel(connection) as channel:
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

def query_service(network, service_id, skip_on_error=True):
    try:
        connection_string = ""
        result = subprocess.run(
        ["micronCLI","query","--network", network,"--service-id", service_id],
        capture_output=True,
        text=True)
        if result.returncode == 0:
            connection_string = result.stdout.strip()
            # print(f"Successfully queried the connection string for micron service '{service_id}' in network '{network}': '{connection_string}'")
            return connection_string
        else:
            print(f"Failed to query the connection string from the micron service '{service_id}' in network '{network}' with micronCLI")
            print("Status :", result.returncode, result.stdout, result.stderr)
            return connection_string
    except Exception as e:
        if skip_on_error:
            pass
        else:
            raise e

if __name__ == "__main__":
    logging.basicConfig()
    run()