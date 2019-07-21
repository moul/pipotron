package pipotron // import "moul.io/pipotron/pipotron"

import (
	"encoding/json"

	"github.com/gohugoio/hugo/common/maps"
)

type Context struct {
	Scratch *maps.Scratch
	Dict    map[string][]string
}

// FIXME: an option should be a string or a string with a weight

func (d Context) String() string {
	out, _ := json.MarshalIndent(d, "", "  ")
	return string(out)
}
