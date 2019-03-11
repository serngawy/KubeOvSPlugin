#!/usr/bin/env bash

br=$1

sudo systemctl stop kubelet
sudo rm -rf /var/lib/cni/
sudo rm -rf /var/lib/kubelet/*
sudo rm -rf /etc/cni/net.d/*
sudo ovs-vsctl del-br $br
sudo ovs-vsctl del-manager
