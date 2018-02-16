//go:generate pigeon -o main_gen.go hosts.peg

package main

import (
	"fmt"
	"log"
	"os"
)

func toIfaceSlice(v interface{}) []interface{} {
	if v == nil {
		return nil
	}
	return v.([]interface{})
}

func toStringSlice(v interface{}) []string {
	if v == nil {
		return nil
	}
	vv := v.([]interface{})
	strs := make([]string, len(vv))
	for i := range vv {
		s, ok := vv[i].(string)
		if !ok {
			continue
		}
		strs = append(strs, s)
	}
	return strs
}

func toASTHostLineSlice(v interface{}) []ASTHostLine {
	if v == nil {
		return nil
	}
	vv := v.([]interface{})
	lns := make([]ASTHostLine, 0, len(vv))
	for i := range vv {
		l, ok := vv[i].(ASTHostLine)
		if !ok {
			continue
		}
		lns = append(lns, l)
	}
	return lns
}

func main() {
	log.Println("Parsing hosts")
	h, err := ParseFile(os.Args[1])

	if err != nil {
		log.Println("Got Error:", err)
		os.Exit(1)
	}

	hlns, ok := h.(ASTHosts)
	if !ok {
		log.Fatal("nil")
	}

	for _, hl := range hlns {
		fmt.Printf("%s", hl.String())

		// NOTE: it is also doable via type switch,
		// however let the entity that's being printed be
		// in charge of it's representation

		// switch hlClass := hl.(type) {
		// case ASTComment:

		// case ASTHost:
		// }
	}
}
