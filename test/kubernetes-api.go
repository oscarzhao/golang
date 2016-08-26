package main

import (
	"k8s.io/kubernetes/pkg/api"
	k8sclientapi "k8s.io/kubernetes/pkg/client/unversioned"
)

func main() {
	// init k8s api server client
	if k8sclient, err := k8sclientapi.New(&k8sclientapi.Config{
		Host:        c.K8sServerConfig.ApiServer,
		BearerToken: c.K8sServerConfig.BearerToken,
		Version:     c.K8sServerConfig.ApiVersion,
		Insecure:    true,
	}); err != nil {
		glog.Errorf("init k8s api client failed, error info:%s\n", err)
	} else {
		c.K8sclient = k8sclient
	}
}
