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
	systemCommand = "system"
)

// infoCmd represents the info command
var systemCmd = &cobra.Command{
	Use:   systemCommand,
	Short: "Call /v1/system",
	Long:  "Get system information about the golook server",
	Run: func(_ *cobra.Command, _ []string) {
		result, err := client.GetSystem()
		failOnError(err, "Cannot retrieve System")
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(systemCmd)
}
