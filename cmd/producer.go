/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

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
	"github.com/sirupsen/logrus"
	"kafkapoc/services"

	"github.com/spf13/cobra"
)

var message string
var topicName string

// producerCmd represents the producer command
var producerCmd = &cobra.Command{
	Use:   "producer",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		check()
		services.Produce(brokers, message, topicName)
	},
}

func init() {
	rootCmd.AddCommand(producerCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// producerCmd.PersistentFlags().String("foo", "", "A help for foo")
	producerCmd.PersistentFlags().StringVar(&message, "message", "Hello world", "Message data to be sent via the kafka topic")
	producerCmd.PersistentFlags().StringVar(&topicName, "topic", "test", "Topic to which message has to be sent")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// producerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func check() {
	exists := false
	// checking if provided topic exists in predefined topic
	for _, tp := range topicList {
		if topicName == tp {
			exists=true
			logrus.Infof("Found %s", topicName)
		}
	}
	if !exists {
		logrus.Errorf("Topic used %s does not exist in the pre-defined topics", topicName)
	}
}