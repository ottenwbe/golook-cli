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

	"github.com/ottenwbe/golook-cli/client"
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
	searchString string
)

var queryCmd = &cobra.Command{
	Use:   QRY_COMMAND,
	Short: "Query the golook server.",
	Long:  "Query system, files, and file locations from a golook server.",
	Run: func(_ *cobra.Command, _ []string) {
		log.WithField("command", QRY_COMMAND).Debug("Query for all files.")

		systemFiles, err := client.GetFiles(searchString)
		failOnError(err, "Could not query for files.")

		b, err := json.Marshal(systemFiles)
		failOnError(err, "Could not marshal files.")

		fmt.Print(string(b))
	},
}

func failOnError(err error, errorDescription string) {
	if err != nil {
		log.WithError(err).Fatal(errorDescription)
	}
}

func init() {

	queryCmd.Flags().StringVarP(&searchString, "file", "f", "", "(required) file you are looking for")

	RootCmd.AddCommand(queryCmd)
}
