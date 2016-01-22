// Package decompose lets you arbitrarily de-nest JSON configuration into multiple files.
package decompose

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// Decompose starting at `root` path, returning a JSON blob. If `root` is a JSON file itself
// and not a directory, its contents is simply returned.
func Decompose(root string) ([]byte, error) {
	m := make(map[string]interface{})

	// TODO: actually support arbitrary depths :)

	info, err := os.Stat(root)
	if err != nil {
		return nil, err
	}

	if info.Mode().IsRegular() {
		return ioutil.ReadFile(root)
	}

	err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		base, err := filepath.Rel(root, path)
		if err != nil {
			return err
		}

		if !info.Mode().IsRegular() {
			return nil
		}

		if !strings.HasSuffix(path, ".json") {
			return nil
		}

		f, err := os.Open(path)
		if err != nil {
			return err
		}

		var v map[string]interface{}
		if err := json.NewDecoder(f).Decode(&v); err != nil {
			return err
		}

		name := strings.Replace(base, ".json", "", -1)

		if "index" == name {
			for k, v := range v {
				m[k] = v
			}
			return nil
		}

		m[name] = v
		return nil
	})

	if err != nil {
		return nil, err
	}

	return json.MarshalIndent(m, "", "  ")
}
