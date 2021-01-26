package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	cmd := &cobra.Command{}

	result := execGenerate(cmd, []string{})
	t.Logf("%s", result)
	assert.Error(t, result)

	result = execGenerate(cmd, []string{""})
	assert.Error(t, result, "User name cannot be empty! Please provide a user name")
}
