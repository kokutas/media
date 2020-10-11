package ui

// 布局相关
import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/flopp/go-findfont"
	"math"
	"media/service/resolution"
	"os"
	"strings"
)

//初始化操作实现字体支持中文
func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		if strings.Contains(path, "SourceHanSerifCN-Bold.ttf") {
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

// 设置全局变量窗口
var (
	window fyne.Window
)

func Application() {
	// 延迟关闭字体设置环境变量
	defer os.Unsetenv("FYNE_FONT")
	// 新建一个应用
	application := app.New()
	application.SetIcon(theme.FyneLogo())

	// 设置应用默认的主题：明亮/白色主题
	application.Settings().SetTheme(theme.DarkTheme())
	// 在应用上初始化窗口，窗口title=”工具"
	window = application.NewWindow("工具")
	// 设置窗口的尺寸为：width：%50，height：50%
	width, hight := resolution.GetWindowsResolution()
	w := int(math.Ceil(0.5 * float64(width)))
	h := int(math.Ceil(0.5 * float64(hight)))
	window.Resize(fyne.NewSize(w, h))

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
