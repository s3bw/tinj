/*Package cmd provides and interface to tinj cli.
- Provide format

`[(cyan|service)]|(blue|severity)|"(message)" |(exc_info)|"(red|httpRequest.status)`

Severity: INFO, WARNING, ERROR, DEBUG. Colours?
*/
package cmd

import (
	"fmt"
	"os"

	tinj "github.com/foxyblue/tinj/pkg"
	"github.com/spf13/cobra"
)

// DefaultFormat for logs
const DefaultFormat = `[(green|service)]|(blue|severity)|(red|httpRequest.status)|"(message)"|(exc_info)|`

func init() {
	rootCmd.Flags().StringP("format", "f", "", "Supply a format string")
}

var rootCmd = &cobra.Command{
	Use: "tinj",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		if format == "" {
			format = DefaultFormat
		}

		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}

		// Help Text
		if info.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("The command is intended to work with pipes.")
			fmt.Println("Usage: cat file.json | tinj")
			return
		}
		tinj.ReadStdin(format)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
