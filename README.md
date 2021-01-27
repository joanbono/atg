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

Read from `role.json` file:

```sh
$> aws sts get-caller-identity
{
    "UserId": "AIDA11111111111111111",
    "Account": "111111111111",
    "Arn": "arn:aws:iam::111111111111:user/myuser"
}

$> eval $(atg -json role.json)
$> aws sts get-caller-identity
{
    "UserId": "AROA22222222222222222:${SESSION_NAME}",
    "Account": "222222222222",
    "Arn": "arn:aws:sts::222222222222:assumed-role/role-name/${SESSION_NAME}"
}
```
