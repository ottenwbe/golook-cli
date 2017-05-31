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
	"github.com/spf13/cobra"
)

const (
	logCommand = "log"
)

// logCmd represents the log command
var logCmd = &cobra.Command{
	Use:   logCommand,
	Short: "Log output from the server",
	Long:  "Get the server's log. Might take a while",
	Run: func(cmd *cobra.Command, args []string) {
		result, err := client.GetLog()
		failOnError(err, "Cannot retrieve Log")
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(logCmd)
}
