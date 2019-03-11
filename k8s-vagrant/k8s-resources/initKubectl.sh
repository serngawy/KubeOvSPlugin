#!/usr/bin/env bash

yes | sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config
sudo kubectl taint nodes --all node-role.kubernetes.io/master-
#sudo kubectl create -f examples/busybox.yml
