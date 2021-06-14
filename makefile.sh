#!/bin/sh
docker-compose build db
docker-compose build backend

docker-compose up -d db

sleep 1m
docker-compose up -d backend

docker-compose logs -f backend 
sleep 5m