package main

import (
	"google.golang.org/api/cloudresourcemanager/v1"
	"log"
	"os"
	"strings"
	"text/template"
)

type IamResourceInput struct {
	Name      string
	Role      string
	Members   []string
	Condition *IamCondition
}

type IamCondition struct {
	Title       string
	Description string
	Expression  string
}

type ResourceTemplate struct {
	tpl       *template.Template
	projectId string
}

func (t ResourceTemplate) resolve(binding *cloudresourcemanager.Binding) {
	input := convertToInput(binding, t.projectId)
	err := t.tpl.Execute(os.Stdout, input)
	if err != nil {
		log.Fatalf("error with %+v", input)
	}
}

func NewResourceTemplate(filename string) *ResourceTemplate {
	tpl := template.Must(template.New(filename).ParseFiles(filename))
	return &ResourceTemplate{
		tpl: tpl,
	}

}

func convertToInput(binding *cloudresourcemanager.Binding, projectId string) IamResourceInput {

	var condition *IamCondition
	if binding.Condition != nil {
		condition = &IamCondition{
			Title:       binding.Condition.Title,
			Description: binding.Condition.Description,
			Expression:  strings.ReplaceAll(binding.Condition.Expression, "\"", "\\\""),
		}
	}

	return IamResourceInput{
		Name:      generateResourceName(binding.Role, projectId),
		Role:      binding.Role,
		Members:   binding.Members,
		Condition: condition,
	}
}
func generateResourceName(role string, projectId string) string {
	replacer := strings.NewReplacer(
		"roles/", "",
		".", "_",
		"projects/", "",
		projectId+"/", "", //remove project id from name when custom project role
	)
	return replacer.Replace(strings.ToLower(role))
}
