package grss

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/chyroc/grss/internal/helper"
	"github.com/chyroc/grss/internal/routers"
)

func updateReadme() error {
	items := []string{}
	for _, v := range routers.Get() {
		if len(v.Args) == 0 {
			v.Args = []map[string]string{nil}
		}
		for _, args := range v.Args {
			source, err := v.Source(args)
			if err != nil {
				panic(err)
			}
			path := helper.JoinArgsURL(v.Path, args)
			items = append(items, fmt.Sprintf("- [%s](/xml%s.xml): %s", path, path, source.Title))
		}
	}

	data := fmt.Sprintf(`# Git RSS

## Routers

%s
`, strings.Join(items, "\n"))
	return ioutil.WriteFile("./index.md", []byte(data), 0o666)
}
