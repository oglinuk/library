package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() && strings.HasSuffix(path, ".md") {
			// TODO: Replace with PEGN-based parser
			fmt.Printf("Rendering %s ...\n", path)
			out := fmt.Sprintf("%s/index.html", filepath.Dir(path))
			err = exec.Command("pandoc", path, "-o", out).Run()
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
}
