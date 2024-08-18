package K8S

import (
	"context"
	"time"
)

type K8sCollector struct {
	Ctx                 context.Context
	InformersFactory    informers.SharedInformerFactory
	Watchers            map[K8SResourceType]cache.SharedIndexInformer
	Stopper             chan struct{}
	DoneChan            chan struct{}
	PodInformer         v1.PodInformer
	ServiceInformer     v1.ServiceInformer
	ReplicasetInformer  appsv1.ReplicaSetInformer
	DeploymentInformer  appsv1.DeploymentInformer
	EndpointsInformer   v1.EndpointsInformer
	DaemonsetInformer   appsv1.DaemonSetInformer
	StatefulSetInformer appsv1.StatefulSetInformer
	Events              chan interface{}
}

type K8sNamespaceResources struct {
	Pods     map[string]corev1.Pod     `json:"pods"`
	Services map[string]Corev1.Service `json:"services"`
}

type K8sResourceMessage struct {
	ResourceType string      `json:"type"`
	EventType    string      `json:"eventType"`
	Object       interface{} `json:"object"`
}

type Container struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	PodUID    string `json:"pod"`
	Image     string `json:"image"`
	Ports     []struct {
		Port     int32  `json:"port"`
		Protocol string `json:"protocol"`
	} `json:"ports"`
}

type K8SResourceType string

const (
	SERVICE     = "Service"
	POD         = "Pod"
	REPLICASET  = "ReplicaSet"
	DEPLOYMENT  = "Deployment"
	ENDPOINTS   = "Endpoints"
	CONTAINER   = "Container"
	DAEMONSET   = "DaemonSet"
	STATEFULSET = "StatefulSet"
	ADD         = "Add"
	UPDATE      = "Update"
	DELETE      = "Delete"
)

var (
	K8sVersion   string
	ResyncPeriod time.Duration = 120 * time.Second
)
