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
	"encoding/json"
	"fmt"

	"github.com/ottenwbe/golook/control"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	file   string
	system string
)

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the golook server.",
	Long:  "Query system, files, and file locations from a golook server.",
	Run: func(_ *cobra.Command, _ []string) {
		golookIfce := control.NewController()
		selectQueryAction(golookIfce)
	},
}

func selectQueryAction(golookIfce control.LookController) {
	switch system {
	case "this":
		queryForReportedFiles(golookIfce)
	case "":
		queryForAllFiles(golookIfce)
	default:
		log.Fatal("System filtering  not supported yet")
	}
}

func queryForAllFiles(golookIfce control.LookController) {
	systemFiles, err := golookIfce.QueryAllSystemsForFile(file)
	if err != nil {
		log.WithError(err).Fatal("Could not query for files.")
	}
	b, err := json.Marshal(systemFiles)
	if err != nil {
		log.WithError(err).Fatal("Could not marshal files.")
		fmt.Print(string(b))
	}
}

func queryForReportedFiles(golookIfce control.LookController) {
	files, err := golookIfce.QueryReportedFiles()
	if err != nil {
		log.WithError(err).Fatal("Could not query for reported files.")
	}
	b, err := json.Marshal(files)
	if err != nil {
		log.WithError(err).Fatal("Could not marshal reported files.")
	}
	fmt.Print(string(b))
}

func init() {
	RootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVarP(&file, "file", "f", "", "(required) file you are looking for")
	queryCmd.Flags().StringVarP(&system, "system", "s", "this", "(optional) only look at the given the system for the file")
}
