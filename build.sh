#!/bin/bash

# Create build folder if it doesn't exist
if [ ! -d "build" ]; then
  mkdir -p build
fi

echo "Building Maple Pie..."

go build -o build/maplepie cmd/maple-pie/main.go

chmod +x build/maplepie

echo "Starting Maple Pie..."
./build/maplepie
