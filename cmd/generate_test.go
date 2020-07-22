package cmd

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	cmd := &cobra.Command{}

	result := runGenerate(cmd, []string{})
	t.Logf("%s", result)
	assert.Error(t, result)

	result = run(cmd, []string{""})
	assert.Equal(t, result, fmt.Errorf("User name cannot be empty! Please provide a user name"))
}
