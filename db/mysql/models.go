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
}

type Report struct {
	ID          int    `json:"id" gorm:"primary_key"`
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	IssueNum    int    `json:"issue_num"`
	HotspotNum  int    `json:"hotspot_num"`
	Duplication string `json:"duplication"`
	Coverage    string `json:"coverage"`
}

type Issue struct {
	ID          int    `json:"id" gorm:"primary_key"`
	ProjectName string `json:"project_name"`
	Type        string `json:"type"`
	File        string `json:"file"`
	StartLine   int    `json:"start_line"`
	EndLine     int    `json:"end_line"`
	StartOffset int    `json:"start_offset"`
	EndOffset   int    `json:"end_offset"`
	Message     string `json:"message"`
}

type Advice struct {
	ID          int    `json:"id" gorm:"primary_key"`
	IssueID     int    `json:"issue_id"`
	ProjectName string `json:"project_name"`
	Advice      string `json:"advice"`
}
