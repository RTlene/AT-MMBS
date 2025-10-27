# AT-MMBS Simple Cleanup Script
Write-Host "AT-MMBS Simple Cleanup Script" -ForegroundColor Green
Write-Host "=============================" -ForegroundColor Green

# Files to delete
$patterns = @(
    "test*.html",
    "*.exe",
    "*.log",
    "test.db",
    "*_old.go.bak",
    "at-mmbs-test",
    "mmbs-modular",
    "comprehensive_test.py",
    "test_report.md",
    "deployment_status.txt",
    "improvement_suggestions.md",
    "WORK_COMPLETED.md",
    "fix_database*.sql",
    "container.config.json"
)

$totalDeleted = 0

Write-Host ""
Write-Host "Starting cleanup..." -ForegroundColor Yellow

foreach ($pattern in $patterns) {
    $files = Get-ChildItem -Path . -Filter $pattern -ErrorAction SilentlyContinue
    foreach ($file in $files) {
        try {
            Remove-Item $file.FullName -Force
            Write-Host "Deleted: $($file.Name)" -ForegroundColor Red
            $totalDeleted++
        }
        catch {
            Write-Host "Failed to delete: $($file.Name)" -ForegroundColor DarkRed
        }
    }
}

Write-Host ""
Write-Host "Cleanup completed. Total files deleted: $totalDeleted" -ForegroundColor Green