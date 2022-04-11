# An offline K8S benchmark extractor
## Program Logic :
1. User set the \<output path\>
1. Create a folder \<output path\>/fetcher_result/
1. Ready the kubehunter and kubebench repo & yaml
    - Please ready the image: aquasec/kube-hunter in your offline env
    - Please ready the image: aquasec/kube-bench:latest in your offline env
    - Please update the image tag in ./assets/yamls/job-eks.yaml
    - Please update the image tag in ./assets/yamls/kube-hunter-job.yaml

### A. Get the kubehunter result by $kubectl apply :
1. Kubectl apply ==>  kubehunter yaml
1. Get the kubehunter's result and save to output path
1. Remove kubehunter by $kubectl

### B. Get the kubebench result by $kubectl apply :

1. Kubectl apply ==> kubebench yaml
1. Get the kubebench's result and save to output path
1. Remove kubebench by $kubectl


### C. Get and save all the yaml from k8s cluster :
1. Run a bash script to get all yaml from k8s cluster!
1. Save to output path
* Backend will receive all the yamls. Run Kube-linter and Kube-score to analysis those yamls.


### D. The output of the program is an archive file :
1. tar cf \<output path\>/fetcher_result.tar \<output path\>/fetcher_result/
* User need to import the archive file to backend through website.
