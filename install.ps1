# iplookup Windows installation script
# https://github.com/gofika/iplookup

$ErrorActionPreference = 'Stop'
$ProgressPreference = 'SilentlyContinue'

$REPO = "gofika/iplookup"
$INSTALL_DIR = "$env:LOCALAPPDATA\Programs\iplookup"
$BINARY_NAME = "iplookup.exe"

function Get-LatestVersion {
    $release = Invoke-RestMethod -Uri "https://api.github.com/repos/$REPO/releases/latest"
    return $release.tag_name
}

function Get-Architecture {
    $arch = [System.Environment]::GetEnvironmentVariable("PROCESSOR_ARCHITECTURE")
    switch ($arch) {
        "AMD64" { return "x86_64" }
        "x86" { return "i386" }
        default { 
            Write-Error "Unsupported architecture: $arch"
            exit 1
        }
    }
}

function Install-IpLookup {
    Write-Host "=== iplookup Windows Installation Script ===" -ForegroundColor Cyan
    Write-Host ""
    
    # Get version and architecture
    $version = Get-LatestVersion
    $arch = Get-Architecture
    
    Write-Host "Installing iplookup $version for windows/$arch..." -ForegroundColor Green
    
    # Build download URL
    $filename = "iplookup_$($version.TrimStart('v'))_windows_$arch.zip"
    $downloadUrl = "https://github.com/$REPO/releases/download/$version/$filename"
    
    # Create installation directory
    if (!(Test-Path $INSTALL_DIR)) {
        New-Item -ItemType Directory -Path $INSTALL_DIR -Force | Out-Null
    }
    
    # Download file
    $tempFile = Join-Path $env:TEMP $filename
    Write-Host "Downloading $downloadUrl..." -ForegroundColor Yellow
    Invoke-WebRequest -Uri $downloadUrl -OutFile $tempFile
    
    # Extract file
    Write-Host "Extracting..." -ForegroundColor Yellow
    Expand-Archive -Path $tempFile -DestinationPath $INSTALL_DIR -Force
    
    # Clean up temporary file
    Remove-Item $tempFile -Force
    
    # Add to PATH
    $currentPath = [System.Environment]::GetEnvironmentVariable("Path", "User")
    if ($currentPath -notlike "*$INSTALL_DIR*") {
        Write-Host "Adding to PATH..." -ForegroundColor Yellow
        [System.Environment]::SetEnvironmentVariable(
            "Path",
            "$currentPath;$INSTALL_DIR",
            "User"
        )
        $env:Path = "$env:Path;$INSTALL_DIR"
    }
    
    Write-Host ""
    Write-Host "iplookup has been successfully installed to $INSTALL_DIR\$BINARY_NAME" -ForegroundColor Green
    Write-Host ""
    Write-Host "Usage:" -ForegroundColor Cyan
    Write-Host "  iplookup 8.8.8.8"
    Write-Host ""
    Write-Host "Show help:" -ForegroundColor Cyan
    Write-Host "  iplookup -h"
    Write-Host ""
    Write-Host "Note: You may need to restart your terminal to use the iplookup command" -ForegroundColor Yellow
}

# Check administrator privileges (optional)
function Test-Administrator {
    $currentUser = [Security.Principal.WindowsIdentity]::GetCurrent()
    $principal = New-Object Security.Principal.WindowsPrincipal($currentUser)
    return $principal.IsInRole([Security.Principal.WindowsBuiltInRole]::Administrator)
}

# Main function
try {
    Install-IpLookup
} catch {
    Write-Error "Installation failed: $_"
    exit 1
}