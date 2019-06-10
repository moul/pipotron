package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ultreme/pipotron/dict"
	"github.com/ultreme/pipotron/pipotron"
	yaml "gopkg.in/yaml.v2"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	dictName := os.Args[1]
	if !strings.HasSuffix(dictName, ".yml") {
		dictName += ".yml"
	}

	// try to get the file content from the packr box first, and fallback to the disk
	dictFile, err := dict.Box.Find(dictName)
	if err != nil {
		dictFile, err = ioutil.ReadFile(dictName)
		if err != nil {
			log.Fatal("failed to open file: %v", err)
		}
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
