package main // import "moul.io/pipotron"

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"

	"github.com/gohugoio/hugo/common/maps"
	yaml "gopkg.in/yaml.v2"
	"moul.io/pipotron/dict"
	"moul.io/pipotron/pipotron"
	"moul.io/srand"
)

func main() {
	rand.Seed(srand.Fast())

	if len(os.Args) < 2 {
		for _, file := range dict.Box.List() {
			if strings.HasSuffix(file, ".yml") {
				fmt.Println(file[:len(file)-4])
			}
		}
		return
	}

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

	var context pipotron.Context
	context.Scratch = maps.NewScratch()
	if err = yaml.Unmarshal(dictFile, &context.Dict); err != nil {
		log.Fatal("failed to unmarshal yaml: %v", err)
	}

	out, err := pipotron.Generate(&context)
	if err != nil {
		log.Fatal("failed to generate %v", err)
	}

	fmt.Println(out)
}
