/*
Copyright © 2024 NAME HERE <felipe.1magrassi@gmail.com>
*/
package cmd

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/felipemagrassi/angry-customer/internal"
	"github.com/spf13/cobra"
)

var (
	URL         string
	Concurrency uint64
	Requests    uint64
)

var rootCmd = &cobra.Command{
	Use:   "angry-customer",
	Short: "A simple application that simulates an angry customer",
	Long: `A simple application that simulates an angry customer making 
	infinite requests to a server. It is used to test the resilience of a server.`,
	Run: func(cmd *cobra.Command, args []string) {
		url, _ := cmd.Flags().GetString("url")
		requests, _ := cmd.Flags().GetUint64("requests")
		concurrency, _ := cmd.Flags().GetUint64("concurrency")

		if url == "" || requests == 0 || concurrency == 0 {
			url, requests, concurrency = buildFromArgs(args)
		}

		err := internal.RunStresser(url, requests, concurrency)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func buildFromArgs(args []string) (string, uint64, uint64) {
	url := ""
	requestsStr := ""
	concurrencyStr := ""

	requestsRegex := regexp.MustCompile(`-r.*`)
	concurrencyRegex := regexp.MustCompile(`-c.*`)
	httpRegex := regexp.MustCompile(`https?://.*`)
	numberRegex := regexp.MustCompile(`[0-9]+`)

	for _, arg := range args {
		fmt.Println(arg)
		if httpRegex.MatchString(arg) {
			url = httpRegex.FindString(arg)
		}

		if requestsRegex.MatchString(arg) {
			requestsStr = numberRegex.FindString(arg)
		}

		if concurrencyRegex.MatchString(arg) {
			concurrencyStr = numberRegex.FindString(arg)
		}
	}

	requests, err := strconv.ParseUint(requestsStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	concurrency, err := strconv.ParseUint(concurrencyStr, 10, 64)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return url, requests, concurrency
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&URL, "url", "u", "", "URL to be requested")
	rootCmd.PersistentFlags().Uint64VarP(&Requests, "requests", "r", 0, "Number of requests to be made")
	rootCmd.PersistentFlags().Uint64VarP(&Concurrency, "concurrency", "c", 0, "Number of concurrent requests to be made")

	// rootCmd.MarkPersistentFlagRequired("url")
	// rootCmd.MarkPersistentFlagRequired("requests")
	// rootCmd.MarkPersistentFlagRequired("concurrency")
}
