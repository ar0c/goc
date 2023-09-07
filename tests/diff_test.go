package tests

import "testing"

func Test_diff(t *testing.T) {
	type args struct {
		target string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "test",
			args: args{
				target: "1a3484",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			diff(tt.args.target)
		})
	}
}
