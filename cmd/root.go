package cmd

import (
	"capi-bootstrap/cloud"
	"capi-bootstrap/cloud/linode"
	"os"

	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "capi-bootstrap",
	Short: "",
	Long:  ``,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.capi-bootstrap.yaml)")
}

func CloudProviderFromClusterKind(kind string) cloud.Provider {
	switch kind {
	case "LinodeCluster":
		return linode.NewProvider()
	default:
		return nil
	}
}
