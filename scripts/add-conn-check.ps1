# Script to add EnsureConnected check to all API functions
# Run from project root: powershell -ExecutionPolicy Bypass -Path scripts\add-conn-check.ps1

$ErrorActionPreference = "Stop"

# Functions that return (something, error)
$files = @(
    "pkg\qot\quote.go",
    "pkg\qot\market.go",
    "pkg\trd\trade.go",
    "pkg\sys\system.go"
)

foreach ($file in $files) {
    $content = Get-Content $file -Raw
    $original = $content
    
    # Pattern: func SomeFunc(c *futuapi.Client, ...) (ReturnType, error) {
    # Add: if err := c.EnsureConnected(); err != nil { return nil/zero, err }
    
    # Match function declarations that take *futuapi.Client
    $pattern = '(func \w+\(c \*futuapi\.Client[^)]*\) \([^)]+error\) \{)'
    
    $content = [regex]::Replace($content, $pattern, {
        param($match)
        $funcDecl = $match.Value
        # Check if it already has EnsureConnected
        if ($funcDecl -match 'EnsureConnected') {
            return $funcDecl
        }
        
        # Determine zero value for return type
        # Most return (*SomeType, error) so we use nil, err
        # Some return just error - handle those differently
        if ($funcDecl -match '\) \(error\) \{') {
            return "$funcDecl`n	if err := c.EnsureConnected(); err != nil { return err }"
        } else {
            return "$funcDecl`n	if err := c.EnsureConnected(); err != nil { return nil, err }"
        }
    }, [System.Text.RegularExpressions.RegexOptions]::Singleline)
    
    if ($content -ne $original) {
        Set-Content $file $content -NoNewline
        Write-Host "Updated: $file"
    } else {
        Write-Host "No changes: $file"
    }
}

Write-Host "`nDone! Connection checks added."
