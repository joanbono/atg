package structs

import "time"

type AssumeRole struct {
	Credentials     Credential       `json:"Credentials"`
	AssumedRoleUser AssumedRoleUsers `json:"AssumedRoleUser"`
}

type Credential struct {
	AccessKeyID     string    `json:"AccessKeyId"`
	SecretAccessKey string    `json:"SecretAccessKey"`
	SessionToken    string    `json:"SessionToken"`
	Expiration      time.Time `json:"Expiration"`
}

type AssumedRoleUsers struct {
	AssumedRoleID string `json:"AssumedRoleId"`
	Arn           string `json:"Arn"`
}
