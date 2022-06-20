package backend

import (
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
)

type MainBlock struct {
	Terraform TerraformBlock `hcl:"terraform,block"`
}

type TerraformBlock struct {
	Backend BackendBlock `hcl:"backend,block"`
}

type BackendBlock struct {
	Name         string `hcl:"name,label"`
	Path         string `hcl:"path,optional"`
	WorkspaceDir string `hcl:"workspace_dir,optional"`
	Bucket       string `hcl:"bucket,optional"`
	Key          string `hcl:"key,optional"`
	Region       string `hcl:"region,optional"`
}

func ReadBackendFromFile(filename string) (*BackendBlock, error) {
	var tfBlock MainBlock

	parser := hclparse.NewParser()
	f, diags := parser.ParseHCLFile(filename)
	if diags.HasErrors() {
		return nil, diags
	}

	diags = gohcl.DecodeBody(f.Body, nil, &tfBlock)
	if diags.HasErrors() {
		return nil, diags
	}

	return &tfBlock.Terraform.Backend, nil
}
