package main

import (
	"fmt"
	"github.com/info24/eva/serve"
	"github.com/spf13/cobra"
	"os"
)

var Version string

func main() {
	root.AddCommand(addUser)
	root.AddCommand(versionCmd)
	err := root.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var root = &cobra.Command{
	Use:   "eva",
	Short: "run web server",
	Long:  "web server run default 9999",
	Run: func(cmd *cobra.Command, args []string) {
		serve.NewServer()
	},
}

var addUser = &cobra.Command{
	Use:     "add",
	Aliases: []string{"user"},
	Short:   "add user",
	Long:    "add user by eva cli",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			panic("user name is required")
		}
		serve.AddUser(args[0])
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(Version)
	},
}
