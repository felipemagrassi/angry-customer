/*
Copyright © 2024 NAME HERE <felipe.1magrassi@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/felipemagrassi/angry-customer/internal"
	"github.com/spf13/cobra"
)

var (
	URL         string
	Requests    uint64
	Concurrency uint64
)

var rootCmd = &cobra.Command{
	Use:   "angry-customer",
	Short: "A simple application that simulates an angry customer",
	Long: `A simple application that simulates an angry customer making 
	infinite requests to a server. It is used to test the resilience of a server.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("customer getting angry...")
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetUint64("requests")
		concurrency, _ := cmd.Flags().GetUint64("concurrency")

		err := internal.RunStresser(url, requests, concurrency)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&URL, "url", "u", "", "URL to make requests (required)")
	rootCmd.Flags().Uint64VarP(&Requests, "requests", "r", 1, "Number of requests to be made")
	rootCmd.Flags().Uint64VarP(&Concurrency, "concurrency", "c", 1, "Number of concurrent requests to be made")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.MarkFlagRequired("url")
	rootCmd.MarkFlagRequired("requests")
	rootCmd.MarkFlagRequired("concurrency")

}
