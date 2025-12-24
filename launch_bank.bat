@echo off
title ELITEBANK COMMAND CENTER
echo Launching 5-Language Infrastructure...

start "Node_Dashboard" cmd /k "node dashboard.js"
start "CPP_Terminal" cmd /k "receiver.exe"
start "Java_Core" cmd /k "java MyOwnPaystack"
start "Go_Watchdog" cmd /k "go run verifier.go"
start "" python auditor.py

echo ------------------------------------------
echo ALL SYSTEMS ONLINE.
pause