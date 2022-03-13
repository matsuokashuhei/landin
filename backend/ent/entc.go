//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate(
		"./schema",
		&gen.Config{},
		entc.Extensions(ex),
		entc.FeatureNames("sql/versioned-migration"),
	); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
