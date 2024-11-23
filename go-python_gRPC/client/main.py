# client/main.py
import grpc
import user_pb2
import user_pb2_grpc

def run():
    # Create a channel and stub
    with grpc.insecure_channel('localhost:50051') as channel:
        stub = user_pb2_grpc.UserServiceStub(channel)
        
        # Get a single user
        try:
            response = stub.GetUser(user_pb2.UserRequest(user_id=1))
            print("User found:", response)
        except grpc.RpcError as e:
            print(f"Error getting user: {e.details()}")
        
        # List all users
        try:
            response = stub.ListUsers(user_pb2.Empty())
            print("\nAll users:")
            for user in response.users:
                print(f"ID: {user.user_id}, Name: {user.name}, Email: {user.email}")
        except grpc.RpcError as e:
            print(f"Error listing users: {e.details()}")

if __name__ == '__main__':
    run()