@echo off
setlocal

echo Building Maple Pie for Windows...

REM Create build folder if it doesn't exist
if not exist build (
    mkdir build
)

REM Build the Go binary
go build -o build\maplepie.exe ./cmd/maple-pie

if %ERRORLEVEL% NEQ 0 (
    echo Build failed!
    exit /b %ERRORLEVEL%
)

echo Build completed: build\maplepie.exe

echo Starting Maple Pie...
build\maplepie.exe

pause