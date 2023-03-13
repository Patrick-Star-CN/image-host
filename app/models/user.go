package models

import "time"

type User struct {
	ID           int    // 登录编号
	Username     string // 登录编号
	JHPassword   string // 密码
	WechatOpenID string // 小程序 openID
	CreateTime   time.Time
}
