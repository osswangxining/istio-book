#!/bin/bash
## kubectl port-forward svc/kiali 20001:20001 -n istio-system
for i in `seq 1000`
do
curl -I http://service1.servicedependency:5678; 
sleep 1; 
done;