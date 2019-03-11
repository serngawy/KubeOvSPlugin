#!/bin/sh
cp symOvSCNI /opt/cni/bin/
SERVER=$(hostname)
var=$(kubectl describe node $(echo $SERVER | tr '[:upper:]' '[:lower:]') | grep PodCIDR | awk '{ print $2 }' | cut -d"." -f1-3)
echo "node-name: " $SERVER " Podcidr: " $var .0/24
cat << CNI > /etc/cni/net.d/sym-cni.conf
{
    "cniVersion":"0.3.0",
    "name":"sym-cni",
    "type":"symOvSCNI",
    "ovsBridge":"br-int",
    "ovsExtBridge":"br-ext",
    "ctlrPort":6653,
    "ctlrActive":true,
    "controller":"",
    "externalIntf":"",
    "externalIp":"",
    "ipam":{
        "type":"host-local",
        "subnet":"$var.0/24",
        "routes":[{
            "dst":"0.0.0.0/0"
        }],
        "gateway":"$var.1"
    }
}
CNI