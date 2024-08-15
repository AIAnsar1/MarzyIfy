package Cri

import "regexp"

type listOptions struct {
	Id                 string
	PodID              string
	NameRegexp         string
	PodNamespaceRegexp string
	State              string
	Verbose            bool
	Labels             map[string]string
	Quiet              bool
	Output             string
	All                bool
	Latest             bool
	Last               int
	NoTrunc            bool
	Image              string
	ResolveImagePath   bool
}

type ContainerPodInfo struct {
	PodUid  string
	PodName string
	PodNs   string
}

type CRITool struct {
	Rs         InternalApi.RuntimeService
	NsFilterRx *regexp.Regexp
}

type ContainerByCreated []*Pb.Container

var defaultRuntimeEndpoints = []string{
	"unix:///proc/1/root/run/containerd/containerd.sock",
	"unix:///proc/1/root/var/run/containerd/containerd.sock",
	"unix:///proc/1/root/var/run/crio/crio.sock",
	"unix:///proc/1/root/run/crio/crio.sock",
	"unix:///proc/1/root/run/cri-dockerd.sock",
	"unix:///proc/1/root/var/run/cri-dockerd.sock",
}
