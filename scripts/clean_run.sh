#!/bin/bash

cd ../infrastructure

docker-compose down -v

docker-compose up -d --build