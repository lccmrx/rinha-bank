package cmd

import (
	"fmt"
	"slices"

	"github.com/lccmrx/rinha-bank/internal/app"
	"github.com/lccmrx/rinha-bank/internal/cmd/server"
	"github.com/spf13/cobra"
)

func NewRoot() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "rinha-bank",
		Short: "rinha-bank is a simple bank system",
		Long: `Rinha Bank is a simple bank system that allows you to transact ephremeral money, 
using debit or credit transactions.`,
		CompletionOptions: cobra.CompletionOptions{
			HiddenDefaultCmd:  true,
			DisableDefaultCmd: true,
		},
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			level, err := cmd.Flags().GetString("level")
			if err != nil {
				return err
			}

			validLevels := []string{"debug", "info", "warning", "error"}
			if !slices.Contains(validLevels, level) {
				return fmt.Errorf("invalid mode: %s", level)
			}

			return nil
		},
	}

	setGloblalFlags(cmd)

	for _, command := range []*cobra.Command{
		version(),
		server.Server(),
	} {
		cmd.AddCommand(command)
	}
	return cmd
}

func setGloblalFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().BoolP("verbose", "v", false, "enables verbose mode")
	cmd.PersistentFlags().String("level", "info", "set logger level")
}

func version() *cobra.Command {
	return &cobra.Command{
		Use: "version",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Printf("Rinha Bank %s\n", app.Version)
		},
	}
}
