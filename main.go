package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	yaml "gopkg.in/yaml.v2"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	dictFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("failed to open file: %v", err)
	}

	var dict Dict
	if err = yaml.Unmarshal(dictFile, &dict); err != nil {
		log.Fatal("failed to unmarshal yaml: %v", err)
	}

	out, err := executeTemplate("{{pick .output}}", &dict)
	if err != nil {
		log.Fatal("template error: %v", err)
	}
	fmt.Println(strings.Title(out))
}
