@echo off
set CHANNEL_NAME=Channel 4
set CHANNEL_PORT=55904
set TCP_PORT=9104
set UDP_PORT=9204
cd ..\mu-server
title Channel 4 - TCP:9104 UDP:9204 PORT:55904
go run cmd\server\main.go --port=55904
pause
