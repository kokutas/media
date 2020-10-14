package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/cmd/fyne_demo/data"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

// 欢迎页面

func welcome(application fyne.App) fyne.CanvasObject {
	// 创建一个logo的图片资源，使用data.fynescene
	logo := canvas.NewImageFromResource(data.FyneScene)
	// 如果当前设备是移动端
	if fyne.CurrentDevice().IsMobile() {
		// 设置logo图片资源的最小尺寸
		logo.SetMinSize(fyne.NewSize(171, 125))
	} else {
		// 设置logo图片资源的最小尺寸
		logo.SetMinSize(fyne.NewSize(228, 167))
	}
	// 返回一个新建的盒子
	return widget.NewVBox(
		// NewSpacer返回一个可以填充垂直和水平空间的spacer对象。 这主要用于盒子布局。
		layout.NewSpacer(),
		// 设置一个带[字体加粗、水平居中的label]
		widget.NewLabelWithStyle("欢迎使用golang开发的客户端工具v2.0", fyne.TextAlignCenter, fyne.TextStyle{Bold: true}),
		// 设置logo图片资源
		widget.NewHBox(layout.NewSpacer(), logo, layout.NewSpacer()),
		// 设置 超链接的盒子
		widget.NewHBox(layout.NewSpacer(),
			widget.NewHyperlink("官网", parseURL("https://www.baidu.com/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("使用说明", parseURL("https://www.baidu.com/")),
			widget.NewLabel("-"),
			widget.NewHyperlink("联系人", parseURL("https://www.baidu.com/")),
			layout.NewSpacer(),
		),
		layout.NewSpacer(),
		// 设置分组
		widget.NewGroup("主题",
			// 设置网格布局的容器布局：2列
			fyne.NewContainerWithLayout(layout.NewGridLayout(2),
				// 设置按钮
				widget.NewButton("雅黑", func() {
					// 设置应用的主题为黑色
					application.Settings().SetTheme(theme.DarkTheme())
				}),
				widget.NewButton("高亮", func() {
					// 设置应用的主题为白色
					application.Settings().SetTheme(theme.LightTheme())
				}),
			),
		),
	)
}
