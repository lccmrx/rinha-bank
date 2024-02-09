package grpc

import (
	"github.com/spf13/cobra"
)

func Grpc() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "grpc",
		Short: "GRPC commands",
	}

	for _, command := range []*cobra.Command{
		start(),
	} {
		cmd.AddCommand(command)
	}

	return cmd
}

func start() *cobra.Command {
	return &cobra.Command{
		Use: "start",
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
}
