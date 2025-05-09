@echo off
set CHANNEL_NAME=Channel 1
set CHANNEL_PORT=55900
set TCP_PORT=55900
set UDP_PORT=55900
cd ..\mu-server
title Channel 1 - TCP:55900 UDP:55900 PORT:55900
go run cmd\server\main.go --port=55900
pause
