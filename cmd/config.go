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
	configCommand = "config"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   configCommand,
	Short: "Get /v1/config",
	Long:  "Returns the current configuration of a golook server",
	Run: func(_ *cobra.Command, _ []string) {
		result, err := client.GetConfig()
		failOnError(err, "Cannot retrieve configuration of golook server")
		fmt.Println(result)
	},
}

func init() {
	RootCmd.AddCommand(configCmd)
}
