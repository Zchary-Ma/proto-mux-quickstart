#!/bin/bash

protoc -I=proto --go_out=app/schema \
  --go_opt=paths=source_relative \
  proto/*.proto