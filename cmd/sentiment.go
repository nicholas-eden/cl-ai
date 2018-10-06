// Copyright © 2018 Nicholas Eden <nicholas.eden@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"github.com/nicholas-eden/clai/sentiment"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

// sentimentCmd represents the sentiment command
var sentimentCmd = &cobra.Command{
	Use:   "sentiment",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: func(cmd *cobra.Command, args []string) error {
		fi, _ := os.Stdin.Stat() // get the FileInfo struct describing the standard input.

		if (fi.Mode() & os.ModeCharDevice) == 0 {
			return nil
		}

		if len(args) < 1 {
			return fmt.Errorf("requires at least %d arg(s), only received %d", 1, len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		key := viper.Get("algorithmia.key").(string)
		sentiment.Execute(args, key)
	},
}

func init() {
	rootCmd.AddCommand(sentimentCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sentimentCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sentimentCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
