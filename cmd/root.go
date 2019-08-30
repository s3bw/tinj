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

const (
	// DefaultFormat for logs
	DefaultFormat = `[(service|yellow)]|(severity|blue)|(httpRequest.status|red)|"(message)"|(exc_info)|`
	// DefaultSeparator between fields
	DefaultSeparator = ` | `
)

func init() {
	rootCmd.Flags().StringP("format", "f", "", "Supply a format string")
	rootCmd.Flags().StringP("separator", "s", "", "Separate fields by supplied character")
}

var rootCmd = &cobra.Command{
	Use: "tinj",
	Run: func(cmd *cobra.Command, args []string) {
		format, _ := cmd.Flags().GetString("format")
		if format == "" {
			format = DefaultFormat
		}

		separator, _ := cmd.Flags().GetString("separator")
		if separator == "" {
			separator = DefaultSeparator
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
		tinj.ReadStdin(format, separator)
	},
}

var subCmd = &cobra.Command{
	Use:   "count [no options!]",
	Short: "My counter",
	Run: func(cmd *cobra.Command, args []string) {
		metrics := `count(httpRequest.status)`
		separator := ` | `

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
		tinj.AggregateStdin(metrics, separator)
	},
}

func Execute() {
	rootCmd.AddCommand(subCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
