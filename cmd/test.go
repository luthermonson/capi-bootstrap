package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"k8s.io/apimachinery/pkg/runtime"

	v1 "k8s.io/client-go/tools/clientcmd/api/v1"

	"github.com/spf13/cobra"
	k8syaml "sigs.k8s.io/yaml"
)

// listCmd represents the list command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {

		kconf, err := os.ReadFile("./kubeconfig.yaml")
		if err != nil {
			return err
		}
		j1, err := k8syaml.YAMLToJSON(kconf)
		if err != nil {
			return err
		}

		var config v1.Config
		if err := json.Unmarshal(j1, &config); err != nil {
			return err
		}

		config.Extensions = append(config.Extensions, v1.NamedExtension{
			Name: "capi-bootstrap",
			Extension: runtime.RawExtension{
				Raw: []byte("{\"test-capi-bootstrap-cluster\":\"test value\"}"),
			},
		})
		fmt.Printf("%+v", config)

		j2, err := json.Marshal(config)
		if err != nil {
			return err
		}

		y2, err := k8syaml.JSONToYAML(j2)
		if err != nil {
			return err
		}
		fmt.Println(string(y2))

		return nil
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
