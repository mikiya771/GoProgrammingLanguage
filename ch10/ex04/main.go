package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		return
	}
	partialList, _ := excuteGoList(args[0])
	allList, _ := excuteGoList("...")
	depNames := map[string]bool{}
	for _, pdep := range partialList {
		for _, depTitle := range pdep.Deps {
			depNames[depTitle] = true
		}
	}
	for _, dep := range allList {
		if depNames[dep.Name] {
			fmt.Printf("%v\n", dep)
		}
	}

}

type packageInfo struct {
	Dir  string
	Name string
	Deps []string
}

func excuteGoList(path string) ([]packageInfo, error) {
	args := []string{"list", "-json", path}
	cmd := exec.Command("go", args...)
	out, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}
	defer out.Close()
	if err = cmd.Start(); err != nil {
		return nil, err
	}
	decoder := json.NewDecoder(out)
	var infos []packageInfo
	for {
		var info packageInfo
		err := decoder.Decode(&info)
		if err != nil {
			if err != io.EOF {
				log.Printf("err: %v\n", err)
			}
			return infos, nil
		}
		infos = append(infos, info)
	}

}
