package cmd

import "github.com/spf13/cobra"

var restore = &cobra.Command{
	Use:   "r",
	Short: "restore command",
	Long:  "",
	Run:   restoreCmd,
}

func init() {
	rootCmd.AddCommand(restore)
}

func restoreCmd(c *cobra.Command, args []string) {

}
