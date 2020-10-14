客户端工具
# 说明
使用golang开发的流媒体application客户端应用程序

# 编译
```shell script
go build -o media.exe -ldflags -H=windowsgui main.go 
```

1.使用robotgo获取屏幕大小

Golang 跨平台自动化系统，控制键盘鼠标位图和读取屏幕，窗口句柄以及全局事件监听
```http request
https://gitee.com/veni0/robotgo
```
or
```http request
https://github.com/go-vgo/robotgo
```

注意：
使用go mod vendor的会出现c的文件引入不了的异常，不建议使用go mod vendor

# fyne图标
## 1.默认图标
设置
```shell script
theme.SettingsIcon()
```

