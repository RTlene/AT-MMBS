@echo off
echo [NOTICE] This script has been replaced due to encoding issues.
echo Please use one of the following scripts instead:
echo.
echo   - quick-deploy.bat    (Recommended - one-click deployment)
echo   - deploy-docker.bat   (Full deployment with options)
echo   - build-docker.bat    (Test build only)
echo.
echo Redirecting to quick-deploy.bat in 5 seconds...
timeout /t 5 /nobreak >nul
call quick-deploy.bat
exit /b %errorlevel%