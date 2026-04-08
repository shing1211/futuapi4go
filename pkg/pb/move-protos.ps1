# Script to move protobuf files to correct package directories
# Run from pkg/pb directory

$files = Get-ChildItem -Filter "*.pb.go" -File

foreach ($file in $files) {
    $name = $file.Name -replace '\.pb\.go$', ''
    $dir = $name.ToLower()
    
    if (Test-Path $dir -PathType Container) {
        Move-Item $file.Name -Destination "$dir\" -Force
        Write-Host "Moved: $($file.Name) -> $dir/"
    } else {
        Write-Host "Warning: Directory $dir not found for $($file.Name)"
    }
}

Write-Host "`nDone! Moved all remaining .pb.go files to their directories."
