package ui

import (
	"fyne.io/fyne"
	"net/url"
)

// ui布局工具类

// 1.解析url
func parseURL(urlStr string) *url.URL {
	link, err := url.Parse(urlStr)
	if err != nil {
		fyne.LogError("Could not parse URL", err)
	}

	return link
}