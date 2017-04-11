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

	"github.com/ottenwbe/golook/routing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	QUERY_THIS        = "this"
	QUERY_ALL         = "all"
	QUERY_ALL_DEFAULT = ""
)

const (
	QRY_COMMAND = "query"
)

var (
	file       string
	system     string
	golookIfce routing.LookRouter
)

var queryCmd = &cobra.Command{
	Use:   QRY_COMMAND,
	Short: "Query the golook server.",
	Long:  "Query system, files, and file locations from a golook server.",
	Run: func(_ *cobra.Command, _ []string) {
		selectQueryAction(golookIfce)
	},
}

func selectQueryAction(golookIfce routing.LookRouter) {
	log.WithField("command", QRY_COMMAND).WithField("system", system).Debug("Selecting query action.")
	switch system {
	case QUERY_THIS:
		queryForReportedFiles()
	case QUERY_ALL, QUERY_ALL_DEFAULT:
		queryForAllFiles()
	default:
		log.Fatal("System filtering not supported yet")
	}
}

func queryForAllFiles() {
	log.WithField("command", QRY_COMMAND).WithField("system", system).Debug("Query for all files.")

	systemFiles, err := golookIfce.QueryAllSystemsForFile(file)
	failOnError(err, "Could not query for files.")

	b, err := json.Marshal(systemFiles)
	failOnError(err, "Could not marshal files.")

	fmt.Print(string(b))
}

func queryForReportedFiles() {

	log.WithField("command", QRY_COMMAND).WithField("system", system).Debug("Query reported files.")

	files, err := golookIfce.QueryReportedFiles()
	failOnError(err, "Could not query for reported files.")

	b, err := json.Marshal(files)
	failOnError(err, "Could not marshal reported files.")

	fmt.Print(string(b))
}

func failOnError(err error, errorDescription string) {
	if err != nil {
		log.WithError(err).Fatal(errorDescription)
	}
}

func init() {
	golookIfce = routing.NewRouter()

	queryCmd.Flags().StringVarP(&file, "file", "f", "", "(required) file you are looking for")
	queryCmd.Flags().StringVarP(&system, "system", "s", "this", "(optional) only look at the given the system for the file")

	RootCmd.AddCommand(queryCmd)
}
