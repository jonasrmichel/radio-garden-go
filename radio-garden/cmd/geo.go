package cmd

import (
	"context"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// geoCmd represents the geo command.
var geoCmd = &cobra.Command{
	Use:   "geo",
	Short: "Get your geolocation",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.GetGeoWithResponse(context.Background())
		if err != radiogarden.CheckResponse(res, err) {
			return err
		}

		print(res.JSON200)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(geoCmd)
}
