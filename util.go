// Package reachability is a generic utility package for the entire codebase. Ideally very smol.
package reachability

import (
	"log"
	"os"
)

// EnvDef is a simple env or default function
func EnvDef(key, def string) (o string) {
	o = os.Getenv(key)
	if o == "" {
		o = def
	}

	return
}

// B or Bytes or Fatal
func B(b []byte, err error) []byte {
	if err != nil {
		log.Fatalln(err)
	}

	return b
}
