package test

import (
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestModule(t *testing.T) {
	ctx := t.Context()

	terraformOptions := terraform.WithDefaultRetryableErrors(t, &terraform.Options{
		TerraformDir: "../examples/complete",
	})

	defer terraform.DestroyContext(t, ctx, terraformOptions)

	terraform.InitAndApplyContext(t, ctx, terraformOptions)

	output := terraform.OutputContext(t, ctx, terraformOptions, "value")
	assert.Equal(t, "Hello!", output)
}
