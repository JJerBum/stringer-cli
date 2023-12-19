package stringer

import (
	"fmt"

	"github.com/JJerBum/stringer-cli/pkg/stringer"
	"github.com/spf13/cobra"
)

var reverseCmd = &cobra.Command{
	Use:     "reverse",
	Short:   "Reverses a string",
	Aliases: []string{"rev"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := stringer.Reverse(args[0])
		fmt.Println(res)
	},
}

func init() {
	rootCmd.AddCommand(reverseCmd)
}
