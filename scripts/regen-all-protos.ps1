# Complete Protobuf Regeneration Script
# This script regenerates all Go protobuf files from .proto definitions

$ErrorActionPreference = "Stop"

Write-Host "╔════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   Regenerating All Protobuf Go Files               ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════════════════╝" -ForegroundColor Cyan
Write-Host ""

$PROTO_DIR = "api\proto"
$PB_DIR = "pkg\pb"
$PROJECT = "gitee.com/shing1211/futuapi4go"

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

$totalFiles = $commonFiles.Count + $systemFiles.Count + $qotFiles.Count + $trdFiles.Count
Write-Host "Step 2: Generating $totalFiles protobuf files..." -ForegroundColor Yellow
Write-Host ""

$generated = 0

# Generate common files
Write-Host "  [Common Files]" -ForegroundColor Cyan
foreach ($file in $commonFiles) {
    Write-Host "    Generating $file..." -NoNewline
    & protoc --go_out=$PB_DIR --go_opt=paths=source_relative --proto_path=$PROTO_DIR $file 2>$null
    if ($LASTEXITCODE -eq 0) { 
        Write-Host " ✓" -ForegroundColor Green
        $generated++
    } else { 
        Write-Host " ✗" -ForegroundColor Red 
    }
}

# Generate system files  
Write-Host "  [System Files]" -ForegroundColor Cyan
foreach ($file in $systemFiles) {
    Write-Host "    Generating $file..." -NoNewline
    & protoc --go_out=$PB_DIR --go_opt=paths=source_relative --proto_path=$PROTO_DIR $file 2>$null
    if ($LASTEXITCODE -eq 0) { 
        Write-Host " ✓" -ForegroundColor Green
        $generated++
    } else { 
        Write-Host " ✗" -ForegroundColor Red 
    }
}

# Generate Qot files
Write-Host "  [Qot Files - $($qotFiles.Count) files]" -ForegroundColor Cyan
foreach ($file in $qotFiles) {
    Write-Host "    Generating $file..." -NoNewline
    & protoc --go_out=$PB_DIR --go_opt=paths=source_relative --proto_path=$PROTO_DIR $file 2>$null
    if ($LASTEXITCODE -eq 0) { 
        Write-Host " ✓" -ForegroundColor Green
        $generated++
    } else { 
        Write-Host " ✗" -ForegroundColor Red 
    }
}

# Generate Trd files
Write-Host "  [Trd Files - $($trdFiles.Count) files]" -ForegroundColor Cyan
foreach ($file in $trdFiles) {
    Write-Host "    Generating $file..." -NoNewline
    & protoc --go_out=$PB_DIR --go_opt=paths=source_relative --proto_path=$PROTO_DIR $file 2>$null
    if ($LASTEXITCODE -eq 0) { 
        Write-Host " ✓" -ForegroundColor Green
        $generated++
    } else { 
        Write-Host " ✗" -ForegroundColor Red 
    }
}

Write-Host ""
Write-Host "╔════════════════════════════════════════════════════╗" -ForegroundColor Green
Write-Host "   ✓ Generated $generated / $totalFiles protobuf files" -ForegroundColor Green
Write-Host "╚════════════════════════════════════════════════════╝" -ForegroundColor Green
