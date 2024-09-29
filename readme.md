# Alpacaproj
This is a project generator that uses templates i like/use, if oyu want to contrib just make a pull request and i will look at it.
# Install
## Linux
```bash
curl -L $(curl -s https://api.github.com/repos/JensvandeWiel/alpacaproj/releases/latest | grep "browser_download_url.*Linux_x86_64.tar.gz" | cut -d '"' -f 4) | sudo tar -xz -C /usr/local/bin
```
## Windows
```powershell
$downloadUrl = (Invoke-RestMethod https://api.github.com/repos/JensvandeWiel/alpacaproj/releases/latest).assets | Where-Object { $_.name -like "*Windows_x86_64.zip" } | Select-Object -ExpandProperty browser_download_url; $destinationPath = "C:\Program Files\alpacaproj"; if (-Not (Test-Path $destinationPath)) { New-Item -ItemType Directory -Path $destinationPath }; Invoke-WebRequest -Uri $downloadUrl -OutFile "$destinationPath\alpacaproj.zip"; Expand-Archive -Path "$destinationPath\alpacaproj.zip" -DestinationPath $destinationPath -Force; Remove-Item "$destinationPath\alpacaproj.zip"; [System.Environment]::SetEnvironmentVariable("Path", $env:Path + ";C:\Program Files\alpacaproj", [System.EnvironmentVariableTarget]::Machine)
```
### Arm
You can change the `x86_64` to `arm64` in the download url