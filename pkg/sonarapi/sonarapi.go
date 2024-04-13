package sonarapi

import (
	"github.com/dkZzzz/quality_hub/config"
	"github.com/dkZzzz/quality_hub/pkg/sentreq"
)

func CreateProject(name, projectKey string) (map[string]string, error) {
	sonarHost := config.Cfg.SonarHost
	url := sonarHost + "/api/projects/create"
	username := config.Cfg.SonarUser
	password := config.Cfg.SonarPassword

	formData := map[string]string{
		"project":    projectKey,
		"name":       name,
		"mainBranch": "main",
	}
	response, err := sentreq.FormDataReq(url, username, password, formData)
	if err != nil {
		return nil, err
	}
	return response["project"].(map[string]string), nil
}
