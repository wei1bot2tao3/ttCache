package OSCache

import "testing"

func TestNewBuildInMapCache(t *testing.T) {
	testCass := []struct {
		name string

		wantErr error
		wantRes map[string]any
	}{
		{},
	}
	for _, tc := range testCass {
		t.Run(tc.name, func(t *testing.T) {

		})
	}
}
