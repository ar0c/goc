package cmd

import (
	"github.com/ar0c/goc/v2/pkg/log"
	"github.com/spf13/cobra"
	"testing"
)

func Test_serve(t *testing.T) {
	type args struct {
		cmd  *cobra.Command
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "",
			args: args{},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		log.NewLogger(true)
		t.Run(tt.name, func(t *testing.T) {
			serve(tt.args.cmd, tt.args.args)
		})
	}
}
