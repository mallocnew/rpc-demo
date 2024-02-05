// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

#include <iostream>
#include <memory>
#include <string>

#include "grpcpp/grpcpp.h"

#include "hello.grpc.pb.h"  // NOLINT

class Client {
 public:
  Client() {
    stub_ = hello::Greeter::NewStub(grpc::CreateChannel(
        "localhost:8080", grpc::InsecureChannelCredentials()));
  }

  std::string SayHello(const std::string& name) {
    hello::HelloRequest request;
    request.set_name(name);
    hello::HelloReply reply;
    grpc::ClientContext context;
    grpc::Status status = stub_->SayHello(&context, request, &reply);
    if (status.ok()) {
      return reply.message();
    } else {
      std::cerr << status.error_code() << ": " << status.error_message()
                << std::endl;
      return "gRPC failed!";
    }
  }

 private:
  std::unique_ptr<hello::Greeter::Stub> stub_;
};

int main(int argc, char** argv) {
  Client client;
  std::cout << client.SayHello("Joker");
  return 0;
}
