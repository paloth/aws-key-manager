package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	// result := run(cmd, []string{})
	// assert.Error(t, result)
	cmd := &cobra.Command{}

	result := run(cmd, []string{})
	t.Logf("%s", result)
	assert.Error(t, result)
}
