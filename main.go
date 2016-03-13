package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/cobra"
)

func main() {
	envCmd := &cobra.Command{
		Use:   "env",
		Short: "export AWS variables from an AWS_PROFILE",
		Run: func(cmd *cobra.Command, args []string) {
			if os.Getenv("AWS_PROFILE") != "" {
				profile := os.Getenv("AWS_PROFILE")

				credsValue, err := credentials.NewSharedCredentials("", profile).Get()

				if err != nil {
					fmt.Printf("No exported. Review your AWS credentials for %s", profile)
					os.Exit(1)
				}

				fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", credsValue.AccessKeyID)
				fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", credsValue.SecretAccessKey)

				fmt.Printf("# Run this command to set the variables:\n")
				fmt.Printf("eval $(aws-profile env)\n")
			}
		},
	}
	unenvCmd := &cobra.Command{
		Use:   "unenv",
		Short: "unset AWS variables",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("unset AWS_ACCESS_KEY_ID\n")
			fmt.Printf("unset AWS_SECRET_ACCESS_KEY\n")
			fmt.Printf("unset AWS_PROFILE\n")

			fmt.Printf("# Run this command to unset the variables:\n")
			fmt.Printf("eval $(aws-profile unenv)\n")
		},
	}

	rootCmd := &cobra.Command{Use: "aws-profile"}
	rootCmd.AddCommand(envCmd, unenvCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
