#!/usr/bin/env bash
go build -o api
sudo docker rmi apitest:0.1
sudo docker build -t apitest:0.1 .