/*

pushover-cli

A simple cli client for sending Pushover push messages.

Copyright © 2022 Staffan Linnaeus staffan.linnaeus@gmail.com

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/gregdel/pushover"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pushover-cli",
	Short: "A simple CLI tool to send Pushover messages.",
	Long: `	
You need to have an API token and user key, which you can get from https://pushover.net.
In your shell, set these variables prior to running the binary: 
export PUSHOVER_API_TOKEN=<your token here>
export PUSHOVER_USER_KEY=<your user key here>`,

	Run: func(cmd *cobra.Command, args []string) {

		token, user := getEnvVars()

		msg, _ := cmd.Flags().GetString("message")
		msgTitle, _ := cmd.Flags().GetString("title")
		msgPrio, _ := cmd.Flags().GetString("prio")
		msgPrioInt, _ := strconv.Atoi(msgPrio)
		msgUrl, _ := cmd.Flags().GetString("url")
		msgUrlTitle, _ := cmd.Flags().GetString("urltitle")
		msgTimestamp, _ := cmd.Flags().GetString("timestamp")
		msgTimestampInt64, _ := strconv.ParseInt(msgTimestamp, 10, 64)
		msgRetry, _ := cmd.Flags().GetString("retry")
		msgRetryTimeDuration, _ := time.ParseDuration(msgRetry)
		msgExpire, _ := cmd.Flags().GetString("expire")
		msgExpireTimeDuration, _ := time.ParseDuration(msgExpire)
		msgCallbackURL, _ := cmd.Flags().GetString("callbackurl")
		msgDevice, _ := cmd.Flags().GetString("device")
		msgSound, _ := cmd.Flags().GetString("sound")
		msgHTML, _ := cmd.Flags().GetString("html")
		msgHTMBool, _ := strconv.ParseBool(msgHTML)
		msgMonospace, _ := cmd.Flags().GetString("monospace")
		msgMonospaceBool, _ := strconv.ParseBool(msgMonospace)

		if msg == "" {
			log.Println("ERROR No message given!")
			os.Exit(1)
		}

		if msgPrioInt == 2 && msgRetry == "" && msgExpire == "" {
			log.Println("ERROR: Emergency Priority but --retry and --expire not set!")
			os.Exit(1)
		}

		push := pushover.New(token)
		recipient := pushover.NewRecipient(user)

		PushMessage := pushover.Message{
			Message:     msg,
			Title:       msgTitle,
			Priority:    msgPrioInt,
			URL:         msgUrl,
			URLTitle:    msgUrlTitle,
			Timestamp:   msgTimestampInt64,
			Retry:       msgRetryTimeDuration,
			Expire:      msgExpireTimeDuration,
			CallbackURL: msgCallbackURL,
			DeviceName:  msgDevice,
			Sound:       msgSound,
			HTML:        msgHTMBool,
			Monospace:   msgMonospaceBool}

		response, err := push.SendMessage(
			&PushMessage,
			recipient)

		check(err)
		log.Println(response)

	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

	rootCmd.PersistentFlags().String("message", "", "Your message (required)")
	rootCmd.PersistentFlags().String("title", "", "A title for your message (optional)")
	rootCmd.PersistentFlags().String("prio", "", "Your message's priority (-2 - 2) (optional)")
	rootCmd.PersistentFlags().String("url", "", "An URL to send with your message (optional)")
	rootCmd.PersistentFlags().String("urltitle", "", "A title for your URL (optional)")
	rootCmd.PersistentFlags().String("timestamp", "", "Include a timestamp in your message (optional)")
	rootCmd.PersistentFlags().String("retry", "", "Time to retry (optional)")
	rootCmd.PersistentFlags().String("expire", "", "Time to expire (optional)")
	rootCmd.PersistentFlags().String("callbackurl", "", "A callback URL for your message. (optional)")
	rootCmd.PersistentFlags().String("device", "", "Device to send your message to (optional)(if left blank all devices will be messaged)")
	rootCmd.PersistentFlags().String("sound", "", "A sound for your message (optional)")
	rootCmd.PersistentFlags().String("html", "", "Parse HTML in message (True/False)(optional)")
	rootCmd.PersistentFlags().String("monospace", "", "Monospaced (True/False)(optional)")

}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getEnvVars() (APIToken string, userKey string) {

	APIToken, token_present := os.LookupEnv("PUSHOVER_API_TOKEN")
	if !token_present {
		fmt.Println("ERROR: PUSHOVER_API_TOKEN not set.")
	}

	userKey, key_present := os.LookupEnv("PUSHOVER_USER_KEY")
	if !key_present {
		fmt.Println("ERROR: PUSHOVER_USER_KEY not set.")
	}

	if !token_present || !key_present {
		os.Exit(1)
	}

	return APIToken, userKey
}
