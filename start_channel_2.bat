@echo off
set CHANNEL_NAME=Channel 2
set CHANNEL_PORT=55901
set TCP_PORT=55901
set UDP_PORT=55901
cd ..\mu-server
title Channel 2 - TCP:55901 UDP:55901 PORT:55901
go run cmd\server\main.go --port=55901
pause
