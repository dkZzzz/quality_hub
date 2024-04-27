package mysql

import (
	"context"
	"errors"
)

type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

var (
	Argon2ParamVar = &Argon2Params{
		Memory:      64 * 1024,
		Iterations:  3,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	}
)

// user模块创建用户
func CreateUser(ctx context.Context, username, password, email string, argon2Params *Argon2Params) (int, error) {
	var users []User
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&users).Error
	if err != nil {
		return 0, err
	}

	if len(users) != 0 {
		return 0, errors.New("user already exists")
	}

	passWord, err := generateFromPassword(password, argon2Params)
	if err != nil {
		return 0, err
	}

	err = DB.WithContext(ctx).Create(&User{Username: username, Password: passWord, Email: email}).Error
	if err != nil {
		return 0, err
	}

	err = DB.WithContext(ctx).Where("username = ?", username).Find(&users).Error
	if err != nil {
		return 0, err
	}
	if len(users) == 0 {
		return 0, errors.New("user does not exist")
	}
	return users[0].ID, nil
}

// user模块检查用户密码
func CheckUser(ctx context.Context, username, password string) (int, error) {
	var user User
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	if err != nil {
		return 0, err
	}
	if user.ID == 0 {
		return 0, errors.New("user does not exist")
	}

	match, err := comparePasswordAndHash(password, user.Password)
	if err != nil {
		return 0, err
	}
	if !match {
		return 0, errors.New("password does not match")
	}

	return user.ID, nil
}

// user模块删除用户
func ModifyUsername(ctx context.Context, oldUsername, NewUsername string) error {
	var user User
	err := DB.WithContext(ctx).Where("username = ?", oldUsername).First(user).Update("username", NewUsername).Error
	if err != nil {
		return err
	}
	return nil
}

// user模块修改密码
func ModifyPassword(ctx context.Context, username, newPassword string) error {
	var user User
	password, err := generateFromPassword(newPassword, Argon2ParamVar)
	if err != nil {
		return err
	}

	err = DB.WithContext(ctx).Where("username = ?", username).First(user).Update("password", password).Error
	if err != nil {
		return err
	}
	return nil
}

// user模块修改邮箱
func ModifyEmail(ctx context.Context, username, newEmail string) error {
	var user User
	err := DB.WithContext(ctx).Where("username = ?", username).First(user).Update("email", newEmail).Error
	if err != nil {
		return err
	}
	return nil
}

// sonarqube模块创建项目
func CreateProject(ctx context.Context, username, projectName, branchName, url, token string) (Project, error) {
	var user User
	var project Project
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&user).Error
	if err != nil {
		return project, err
	}
	if user.ID == 0 {
		return Project{}, errors.New("user does not exist")
	}

	err = DB.WithContext(ctx).Where("project_name = ?", projectName).Find(&project).Error
	if err != nil {
		return project, err
	}
	if project.ID != 0 {
		return project, errors.New("project already exists")
	}

	err = DB.WithContext(ctx).Create(&Project{ProjectName: projectName, Username: username, BranchName: branchName, Url: url}).Error
	if err != nil {
		return project, err
	}
	return project, nil
}

// sonarqube模块创建issue
func CreateIssue(ctx context.Context, ProjectName, Type, File string, StartLine, EndLine, StartOffset, EndOffset int, Message string) error {
	err := DB.WithContext(ctx).Create(&Issue{ProjectName: ProjectName, Type: Type, File: File, StartLine: StartLine, EndLine: EndLine, StartOffset: StartOffset, EndOffset: EndOffset, Message: Message}).Error
	if err != nil {
		return err
	}
	return nil
}

// sonarqube模块创建report
func CreateReport(ctx context.Context, Username, ProjectName string, issueCnt, hotspotCnt int, duplication, coverage string) error {
	err := DB.WithContext(ctx).Create(&Report{Username: Username, ProjectName: ProjectName, IssueNum: issueCnt, HotspotNum: hotspotCnt, Duplication: duplication, Coverage: coverage}).Error
	if err != nil {
		return err
	}
	return nil
}

// sonarqube模块获取项目列表
func GetProjectList(ctx context.Context, username string) ([]Project, error) {
	var projects []Project
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&projects).Error
	if err != nil {
		return nil, err
	}
	return projects, nil
}

// sonarqube模块获取项目
func GetProject(ctx context.Context, projectName string) (*Project, error) {
	var project Project
	err := DB.WithContext(ctx).Where("project_name = ?", projectName).Find(&project).Error
	if err != nil {
		return nil, err
	}
	return &project, nil
}

// sonarqube模块获取报告列表
func GetReportList(ctx context.Context, username string) ([]Report, error) {
	var reports []Report
	err := DB.WithContext(ctx).Where("username = ?", username).Find(&reports).Error
	if err != nil {
		return nil, err
	}
	return reports, nil
}

// sonarqube模块获取报告
func GetReport(ctx context.Context, reportID int) (*Report, error) {
	var report Report
	err := DB.WithContext(ctx).Where("id = ?", reportID).Find(&report).Error
	if err != nil {
		return nil, err
	}
	return &report, nil
}

// sonarqube模块获取issue
func GetIssue(ctx context.Context, projectName string) ([]Issue, error) {
	var issues []Issue
	err := DB.WithContext(ctx).Where("project_name = ?", projectName).Find(&issues).Error
	if err != nil {
		return nil, err
	}
	return issues, nil
}

func GetIssueByID(ctx context.Context, issueID int) (*Issue, error) {
	var issue Issue
	err := DB.WithContext(ctx).Where("id = ?", issueID).Find(&issue).Error
	if err != nil {
		return nil, err
	}
	return &issue, nil
}

func CraeteAdvice(ctx context.Context, issueID int, projectName, advice string) (int, error) {
	var ad Advice
	err := DB.WithContext(ctx).Create(&Advice{IssueID: issueID, ProjectName: projectName, Advice: advice}).Error
	if err != nil {
		return 0, err
	}
	return ad.ID, nil
}
