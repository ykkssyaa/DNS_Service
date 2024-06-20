package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// hostnameCmd represents the hostname command
var hostnameCmd = &cobra.Command{
	Use:   "hostname",
	Short: "Get the hostname of a machine",
	Long:  `Get the hostname of a machine`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hostname called")
	},
}

func init() {
	rootCmd.AddCommand(hostnameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// hostnameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// hostnameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
