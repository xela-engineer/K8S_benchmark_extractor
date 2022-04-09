# An offline K8S benchmark extractor
## Program Logic :
1. User set the \<output path\>
1. Create a folder \<output path\>/result/
1. Ready the kubehunter and kubebench repo & yaml

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
1. tar cf \<output path\>/result.tar \<output path\>/result/
* User need to import the archive file to backend through website.
