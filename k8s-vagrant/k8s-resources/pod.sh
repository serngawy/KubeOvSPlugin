#!/bin/bash -e

#sudo kubectl create namespace sock-shop
#sudo kubectl create -f https://raw.githubusercontent.com/microservices-demo/microservices-demo/master/deploy/kubernetes/complete-demo.yaml
#sleep 30s
pod_status=$(sudo kubectl -n sock-shop get pods | grep front-end | awk '{print $3}')
until [ $pod_status == 'Running' ]; do
    echo 'front-end pod status is ' $pod_status
    echo 'waiting for front-end  pod get created'
    sleep 10s
    pod_status=$(sudo kubectl -n sock-shop get pods | grep front-end | awk '{print $3}')
done
pod_name=$(sudo kubectl -n sock-shop get pods | grep front-end | awk '{print $1}')
node_ip=$(sudo kubectl -n sock-shop get pods $pod_name -o json | jq -r '.status|.hostIP')
port_no=30001
status_code=$(curl -s -o /dev/null -I -w %{http_code} http://$node_ip':'$port_no)
#sudo kubectl delete namespace sock-shop
echo $status_code
