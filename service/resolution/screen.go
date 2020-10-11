package resolution

import "syscall"

// 屏幕分辨率
// 获取Windows分辨率
func GetWindowsResolution() (width int,height int) {
	return getWindowsSystemMetrics(0),getWindowsSystemMetrics(1)
}
func getWindowsSystemMetrics(index int)int  {
	ret, _, _ := syscall.NewLazyDLL(`User32.dll`).NewProc(`GetSystemMetrics`).Call(uintptr(index))
	return int(ret)
}