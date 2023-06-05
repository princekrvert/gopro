/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
        "log"
        "github.com/google/gopacket/pcap"
	"github.com/spf13/cobra"
)

// showallCmd represents the showall command
var showallCmd = &cobra.Command{
	Use:   "showall",
	Short: "This command will show all interfaces ",
	Long: `Run this command to see the all interfaces along with the details.`,
	Run: func(cmd *cobra.Command, args []string) {
		devices, err := pcap.FindAllDevs()
    if err != nil {
        log.Fatal(err)
    }
    for _,devies := range devices {
    fmt.Println(devies)}
	},
}

func init() {
	rootCmd.AddCommand(showallCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showallCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showallCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
