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
	"github.com/ottenwbe/golook/broker/models"
	"github.com/ottenwbe/golook/client"
	"github.com/spf13/cobra"
	"path/filepath"
)

const (
	reportCommand = "report"
)

var (
	report = models.FileReport{}
)

var reportCmd = &cobra.Command{
	Use:   reportCommand,
	Short: "Report files",
	Long:  "Report files and folders to the uplink server in order to be able to search for them from other devices.",
	Run: func(_ *cobra.Command, _ []string) {
		report.Path, _ = filepath.Abs(report.Path)
		result, err := client.ReportFiles(report)
		failOnError(err, "Cannot report file or folder.")
		fmt.Print(result)
	},
}

func init() {

	reportCmd.Flags().StringVarP(&report.Path, "path", "p", ".", "(required) path to the folder or file you want to report (default is the current working directory)")
	reportCmd.Flags().BoolVarP(&report.Delete, "delete", "d", false, "(optional) remove files from being monitored")

	RootCmd.AddCommand(reportCmd)
}
