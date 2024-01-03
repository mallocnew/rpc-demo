cmake_minimum_required(VERSION 3.27)
set(CMAKE_TOOLCHAIN_FILE ${VCPKG_ROOT}/vcpkg/scripts/buildsystems/vcpkg.cmake)
message(STATUS "Using toolchain file: ${CMAKE_TOOLCHAIN_FILE}")

project("grpc demo")

set(CMAKE_CXX_STANDARD 23)
set(CMAKE_EXPORT_COMPILE_COMMANDS ON)
set(CMAKE_VERBOSE_MAKEFILE ON)

find_package(gRPC CONFIG REQUIRED)

add_library(grpc_proto ${CMAKE_CURRENT_SOURCE_DIR}/hello.grpc.pb.cc
                       ${CMAKE_CURRENT_SOURCE_DIR}/hello.pb.cc)

target_link_libraries(
  grpc_proto PUBLIC gRPC::gpr gRPC::grpc gRPC::grpc++ gRPC::grpc++_alts
                    gRPC::grpc++_reflection)

add_executable(client ${CMAKE_CURRENT_SOURCE_DIR}/client.cc)
add_executable(server ${CMAKE_CURRENT_SOURCE_DIR}/server.cc)
target_link_libraries(client PRIVATE grpc_proto)
target_link_libraries(server PRIVATE grpc_proto)
