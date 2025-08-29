# 项目清理指南

我已经为您创建了两个文件来帮助清理项目：

## 1. NECESSARY_FILES.md
详细列出了项目运行所需的所有必要文件，不在列表中的文件都可以安全删除。

## 2. cleanup-project.ps1
PowerShell自动清理脚本，可以帮助您快速清理项目。

### 使用方法：

#### 步骤 1：拉取最新代码
```powershell
git pull origin cursor/explain-background-mode-concept-201c
```

#### 步骤 2：恢复所有文件（如果之前有删除状态的文件）
```powershell
git checkout HEAD -- .
```

#### 步骤 3：预览将要删除的文件（推荐先执行）
```powershell
.\cleanup-project.ps1 -DryRun
```

#### 步骤 4：执行清理
```powershell
# 基本清理（删除所有不必要的文件）
.\cleanup-project.ps1

# 保留测试文件
.\cleanup-project.ps1 -KeepTests

# 保留日志文件
.\cleanup-project.ps1 -KeepLogs

# 保留测试和日志文件
.\cleanup-project.ps1 -KeepTests -KeepLogs
```

### 清理后验证：

1. 检查项目结构是否正确
2. 运行Docker构建测试：
   ```powershell
   .\test-docker-build.bat
   ```

### 注意事项：
- 建议在清理前备份项目
- 清理脚本会跳过 .git 目录
- 如果有自定义的重要文件，请先移到安全位置