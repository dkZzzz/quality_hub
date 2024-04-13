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
	ID          int    `json:"id" gorm:"primary_key"`
	ProjectName string `json:"project_name"`
	Username    string `json:"username"`
	BranchName  string `json:"branch_name"`
	Url         string `json:"url"`
	ReportID    int    `json:"report_id"`
}

type Report struct {
	ID        int `json:"id" gorm:"primary_key"`
	ProjectID int `json:"project_id"`
}

type Advice struct {
	ID int `json:"id" gorm:"primary_key"`
}
