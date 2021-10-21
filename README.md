# Save

`go get github.com/gabeduke/save`

This utility saves aws credentials from the environment to the AWS Credentials file. For use with AWS-Keycloak.

How does this work? AWS-Keycloak exports AWS session tokens to the environment. This tool simply creates an AWS session from the environment and saves it to the normal file. While the token is not expired, aws calls will work without the keycloak wrapper.

Usage:

```bash
$ aws-keycloak -p [profile] -- save
INFO[0000] Write AWS Session to file successful!
```
