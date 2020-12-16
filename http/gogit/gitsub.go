package gogit

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetModuleList(dir string) ([]string, error) {
	root, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	path, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}

	fmt.Println(path)
	// todo fix ...
	return nil, nil
}
