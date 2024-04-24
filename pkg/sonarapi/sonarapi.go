package sonarapi

import (
	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/pkg/sentreq"
)

// sonarqube创建项目
func CreateProject(name, projectKey string) error {
	url := config.Cfg.SonarHost + "/api/projects/create"
	username := config.Cfg.SonarUser
	password := config.Cfg.SonarPassword

	formData := map[string]string{
		"project":      projectKey,
		"name":         name,
		"mainBranch":   "main",
		"creationMode": "manual",
	}
	_, err := sentreq.FormDataReq(url, username, password, formData)
	if err != nil {
		return err
	}
	return nil
}

// sonarqube创建项目token
func GenerateProjectToken(name, projectKey string) (string, error) {
	url := config.Cfg.SonarHost + "/api/user_tokens/generate"
	username := config.Cfg.SonarUser
	password := config.Cfg.SonarPassword

	formData := map[string]string{
		"name":       name,
		"type":       "PROJECT_ANALYSIS_TOKEN",
		"projectKey": projectKey,
	}
	response, err := sentreq.FormDataReq(url, username, password, formData)
	if err != nil {
		return "", err
	}

	return response["token"].(string), nil
}

// sonarqube获取项目问题
func GetIssueByProject(projectName string) (map[string]interface{}, error) {
	url := config.Cfg.SonarHost + "/api/issues/search"
	username := config.Cfg.SonarUser
	password := config.Cfg.SonarPassword

	formData := map[string]string{
		"components":              projectName,
		"s":                       "FILE_LINE",
		"impactSoftwareQualities": "MAINTAINABILITY",
		"issueStatuses":           "OPEN,CONFIRMED",
		"ps":                      "500",
		"facets":                  "cleanCodeAttributeCategories,impactSoftwareQualities,codeVariants",
		"additionalFields":        "_all",
	}

	response, err := sentreq.GET(url, username, password, formData)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// sonarqube获取项目安全热点
func GetHotspotByProject(projectName string) (map[string]interface{}, error) {
	url := config.Cfg.SonarHost + "/api/hotspots/search"
	username := config.Cfg.SonarUser
	password := config.Cfg.SonarPassword

	formData := map[string]string{
		"project": projectName,
		"ps":      "500",
		"p":       "1",
		"status":  "TO_REVIEW",
	}

	response, err := sentreq.GET(url, username, password, formData)
	if err != nil {
		return nil, err
	}
	return response, nil
}
