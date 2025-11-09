from concurrent import futures
from datetime import datetime
import logging
import time
import subprocess

import fire

import grpc
import generated.micron_pb2 as micron_pb2
import generated.micron_pb2_grpc as micron_pb2_grpc

# Flags
stop_flag = False

class Micron(micron_pb2_grpc.MicronGRPCServicer):
    def Message(self, request, context):
        # Inputs
        cmd = request.command.lower() or ""
        payload = request.payload or ""

        # Initialize Outputs
        reply = micron_pb2.MessageReply()

        if cmd == "stop":
            global stop_flag
            reply.payload = "Stopping..."
            stop_flag = True
        elif cmd == "ping":
            reply.payload = f"Hello {payload}"
        else:
            reply.payload = f"Echo {cmd}:{payload}"

        print(f"{datetime.now()} {request.command}:{payload} -> {reply.payload}")
        return reply

def serve(network="default", service_id="pymicron", port=50052, **kwargs):
    # print(f"Received Args: network: {network}, service_id: {service_id}, port={port}, kwargs={kwargs}")
    skip_register_on_error = False
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    micron_pb2_grpc.add_MicronGRPCServicer_to_server(Micron(), server)
    server.add_insecure_port("[::]:" + str(port))
    server.start()
    print(f"Server started, listening on {port}")
    register_service(network, service_id, f"localhost:{port}", skip_register_on_error)
    print("Waiting for commands from client...")

    while True:
        if stop_flag:
            time.sleep(1)
            server.stop(grace=True)
            break
        else:
            time.sleep(1)

    unregister_service(network, service_id, skip_register_on_error)
    server.wait_for_termination()

def register_service(network, service_id, connection, skip_on_error=True):
    try:
        result = subprocess.run(
        ["micronCLI","register","--network", network,"--service-id", service_id,"--connection", connection],
        capture_output=True,
        text=True)
        if result.returncode == 0:
            print(f"Successfully registered the micron service '{service_id}' in network '{network}' with connection string '{connection}'")
        else:
            print(f"Failed to register the micron service '{service_id}' in network '{network}' with micronCLI")
            print("Status :", result.returncode, result.stdout, result.stderr)
    except Exception as e:
        if skip_on_error:
            pass
        else:
            raise e
        
def unregister_service(network, service_id, skip_on_error=True):
    try:
        result = subprocess.run(
        ["micronCLI","unregister","--network", network,"--service-id", service_id],
        capture_output=True,
        text=True
        )
        if result.returncode == 0:
            print(f"Successfully unregistered the micron service '{service_id}' in network '{network}'")
        else:
            print(f"Failed to unregister the micron service '{service_id}' in network '{network}' with micronCLI")
            print("Status :", result.returncode, result.stdout, result.stderr)
    except Exception as e:
        if skip_on_error:
            pass
        else:
            raise e

if __name__ == "__main__":
    logging.basicConfig()
    fire.Fire(serve)