#!/bin/bash
cd /home/ec2-user/aws-go-demo
go build -o app
./app &
