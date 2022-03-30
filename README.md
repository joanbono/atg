# atg
Assume to go is a utility to export variables from `assume-role` and be ready to go.

From this:

```json
{
    "Credentials": {
        "AccessKeyId": "AAAAAAAAAAAAAAA",
        "SecretAccessKey": "BBBBBBBBBBBBBBBBBB",
        "SessionToken": "CCCCCCCCCCCCCCCCCCCCCCCC",
        "Expiration": "2021-01-27T13:10:39+00:00"
    },
    "AssumedRoleUser": {
        "AssumedRoleId": "DDDD:name",
        "Arn": "ARN"
    }
}
```

To this:

```sh
export AWS_ACCESS_KEY_ID="AAAAAAAAAAAAAAA"
export AWS_SECRET_ACCESS_KEY="BBBBBBBBBBBBBBBBBB"
export AWS_SESSION_TOKEN="CCCCCCCCCCCCCCCCCCCCCCCC"
```

As an alternative for people using [`fish`](https://fishshell.com/), [`cmd`](https://en.wikipedia.org/wiki/Cmd.exe) or [`PowerShell`](https://en.wikipedia.org/wiki/PowerShell) (defaults to `bash`), it is possible to set which output is preferred. As example for `PowerShell`:

```powershell
PS> aws sts assume-role --role-arn ${ROLE_ARN} --role-session-name ${SESSION_NAME} --external-id ${EXTERNAL_ID} | atg -powershell

$Env:AWS_ACCESS_KEY_ID="AAAAAAAAAAAAAAA"
$Env:AWS_SECRET_ACCESS_KEY="BBBBBBBBBBBBBBBBBB"
$Env:AWS_SESSION_TOKEN="CCCCCCCCCCCCCCCCCCCCCCCC"
```

Read from `stdin` and import 

```sh
$> aws sts get-caller-identity
{
    "UserId": "AIDA11111111111111111",
    "Account": "111111111111",
    "Arn": "arn:aws:iam::111111111111:user/myuser"
}

$> eval $(aws sts assume-role --role-arn ${ROLE_ARN} --role-session-name ${SESSION_NAME} --external-id ${EXTERNAL_ID} | atg)
$> aws sts get-caller-identity
{
    "UserId": "AROA22222222222222222:${SESSION_NAME}",
    "Account": "222222222222",
    "Arn": "arn:aws:sts::222222222222:assumed-role/role-name/${SESSION_NAME}"
}
```

Read from `role.json` file for `fish`:

```sh
~> aws sts get-caller-identity
{
    "UserId": "AIDA11111111111111111",
    "Account": "111111111111",
    "Arn": "arn:aws:iam::111111111111:user/myuser"
}

~> eval $(atg -json role.json -fish)
~> aws sts get-caller-identity
{
    "UserId": "AROA22222222222222222:${SESSION_NAME}",
    "Account": "222222222222",
    "Arn": "arn:aws:sts::222222222222:assumed-role/role-name/${SESSION_NAME}"
}
```

## Using MFA

If an MFA device is [required to authenticate](https://aws.amazon.com/premiumsupport/knowledge-center/authenticate-mfa-cli/), follow these instructions:

```sh
# Account with MFA configured
$ eval $(aws --profile mfa sts get-session-token --serial-number arn:aws:iam::111111111111:mfa/userMFA --token-code 123123 | atg)

# Assume the role from the MFA account to a different account
$ eval $(aws sts assume-role --role-arn arn:aws:iam::222222222222:role/role-to-assume --role-session-name assume-with-mfa | atg)
```
