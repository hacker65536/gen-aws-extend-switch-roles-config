/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/hacker65536/gen-aws-extend-switch-roles-config/pkg/myaws"
	"github.com/hacker65536/gen-aws-extend-switch-roles-config/pkg/utils"
	"github.com/spf13/cobra"
)

// genssoCmd represents the gensso command
var genssoCmd = &cobra.Command{
	Use:   "gensso",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		inputProfilePrefix := utils.StringPrompt("Input Profile prefix [default awsso]: ")
		if inputProfilePrefix == "" {
			inputProfilePrefix = "awsso"
		}
		fmt.Println("ProfilePrefix: ", inputProfilePrefix)

		inputSessionName := utils.StringPrompt("Input Session Name [default awsso]: ")
		if inputSessionName == "" {
			inputSessionName = "awsso"
		}
		fmt.Println("SessionName: ", inputSessionName)

		inputRoleName := utils.StringPrompt("Input Role Name [default AWSAdministratorAccess]: ")
		if inputRoleName == "" {
			inputRoleName = "AWSAdministratorAccess"
		}
		fmt.Println("RoleName: ", inputRoleName)

		fmt.Println()

		myaws.GenSSOConfig(inputProfilePrefix, inputSessionName, inputRoleName)

	},
}

func init() {
	rootCmd.AddCommand(genssoCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genssoCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genssoCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
