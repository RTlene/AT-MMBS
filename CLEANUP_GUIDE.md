# Project Cleanup Guide

We have created several cleanup scripts to help you clean up unnecessary files from your project. All scripts are now in English to avoid encoding issues.

## Available Scripts:

### 1. quick-clean.bat (Recommended for beginners)
The simplest and most direct cleanup method. Just double-click to run.

```bash
quick-clean.bat
```

### 2. cleanup.bat + cleanup.ps1
Interactive cleanup with multiple options.

```bash
cleanup.bat
```

### 3. simple-cleanup.ps1
A simple PowerShell script that directly deletes unnecessary files.

```powershell
# Run directly
powershell -ExecutionPolicy Bypass -File .\simple-cleanup.ps1
```

## Steps to Use:

### Step 1: Pull Latest Code
```powershell
git pull origin cursor/explain-background-mode-concept-201c
```

### Step 2: Restore Files (if needed)
If you see deleted files in git status:
```powershell
git checkout HEAD -- .
```

### Step 3: Test PowerShell Works
```powershell
powershell -ExecutionPolicy Bypass -File .\test-cleanup.ps1
```

### Step 4: Run Cleanup

#### Option A: Use quick-clean.bat (Easiest)
Just double-click `quick-clean.bat` or run:
```bash
quick-clean.bat
```

#### Option B: Use interactive cleanup
```bash
cleanup.bat
```

#### Option C: Use PowerShell directly
```powershell
powershell -ExecutionPolicy Bypass -File .\simple-cleanup.ps1
```

## After Cleanup:

1. Verify project structure
2. Run Docker build test:
   ```bash
   test-docker-build.bat
   ```

## Files That Will Be Deleted:

- Test files: test*.html, test.db, *.py
- Build artifacts: *.exe, mmbs-modular, at-mmbs-test
- Temporary files: *.log, *_old.go.bak
- Documentation: improvement_suggestions.md, WORK_COMPLETED.md
- Database fixes: fix_database*.sql
- Config: container.config.json

## Important Notes:

- These scripts will NOT delete files in .git directory
- Source code and essential files are preserved
- Always backup your project before cleaning