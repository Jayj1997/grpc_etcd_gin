#!/bin/bash
set -ue
cd `dirname $0`
protoc -I . --go_out=plugins=grpc:. proto/pb/*.proto
# protoc --grpc-gateway_out=logtostderr=true:. proto/pb/*.proto