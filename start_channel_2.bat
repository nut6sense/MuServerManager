@echo off
set CHANNEL_NAME=Channel 2
set CHANNEL_PORT=55902
set TCP_PORT=9102
set UDP_PORT=9202
cd ..\mu-server
title Channel 2 - TCP:9102 UDP:9202 PORT:55902
go run cmd\server\main.go --port=55902
pause
