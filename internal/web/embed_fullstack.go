//go:build fullstack

package web

import (
	"embed"
	"io/fs"
)

//go:embed all:dist/admin all:dist/user
var assets embed.FS

// Enabled 报告是否有嵌入式 Web 资源。
func Enabled() bool { return true }

// UserEnabled 报告用户商城 SPA 是否被嵌入。
func UserEnabled() bool { return true }

// AdminFS 返回 admin SPA 的子文件系统。
func AdminFS() fs.FS {
	sub, err := fs.Sub(assets, "dist/admin")
	if err != nil {
		panic("embed: dist/admin missing: " + err.Error())
	}
	return sub
}

// UserFS 返回 user SPA 的子文件系统。
func UserFS() fs.FS {
	sub, err := fs.Sub(assets, "dist/user")
	if err != nil {
		panic("embed: dist/user missing: " + err.Error())
	}
	return sub
}
