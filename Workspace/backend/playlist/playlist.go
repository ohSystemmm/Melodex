package playlist

import (
	"os"
	"path/filepath"
	"strings"

	// "github.com/hcl/audioduration"

	"github.com/go-audio/audio"
	"github.com/go-audio/wav"
	"github.com/mewkiz/flac"
)

var playlistname string

func AddPlaylist(playlistname string) string {
	return playlistname
}

func GetAllSongs() []string {
	var filesWithExt []string
	ext := ".m4a"
	tempPath := "/home/" + os.Getenv("USER") + "/TempSongs"

	files, err := os.ReadDir(tempPath)
	if err != nil {
		return nil
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ext) {
			filesWithExt = append(filesWithExt, file.Name())
		}
	}

	return filesWithExt
}

// func DurationOfAllSongs(dir string) (map[string]int, error) {
// 	durations := make(map[string]int)
// 	exts := []string{".mp3", ".m4a"} // Supported audio formats

// 	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
// 		if err != nil {
// 			return err
// 		}

// 		if !info.IsDir() {
// 			ext := strings.ToLower(filepath.Ext(path))
// 			for _, supportedExt := range exts {
// 				if ext == supportedExt {
// 					file, err := os.Open(path)
// 					if err != nil {
// 						return err
// 					}
// 					defer file.Close()

// 					var duration int
// 					var err2 error

// 					switch ext {
// 					case ".mp3":
// 						duration, err2 = audioduration.Mp3(file)
// 					case ".m4a":
// 						duration, err2 = audioduration.M4a(file)
// 					case ".wav":
// 						duration, err2 = audioduration.FLAC(file)
// 					}

// 					if err2 != nil {
// 						return err2
// 					}

// 					durations[filepath.Base(path)] = duration
// 					break
// 				}
// 			}
// 		}
// 		return nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}
// 	return durations, nil
// }

// DurationOfAllSongs returns song durations in seconds for supported formats
func DurationOfAllSongs(dir string) (map[string]float64, error) {
	durations := make(map[string]float64)
	supportedExts := []string{".mp3", ".m4a", ".flac", ".wav"}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return err
		}

		ext := strings.ToLower(filepath.Ext(path))
		for _, supportedExt := range supportedExts {
			if ext == supportedExt {
				duration, err := getDuration(path, ext)
				if err != nil {
					return err
				}
				durations[filepath.Base(path)] = duration
				break
			}
		}
		return nil
	})

	return durations, err
}

// getDuration handles format-specific duration extraction
func getDuration(path string, ext string) (float64, error) {
	file, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	switch ext {
	case ".mp3":
		decoder, err := audio.NewDecoder(file)
		if err != nil {
			return 0, err
		}
		return float64(decoder.Length()) / float64(decoder.SampleRate()*2*2), nil

	case ".m4a":
		// M4A/AAC support via go-audio
		decoder := audio.NewDecoder(file)
		if err := decoder.Decode(); err != nil {
			return 0, err
		}
		return float64(decoder.NumFrames()) / float64(decoder.SampleRate()), nil

	case ".flac":
		stream, err := flac.ParseFile(path)
		if err != nil {
			return 0, err
		}
		return float64(stream.Info.NSamples) / float64(stream.Info.SampleRate), nil

	case ".wav":
		decoder := wav.NewDecoder(file)
		if !decoder.IsValidFile() {
			return 0, err
		}
		dur, _ := decoder.Duration()
		return dur.Seconds(), nil

	default:
		return 0, nil
	}
}

func ShufflePlaylist(playlist []string) []string {
	shuffeledplaylist := playlist
	return shuffeledplaylist

}

func RemovePlaylist(config string) bool {
	return false
}

func GetPlaylistName(directory string) string {
	return directory
}
