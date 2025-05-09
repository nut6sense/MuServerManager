@echo off
set DELAY_MS=300

start start_channel_main.bat
powershell -Command "Start-Sleep -Milliseconds %DELAY_MS%"

start start_channel_1.bat
powershell -Command "Start-Sleep -Milliseconds %DELAY_MS%"

start start_channel_2.bat
powershell -Command "Start-Sleep -Milliseconds %DELAY_MS%"

start start_channel_3.bat
powershell -Command "Start-Sleep -Milliseconds %DELAY_MS%"

start start_channel_4.bat
powershell -Command "Start-Sleep -Milliseconds %DELAY_MS%"

echo All channels started.