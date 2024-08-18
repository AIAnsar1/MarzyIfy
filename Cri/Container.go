package Cri

import (
	InternalApi "k8s.io/cri-api/pkg/apis"
)

func MarzyGetContainerList(ContainerList []*PbContainer, Opts ListOptions) []PbOptions {

}

func MarzyListContainers(RuntimeClient InternalApi.RunTimeService, Opts listOptions) error {

}



func MarzyContainerStatus(Client InternalApi.RunTimeService, Id string Output string, TmplStr string,Quiet bool) error {

}


func MarzyMarshalContainerStatus(Cs *Pb.ContainerStatus) (string, error) {

}