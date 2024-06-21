package cmd

import (
	"bufio"
	_ "embed"
	"fmt"
	"github.com/spf13/cobra"
	"log"
	"os"
	"path/filepath"
	"strings"
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "init config",
	Run:   runInit,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func runInit(c *cobra.Command, args []string) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	dir, err := ExecName()
	if err != nil {
		log.Fatal(err)
	}
	dir = filepath.Join(userHome, dir)
	configPath := filepath.Join(dir, configName)
	configFile, err := os.Open(configPath)
	defer configFile.Close()
	if err != nil {
		if os.IsNotExist(err) {
			err = os.MkdirAll(filepath.Dir(configPath), os.ModePerm)
			if err != nil {
				log.Fatal(err)
			}
			configFile, err = os.Create(configPath)
			if err != nil {
				log.Fatal(err)
			}
			w := bufio.NewWriter(configFile)
			if len(args) > 0 {
				defaultConfig, err = NewConfigBytes(args...)
				if err != nil {
					log.Fatal(err)
				}
			}
			w.Write(defaultConfig)
			w.Flush()
			return
		}
	}
	fmt.Println(configFile.Name())
}

func ExecName() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	exec := filepath.Base(path)
	ext := filepath.Ext(exec)
	if !strings.EqualFold(ext, "") {
		exec = strings.TrimSuffix(exec, ext)
	}
	return exec, nil
}
