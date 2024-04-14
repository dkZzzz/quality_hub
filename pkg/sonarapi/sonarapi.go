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
