/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"journal/users"

	"github.com/spf13/cobra"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login into existing account",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		//fmt.Println("login called")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		fmt.Printf("Logging in user: %s \n", username)
		fmt.Printf("Setting up the journal for %s \n", username)
		loginUser := login(username, password)
		if loginUser != nil {
			//users.CreateJournal(loginUser)
			//Start(loginUser)
			fmt.Println("User logged in!")
		}
	},
}

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringP("username", "u", "", "enter a username for user")
	loginCmd.Flags().StringP("password", "p", "", "enter a password for user")
	loginCmd.MarkFlagRequired("username")
	loginCmd.MarkFlagRequired("password")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loginCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loginCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func login(username, password string) *users.User {
	if users.Auth(username, password) {
		authUser := users.GetValue(username, password)
		return authUser
	} else {
		return nil
	}
}
