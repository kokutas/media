package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/cmd/fyne_demo/screens"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/flopp/go-findfont"
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

const preferenceCurrentTab = "currentTab"



// 客户端ui界面相关
func Run() {
	// 延迟关闭字体设置环境变量
	defer os.Unsetenv("FYNE_FONT")

	a := app.NewWithID("我的世界")
	a.SetIcon(theme.FyneLogo())

	w := a.NewWindow("Fyne Demo")

	// 创建一个tab容器
	tabs := widget.NewTabContainer(
		// 新建一个首页的tab
		widget.NewTabItemWithIcon("欢迎", theme.HomeIcon(), welcome(a)),
		widget.NewTabItemWithIcon("编辑", theme.DocumentCreateIcon(), screens.GraphicsScreen()),
		widget.NewTabItemWithIcon("Widgets", theme.CheckButtonCheckedIcon(), screens.WidgetScreen()),
		widget.NewTabItemWithIcon("Containers", theme.ViewRestoreIcon(), screens.ContainerScreen()),
		widget.NewTabItemWithIcon("Windows", theme.ViewFullScreenIcon(), screens.DialogScreen(w)))

	if !fyne.CurrentDevice().IsMobile() {
		tabs.Append(widget.NewTabItemWithIcon("高级", theme.SettingsIcon(), screens.AdvancedScreen(w)))
	}
	tabs.SetTabLocation(widget.TabLocationLeading)
	tabs.SelectTabIndex(a.Preferences().Int(preferenceCurrentTab))
	w.SetContent(tabs)

	w.ShowAndRun()
	a.Preferences().SetInt(preferenceCurrentTab, tabs.CurrentTabIndex())
}

/*
import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/go-vgo/robotgo"
	"math"
	"os"
)

var (
	// 设置全局变量窗口
	window fyne.Window
	// 设置全局变量屏幕尺寸
	sx int
	sy int
)

//初始化操作实现字体支持中文
func init() {
	//设置中文字体
	dir, err := os.UserHomeDir()
	if err != nil {
		panic(err.Error())
	}
	os.Setenv("FYNE_FONT", fmt.Sprintf("%s/%s", dir, "documents/fonts/YaHeiMonacoHybrid.ttf"))
}

// 客户端ui界面相关
func Run() {
	// 延迟关闭字体设置环境变量
	defer os.Unsetenv("FYNE_FONT")

	// 新建一个应用
	application := app.New()
	// 设置应用的icon为主题中的fyneLogo
	application.SetIcon(theme.FyneLogo())
	// 设置应用默认的主题：明亮/白色主题
	application.Settings().SetTheme(theme.DarkTheme())
	// 在应用上初始化窗口，窗口title=”工具"
	window = application.NewWindow("工具")
	// 通过robotgo获取屏幕的尺寸信息
	sx, sy = robotgo.GetScreenSize()
	// 设置窗口的尺寸为：width：%50，height：50%
	width := int(math.Ceil(0.5 * float64(sx)))
	hight := int(math.Ceil(0.5 * float64(sy)))
	window.Resize(fyne.NewSize(width, hight))
	// 设置窗口的内容
	window.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewGridLayoutWithColumns(2),
			// 设置一个按钮，并设置按钮触发事件
			widget.NewButton("Quit", func() {
				// 应用退出
				application.Quit()
			}),
		))

	// 展示并运行窗口
	window.ShowAndRun()
}
*/
