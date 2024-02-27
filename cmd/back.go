package cmd

import (
	"github.com/leigme/wow-helper/file"
	"github.com/spf13/cobra"
	"log"
	"path/filepath"
)

var backup = &cobra.Command{
	Use:   "b",
	Short: "backup command",
	Long:  filepath.Join(confPath, confName),
	Run:   backupCmd,
}

func init() {
	rootCmd.AddCommand(backup)
}

var backupCmd = func(cmd *cobra.Command, args []string) {
	for _, bd := range conf.GetStringSlice(backupDirs) {
		err := file.Zip(filepath.Join(conf.GetString(srcDir), bd), conf.GetString(destDir))
		if err != nil {
			log.Println(err)
		}
	}
}
