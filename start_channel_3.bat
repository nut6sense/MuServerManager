@echo off
set CHANNEL_NAME=Channel 3
set CHANNEL_PORT=55903
set TCP_PORT=9103
set UDP_PORT=9203
cd ..\mu-server
title Channel 3 - TCP:9103 UDP:9203 PORT:55903
go run cmd\server\main.go --port=55903
pause
