package cmd

import (
	"github.com/spf13/cobra"
	"github.com/ubcctf/ctfci/ctfd"
)

func init() {
	rootCmd.AddCommand(cmdCTFd)
	cmdCTFd.AddCommand(cmdCTFdSync)
	cmdCTFd.AddCommand(cmdCTFdQuery)
}

var cmdCTFd = &cobra.Command{
	Use:   "ctfd",
	Short: "CTFd API client",
	Long:  `CTFd API client`,
}

var cmdCTFdSync = &cobra.Command{
	Use:   "sync",
	Short: "synchronize challenges",
	Long:  "synchronize challenges",
	Run:   ctfd.Sync,
}

var cmdCTFdQuery = &cobra.Command{
	Use:   "query",
	Short: "query challenges",
	Long:  "query challenges",
	Run:   ctfd.Query,
}
