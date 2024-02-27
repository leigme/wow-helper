/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const (
	confName   = "conf.json"
	srcDir     = "srcDir"
	destDir    = "destDir"
	backupDirs = "backupDirs"
)

var (
	conf     *viper.Viper
	confPath string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "wow-helper",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cmdName := "." + strings.TrimSuffix(filepath.Base(os.Args[0]), filepath.Ext(os.Args[0]))
	rootPath, err := os.UserHomeDir()
	if err != nil {
		rootPath, err = filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			log.Fatalln(err)
		}
	}

	workDir := filepath.Join(rootPath, cmdName)

	err = os.MkdirAll(workDir, os.ModePerm)

	if err != nil {
		log.Fatalln(err)
	}

	rootCmd.Flags().StringVar(&confPath, "c", filepath.Join(workDir, confName), "config file path")

	conf = viper.New()

	conf.SetDefault(srcDir, "")
	conf.SetDefault(destDir, workDir)
	conf.SetDefault(backupDirs, defaultBackupDirs())

	conf.SetConfigName(confName)
	conf.SetConfigType("json")
	conf.AddConfigPath(workDir)

	if err = conf.ReadInConfig(); err != nil {
		err = conf.SafeWriteConfigAs(filepath.Join(workDir, confName))
		if err != nil {
			log.Fatalln(err)
		}
	}

}

func defaultBackupDirs() []string {
	files := make([]string, 0)
	files = append(files, "Fonts", "Interface", "WTF")
	return files
}
