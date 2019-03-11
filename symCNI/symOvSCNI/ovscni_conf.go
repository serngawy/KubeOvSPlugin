/*
 * Copyright (c) 2018 Kontron Canada Company and others.  All rights reserved.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License v1.0 which accompanies this distribution,
 * and is available at http://www.eclipse.org/legal/epl-v10.html
 */

package main

import (
	"encoding/json"
	"fmt"
	"github.com/containernetworking/cni/pkg/types"
	"net"
)

//Example of the expected json
//{
//    "cniVersion":"0.3.0",
//    "name":"sym-cni",
//    "type":"symOvSCNI",
//    "ovsBridge":"br-int",
//    "ovsExtBridge":"br-ext",
//    "ctlrPort":6653,
//    "ctlrActive":true,
//    "controller":"192.168.33.1",
//    "externalIntf":"eth2",
//    "externalIp":"192.168.50.11",
//    "ipam":{
//        "type":"host-local",
//        "subnet":"10.11.1.0/24",
//        "routes":[{
//        "dst":"0.0.0.0/0"
//        }],
//        "gateway":"10.11.1.1"
//    }
//}

// The symkloud cni config type for OVS
type CniConf struct {
	types.NetConf
	OvsBridge    string `json:"ovsBridge"`
	OvsExtBridge string `json:"ovsExtBridge"`
	CtlrPort     int    `json:"ctlrPort"`
	CtlrActive   bool   `json:"ctlrActive"`
	Controller   net.IP `json:"controller"`
	ExternalIntf string `json:"externalIntf"`
	ExternalIp   net.IP `json:"externalIp"`
}

// K8sArgs is the CNI_ARGS used by Kubernetes
type K8sArgs struct {
	types.CommonArgs
	K8S_POD_NAME      types.UnmarshallableString
	K8S_POD_NAMESPACE types.UnmarshallableString
}

// parse cni conf
func parseCniConf(stdin []byte) (CniConf, error) {
	cniConf := CniConf{}
	err := json.Unmarshal(stdin, &cniConf)
	if err != nil {
		fmt.Errorf("failed to parse cni configurations: %v", err)
	}
	return cniConf, nil
}
