package ui

import (
	"fyne.io/fyne"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// 设置界面
func setting(window fyne.Window) fyne.CanvasObject {
	// 返回一个新建的盒子
	return widget.NewVBox(
		// NewSpacer返回一个可以填充垂直和水平空间的spacer对象。 这主要用于盒子布局。
		//layout.NewSpacer(),
		// 设置分组
		widget.NewGroup("窗口",
			// 设置网格布局的容器布局：1列
			fyne.NewContainerWithLayout(layout.NewGridLayout(1),
				widget.NewButton("全屏/还原", func() {
					// 设置全屏
					window.SetFullScreen(!window.FullScreen())
				}),
			),
		),
		// NewSpacer返回一个可以填充垂直和水平空间的spacer对象。 这主要用于盒子布局。
		//layout.NewSpacer(),
	)
}
