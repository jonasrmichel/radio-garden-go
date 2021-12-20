package cmd

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// globalOpts contains the globally available settings.
var globalOpts struct {
	Server string
	Format bool
}

// rootCmd represents the entrypoint into the app.
var rootCmd = &cobra.Command{
	Use:   "radio-garden",
	Short: "An application for using the Radio Garden API",
}

// Execute is called to invoke rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	cobra.OnInitialize(initClient)

	rootCmd.PersistentFlags().StringVarP(
		&globalOpts.Server,
		"server",
		"s",
		"https://radio.garden/api",
		"the Radio Garden server",
	)
	rootCmd.PersistentFlags().BoolVarP(
		&globalOpts.Format,
		"format",
		"f",
		false,
		"format the output",
	)
}

// client is a global Radio Garden client that is initialized before any command
// is run.
var client *radiogarden.ClientWithResponses

// initClient initializes the global Radio Garden client.
func initClient() {
	// Construct a client with automatic redirect following disabled to avoid
	// downloading station audio streams via the station stream command.
	c, err := radiogarden.NewClientWithResponses(
		globalOpts.Server,
		radiogarden.WithFollowRedirectsDisabled(),
	)
	cobra.CheckErr(err)

	client = c
}

// print prints a string representation of v.
func print(v interface{}) (err error) {
	var s string

	if globalOpts.Format {
		s, err = renderFormatted(v)
	} else {
		s, err = render(v)
	}

	if err != nil {
		return
	}

	fmt.Println(s)

	return
}

// render returns a JSON string representation of v.
func render(v interface{}) (string, error) {
	bs, err := json.Marshal(v)
	if err != nil {
		return "", err
	}

	return string(bs), nil
}

// renderFormatted returns a formatted JSON string representation of v.
func renderFormatted(v interface{}) (string, error) {
	bs, err := json.MarshalIndent(v, "", " ")
	if err != nil {
		return "", err
	}

	return string(bs), nil
}
