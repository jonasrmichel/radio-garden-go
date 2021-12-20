package cmd

import (
	"context"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// stationOpts contains the settings available for the station command.
var stationOpts struct {
	StationID string
}

// stationCmd represents the station command.
var stationCmd = &cobra.Command{
	Use:   "station",
	Short: "Get a radio station's details",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.GetAraContentChannelChannelIdWithResponse(
			context.Background(),
			stationOpts.StationID,
		)
		if err != radiogarden.CheckResponse(res, err) {
			return err
		}

		print(res.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(stationCmd)

	stationCmd.PersistentFlags().StringVarP(
		&stationOpts.StationID,
		"station-id",
		"i",
		"",
		"the station ID",
	)
	stationCmd.MarkPersistentFlagRequired("station-id")
}
