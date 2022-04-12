#!/bin/bash
objectDir=$1

#rm -r ${objectDir} 2>/dev/null
mkdir -p ${objectDir} 2>/dev/null

getYamlInNamespace() {
	namespace=$1
	objectDir=$2
	kubectl api-resources --namespaced=true 2>/dev/null | tail -n +2 | awk '{print $1}' | sort | uniq | sed 's/events//g' | sed '/^[[:space:]]*$/d' | xargs  -P 8 -I {resource} sh -c \
 			"kubectl get {resource} -n ${namespace} 2>&1 | tail -n +2 | awk '{print \$1}' | xargs  -P 8 -I {obj} sh -c \
			  \"kubectl get {resource} -n ${namespace} {obj} -o yaml > ${objectDir}/${namespace}-{resource}-{obj}.yaml \""
}

export -f getYamlInNamespace

kubectl get namespaces | tail -n +2 | awk '{print $1}' | sed 's/^\(kube-public\|kube-system\|kube-node-lease\)//g' | sed '/^[[:space:]]*$/d' | xargs  -n1 -I {} -P 8 bash -c 'getYamlInNamespace "$@"' _ {} $objectDir	  
