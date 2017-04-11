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
	"github.com/spf13/cobra"
)

const (
	REPORT_COMMAND = "report"
)

var (
	reportFile        string
	reportFolder      string
	reportPermanent   bool //TODO: Implement permanent reporting
	reportReplacement bool
)

var reportCmd = &cobra.Command{
	Use:   REPORT_COMMAND,
	Short: "Report files",
	Long:  "Report files and folders to the uplink server in order to be able to search them.",
	Run: func(_ *cobra.Command, _ []string) {

		if reportFile != "" {
			doReportFile()
		}

		if reportFolder != "" {
			doReportFolder()
		}
	},
}

func doReportFile() {
	//TODO ReportFileF in golookinterface
	golookIfce.ReportFile(reportFile)
}

func doReportFolder() {
	if reportReplacement {
		golookIfce.ReportFolder(reportFolder)
	} else {
		golookIfce.ReportFolderR(reportFolder)
	}
}

func init() {

	reportCmd.Flags().StringVarP(&reportFile, "file", "f", "", "(optional) file you want to report")
	reportCmd.Flags().StringVarP(&reportFolder, "folder", "o", "", "(optional) folder you want to report")
	reportCmd.Flags().BoolVarP(&reportPermanent, "monitor", "m", true, "(optional) allow server to monitor continously the file or folder")
	reportCmd.Flags().BoolVarP(&reportReplacement, "replace", "m", true, "(optional) enforces a replacement of all files or folders on the server")

	RootCmd.AddCommand(reportCmd)
}
