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

Read from `stdin`

```sh
aws sts assume-role --role-arn ${ROLE_ARN} --role-session-name ${SESSION_NAME} --external-id ${EXTERNAL_ID} | atg
```

Read from `role.json` file:

```sh
atg -json role.json
```
