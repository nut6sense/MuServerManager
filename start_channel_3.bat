@echo off
set CHANNEL_NAME=Channel 3
set CHANNEL_PORT=55902
set TCP_PORT=55902
set UDP_PORT=55902
cd ..\mu-server
title Channel 3 - TCP:55902 UDP:55902 PORT:55902
go run cmd\server\main.go --port=55902
pause
