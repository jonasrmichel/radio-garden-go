package cmd

import (
	"context"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// searchOpts contains the settings available for the search command.
var searchOpts struct {
	Query string
}

// searchCmd represents the search command.
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for countries, places, and radio stations",
	RunE: func(cmd *cobra.Command, args []string) error {
		params := &radiogarden.GetSearchParams{
			Q: searchOpts.Query,
		}
		res, err := client.GetSearchWithResponse(context.Background(), params)
		if err != radiogarden.CheckResponse(res, err) {
			return err
		}

		print(res.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.PersistentFlags().StringVarP(
		&searchOpts.Query,
		"query",
		"q",
		"",
		"the search query",
	)
	searchCmd.MarkPersistentFlagRequired("query")
}
