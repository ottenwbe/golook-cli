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
package main

import (
	"fmt"

	"github.com/ottenwbe/golook/control"
	"github.com/ottenwbe/golook/routing"

	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	PROGRAM_NAME = "golook-cli"
	VERSION      = "v0.1.0-dev"
)

var (
	host string
	port int

	file   string
	system string
)

var RootCmd = &cobra.Command{
	Use:   PROGRAM_NAME,
	Short: "Golook cli",
	Long:  "Cli for the golook distributed file search.",
}

var cmdVersion = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", PROGRAM_NAME),
	Long:  fmt.Sprintf("All software has versions. This is %s's", PROGRAM_NAME),
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Print(VERSION)
	},
}

var queryCmd = &cobra.Command{
	Use:   "query",
	Short: "Query the golook server.",
	Long:  "Query system, files, and file locations from a golook server.",
	Run: func(_ *cobra.Command, _ []string) {
		golookIfce := control.NewController()

		switch system {
		case "this":
			files, err := golookIfce.QueryReportedFiles()
			if err != nil {
				log.WithError(err).Fatal("Could not query for reported files.")
			}

			b, err := json.Marshal(files)
			if err != nil {
				log.WithError(err).Fatal("Could not marshal reported files.")
			}

			fmt.Print(string(b))
		case "":
			systemFiles, err := golookIfce.QueryAllSystemsForFile(file)
			if err != nil {
				log.WithError(err).Fatal("Could not query for files.")
			}

			b, err := json.Marshal(systemFiles)
			if err != nil {
				log.WithError(err).Fatal("Could not marshal files.")
			}

			fmt.Print(string(b))
		default:
			log.Fatal("System filtering  not supported yet")
		}
	},
}

var reportCmd = &cobra.Command{
	Use:   "report",
	Short: "",
	Long:  "",
	Run: func(_ *cobra.Command, _ []string) {

	},
}

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "",
	Long:  "",
	Run: func(_ *cobra.Command, _ []string) {

	},
}

func Run() {

	configNetwork()

	if err := RootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Executing root command failed")
	}
}

func configNetwork() {
	routing.ConfigLookClient(host, port)
}

func init() {
	RootCmd.AddCommand(cmdVersion)
	RootCmd.PersistentFlags().StringVarP(&host, "uplink", "u", "http://127.0.0.1", "(optional) Address of the uplink host (default is 127.0.0.1)")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8383, "(optional) Address of the server (default is 8383)")

	RootCmd.AddCommand(queryCmd)
	queryCmd.Flags().StringVarP(&file, "file", "f", "", "(required) file you are looking for")
	queryCmd.Flags().StringVarP(&system, "system", "s", "this", "(optional) only look at the given the system for the file")

	RootCmd.AddCommand(reportCmd)
	RootCmd.AddCommand(systemCmd)
}
