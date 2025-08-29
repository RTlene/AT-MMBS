# AT-MMBS 项目清理脚本
# 此脚本将删除不必要的文件，只保留项目运行所需的文件

param(
    [switch]$DryRun = $false,
    [switch]$KeepTests = $false,
    [switch]$KeepLogs = $false
)

Write-Host "🧹 AT-MMBS 项目清理工具" -ForegroundColor Cyan
Write-Host "=========================" -ForegroundColor Cyan

if ($DryRun) {
    Write-Host "⚠️  运行在预览模式 - 不会真正删除文件" -ForegroundColor Yellow
}

# 定义要删除的文件模式
$filesToDelete = @(
    # 测试文件
    "test.html",
    "test_*.html",
    "test.db",
    "comprehensive_test.py",
    "test_report.md",
    "simple_test.sh",
    "test-deployment.sh",
    "test-modular.sh",
    
    # 构建产物
    "at-mmbs-test",
    "mmbs-modular",
    "*.exe",
    
    # 临时和文档文件
    "deployment_status.txt",
    "improvement_suggestions.md",
    "WORK_COMPLETED.md",
    "fix_database*.sql",
    "container.config.json",
    "*_old.go.bak",
    
    # 日志文件
    "*.log",
    
    # 系统文件
    ".DS_Store",
    "Thumbs.db"
)

# 如果保留测试文件
if ($KeepTests) {
    Write-Host "ℹ️  保留测试文件" -ForegroundColor Blue
    $filesToDelete = $filesToDelete | Where-Object { $_ -notmatch "test" }
}

# 如果保留日志文件
if ($KeepLogs) {
    Write-Host "ℹ️  保留日志文件" -ForegroundColor Blue
    $filesToDelete = $filesToDelete | Where-Object { $_ -ne "*.log" }
}

$deletedCount = 0
$totalSize = 0

Write-Host "`n开始扫描文件..." -ForegroundColor Green

foreach ($pattern in $filesToDelete) {
    $files = Get-ChildItem -Path . -Filter $pattern -Recurse -ErrorAction SilentlyContinue
    
    foreach ($file in $files) {
        # 跳过 .git 目录中的文件
        if ($file.FullName -match "\.git\\") {
            continue
        }
        
        $relativePath = $file.FullName.Replace($PWD.Path + "\", "")
        $fileSize = $file.Length
        $totalSize += $fileSize
        
        if ($DryRun) {
            Write-Host "  [预览] 将删除: $relativePath ($('{0:N2}' -f ($fileSize/1KB)) KB)" -ForegroundColor DarkGray
        } else {
            try {
                Remove-Item $file.FullName -Force
                Write-Host "  ✓ 已删除: $relativePath ($('{0:N2}' -f ($fileSize/1KB)) KB)" -ForegroundColor Red
                $deletedCount++
            } catch {
                Write-Host "  ✗ 删除失败: $relativePath - $_" -ForegroundColor DarkRed
            }
        }
    }
}

# 清理空目录
Write-Host "`n清理空目录..." -ForegroundColor Green
$emptyDirs = Get-ChildItem -Path . -Recurse -Directory | 
    Where-Object { 
        $_.FullName -notmatch "\.git" -and 
        (Get-ChildItem $_.FullName -Force).Count -eq 0 
    } | 
    Sort-Object -Property FullName -Descending

foreach ($dir in $emptyDirs) {
    $relativePath = $dir.FullName.Replace($PWD.Path + "\", "")
    if ($DryRun) {
        Write-Host "  [预览] 将删除空目录: $relativePath" -ForegroundColor DarkGray
    } else {
        try {
            Remove-Item $dir.FullName -Force
            Write-Host "  ✓ 已删除空目录: $relativePath" -ForegroundColor Red
        } catch {
            Write-Host "  ✗ 删除目录失败: $relativePath - $_" -ForegroundColor DarkRed
        }
    }
}

# 显示统计信息
Write-Host "`n📊 清理统计" -ForegroundColor Cyan
Write-Host "=========================" -ForegroundColor Cyan
if ($DryRun) {
    Write-Host "预计删除文件数: $deletedCount" -ForegroundColor Yellow
    Write-Host "预计释放空间: $('{0:N2}' -f ($totalSize/1MB)) MB" -ForegroundColor Yellow
    Write-Host "`n💡 提示: 去掉 -DryRun 参数来真正执行删除操作" -ForegroundColor Green
} else {
    Write-Host "已删除文件数: $deletedCount" -ForegroundColor Green
    Write-Host "已释放空间: $('{0:N2}' -f ($totalSize/1MB)) MB" -ForegroundColor Green
}

Write-Host "`n✅ 清理完成！" -ForegroundColor Green

# 显示剩余文件统计
$remainingFiles = Get-ChildItem -Path . -File -Recurse | Where-Object { $_.FullName -notmatch "\.git\\" }
$remainingSize = ($remainingFiles | Measure-Object -Property Length -Sum).Sum

Write-Host "`n📁 项目当前状态:" -ForegroundColor Cyan
Write-Host "剩余文件数: $($remainingFiles.Count)" -ForegroundColor White
Write-Host "项目大小: $('{0:N2}' -f ($remainingSize/1MB)) MB" -ForegroundColor White