package mvfiles

import (
	"fmt"
	"os"
	"path/filepath"
)

func MakeDirs(path string) error {
	if f, err := os.Stat(path); os.IsNotExist(err) || !f.IsDir() {
		fmt.Println("Make dir", path)
		if err := os.MkdirAll(path, 0777); err != nil {
			return err
		}
	}
	return nil
}

func MoveFiles(files []string, moveto string) {
	fmt.Println("Move to --> ", moveto)
	if len(files) == 0 {
		fmt.Println("  [-] Target File does not exist.")
		return
	}
	for _, path := range files {
		current, _ := filepath.Rel(".", path)
		filename := filepath.Base(current)
		moveto := filepath.Join(moveto, filename)
		if f, err := os.Stat(moveto); err == nil {
			fmt.Println("  [=] Already exsists:", f.Name())
			continue
		}
		if err := os.Rename(current, moveto); err != nil {
			fmt.Println("  [x] Error: ", filename)
			fmt.Println(err)
		} else {
			fmt.Println("  [o] Moved: ", filename)
		}
	}
}
