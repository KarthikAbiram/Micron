from concurrent import futures
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
        else:
            reply.payload = f"Hello {payload}"

        print(f"{request.command}:{payload} -> {reply.payload}")
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
        print("Exit code:", result.returncode, result.stdout, result.stderr)
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
        print("Exit code:", result.returncode, result.stdout, result.stderr)
    except Exception as e:
        if skip_on_error:
            pass
        else:
            raise e

if __name__ == "__main__":
    logging.basicConfig()
    fire.Fire(serve)