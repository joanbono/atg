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

func CheckStdin() bool {
	if termutil.Isatty(os.Stdin.Fd()) {
		return false
	}
	return true
}

func ParseStdin(stdin string) {
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

	KeyPrinter(data.Credentials.AccessKeyID, data.Credentials.SecretAccessKey, data.Credentials.SessionToken)

}

func ParseJSON(jsonfile string) {

	jsonFile, err := os.Open(jsonfile)
	CheckErr(err)

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data structs.AssumeRole
	json.Unmarshal(byteValue, &data)
	//pp.Print(data)
	KeyPrinter(data.Credentials.AccessKeyID, data.Credentials.SecretAccessKey, data.Credentials.SessionToken)
}

func KeyPrinter(access, secret, token string) {
	fmt.Printf("export AWS_ACCESS_KEY_ID=\"%v\"\n", access)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=\"%v\"\n", secret)
	fmt.Printf("export AWS_SESSION_TOKEN=\"%v\"\n", token)
}

// CheckErr will handle errors
// for the entire program
func CheckErr(err error) {
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
}
