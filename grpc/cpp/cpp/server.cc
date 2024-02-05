// Copyright 2024 JOK Inc. All Rights Reserved.
// Author: easytojoin@163.com (jok)

#include <iostream>
#include <memory>
#include <string>

#include "grpcpp/ext/proto_server_reflection_plugin.h"
#include "grpcpp/grpcpp.h"

#include "hello.grpc.pb.h"  // NOLINT

class Service final : public hello::Greeter::Service {
 public:
  ::grpc::Status SayHello(::grpc::ServerContext* context,
                          const ::hello::HelloRequest* request,
                          ::hello::HelloReply* response) {
    std::cout << "invoke from " << request->SerializeAsString() << std::endl;
    std::string prefix("Hello ");
    response->set_message(prefix + request->name());
    return ::grpc::Status::OK;
  }
};

int main(int argc, char** argv) {
  Service service;
  grpc::EnableDefaultHealthCheckService(true);
  grpc::reflection::InitProtoReflectionServerBuilderPlugin();
  grpc::ServerBuilder builder;
  builder.AddListeningPort("localhost:8080",
                           grpc::InsecureServerCredentials());
  builder.RegisterService(&service);
  std::unique_ptr<grpc::Server> server(builder.BuildAndStart());
  std::cout << "Server listening on localhost:8080\n";
  server->Wait();
  std::cout << "exit\n";
  return 0;
}
