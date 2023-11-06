package gapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestResourceIdent(t *testing.T) {
	require.Equal(t, "1", ResourceID(1).String())
	require.Equal(t, ResourceUID("testing").String(), "testing")
}
