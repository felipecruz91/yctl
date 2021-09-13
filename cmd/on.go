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
	"log"

	"github.com/akominch/yeelight"
	"github.com/spf13/cobra"
)

// onCmd represents the on command
var onCmd = &cobra.Command{
	Use:   "on",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		ip, _ := cmd.Flags().GetString("ip")
		config := yeelight.BulbConfig{
			Ip:     ip,
			Effect: yeelight.Smooth,
		}
		bulb := yeelight.New(config)

		isOn, err := bulb.IsOn()
		if err != nil {
			log.Fatalf("failed to check whether yeelight device is on or off: %s\n", err.Error())
		}

		if isOn {
			log.Println("➖ lighyeelight device is already on")
			return
		}

		_, err = bulb.TurnOn()
		if err != nil {
			log.Fatalf("failed to turn on the yeelight device: %s\n", err.Error())
		}

		log.Println("💡 yeelight device switched on")
	},
}

func init() {
	rootCmd.AddCommand(onCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// onCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// onCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
