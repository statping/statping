package main

import (
	"github.com/hunterlong/statup/plugin"
)

type Example struct {
	*plugin.Plugin
}

var example = &Example{&plugin.Plugin{
	Name:        "Example",
	Description: "This is an example plugin",
}}

func main() {

}

func (e *Example) Select() *plugin.Plugin {
	return e.Plugin
}
