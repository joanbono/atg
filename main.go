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

package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gocaio/atg/modules/parser"
)

var (
	jsonfile   string
	powershell bool
	cmd        bool
	fish       bool
	bash       bool
)

func init() {
	flag.StringVar(&jsonfile, "json", "", "JSON file to use")
	flag.BoolVar(&bash, "bash", false, "Export bash/zsh environmental variables [default]")
	flag.BoolVar(&cmd, "cmd", false, "Export Windows CMD environmental variables")
	flag.BoolVar(&fish, "fish", false, "Export fish environmental variables")
	flag.BoolVar(&powershell, "powershell", false, "Export Powershell environmental variables")

	flag.Parse()
}

func main() {
	boolElements := [4]bool{bash, cmd, powershell, fish}
	if Count(boolElements) > 1 {
		fmt.Println("Choose one [bash, cmd, fish, powershell]")
		fmt.Println("Use atg -h")
		//flag.PrintDefaults()
		os.Exit(0)
	}

	if jsonfile == "" {
		onStdin := parser.CheckStdin()
		if !onStdin {
			flag.PrintDefaults()
		} else {
			access, secret, token := parser.ParseStdin()
			if cmd == true {
				parser.KeyPrinter(access, secret, token, "cmd")
			} else if fish == true {
				parser.KeyPrinter(access, secret, token, "fish")
			} else if powershell == true {
				parser.KeyPrinter(access, secret, token, "powershell")
			} else {
				parser.KeyPrinter(access, secret, token, "bash")
			}
		}
	} else {
		access, secret, token := parser.ParseJSON(jsonfile)
		if cmd == true {
			parser.KeyPrinter(access, secret, token, "cmd")
		} else if fish == true {
			parser.KeyPrinter(access, secret, token, "fish")
		} else if powershell == true {
			parser.KeyPrinter(access, secret, token, "powershell")
		} else {
			parser.KeyPrinter(access, secret, token, "bash")
		}
	}
}

// Count will count how many
// bool flags are there in total
func Count(boolElements [4]bool) int {
	var True = 0
	for i := range boolElements {
		if boolElements[i] == true {
			True = True + 1
		}
	}
	return True
}
