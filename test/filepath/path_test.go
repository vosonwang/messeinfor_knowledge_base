package filepath

import (
	"testing"
	"path/filepath"
	"os"
	"log"
	"fmt"
)

func TestPath(t *testing.T)  {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	execpath, err := os.Executable() // 获得程序路径


	fmt.Println(dir)
	fmt.Println(execpath)
}