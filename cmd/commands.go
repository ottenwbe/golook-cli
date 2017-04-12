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

	"github.com/ottenwbe/golook/communication"
	"github.com/ottenwbe/golook/routing"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	PROGRAM_NAME = "golook-cli"
	VERSION      = "v0.1.0-dev"
)

var (
	host       string
	port       int
	GolookIfce routing.LookRouter
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
	communication.ConfigLookClient(host, port)
	GolookIfce = routing.NewRouter()
}

func init() {
	RootCmd.AddCommand(cmdVersion)
	RootCmd.PersistentFlags().StringVarP(&host, "uplink", "u", "http://127.0.0.1", "(optional) Address of the uplink host (default is 127.0.0.1)")
	RootCmd.PersistentFlags().IntVarP(&port, "port", "p", 8383, "(optional) Port of the uplink host (default is 8383)")

	RootCmd.AddCommand(systemCmd)
}
