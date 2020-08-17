## How to use this tool

### Prerequisites 

Default behaviour of this tool is to look in .kube directory for config file and it should have permissions to read namespaces, secrets, and serviceaccounts at cluster level.
It can also get credentials of the K8 cluster you want to work with from Environment variable KUBECONFIG just like kubectl would.

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
Tested on windows powershell (not PowerShell ISE) and [Terminal](https://github.com/microsoft/terminal) app

For Mac

```
wget https://github.com/bit-cloner/kcg/releases/download/0.9/darwin-amd64-kcg
chmod +x darwin-amd64-kcg
./darwin-amd64-kcg

```