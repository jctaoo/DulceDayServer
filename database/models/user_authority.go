package models

type AuthorityLevel int

// 用户角色（权限）
const (
	// 一般用户
	AuthorityUser = iota
	// 系统最高管理人员
	AuthorityRoot
)





