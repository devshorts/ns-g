package main

import (
	"github.com/paradoxical/ns-go/core/model"
	"github.com/tkrajina/typescriptify-golang-structs/typescriptify"
)

func main() {
	converter := typescriptify.New()
	converter.CreateFromMethod = true
	converter.Indent = "    "
	converter.BackupDir = ""

	converter.Add(model.Bank{})

	err := converter.ConvertToFile("example_output.ts")
	if err != nil {
		panic(err.Error())
	}
}

