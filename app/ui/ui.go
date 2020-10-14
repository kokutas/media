package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/flopp/go-findfont"
	"github.com/go-vgo/robotgo"
	"math"
	"os"
	"strings"
)

//初始化操作实现字体支持中文
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		//fmt.Println(path)
		//楷体:simkai.ttf
		//黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		} else if strings.Contains(path, "simhei.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

// 设置上一次的操作为当前选项tab
//const preferenceCurrentTab = "currentTab"

// 客户端ui界面相关
func Run() {
	// 延迟关闭字体设置环境变量
	defer os.Unsetenv("FYNE_FONT")

	// 新建一个应用
	application := app.NewWithID("www.kokutas.com")
	// 设置应用的icon为主题中的fyneLogo
	application.SetIcon(theme.FyneLogo())
	// 设置应用默认的主题：明亮/白色主题
	application.Settings().SetTheme(theme.DarkTheme())
	// 在应用上初始化窗口，窗口title=”工具"
	window := application.NewWindow("客户端工具")

	// 通过robotgo获取屏幕的尺寸信息
	sx, sy := robotgo.GetScreenSize()
	// 如果当前设备是移动端
	if fyne.CurrentDevice().IsMobile() {
		// 设置窗口的尺寸为：width：%95，height：95%
		window.Resize(fyne.NewSize(int(math.Ceil(0.95*float64(sx))), int(math.Ceil(0.5*float64(sy)))))
	} else {
		// 设置窗口的尺寸为：width：%40，height：50%
		window.Resize(fyne.NewSize(int(math.Ceil(0.4*float64(sx))), int(math.Ceil(0.5*float64(sy)))))
	}
	// 设置窗口居中
	window.CenterOnScreen()

	// 创建一个tab容器
	tabs := widget.NewTabContainer(
		// 新建一个欢迎的tab
		widget.NewTabItemWithIcon("欢迎", theme.HomeIcon(), welcome(application)),
		// 新建一个设置的tab
		widget.NewTabItemWithIcon("设置", theme.SettingsIcon(), setting(window)),
	)
	tabs.SetTabLocation(widget.TabLocationLeading)
	// 设置上一次的操作为当前
	//tabs.SelectTabIndex(application.Preferences().Int(preferenceCurrentTab))
	window.SetContent(tabs)

	// 展示并运行窗口
	window.ShowAndRun()
	// 应用tab初始化
	// 设置上一次的操作的tab为第一个选项tab
	//application.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}
