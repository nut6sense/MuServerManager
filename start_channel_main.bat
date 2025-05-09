@echo off
set CHANNEL_NAME=Channel Main
set CHANNEL_PORT=44405
set TCP_PORT=44405
set UDP_PORT=44405
cd ..\mu-server
title Channel Main - TCP:44405 UDP:44405 PORT:44405
go run cmd\server\main.go --port=44405
pause
