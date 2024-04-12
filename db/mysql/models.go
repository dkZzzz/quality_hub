package mysql

// 表结构定义

// User 用户表
type User struct {
	ID       int    `json:"id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Project struct {
	ID int `json:"id" gorm:"primary_key"`
}

type Report struct {
	ID int `json:"id" gorm:"primary_key"`
}

type Advice struct {
	ID int `json:"id" gorm:"primary_key"`
}
