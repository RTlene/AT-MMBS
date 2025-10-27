# AT-MMBS Project Cleanup Script
# This script will remove unnecessary files while keeping essential ones

param(
    [switch]$DryRun = $false,
    [switch]$KeepTests = $false,
    [switch]$KeepLogs = $false
)

Write-Host "AT-MMBS Project Cleanup Tool" -ForegroundColor Cyan
Write-Host "=============================" -ForegroundColor Cyan

if ($DryRun) {
    Write-Host "WARNING: Running in preview mode - no files will be deleted" -ForegroundColor Yellow
}

# Define files to delete
$filesToDelete = @(
    # Test files
    "test.html",
    "test_*.html",
    "test.db",
    "comprehensive_test.py",
    "test_report.md",
    "simple_test.sh",
    "test-deployment.sh",
    "test-modular.sh",
    
    # Build artifacts
    "at-mmbs-test",
    "mmbs-modular",
    "*.exe",
    
    # Temporary and documentation files
    "deployment_status.txt",
    "improvement_suggestions.md",
    "WORK_COMPLETED.md",
    "fix_database*.sql",
    "container.config.json",
    "*_old.go.bak",
    
    # Log files
    "*.log",
    
    # System files
    ".DS_Store",
    "Thumbs.db"
)

# Filter files based on options
if ($KeepTests) {
    Write-Host "INFO: Keeping test files" -ForegroundColor Blue
    $filesToDelete = $filesToDelete | Where-Object { $_ -notmatch "test" }
}

if ($KeepLogs) {
    Write-Host "INFO: Keeping log files" -ForegroundColor Blue
    $filesToDelete = $filesToDelete | Where-Object { $_ -ne "*.log" }
}

$deletedCount = 0
$totalSize = 0

Write-Host ""
Write-Host "Scanning files..." -ForegroundColor Green

foreach ($pattern in $filesToDelete) {
    $files = Get-ChildItem -Path . -Filter $pattern -Recurse -ErrorAction SilentlyContinue
    
    foreach ($file in $files) {
        # Skip files in .git directory
        if ($file.FullName -match "\.git\\") {
            continue
        }
        
        $relativePath = $file.FullName.Replace($PWD.Path + "\", "")
        $fileSize = $file.Length
        
        if ($DryRun) {
            $totalSize += $fileSize
            $sizeKB = "{0:N2}" -f ($fileSize/1KB)
            Write-Host "  [PREVIEW] Would delete: $relativePath ($sizeKB KB)" -ForegroundColor DarkGray
        }
        else {
            try {
                Remove-Item $file.FullName -Force
                $totalSize += $fileSize
                $deletedCount++
                $sizeKB = "{0:N2}" -f ($fileSize/1KB)
                Write-Host "  [DELETED] $relativePath ($sizeKB KB)" -ForegroundColor Red
            }
            catch {
                Write-Host "  [ERROR] Failed to delete: $relativePath - $_" -ForegroundColor DarkRed
            }
        }
    }
}

# Clean empty directories
Write-Host ""
Write-Host "Cleaning empty directories..." -ForegroundColor Green
$emptyDirs = Get-ChildItem -Path . -Recurse -Directory | 
    Where-Object { 
        $_.FullName -notmatch "\.git" -and 
        (Get-ChildItem $_.FullName -Force).Count -eq 0 
    } | 
    Sort-Object -Property FullName -Descending

foreach ($dir in $emptyDirs) {
    $relativePath = $dir.FullName.Replace($PWD.Path + "\", "")
    if ($DryRun) {
        Write-Host "  [PREVIEW] Would remove empty directory: $relativePath" -ForegroundColor DarkGray
    }
    else {
        try {
            Remove-Item $dir.FullName -Force
            Write-Host "  [REMOVED] Empty directory: $relativePath" -ForegroundColor Red
        }
        catch {
            Write-Host "  [ERROR] Failed to remove directory: $relativePath - $_" -ForegroundColor DarkRed
        }
    }
}

# Display statistics
Write-Host ""
Write-Host "Cleanup Statistics" -ForegroundColor Cyan
Write-Host "==================" -ForegroundColor Cyan

if ($DryRun) {
    $sizeMB = "{0:N2}" -f ($totalSize/1MB)
    Write-Host "Files that would be deleted: $deletedCount" -ForegroundColor Yellow
    Write-Host "Space that would be freed: $sizeMB MB" -ForegroundColor Yellow
    Write-Host ""
    Write-Host "TIP: Remove -DryRun parameter to actually delete files" -ForegroundColor Green
}
else {
    $sizeMB = "{0:N2}" -f ($totalSize/1MB)
    Write-Host "Files deleted: $deletedCount" -ForegroundColor Green
    Write-Host "Space freed: $sizeMB MB" -ForegroundColor Green
}

Write-Host ""
Write-Host "Cleanup completed!" -ForegroundColor Green

# Show remaining files statistics
$remainingFiles = Get-ChildItem -Path . -File -Recurse | Where-Object { $_.FullName -notmatch "\.git\\" }
if ($remainingFiles) {
    $remainingSize = ($remainingFiles | Measure-Object -Property Length -Sum).Sum
    $remainingSizeMB = "{0:N2}" -f ($remainingSize/1MB)
    
    Write-Host ""
    Write-Host "Project Status:" -ForegroundColor Cyan
    Write-Host "Remaining files: $($remainingFiles.Count)" -ForegroundColor White
    Write-Host "Project size: $remainingSizeMB MB" -ForegroundColor White
}