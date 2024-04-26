package handlers

// user模块参数
type LoginParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RegisterParam struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type LogoutParam struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type ModifyUsernameParam struct {
	Username    string `json:"username"`
	NewUsername string `json:"new_username"`
	Token       string `json:"token"`
}

type ModifyEmailParam struct {
	Username string `json:"username"`
	NewEmail string `json:"new_email"`
	Token    string `json:"token"`
}

type ModifyPasswordParam struct {
	Username    string `json:"username"`
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
	Token       string `json:"token"`
}

// sonarqube模块参数
type CreateProjectParam struct {
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	BranchName  string `json:"branch_name"`
	Url         string `json:"url"`
	Token       string `json:"token"`
}

type GetProjectParam struct {
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	Token       string `json:"token"`
}

type GetProjectListParam struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetReportParam struct {
	Username string `json:"username"`
	ReportID int    `json:"report_id"`
	Token    string `json:"token"`
}

type GetReportListParam struct {
	Username string `json:"username"`
	Token    string `json:"token"`
}

type GetIssueParam struct {
	Username    string `json:"username"`
	ProjectName string `json:"project_name"`
	Token       string `json:"token"`
}
