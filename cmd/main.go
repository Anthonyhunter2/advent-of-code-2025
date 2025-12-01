package main

import (
	"fmt"
	"os"

	"advent-of-code-2025/cmd/aoc"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(aoc.DayCmd)
}

var rootCmd = &cobra.Command{
	Use:   "aoc",
	Short: "aoc is a tool for solving Advent of Code problems",
	Long:  `aoc is a tool for solving Advent of Code problems and acts as a helper for each day's problem`,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
