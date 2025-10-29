package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "neno",
	Short: "NENO - Never Enough Notes Organizar",
	Long:  "A markdown-first note organiszer for your terminal",
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(newCmd)
	rootCmd.AddCommand(openCmd)
	rootCmd.AddCommand(showCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(tagsCmd)
	rootCmd.AddCommand(searchCmd)
}
