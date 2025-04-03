@echo off
set CHANNEL_NAME=Channel Main
set CHANNEL_PORT=55900
set TCP_PORT=9000
set UDP_PORT=9000
cd ..\mu-server
title Channel Main - TCP:9000 UDP:9000 PORT:55900
go run cmd\server\main.go --port=55900
pause
