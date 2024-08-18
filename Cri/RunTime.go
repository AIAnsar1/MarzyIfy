package Cri

import (
	"context"
	Pb "k8s.io/cri-api/pkg/apis/runtime/v1"
)

func MarzyNewCRITool(Ctx context.Context) (*CRITool, error) {

}

func (this *CRITool) MarzyFilterNamespace(Ns string) bool {

}
func (this *CRITool) MarzyFilterNamespaceWithContainerId(Id string) bool {

}
func (this *CRITool) MarzyGetAllContainers() ([]*Pb.Container, error) {

}
func (this *CRITool) MarzyGetPidsRunningOnContainers() (map[uint32]struct{}, error) {

}
func (this *CRITool) MarzyGetAllRunningProcessInsideContainer(ContainerId string) ([]uint32, error) {

}
func (this *CRITool) MarzyGetLogPath(Id string) (string, error) {

}
func (this *CRITool) MarzyContainerStatus(Id string) (*ContainerPodInfo, error) {

}
func (this *CRITool) MarzyGetContainersOfPod(PodSandBoxId string) ([]*Pb.Container, error) {

}
func (this *CRITool) MarzyGetPod(PodUid string) ([]*Pb.PodSandBox, error) {

}
