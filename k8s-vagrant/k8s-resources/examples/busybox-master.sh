#!/bin/bash -e

sudo kubectl label nodes k8smaster ndName=master1

sudo kubectl create -f examples/busybox.yml
