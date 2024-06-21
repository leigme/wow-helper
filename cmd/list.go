package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "l",
	Short: "list",
	Long:  "",
	Run:   listCmd,
}

func init() {
	rootCmd.AddCommand(list)
}

func listCmd(c *cobra.Command, args []string) {
	fmt.Printf("list %s\n", args)
}
