package Music

import (
	"io/ioutil"
	"path/filepath"
)

func GetFilesWithExtensions(dir string) (map[string]string, error) {
	filesMap := make(map[string]string)
	ext := []string{".wav", ".flac", ".aac", ".ogg", ".m4a", ".wma", ".mp3"}

	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			for _, ext := range ext {
				if filepath.Ext(file.Name()) == ext {
					fullPath := filepath.Join(dir, file.Name())
					filesMap[file.Name()] = fullPath
					break
				}
			}
		}
	}

	return filesMap, nil
}
