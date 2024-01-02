#!/usr/bin/env bash

#-------------------------------------------------------------------------------
# Purpose: gen grpc src code 
# Author : easytojoin@163.com (Jok)
# Date   : Tue Jan 02 16:31:08 CST 2024
#-------------------------------------------------------------------------------

protoc --go_out=. --go-grpc_out=. hello.proto
