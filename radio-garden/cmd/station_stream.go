package cmd

import (
	"context"
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	radiogarden "github.com/jonasrmichel/radio-garden-go"
)

// stationStreamCmd represents the station stream command.
var stationStreamCmd = &cobra.Command{
	Use:   "stream",
	Short: "Get a radio station's live broadcast stream",
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := client.HeadAraContentListenChannelIdChannelMp3WithResponse(
			context.Background(),
			stationOpts.StationID,
		)
		if err != nil {
			return err
		} else if err := radiogarden.EnsureStatusByCode(res, http.StatusFound); err != nil {
			return err
		}

		fmt.Println(res.HTTPResponse.Header.Get("location"))

		return nil
	},
}

func init() {
	stationCmd.AddCommand(stationStreamCmd)
}
