# iam-to-terraform   
Simple tools to read project iam policy and generate terraform resource associated.
- Terraform resource can be customized through the template ```resource_template.tmpl```
- Skip service agent role
- Required get iam policy permission on the project

### Test
go test


### Build:
go build

### Run
iam-to-terraform ```<project-id>``` > iam.tf