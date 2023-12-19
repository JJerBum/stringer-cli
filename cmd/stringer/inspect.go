package stringer

import (
	"fmt"

	"github.com/JJerBum/stringer-cli/pkg/stringer"
	"github.com/spf13/cobra"
)

var onlyDigits bool // local flag
var inspectCmd = &cobra.Command{
	Use:     "inspect",
	Short:   "inspects a string",
	Aliases: []string{"insp"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		i := args[0]
		res, kind := stringer.Inspect(i, onlyDigits)

		pluralS := "s"
		if res == 1 {
			pluralS = ""
		}
		fmt.Printf("'%s' has a %d %s%s.\n", i, res, kind, pluralS)
	},
}

func init() {
	inspectCmd.Flags().BoolVarP(&onlyDigits, "digits", "d", false, "Count only digits")
	rootCmd.AddCommand(inspectCmd)
}
