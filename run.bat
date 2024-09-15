@echo off
setlocal

set ROOT=%1
set KEY=%2

set "WORK_DIR=%~dp0"
@REM echo %WORK_DIR%
set "FILE_PATH=%SCRIPT_DIR%mr-out"

del "%FILE_PATH%"
echo. > %FILE_PATH%

go run .\main.go -path=%ROOT% -key=%KEY%

endlocal