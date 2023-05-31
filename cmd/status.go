// Copyright 2020-2021 Changkun Ou. All rights reserved.
// Use of this source code is governed by a GPL-3.0
// license that can be found in the LICENSE file.

package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"changkun.de/x/midgard/api/daemon"
	"changkun.de/x/midgard/internal/config"
	"changkun.de/x/midgard/internal/term"
	"changkun.de/x/midgard/internal/types"
	"changkun.de/x/midgard/internal/types/proto"
	"changkun.de/x/midgard/internal/utils"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "check midgard setup status",
	Long:  `check midgard setup status`,
	Args:  cobra.ExactArgs(0),
	Run: func(_ *cobra.Command, args []string) {
		var s string

		// check server status
		url := config.Get().Domain+"/midgard/ping"
		fmt.Println(url)
		res, err := utils.Request(http.MethodGet,url, nil)
		if err != nil {
			s += fmt.Sprintf("server status: %s, %v\n",
				term.Red("request error"), err)
		} else {
			var out types.PingOutput
			err = json.Unmarshal(res, &out)
			if err != nil {
				s += fmt.Sprintf("server status: %s, details:\n%v\n",
					term.Red("failed to parse ping response from server"),
					err)
			} else {
				s += fmt.Sprintf("server status: %s\n", term.Green("OK"))
			}
		}

		// check daemon status
		daemon.Connect(func(ctx context.Context, c proto.MidgardClient) {
			_, err := c.Ping(ctx, &proto.PingInput{})
			if err != nil {
				s += fmt.Sprintf("daemon status: %s, details:\n%v\n",
					term.Red("failed to ping daemon"),
					status.Convert(err).Message())
			} else {
				s += fmt.Sprintf("daemon status: %s\n", term.Green("OK"))
			}
		})

		fmt.Println(s)
	},
}
