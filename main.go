package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/ultreme/pipotron/pipotron"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	dictFile, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal("failed to open file: %v", err)
	}

	var dict pipotron.Dict
	if err = yaml.Unmarshal(dictFile, &dict); err != nil {
		log.Fatal("failed to unmarshal yaml: %v", err)
	}

	out, err := pipotron.Generate(&dict)
	if err != nil {
		log.Fatal("failed to generate %v", err)
	}

	fmt.Println(out)
}
