//go:build !fullstack && !adminstack

package web

import "io/fs"

// Enabled 报告是否有嵌入式 Web 资源。
func Enabled() bool { return false }

// UserEnabled 报告用户商城 SPA 是否被嵌入。
func UserEnabled() bool { return false }

// AdminFS 默认构建模式下返回 nil（无 embed 资源）。
func AdminFS() fs.FS { return nil }

// UserFS 默认构建模式下返回 nil。
func UserFS() fs.FS { return nil }
