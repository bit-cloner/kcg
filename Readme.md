## What is KCG?
Kubernetes config generator is an interactive command line tool that lets you create kubeconfig files
related to a service account in a given namespace. The user can interactively chose a namespace and service account from a K8 cluster.
The output is a config file with token authentication that has same RBAC permissions assigned to chosen
service account.

### Demo
![til](./usage.gif)

### Prerequisites 

Default behaviour of this tool is to look in .kube directory for config file and it should have permissions to read namespaces, secrets, and serviceaccounts at cluster level.
It can also look for credentials of the K8 cluster you want to work with from Environment variable KUBECONFIG just like kubectl would.

### Get it

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
Install it as a kubectl plugin

```
wget https://github.com/bit-cloner/kcg/releases/download/0.9/kubectl-kcg
sudo cp kubectl-kcg /usr/local/bin
kubectl plugin list
kubectl kcg
```
<h3 align="left">Support:</h3>
<p><a href="https://www.buymeacoffee.com/welldone"> <img align="left" src="https://cdn.buymeacoffee.com/buttons/v2/default-yellow.png" height="50" width="210" alt="welldone" /></a></p><br><br>
