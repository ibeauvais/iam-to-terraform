package main

import (
	"context"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("You have to specify project id: iamToTerraform <project-id>")
	}

	projectId := os.Args[1]

	resourceService := ResourceService{}

	policy, err := resourceService.GetIamPolicy(projectId, context.Background())
	if err != nil {
		log.Fatalf("Projects.GetIamPolicy: %v", err)
	}

	template := NewResourceTemplate("resource_template.tmpl")

	for _, binding := range policy.Bindings {
		if !strings.Contains(strings.ToLower(binding.Role), "serviceagent") {
			template.resolve(binding)

		}
	}

}
