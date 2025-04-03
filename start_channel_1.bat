@echo off
set CHANNEL_NAME=Channel 1
set CHANNEL_PORT=55901
set TCP_PORT=9101
set UDP_PORT=9201
cd ..\mu-server
title Channel 1 - TCP:9101 UDP:9201 PORT:55901
go run cmd\server\main.go --port=55901
pause
