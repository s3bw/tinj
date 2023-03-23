/*Package cmd provides and interface to tinj cli.
- Provide format

`[(cyan|service)]|(blue|severity)|"(message)" |(exc_info)|"(red|httpRequest.status)`

Severity: INFO, WARNING, ERROR, DEBUG. Colours?
*/
package cmd

import (
	"fmt"
	"os"
	"strings"

	tinj "github.com/foxyblue/tinj/pkg"
	"github.com/spf13/cobra"
)

const (
	// DefaultStyle of logs to parse
	DefaultStyle = `tail`
	// DefaultFormat for logs
	DefaultFormat = `(service|yellow),(severity|blue),(httpRequest.status|red),(message),(exc_info)`
	// DefaultSeparator between fields
	DefaultSeparator = ` | `
	// DefaultMetrics to capture when aggregating
	DefaultMetrics = `(severity|white),(httpRequest.requestMethod|white),(httpRequest.status|green)`
	// DefaultTrim is set to zero
	DefaultTrim = "0"
)

func init() {
	rootCmd.Flags().StringP("format", "f", "", "Supply a format string")
	rootCmd.Flags().StringP("style", "s", "", "Supply a style of log")
	rootCmd.Flags().StringP("separator", "p", "", "Separate fields by supplied character")
	rootCmd.Flags().Int("trim", 0, "Supply an integer")
}

func parseFlag(cmd *cobra.Command, flag, defaultFlag string) string {
	f, _ := cmd.Flags().GetString(flag)
	if f == "" {
		return defaultFlag
	}
	return f
}

func parseStyle(style string) (tinj.Style, error) {
	switch strings.ToLower(style) {
	case `tail`:
		return tinj.Tail, nil
	case `stern`:
		return tinj.Stern, nil
	case `compose`:
		return tinj.Compose, nil
	}
	return tinj.Tail, fmt.Errorf("--style %q not recognised", style)
}

var rootCmd = &cobra.Command{
	Use: "tinj",
	Run: func(cmd *cobra.Command, args []string) {
		format := parseFlag(cmd, "format", DefaultFormat)
		separator := parseFlag(cmd, "separator", DefaultSeparator)
		trim, _ := cmd.Flags().GetInt("trim")
		style := parseFlag(cmd, "style", DefaultStyle)

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

		// Parse provided style flags
		lineStyle, err := parseStyle(style)
		if err != nil {
			fmt.Println(err)
			fmt.Println("Try: tail, stern or compose")
			return
		}
		tinj.ReadStdin(format, separator, lineStyle, trim)
	},
}

var subCmd = &cobra.Command{
	Use:   "count [no options!]",
	Short: "My counter",
	Run: func(cmd *cobra.Command, args []string) {
		metrics := parseFlag(cmd, "metrics", DefaultMetrics)
		separator := parseFlag(cmd, "separator", DefaultSeparator)

		info, err := os.Stdin.Stat()
		if err != nil {
			panic(err)
		}

		// Help Text
		if info.Mode()&os.ModeCharDevice != 0 {
			fmt.Println("The command is intended to work with pipes.")
			fmt.Println("Usage: cat file.json | tinj count")
			return
		}
		tinj.AggregateStdin(metrics, separator)
	},
}

// Execute runs tinj CLI
func Execute() {
	rootCmd.AddCommand(subCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
