package main

import (
	"context"
	"google.golang.org/api/cloudresourcemanager/v1"
	"google.golang.org/api/option"
	"log"
)

type ResourceService struct {
}

func (receiver ResourceService) GetIamPolicy(projectId string, ctx context.Context) (*cloudresourcemanager.Policy, error) {

	resourceManager, err := cloudresourcemanager.NewService(ctx, option.WithScopes(cloudresourcemanager.CloudPlatformReadOnlyScope))
	if err != nil {
		log.Fatalf("can't load configurations %v", err)
	}
	request := &cloudresourcemanager.GetIamPolicyRequest{
		Options: &cloudresourcemanager.GetPolicyOptions{
			RequestedPolicyVersion: 3,
		},
	}

	policy, err := resourceManager.Projects.GetIamPolicy(projectId, request).Do()

	if err != nil {
		return nil, err
	}
	return policy, err
}
