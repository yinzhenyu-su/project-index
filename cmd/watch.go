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
	"project-index/service"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// watchCmd represents the watch command
var watchCmd = &cobra.Command{
	Use:   "watch -t <gitlab api token> -d <watch directory>",
	Short: "展示Gitlab项目状态",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Println(cmd.UsageString())
		} else {
			directory, dirErr := cmd.Flags().GetString("directory")
			token, tokenErr := cmd.Flags().GetString("token")
			if dirErr != nil || tokenErr != nil {
				cmd.Println(cmd.UsageString())
			} else {
				service.Router(token, directory)
			}

		}
	},
}

func init() {
	rootCmd.AddCommand(watchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// watchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	watchCmd.Flags().StringP("directory", "d", "", "watch directory")
	watchCmd.Flags().StringP("token", "t", "", "gitlab api token")
	var dir = viper.GetString("watch.directory")
	var token = viper.GetString("watch.token")
	if token != "" && dir != "" {
		service.Router(token, dir)
	}
}
