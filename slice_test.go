package dailyleetcode

import "testing"

func Test_slice(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test case"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slice()
		})
	}
}
