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
	"github.com/ottenwbe/golook/client"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

const (
	programName = "golook-cli"
	version     = "v0.1.0-dev"
)

/*
RootCmd is the root of all commands
*/
var RootCmd = &cobra.Command{
	Use:   programName,
	Short: "Golook cli",
	Long:  "Cli for the golook distributed file search",
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: fmt.Sprintf("Print the version number of %s", programName),
	Long:  fmt.Sprintf("All software has versions: This is %s's", programName),
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Print(version)
	},
}

/*
Run is called to trigger a user specified command
*/
func Run() {

	if err := RootCmd.Execute(); err != nil {
		log.WithError(err).Fatal("Executing root command failed")
	}
}

func init() {
	RootCmd.AddCommand(versionCmd)
	RootCmd.PersistentFlags().StringVarP(&client.Host, "host", "u", "http://127.0.0.1:8383", "(optional) Address of the uplink host")
}
