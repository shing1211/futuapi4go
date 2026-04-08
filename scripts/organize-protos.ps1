# Organize Protobuf Files into Go Packages
# Run from project root: powershell -ExecutionPolicy Bypass -Path scripts\organize-protos.ps1

$ErrorActionPreference = "Stop"
$PB_DIR = "pkg\pb"

Write-Host "╔════════════════════════════════════════════════════╗" -ForegroundColor Cyan
Write-Host "║   Organizing Protobuf into Go Packages             ║" -ForegroundColor Cyan
Write-Host "╚════════════════════════════════════════════════════╝" -ForegroundColor Cyan
Write-Host ""

# Create all package directories
Write-Host "Step 1: Creating package directories..." -ForegroundColor Yellow

$packages = @(
    "common", "qotcommon", "trdcommon",
    "initconnect", "keepalive", "getglobalstate", "getuserinfo", 
    "getdelaystatistics", "verification", "notify", "usedquota",
    "qotgetbasicqot", "qotgetkl", "qotgetorderbook", "qotgetticker",
    "qotgetrt", "qotgetbroker", "qotgetstaticinfo", "qotgetplateset",
    "qotgetplatesecurity", "qotgetownerplate", "qotgetreference",
    "qotgettradedate", "qotrequesttradedate", "qotrequesthistorykl",
    "qotrequesthistoryklquota", "qotgetsecuritysnapshot", "qotsub",
    "qotgetsubinfo", "qotregqotpush", "qotgetcapitalflow",
    "qotgetcapitaldistribution", "qotstockfilter", "qotgetoptionchain",
    "qotgetoptionexpirationdate", "qotgetwarrant", "qotgetusersecurity",
    "qotgetusersecuritygroup", "qotmodifyusersecurity", "qotgetpricereminder",
    "qotsetpricereminder", "qotrequestrehab", "qotgetsuspend",
    "qotgetcodechange", "qotgetfutureinfo", "qotgetipolist",
    "qotgetholdingchangelist", "qotgetmarketstate", "qotupdatebasicqot",
    "qotupdatekl", "qotupdateorderbook", "qotupdateticker",
    "qotupdatert", "qotupdatebroker", "qotupdatepricereminder",
    "trdgetacclist", "trdunlocktrade", "trdgetfunds", "trdgetorderfee",
    "trdgetmarginratio", "trdgetmaxtrdqtys", "trdplaceorder",
    "trdmodifyorder", "trdgetorderlist", "trdgethistoryorderlist",
    "trdgetorderfilllist", "trdgethistoryorderfilllist", "trdgetpositionlist",
    "trdflowsummary", "trdsubaccpush", "trdreconfirmorder", "trdnotify",
    "trdupdateorder", "trdupdateorderfill"
)

foreach ($pkg in $packages) {
    $path = Join-Path $PB_DIR $pkg
    if (-not (Test-Path $path)) {
        New-Item -ItemType Directory -Path $path -Force | Out-Null
    }
}

Write-Host "✓ Created $($packages.Count) package directories" -ForegroundColor Green
Write-Host ""

# Move files to correct directories
Write-Host "Step 2: Moving files to packages..." -ForegroundColor Yellow

$files = Get-ChildItem "$PB_DIR\*.pb.go" -File
$moved = 0

foreach ($file in $files) {
    $name = $file.Name -replace '\.pb\.go$', ''
    $dir = ""
    
    # Map file name to package directory
    switch -Regex ($name) {
        "^Common$"             { $dir = "common" }
        "^Qot_Common$"         { $dir = "qotcommon" }
        "^Trd_Common$"         { $dir = "trdcommon" }
        "^InitConnect$"        { $dir = "initconnect" }
        "^KeepAlive$"          { $dir = "keepalive" }
        "^GetGlobalState$"     { $dir = "getglobalstate" }
        "^GetUserInfo$"        { $dir = "getuserinfo" }
        "^GetDelayStatistics$" { $dir = "getdelaystatistics" }
        "^Verification$"       { $dir = "verification" }
        "^Notify$"             { $dir = "notify" }
        "^UsedQuota$"          { $dir = "usedquota" }
        "^Qot_GetBasicQot$"    { $dir = "qotgetbasicqot" }
        "^Qot_GetKL$"          { $dir = "qotgetkl" }
        "^Qot_GetOrderBook$"   { $dir = "qotgetorderbook" }
        "^Qot_GetTicker$"      { $dir = "qotgetticker" }
        "^Qot_GetRT$"          { $dir = "qotgetrt" }
        "^Qot_GetBroker$"      { $dir = "qotgetbroker" }
        "^Qot_GetStaticInfo$"  { $dir = "qotgetstaticinfo" }
        "^Qot_GetPlateSet$"    { $dir = "qotgetplateset" }
        "^Qot_GetPlateSecurity$" { $dir = "qotgetplatesecurity" }
        "^Qot_GetOwnerPlate$"  { $dir = "qotgetownerplate" }
        "^Qot_GetReference$"   { $dir = "qotgetreference" }
        "^Qot_GetTradeDate$"   { $dir = "qotgettradedate" }
        "^Qot_RequestTradeDate$" { $dir = "qotrequesttradedate" }
        "^Qot_RequestHistoryKL$" { $dir = "qotrequesthistorykl" }
        "^Qot_RequestHistoryKLQuota$" { $dir = "qotrequesthistoryklquota" }
        "^Qot_GetSecuritySnapshot$" { $dir = "qotgetsecuritysnapshot" }
        "^Qot_Sub$"            { $dir = "qotsub" }
        "^Qot_GetSubInfo$"     { $dir = "qotgetsubinfo" }
        "^Qot_RegQotPush$"     { $dir = "qotregqotpush" }
        "^Qot_GetCapitalFlow$" { $dir = "qotgetcapitalflow" }
        "^Qot_GetCapitalDistribution$" { $dir = "qotgetcapitaldistribution" }
        "^Qot_StockFilter$"    { $dir = "qotstockfilter" }
        "^Qot_GetOptionChain$" { $dir = "qotgetoptionchain" }
        "^Qot_GetOptionExpirationDate$" { $dir = "qotgetoptionexpirationdate" }
        "^Qot_GetWarrant$"     { $dir = "qotgetwarrant" }
        "^Qot_GetUserSecurity$" { $dir = "qotgetusersecurity" }
        "^Qot_GetUserSecurityGroup$" { $dir = "qotgetusersecuritygroup" }
        "^Qot_ModifyUserSecurity$" { $dir = "qotmodifyusersecurity" }
        "^Qot_GetPriceReminder$" { $dir = "qotgetpricereminder" }
        "^Qot_SetPriceReminder$" { $dir = "qotsetpricereminder" }
        "^Qot_RequestRehab$"   { $dir = "qotrequestrehab" }
        "^Qot_GetSuspend$"     { $dir = "qotgetsuspend" }
        "^Qot_GetCodeChange$"  { $dir = "qotgetcodechange" }
        "^Qot_GetFutureInfo$"  { $dir = "qotgetfutureinfo" }
        "^Qot_GetIpoList$"     { $dir = "qotgetipolist" }
        "^Qot_GetHoldingChangeList$" { $dir = "qotgetholdingchangelist" }
        "^Qot_GetMarketState$" { $dir = "qotgetmarketstate" }
        "^Qot_UpdateBasicQot$" { $dir = "qotupdatebasicqot" }
        "^Qot_UpdateKL$"       { $dir = "qotupdatekl" }
        "^Qot_UpdateOrderBook$" { $dir = "qotupdateorderbook" }
        "^Qot_UpdateTicker$"   { $dir = "qotupdateticker" }
        "^Qot_UpdateRT$"       { $dir = "qotupdatert" }
        "^Qot_UpdateBroker$"   { $dir = "qotupdatebroker" }
        "^Qot_UpdatePriceReminder$" { $dir = "qotupdatepricereminder" }
        "^Trd_GetAccList$"     { $dir = "trdgetacclist" }
        "^Trd_UnlockTrade$"    { $dir = "trdunlocktrade" }
        "^Trd_GetFunds$"       { $dir = "trdgetfunds" }
        "^Trd_GetOrderFee$"    { $dir = "trdgetorderfee" }
        "^Trd_GetMarginRatio$" { $dir = "trdgetmarginratio" }
        "^Trd_GetMaxTrdQtys$"  { $dir = "trdgetmaxtrdqtys" }
        "^Trd_PlaceOrder$"     { $dir = "trdplaceorder" }
        "^Trd_ModifyOrder$"    { $dir = "trdmodifyorder" }
        "^Trd_GetOrderList$"   { $dir = "trdgetorderlist" }
        "^Trd_GetHistoryOrderList$" { $dir = "trdgethistoryorderlist" }
        "^Trd_GetOrderFillList$" { $dir = "trdgetorderfilllist" }
        "^Trd_GetHistoryOrderFillList$" { $dir = "trdgethistoryorderfilllist" }
        "^Trd_GetPositionList$" { $dir = "trdgetpositionlist" }
        "^Trd_FlowSummary$"    { $dir = "trdflowsummary" }
        "^Trd_SubAccPush$"     { $dir = "trdsubaccpush" }
        "^Trd_ReconfirmOrder$" { $dir = "trdreconfirmorder" }
        "^Trd_Notify$"         { $dir = "trdnotify" }
        "^Trd_UpdateOrder$"    { $dir = "trdupdateorder" }
        "^Trd_UpdateOrderFill$" { $dir = "trdupdateorderfill" }
    }
    
    if ($dir) {
        $dest = Join-Path $PB_DIR $dir
        Move-Item $file.FullName -Destination $dest -Force
        $moved++
    } else {
        Write-Host "  ⚠ Warning: No mapping for $($file.Name)" -ForegroundColor Yellow
    }
}

Write-Host "✓ Moved $moved files" -ForegroundColor Green
Write-Host ""

Write-Host "╔════════════════════════════════════════════════════╗" -ForegroundColor Green
Write-Host "   ✓ Protobuf organization complete!" -ForegroundColor Green
Write-Host "╚════════════════════════════════════════════════════╝" -ForegroundColor Green
