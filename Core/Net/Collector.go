package Net

import (
	"context"
	"debug/elf"
)

func MarzyNewEbpfCollector(ParentCtx context.Context, Ct *Cri.CRITool) *EbpfCollector {

}
func MarzyFindSSLExecutablesByPid(procfs string, pid uint32) (map[string]*SslLib, error) {

}
func MarzyListenDebugMsgs() {

}
func MarzyGetReturnOffsets(Machine elf.Machine, instructions []byte) []int {

}
func MarzyToBytes(S string) []byte {

}
func MarzyToString(B []byte) string {

}

func (e *EbpfCollector) MarzyDone() chan struct{} {

}

func (e *EbpfCollector) MarzyEbpfEvents() chan interface{} {

}

func (e *EbpfCollector) MarzyEbpfProcEvents() chan interface{} {

}

func (e *EbpfCollector) MarzyEbpfTcpEvents() chan interface{} {

}

func (e *EbpfCollector) MarzyTlsAttachQueue() chan uint32 {

}

func (e *EbpfCollector) MarzyLoad() {

}

func (e *EbpfCollector) MarzyInit() {

}

func (e *EbpfCollector) MarzyListenEvents() {

}

func (e *EbpfCollector) MarzyClose() {

}

func (e *EbpfCollector) MarzyAttachUpRobesForEncrypted() {

}

func (e *EbpfCollector) MarzyAttachGoTlsUpRobesOnProcess(Procfs string, Pid uint32) []error {

}

func (t *EbpfCollector) MarzyAttachSslUpRobesOnProcess(Procfs string, Pid uint32) []error {

}

func (e *EbpfCollector) MarzyAttachSSlUpRobes(Pid uint32, ExecutablePath string, Version string) error {

}
