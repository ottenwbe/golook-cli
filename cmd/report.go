//Copyright 2017 Beate Ottenwälder
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

package cmd

import (
	"fmt"
	"github.com/ottenwbe/golook-cli/client"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	REPORT_COMMAND = "report"
)

var file string

var reportCmd = &cobra.Command{
	Use:   REPORT_COMMAND,
	Short: "Report files",
	Long:  "Report files and folders to the uplink server in order to be able to search for them from other devices.",
	Run: func(_ *cobra.Command, _ []string) {
		result, err := client.ReportFiles(file)
		if err != nil {
			logrus.WithError(err).Fatal("Cannot report file or folder.")
		}
		fmt.Print(result)
	},
}

func init() {

	reportCmd.Flags().StringVarP(&file, "file", "f", ".", "(required) file you want to report (default is the current working directory)")

	RootCmd.AddCommand(reportCmd)
}
