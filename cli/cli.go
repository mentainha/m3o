// Package cli contains source code for the cli
package cli

import (
	"bytes"
	"fmt"
	"encoding/json"
	"os"

	"github.com/m3o/m3o-go/client"
	"github.com/spf13/cobra"
)

var usage = `m3o call [service] [endpoint] [request] e.g m3o call helloworld Call '{"name": "Alice"}'`

var rootCmd = &cobra.Command{
	Use:   "m3o",
	Short: "m3o is an open source public cloud platform",
	Long:  `An open source AWS alternative built for the next generation of developers. See https://m3o.com for more info.`,
}

var callCmd = &cobra.Command{
	Use:   "call",
	Short: "Call a service",
	Long:  `Call a service on the M3O platform`,
	Run: func(cmd *cobra.Command, args []string) {
		// check for a token
		token := os.Getenv("M3O_API_TOKEN")
		if len(token) == 0 {
			fmt.Fprintln(os.Stderr, "Missing M3O_API_TOKEN")
			os.Exit(1)
		}

		// get the args
		if len(args) != 3 {
			fmt.Fprintln(os.Stderr, usage)
			os.Exit(1)
		}

		service := args[0]
		endpoint := args[1]
		request := args[2]

		var req map[string]interface{}

		if err := json.Unmarshal([]byte(request), &req); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse request as JSON: %v\n", err)
			os.Exit(1)
		}

		c := client.NewClient(nil)
		c.SetToken(token)

		var rsp json.RawMessage

		if err := c.Call(service, endpoint, req, &rsp); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

		var out bytes.Buffer
		defer out.Reset()

		if err := json.Indent(&out, rsp, "", "\t"); err != nil {
			fmt.Fprintln(os.Stderr, "Error while trying to format the response", err)
		}

		fmt.Printf("%v\n", string(out.Bytes()))
	},
}

func Execute() {
	rootCmd.AddCommand(callCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
