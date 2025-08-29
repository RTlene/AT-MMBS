@echo off
echo Testing Docker image build...
echo.
echo Cleaning previous build cache...
docker system prune -f

echo.
echo Building application image...
docker build --no-cache -t at-mmbs-test .

if errorlevel 1 (
    echo.
    echo [ERROR] Build failed
    echo Please check the error messages above
    echo.
    echo Common issues:
    echo - Go version mismatch: Already updated to Go 1.24
    echo - Main function duplication: Already fixed
    echo - Dependency download failure: Check network connection
) else (
    echo.
    echo [SUCCESS] Build completed!
    echo Image name: at-mmbs-test
    echo.
    echo You can now run deploy-docker.bat to deploy the application
)

pause