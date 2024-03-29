# An offline K8S benchmark extractor
This is an offline K8S manifest extractor for HKUST FYP project: Security Compliance Checker tools for Cloud Environment v4

## Prerequisites
- Please ready your $Kubectl command
  - Please check if your Kubectl command should be connected to a k8s cluster. (E.G. we should get the k8s’ list of nodes by $kubectl get nodes)
- Install golang 1.17 version in your environment (we should get version name by $go version)
```bash
# Install Golang binary
wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
tar -zxvf go1.17.5.linux-amd64.tar.gz -C /usr/local/
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_profile
mkdir /go
echo "export GOPATH=/go" >> ~/.bash_profile
source ~/.bash_profile
rm -f go1.17.5.linux-amd64.tar.gz
```
- Development environment: Linux Redhat / Centos7

## Quick Start
```bash 
git clone https://github.com/alexshinningsun/K8S_benchmark_extractor.git
export APP_K8SRESULT_PATH=/tmp/K8S_benchmark_extractor
export APP_K8S_CLOUD_PLATFORM=AWS
cd <repo-directory>
go mod download
go build -o app .
./app
```
Please Upload the /tmp/K8S_benchmark_extractor/fetcher_result.tar to Security Compliance Checker tools for Cloud Environment v4 frontend

## Please Note
  * 🚩 **Keep** the tar **file name** as fetcher_result.tar
  * 🚩 **Do not change** the naming inside the fetcher_result.tar
## Program Logic :
1. User set the \<output path\>
    - `export APP_K8SRESULT_PATH= ` E.G. `export APP_K8SRESULT_PATH=/tmp/K8S_benchmark_extractor`
1. Create a folder \<output path\>/fetcher_result/
1. Ready the kubehunter and kubebench repo & yaml
    - Please ready the image: aquasec/kube-hunter in your offline env
    - Please ready the image: aquasec/kube-bench:latest in your offline env
    - Please update the image tag in ./assets/yamls/job-eks.yaml
    - Please update the image tag in ./assets/yamls/kube-hunter-job.yaml
    - Please update the cloud platform of your K8S to environment variable: 
        - `export APP_K8S_CLOUD_PLATFORM= ` E.G. `export APP_K8S_CLOUD_PLATFORM=AWS`
            - AWS
            - Azure
            - GCP
            - Other

### A. Get the kubehunter result by $kubectl apply :
1. Kubectl apply ==>  kubehunter yaml
1. Get the kubehunter's result and save to output path
1. Remove kubehunter by $kubectl

### B. Get the kubebench result by $kubectl apply :

1. Kubectl apply ==> kubebench yaml
1. Get the kubebench's result and save to output path
1. Remove kubebench by $kubectl


### C. Get and save all the K8S Manifest from k8s cluster(k8sExtractor.go) :
1. Run a bash script to get all yaml from k8s cluster!
1. Save to output path
* Backend will receive all the Manifests. Run Kube-linter and Kube-score to analysis those yamls.


### D. The output of the program is an archive file :
1. tar cf \<output path\>/fetcher_result.tar \<output path\>/fetcher_result/
* User need to import the archive file to backend through website.

