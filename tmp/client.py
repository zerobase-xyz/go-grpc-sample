from __future__ import print_function
import grpc
import dog_pb2
import dog_pb2_grpc


def run():
    channel = grpc.insecure_channel('localhost:8080')
    stub = dog_pb2_grpc.DogStub(channel)
    while True:
        response = stub.GetMyDog(dog_pb2.GetMyDogMessage(target_dog="pochi"))
        print(response.name)


if __name__ == '__main__':
    run()
