package resolution

import (
	"fmt"
	"testing"
)

func TestGetWindowsResolution(t *testing.T) {
	// 获取Windows操作系统的屏幕分辨率
	w, h := GetWindowsResolution()
	fmt.Printf("width:%v,height:%v\n", w, h)
}
