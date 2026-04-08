@echo off
REM Regenerate all protobuf files
REM Run from project root

echo ╔════════════════════════════════════════════════════╗
echo ║   Regenerating All Protobuf Go Files               ║  
echo ╚════════════════════════════════════════════════════╝
echo.

set PROTO_DIR=api\proto
set PB_DIR=pkg\pb

echo Cleaning old files...
if exist %PB_DIR% (
    del /S /Q %PB_DIR%\*.pb.go >nul 2>&1
    for /d %%d in (%PB_DIR%\*) do @rd /s /q "%%d" 2>nul
)
mkdir %PB_DIR% >nul 2>&1
echo ✓ Clean complete
echo.

echo Generating protobuf files...
echo.

echo [1/4] Common files...
cd %PROTO_DIR%
protoc --go_out=../../%PB_DIR% --go_opt=paths=source_relative --proto_path=. Common.proto
protoc --go_out=../../%PB_DIR% --go_opt=paths=source_relative --proto_path=. Qot_Common.proto
protoc --go_out=../../%PB_DIR% --go_opt=paths=source_relative --proto_path=. Trd_Common.proto
echo ✓ Common files generated
echo.

echo [2/4] System files...
for %%f in (InitConnect KeepAlive GetGlobalState GetUserInfo GetDelayStatistics Verification Notify UsedQuota) do (
    protoc --go_out=../../%PB_DIR% --go_opt=paths=source_relative --proto_path=. %%f.proto 2>nul
)
echo ✓ System files generated
echo.

echo [3/4] Qot files...
for %%f in (Qot_*.proto) do (
    protoc --go_out=../../%PB_DIR% --go_opt=paths=source_relative --proto_path=. %%f 2>nul
)
echo ✓ Qot files generated
echo.

echo [4/4] Trd files...
for %%f in (Trd_*.proto) do (
    protoc --go_out=../../%PB_DIR% --go_opt=paths=source_relative --proto_path=. %%f 2>nul
)
echo ✓ Trd files generated
echo.

cd ..\..

echo ╔════════════════════════════════════════════════════╗
echo ║   ✓ Protobuf generation complete!                  ║
echo ╚════════════════════════════════════════════════════╝
