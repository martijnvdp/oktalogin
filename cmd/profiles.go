/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/martijnxd/oktalogin/oktalogin"
	"github.com/spf13/cobra"
)

// profilesCmd represents the profiles command
var profilesCmd = &cobra.Command{
	Use:   "profiles",
	Short: "add/list/edit okta login profiles",
	Long: `Add/list/edit okta login profiles stored in the oktalogin.yaml config file located in the home folder. For example:

oktalogin profiles add
oktalogin profiles list.
`,
	Run: func(cmd *cobra.Command, args []string) {
		//l, _ := cmd.Flags().GetBool("list")
		a, _ := cmd.Flags().GetBool("add")

		if a {
			oktalogin.AddProfiles()
		} else {
			oktalogin.ListProfiles()
		}
	},
}

func init() {
	rootCmd.AddCommand(profilesCmd)
	profilesCmd.Flags().BoolP("list", "l", false, "List oktalogin profiles")
	profilesCmd.Flags().BoolP("add", "a", false, "Add oktalogin profiles")
}
