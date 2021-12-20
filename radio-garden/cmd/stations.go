package cmd

import (
	"context"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// stationsOpts contains the settings available for the stations command.
var stationsOpts struct {
	PlaceID string
}

// stationsCmd represents the stations command.
var stationsCmd = &cobra.Command{
	Use:   "stations",
	Short: "Get a place's registered radio stations",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.GetAraContentPagePlaceIdChannelsWithResponse(
			context.Background(),
			stationsOpts.PlaceID,
		)
		if err != radiogarden.CheckResponse(res, err) {
			return err
		}

		print(res.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(stationsCmd)

	stationsCmd.PersistentFlags().StringVarP(
		&stationsOpts.PlaceID,
		"place-id",
		"i",
		"",
		"the place ID",
	)
	stationsCmd.MarkPersistentFlagRequired("place-id")
}
