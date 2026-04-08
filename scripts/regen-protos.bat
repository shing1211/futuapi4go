@echo off
REM Regenerate all protobuf Go files from proto definitions
REM Run this from project root: scripts\regen-protos.bat

echo ╔════════════════════════════════════════════════════╗
echo ║   Regenerating Protobuf Go Files                  ║
echo ╚════════════════════════════════════════════════════╝
echo.

set PROTO_DIR=api\proto
set PB_DIR=pkg\pb
set IMPORT_PATH=gitee.com/shing1211/futuapi4go/pkg/pb

echo Step 1: Creating output directories...
mkdir %PB_DIR%\common 2>nul
mkdir %PB_DIR%\qotcommon 2>nul
mkdir %PB_DIR%\trdcommon 2>nul
mkdir %PB_DIR%\initconnect 2>nul
mkdir %PB_DIR%\keepalive 2>nul
mkdir %PB_DIR%\getglobalstate 2>nul
mkdir %PB_DIR%\getuserinfo 2>nul
mkdir %PB_DIR%\getdelaystatistics 2>nul
mkdir %PB_DIR%\verification 2>nul
mkdir %PB_DIR%\notify 2>nul
mkdir %PB_DIR%\usedquota 2>nul
echo ✓ Directories created

echo.
echo Step 2: Generating protobuf files...
echo.

REM Common files (no dependencies)
echo [1/3] Generating common protos...
for %%f in (%PROTO_DIR%\Common.proto) do (
    protoc --go_out=. --go_opt=paths=source_relative ^
           --proto_path=%PROTO_DIR% ^
           --go_opt=MCommon.proto=%IMPORT_PATH%/common ^
           -o nul ^
           %%f
)

REM Qot files (depend on common)
echo [2/3] Generating Qot protos...
for %%f in (%PROTO_DIR%\Qot_*.proto %PROTO_DIR%\Get*.proto %PROTO_DIR%\Keep*.proto %PROTO_DIR%\Notify.proto %PROTO_DIR%\UsedQuota.proto) do (
    protoc --go_out=. --go_opt=paths=source_relative ^
           --proto_path=%PROTO_DIR% ^
           -o nul ^
           %%f 2>nul
)

REM Trd files (depend on common and trdcommon)  
echo [3/3] Generating Trd protos...
for %%f in (%PROTO_DIR%\Trd_*.proto %PROTO_DIR%\Verification.proto) do (
    protoc --go_out=. --go_opt=paths=source_relative ^
           --proto_path=%PROTO_DIR% ^
           -o nul ^
           %%f 2>nul
)

echo.
echo ╔════════════════════════════════════════════════════╗
echo ║   ✓ Protobuf generation complete!                 ║
echo ╚════════════════════════════════════════════════════╝
