package Music

import (
	"bytes"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
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

func GetDuration(filePath string) string {
	cmd := exec.Command("ffprobe", "-i", filePath, "-show_entries", "format=duration", "-v", "quiet", "-of", "csv=p=0")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return "--:--"
	}

	duration, err := strconv.ParseFloat(strings.TrimSpace(out.String()), 64)
	if err != nil {
		return "--:--"
	}

	minutes := int(duration) / 60
	seconds := int(duration) % 60
	return strconv.Itoa(minutes) + ":" + strconv.Itoa(seconds)
}
