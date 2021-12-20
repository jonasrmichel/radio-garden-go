package cmd

import (
	"context"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// placeOpts contains the settings available for the place command.
var placeOpts struct {
	PlaceID string
}

// placeCmd represents the place command.
var placeCmd = &cobra.Command{
	Use:   "place",
	Short: "Get a place's details",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.GetAraContentPagePlaceIdWithResponse(
			context.Background(),
			placeOpts.PlaceID,
		)
		if err != radiogarden.CheckResponse(res, err) {
			return err
		}

		print(res.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(placeCmd)

	placeCmd.PersistentFlags().StringVarP(
		&placeOpts.PlaceID,
		"place-id",
		"i",
		"",
		"the place ID",
	)
	placeCmd.MarkPersistentFlagRequired("place-id")
}
