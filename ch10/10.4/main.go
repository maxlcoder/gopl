package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

type pkg struct {
	ImportPath string
	Deps       []string
}

func pkgs(args []string) ([]pkg, error) {
	out, err := exec.Command("go",
		append([]string{"list", "-json"},
			args...)...).Output()
	if err != nil {
		return nil, err
	}
	var pkgs []pkg
	dec := json.NewDecoder(bytes.NewBuffer(out))
	for {
		var pkg pkg
		if err := dec.Decode(&pkg); err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		pkgs = append(pkgs, pkg)
	}
	return pkgs, nil
}

func main() {
	ips, err := pkgs(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	deps, err := pkgs([]string{"all"})
	if err != nil {
		log.Fatal(err)
	}
	for _, dep := range deps {
		for _, depip := range dep.Deps {
			for _, ip := range ips {
				if ip.ImportPath == depip {
					fmt.Println(dep.ImportPath)
				}
			}
		}
	}
}
