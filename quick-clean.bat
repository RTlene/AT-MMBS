@echo off
echo Quick Clean - AT-MMBS Project
echo ==============================
echo.
echo This will delete unnecessary files from your project.
echo.
echo Files to be deleted:
echo - Test files (test*.html, *.py, test.db)
echo - Build artifacts (*.exe, mmbs-modular, at-mmbs-test)
echo - Temporary files (*.log, *_old.go.bak)
echo - Documentation (improvement_suggestions.md, WORK_COMPLETED.md)
echo.
set /p confirm="Are you sure you want to continue? (Y/N): "

if /i "%confirm%" neq "Y" (
    echo.
    echo Cleanup cancelled.
    pause
    exit /b 0
)

echo.
echo Starting cleanup...
echo.

REM Delete test files
echo Deleting test files...
del /q test.html 2>nul
del /q test_*.html 2>nul
del /q test.db 2>nul
del /q comprehensive_test.py 2>nul
del /q test_report.md 2>nul
del /q simple_test.sh 2>nul
del /q test-deployment.sh 2>nul
del /q test-modular.sh 2>nul

REM Delete build artifacts
echo Deleting build artifacts...
del /q at-mmbs-test 2>nul
del /q mmbs-modular 2>nul
del /q *.exe 2>nul

REM Delete temporary files
echo Deleting temporary files...
del /q deployment_status.txt 2>nul
del /q improvement_suggestions.md 2>nul
del /q WORK_COMPLETED.md 2>nul
del /q fix_database*.sql 2>nul
del /q container.config.json 2>nul
del /q *_old.go.bak 2>nul
del /q *.log 2>nul

REM Delete system files
echo Deleting system files...
del /q .DS_Store 2>nul
del /q Thumbs.db 2>nul

echo.
echo Cleanup completed!
echo.
echo Tip: Run 'test-docker-build.bat' to verify the project still builds correctly.
echo.
pause