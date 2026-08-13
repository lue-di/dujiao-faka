//go:build adminstack && !fullstack

package web

import (
	"embed"
	"io/fs"
)

// adminstack 用于「用户商城由 Vercel 等静态托管平台部署，服务器只承载 API
// 和管理后台」的场景，刻意不嵌入 user SPA，避免根路径错误地回退成后台服务。
//
//go:embed all:dist/admin
var adminAssets embed.FS

// Enabled 报告是否有嵌入式 Web 资源。
func Enabled() bool { return true }

// UserEnabled 报告用户商城 SPA 是否被嵌入。
func UserEnabled() bool { return false }

// AdminFS 返回 admin SPA 的子文件系统。
func AdminFS() fs.FS {
	sub, err := fs.Sub(adminAssets, "dist/admin")
	if err != nil {
		panic("embed: dist/admin missing: " + err.Error())
	}
	return sub
}

// UserFS 在 adminstack 模式中没有用户商城资源。
func UserFS() fs.FS { return nil }
