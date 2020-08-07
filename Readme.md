## How to use this tool

For Linux

```
wget https://github.com/bit-cloner/kcg/releases/download/0.9/linux-amd64-kcg
chmod +x linux-amd64-kcg
./linux-amd64-kcg
```
One liner

```
wget https://github.com/bit-cloner/kcg/releases/download/0.9/linux-amd64-kcg && sudo chmod +x ./linux-amd64-kcg && ./linux-amd64-kcg
```

For Windows 

```
$url = "https://github.com/bit-cloner/kcg/releases/download/0.9/windows-386-kcg.exe"
$output = "kcg.exe"
Invoke-WebRequest -Uri $url -OutFile $output -UseBasicParsing
./kcg
```
Tested on windows powershell (not PowerShell ISE)
