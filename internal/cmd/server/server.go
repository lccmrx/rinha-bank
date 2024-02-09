package server

import (
	"context"

	"github.com/spf13/cobra"
	"github.com/wyreyx/rinha-bank/internal/api/http"
	"github.com/wyreyx/rinha-bank/internal/app"
)

func Server() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "server",
		Short: "server commands",
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
			verbose, _ := cmd.Flags().GetBool("verbose")
			level, _ := cmd.Flags().GetString("level")

			ctx := context.Background()
			ctx = context.WithValue(ctx, "verbose", verbose)
			ctx = context.WithValue(ctx, "level", level)

			app.Start(ctx, http.NewServer)
		},
	}
}
