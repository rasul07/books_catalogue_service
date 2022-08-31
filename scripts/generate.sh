#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/books_protos/* -type d); do
  protoc -I=${CURRENT_DIR}/books_protos --go_out=${CURRENT_DIR} \
   --go-grpc_out=${CURRENT_DIR} ${x}/*.proto
done