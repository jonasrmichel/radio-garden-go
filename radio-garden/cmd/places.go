package cmd

import (
	"context"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// placesCmd represents the places command.
var placesCmd = &cobra.Command{
	Use:   "places",
	Short: "Get places with registered radio stations",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.GetAraContentPlacesWithResponse(context.Background())
		if err != radiogarden.CheckResponse(res, err) {
			return err
		}

		print(res.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(placesCmd)
}
