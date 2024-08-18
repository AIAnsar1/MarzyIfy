package K8S

import "context"

func MarzyNewK8sCollector(ParentCtx context.Context) (*K8sCollector, error) {

}

func (this *K8sCollector) MarzyInit(Events chan interface{}) error {

}

func (this *K8sCollector) MarzyDone() <-chan struct{} {

}

func (this *K8sCollector) MarzyGetK8sVersion() string {

}

func (this *K8sCollector) MarzyClose() {

}
