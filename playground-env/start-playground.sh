#!/bin/bash
echo "Composing resources for go-playground..."
docker-compose up -d
echo "Starting go-playground..."
docker-compose run go
