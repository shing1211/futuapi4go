# Complete Protobuf Regeneration Script
# This script regenerates all Go protobuf files from .proto definitions

$ErrorActionPreference = "Stop"

Write-Host "╔════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   Regenerating All Protobuf Go Files               ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════════════════╝" -ForegroundColor Cyan
Write-Host ""

$PROTO_DIR = "api\proto"
$PB_DIR = "pkg\pb"
$PROJECT = "github.com/shing1211/futuapi4go"

# Clean old generated files
Write-Host "Step 1: Cleaning old generated files..." -ForegroundColor Yellow
if (Test-Path $PB_DIR) {
    Remove-Item "$PB_DIR\*.pb.go" -Recurse -Force -ErrorAction SilentlyContinue
    Get-ChildItem $PB_DIR -Directory | Remove-Item -Recurse -Force -ErrorAction SilentlyContinue
} else {
    New-Item -ItemType Directory -Path $PB_DIR -Force | Out-Null
}
Write-Host "✓ Clean complete" -ForegroundColor Green

# Define file groups
$commonFiles = @("Common.proto", "Qot_Common.proto", "Trd_Common.proto")
$systemFiles = @("InitConnect.proto", "KeepAlive.proto", "GetGlobalState.proto",
                 "GetUserInfo.proto", "GetDelayStatistics.proto", "Verification.proto",
                 "Notify.proto", "UsedQuota.proto")
$qotFiles = Get-ChildItem "$PROTO_DIR\Qot_*.proto" | Select-Object -ExpandProperty Name
$trdFiles = Get-ChildItem "$PROTO_DIR\Trd_*.proto" | Select-Object -ExpandProperty Name

$allFiles = @($commonFiles) + @($systemFiles) + @($qotFiles) + @($trdFiles)
$totalFiles = $allFiles.Count
Write-Host "Step 2: Generating $totalFiles protobuf files..." -ForegroundColor Yellow
Write-Host ""

# Generate all at once with module option for proper subdirectory structure
Write-Host "  Generating all files..." -ForegroundColor Cyan
$fileList = $allFiles -join " "
& protoc --go_out=. --go_opt=module=$PROJECT --proto_path=$PROTO_DIR $fileList 2>$null
if ($LASTEXITCODE -eq 0) {
    Write-Host "  ✓ All $totalFiles files generated successfully" -ForegroundColor Green
} else {
    Write-Host "  ✗ Generation failed" -ForegroundColor Red
    exit 1
}

Write-Host ""
Write-Host "╔════════════════════════════════════════════════════╗" -ForegroundColor Green
Write-Host "   ✓ Generated $totalFiles protobuf files" -ForegroundColor Green
Write-Host "╚════════════════════════════════════════════════════╝" -ForegroundColor Green
