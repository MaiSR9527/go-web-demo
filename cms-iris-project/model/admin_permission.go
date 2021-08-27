package model

type AdminPermission struct {
	Admin      *Admin      `xorm:"extends"` //不需要映射admin结构体
	Permission *Permission `xorm:"extends"` //不需要映射permission结构体
}
