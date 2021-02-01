/*
 Licensed to the Apache Software Foundation (ASF) under one
 or more contributor license agreements.  See the NOTICE file
 distributed with this work for additional information
 regarding copyright ownership.  The ASF licenses this file
 to you under the Apache License, Version 2.0 (the
 "License"); you may not use this file except in compliance
 with the License.  You may obtain a copy of the License at
   http://www.apache.org/licenses/LICENSE-2.0
 Unless required by applicable law or agreed to in writing,
 software distributed under the License is distributed on an
 "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 KIND, either express or implied.  See the License for the
 specific language governing permissions and limitations
 under the License.
*/

package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/andrew-d/go-termutil"
	"github.com/gocaio/atg/modules/structs"
)

// CheckStdin will check if there is
// something in the stdin buffer
func CheckStdin() bool {
	if termutil.Isatty(os.Stdin.Fd()) {
		return false
	}
	return true
}

// ParseStdin will parse the output from
// the aws sts assume-role command from
// stdin piping ATG
func ParseStdin() (string, string, string) {
	var buf bytes.Buffer
	reader := bufio.NewReader(os.Stdin)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				buf.WriteString(line)
				break // end of the input
			} else {
				fmt.Println(err.Error())
				os.Exit(1) // something bad happened
			}
		}
		buf.WriteString(line)

	}

	var data structs.AssumeRole
	err := json.Unmarshal(buf.Bytes(), &data)
	CheckErr(err)

	return data.Credentials.AccessKeyID, data.Credentials.SecretAccessKey, data.Credentials.SessionToken
}

// ParseJSON will spect the output json from the
// aws sts assume-role command
func ParseJSON(jsonfile string) (string, string, string) {

	jsonFile, err := os.Open(jsonfile)
	CheckErr(err)

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data structs.AssumeRole
	json.Unmarshal(byteValue, &data)
	return data.Credentials.AccessKeyID, data.Credentials.SecretAccessKey, data.Credentials.SessionToken
}

// KeyPrinter will print the environment
// variables depending of the chosen shell
func KeyPrinter(access, secret, token, shell string) {
	switch shell {
	case "cmd":
		fmt.Printf("set AWS_ACCESS_KEY_ID=\"%v\"\n", access)
		fmt.Printf("set AWS_SECRET_ACCESS_KEY=\"%v\"\n", secret)
		fmt.Printf("set AWS_SESSION_TOKEN=\"%v\"\n", token)
	case "fish":
		fmt.Printf("set -x AWS_ACCESS_KEY_ID=\"%v\"\n", access)
		fmt.Printf("set -x AWS_SECRET_ACCESS_KEY=\"%v\"\n", secret)
		fmt.Printf("set -x AWS_SESSION_TOKEN=\"%v\"\n", token)
	case "powershell":
		fmt.Printf("$Env:AWS_ACCESS_KEY_ID=\"%v\"\n", access)
		fmt.Printf("$Env:AWS_SECRET_ACCESS_KEY=\"%v\"\n", secret)
		fmt.Printf("$Env:AWS_SESSION_TOKEN=\"%v\"\n", token)
	default:
		fmt.Printf("export AWS_ACCESS_KEY_ID=\"%v\"\n", access)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=\"%v\"\n", secret)
		fmt.Printf("export AWS_SESSION_TOKEN=\"%v\"\n", token)
	}

}

// CheckErr will handle errors
// for the entire program
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
