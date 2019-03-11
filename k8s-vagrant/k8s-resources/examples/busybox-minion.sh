#!/bin/bash -e

sudo kubectl label nodes k8sminion1 ndName=minion1
#sudo kubectl label nodes k8sminion2 ndName=minion2

sudo kubectl create -f examples/busybox2.yml
#sudo kubectl create -f examples/busybox3.yml
