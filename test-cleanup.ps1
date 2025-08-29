# Simple test to verify PowerShell is working
Write-Host "PowerShell test script is running successfully!" -ForegroundColor Green
Write-Host "Current directory: $PWD"
Write-Host ""
Write-Host "Files in current directory:"
Get-ChildItem | Select-Object Name, Length | Format-Table