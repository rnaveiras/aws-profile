# aws-profile

Tired to use software that need `AWS_` variables but doesn't support AWS profiles. Here a simple and dirty hack that will save you a few types.

Ensure that you shell is properly config and exist a exported profile that you want to use e.g. `export AWS_PROFILE=wadus`
Then just simple `eval $(aws-profile env)` will set `AWS_ACCESS_KEY_ID` and `AWS_SECRET_ACCESS_KEY`

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
