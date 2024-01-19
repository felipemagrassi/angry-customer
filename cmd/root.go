/*
Copyright © 2024 NAME HERE <felipe.1magrassi@gmail.com>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "angry-customer",
	Short: "A simple application that simulates an angry customer",
	Long: `A simple application that simulates an angry customer making 
	infinite requests to a server. It is used to test the resilience of a server.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringP("url", "u", "", "URL to make requests")
	rootCmd.Flags().Uint64P("requests", "r", 0, "Number of requests to make")
	rootCmd.Flags().Uint64P("concurrency", "c", 0, "Number of concurrent requests")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
