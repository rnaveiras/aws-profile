# aws-profile

Tired of using software that need `AWS_` variables but doesn't support AWS profiles. Here a simple and dirty hack that will save you a few keystrokes.

Make sure your shell is properly configured and the profile you want to use exists and is exported. e.g. `export AWS_PROFILE=wadus`
Then just run `eval $(aws-profile env)` and it will set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

## Usage

```
Usage:
  aws-profile [command]

Available Commands:
  env         export AWS_ variables from a AWS_PROFILE
  unenv       unset AWS variables

Flags:
  -h, --help   help for aws-profile

Use "aws-profile [command] --help" for more information about a command.
```
