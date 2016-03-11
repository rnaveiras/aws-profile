package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/cobra"
)

func main() {

	var envCmd = &cobra.Command{
		Use:   "env",
		Short: "export AWS_ variables from a AWS_PROFILE",
		Run: func(cmd *cobra.Command, args []string) {
			if os.Getenv("AWS_PROFILE") != "" {
				profile := os.Getenv("AWS_PROFILE")

				credsValue, err := credentials.NewSharedCredentials("", profile).Get()

				if err != nil {
					panic("oops")
				}

				fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", credsValue.AccessKeyID)
				fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", credsValue.SecretAccessKey)

				fmt.Printf("# Run this command to configure your shell:\n")
				fmt.Printf("eval $(aws-profile env)\n")
			}
		},
	}

	var unenvCmd = &cobra.Command{
		Use:   "unenv",
		Short: "unset AWS variables",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("unset AWS_ACCESS_KEY_ID\n")
			fmt.Printf("unset AWS_SECRET_ACCESS_KEY\n")
			fmt.Printf("unset AWS_PROFILE\n")

			fmt.Printf("# Run this command to configure your shell:\n")
			fmt.Printf("eval $(aws-profile unenv)\n")
		},
	}

	var rootCmd = &cobra.Command{Use: "aws-profile"}
	rootCmd.AddCommand(envCmd, unenvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
