@echo off
REM Script to update import paths after project restructuring
REM Run this from the project root

echo Updating import paths...

REM Update all .go files to use new import paths
powershell -Command "(Get-Content -Recurse -Filter '*.go' -Path .) -replace 'gitee\.com/shing1211/futuapi4go/client', 'github.com/shing1211/futuapi4go/internal/client' | Set-Content -Path _.tmp"

echo Done!
