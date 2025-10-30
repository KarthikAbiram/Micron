from concurrent import futures
import logging
import time

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

def serve():
    port = "50051"
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    micron_pb2_grpc.add_MicronGRPCServicer_to_server(Micron(), server)
    server.add_insecure_port("[::]:" + port)
    server.start()
    print("Server started, listening on " + port)

    while True:
        if stop_flag:
            time.sleep(1)
            server.stop(grace=True)
            break
        else:
            time.sleep(1)

    server.wait_for_termination()


if __name__ == "__main__":
    logging.basicConfig()
    serve()