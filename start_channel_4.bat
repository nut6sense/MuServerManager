@echo off
set CHANNEL_NAME=Channel 4
set CHANNEL_PORT=55903
set TCP_PORT=55903
set UDP_PORT=55903
cd ..\mu-server
title Channel 4 - TCP:55903 UDP:55903 PORT:55903
go run cmd\server\main.go --port=55903
pause
