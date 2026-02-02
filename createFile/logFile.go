package createFile

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
)

func getFile(url, filename, path string) {
	once := sync.Once{}
	once.Do(func() {
		err := os.Mkdir(path, 0755)
		if err != nil {
			return
		}
	})
	file, _ := http.Get(url)
	filePath := filepath.Join(path, filename)
	createFile, err := os.Create(filePath)
	if err != nil {
		return
	}
	_, err = io.Copy(createFile, file.Body)
	if err != nil {
		return
	}

}
