/*
 * Copyright (c) 2018 Kontron Canada Company and others.  All rights reserved.
 *
 * This program and the accompanying materials are made available under the
 * terms of the Eclipse Public License v1.0 which accompanies this distribution,
 * and is available at http://www.eclipse.org/legal/epl-v10.html
 */

package main

import (
	log "github.com/Sirupsen/logrus"
	"fmt"
	"time"
	"bitbucket.org/kci1/coe/symKubeProxy/watchers"
	"bitbucket.org/kci1/coe/symKubeProxy/utils"
	"bitbucket.org/kci1/coe/symKubeProxy/ovs_ctrl"
	"github.com/serngawy/libOpenflow/ofctrl"
	"net"
	"os"
)

const (
	syncTime = 5 * time.Minute
	cniFile = "/etc/cni/net.d/ovs-cni.conf"
)

func main() {

	k8s_client := utils.GetClientSetlocal()
	endpntWatcher, err := watchers.StartEndpointsWatcher(k8s_client, syncTime, "", nil)
	if err != nil {
		log.Error("Endpoint watcher didn't start %v", err)
	}
	srvWatcher, err := watchers.StartServiceWatcher(k8s_client, syncTime, "")
	if err != nil {
		log.Error("Services watcher didn't start %v", err)
	}
	nodeWatcher, err := watchers.StartNodeWatcher(k8s_client, syncTime, nil)
	if err != nil {
		log.Error("Node watcher didn't start %v", err)
	}
	podWatcher, err := watchers.StartPodWatcher(k8s_client,syncTime)
	if err != nil {
		log.Error("Pod watcher didn't start %v", err)
	}
	hostName, err := utils.GetHostName()
	if err !=nil {
		log.Error("Cannot get hostname ", err)
	}
	ndIP, ndList, err := utils.GetHostNodeIP(k8s_client, hostName)
	if err !=nil {
		log.Error("Cannot get host ip ", err)
	}
	log.Println("connecting to Host Name & IP-Address ", hostName, ndIP)

	var confFile string
	if len(os.Args) > 1 && os.Args[1] != "" {
		confFile = os.Args[1]
	} else {
		confFile = cniFile
	}
	log.Debug("Reading symkubeproxy config file at %s", confFile)

	kubeconf := utils.ReadKubeConf(confFile)
	ctrl := ovs_ctrl.NewOvsController(hostName, net.ParseIP(ndIP), kubeconf.OvsBridge, kubeconf.CtlrPort)
	// Start ofctrl
	ctrler := ofctrl.NewController(ctrl)
	podWatcher.RegisterHandler(ctrl)
	endpntWatcher.RegisterHandler(ctrl)
	srvWatcher.RegisterHandler(ctrl)
	nodeWatcher.RegisterHandler(ctrl)
	ctrl.PopulateNodes(ndList)
	ctrler.Listen(fmt.Sprintf(":%d", kubeconf.CtlrPort))
}
