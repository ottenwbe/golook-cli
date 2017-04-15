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

type ReportData struct {
	file    string
	folder  string
	replace bool
	monitor bool
}

var (
	reportData ReportData
)

var reportCmd = &cobra.Command{
	Use:   REPORT_COMMAND,
	Short: "Report files",
	Long:  "Report files and folders to the uplink server in order to be able to search for them from other devices.",
	Run: func(_ *cobra.Command, _ []string) {

		if reportData.file != "" {
			doReportFile()
		}

		if reportData.folder != "" {
			doReportFolder()
		}
	},
}

func doReportFile() {
	if reportData.replace {
		GolookIfce.ReportFile(reportData.file)
	} else {
		GolookIfce.ReportFileR(reportData.file)
	}
}

func doReportFolder() {
	if reportData.replace {
		GolookIfce.ReportFolder(reportData.folder)
	} else {
		GolookIfce.ReportFolderR(reportData.folder)
	}
}

func init() {

	reportData = ReportData{}

	reportCmd.Flags().StringVarP(&reportData.file, "file", "f", "", "(optional) file you want to report")
	reportCmd.Flags().StringVarP(&reportData.folder, "folder", "o", "", "(optional) folder you want to report")
	reportCmd.Flags().BoolVarP(&reportData.monitor, "monitor", "m", true, "(optional) instruct the server to continously monitor the file or folder")
	reportCmd.Flags().BoolVarP(&reportData.replace, "replace", "r", true, "(optional) enforces a replacement of all files or folders on the server")

	RootCmd.AddCommand(reportCmd)
}
