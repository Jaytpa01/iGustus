package util_test

import (
	"fmt"
	"testing"

	"github.com/Jaytpa01/iGustus/pkg/util"
	"github.com/stretchr/testify/assert"
)

func TestHasArgumentFlag(t *testing.T) {
	tests := []struct {
		argString string
		hasArgs   bool
	}{
		{
			argString: "this is a test -i trog -g slog --friggers",
			hasArgs:   true,
		},
		{
			argString: "this has no args",
			hasArgs:   false,
		},
		{
			argString: "only has --thistype",
			hasArgs:   true,
		},
		{
			argString: "has --multiple of --doubles pisser",
			hasArgs:   true,
		},
		{
			argString: "",
			hasArgs:   false,
		},
		{
			argString: "just -seing but this should fail",
			hasArgs:   false,
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.hasArgs, util.HasArgumentFlag(test.argString), fmt.Sprintf("%s failed", test.argString))
	}
}
