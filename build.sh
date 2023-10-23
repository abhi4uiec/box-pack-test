#!/bin/bash

echo "Build the binary"
env GOOS=linux GOARCH=amd64 go build -o main

echo "Create a ZIP file"
zip box-deployment.zip main