package Temperature

func MarzyGetNvmlDriver(SoPath string) (*NvmlDriver, error) {

}
func MarzyBuildId(Id [32]int8) string {

}

func (this *NvmlDriver) Marzyinit(Opts ...nvml.LibraryOption) bool {

}

func (this *NvmlDriver) MarzyClose() {

}

func (this *NvmlDriver) MarzySystemDriverVersion() (string, error) {

}
func (this *NvmlDriver) MarzyDeviceCount() (uint, error) {

}
func (this *NvmlDriver) MarzyCudeVersion() (uint, error) {

}
func (this *NvmlDriver) MarzyDeviceInfoAndStatusByIndex(Index uint) (*DeviceInfo, *DeviceStatus, error) {

}
func (this *NvmlDriver) MarzyGetFanSpeed(Device nvml.Device, FanId int) (uint, error) {

}
func (this *NvmlDriver) MarzyDeviceInfoByIndex(Index uint) (*DeviceInfo, error) {

}
func (this *NvmlDriver) MarzyPrintAllDeviceData() {

}
