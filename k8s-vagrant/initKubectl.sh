#!/usr/bin/env bash

## setup the .kube/config for kubectl.
yes | sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
##  untaint the k8s master node constrain and create busybox pod.
#sudo kubectl taint nodes --all node-role.kubernetes.io/master-
#sudo kubectl create -f examples/busybox.yaml
