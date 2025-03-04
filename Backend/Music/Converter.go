package Music

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func isSupportedExtension(ext string) bool {
	supported := []string{".wav", ".flac", ".aac", ".ogg", ".m4a", ".wma", ".mp3"}
	for _, s := range supported {
		if strings.EqualFold(ext, s) {
			return true
		}
	}
	return false
}

func ConvertToMP3(sourceDir, destDir string) ([]string, error) {
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		return nil, fmt.Errorf("failed to read source directory: %w", err)
	}

	var mp3Files []string

	for _, file := range files {
		ext := filepath.Ext(file.Name())
		if isSupportedExtension(ext) {
			sourcePath := filepath.Join(sourceDir, file.Name())
			destFileName := strings.TrimSuffix(file.Name(), ext) + ".mp3"
			destPath := filepath.Join(destDir, destFileName)

			cmd := exec.Command("ffmpeg", "-i", sourcePath, "-ar", "44100", "-ac", "2", destPath)
			err := cmd.Run()
			if err != nil {
				log.Printf("Failed to convert %s: %v\n", file.Name(), err)
				continue
			}

			fmt.Printf("Converted %s to %s\n", file.Name(), destFileName)
			mp3Files = append(mp3Files, destPath)
		}
	}

	return mp3Files, nil
}
