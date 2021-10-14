package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var rolesAndExpectedNames = []struct {
	role string
	name string
}{
	{"roles/owner", "owner"},
	{"projects/myprojectid/roles/customRole", "customrole"},
	{"roles/appengine.appAdmin", "appengine_appadmin"},
	{"roles/secretmanager.admin", "secretmanager_admin"},
}

func TestGenerateResourceNameWithPrimitiveRole(t *testing.T) {
	for _, testcase := range rolesAndExpectedNames {
		//Given
		role := testcase.role

		//when
		name := generateResourceName(role, "myprojectid")

		//Then
		assert.Equal(t, testcase.name, name)
	}
}
