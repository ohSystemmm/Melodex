package music

import (
	"Melodex/src/logger"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	vlc "github.com/adrg/libvlc-go/v3"
	"github.com/charmbracelet/bubbles/table"
)

var (
	player *vlc.Player
	media  *vlc.Media
)

// FUNCTION: Initializes VLC
func Init() {
	var err error

	if err = vlc.Init("--no-video", "--quiet"); err != nil {
		logger.Log.Errorf("Error initializing VLC: %v", err)
	}

	player, err = vlc.NewPlayer()
	if err != nil {
		logger.Log.Errorf("Error creating VLC player: %v", err)
	}

	muted, err := player.IsMuted()
	if err != nil {
		logger.Log.Errorf("Error checking if player is muted: %v", err)
	} else if muted {
		if err = player.SetMute(false); err != nil {
			logger.Log.Errorf("Error unmuting player: %v", err)
		}
	}

	volume, err := player.Volume()
	if err != nil {
		logger.Log.Errorf("Error checking player volume: %v", err)
	} else if volume == 0 {
		if err = player.SetVolume(100); err != nil {
			logger.Log.Errorf("Error setting volume to 100%%: %v", err)
		}
	}
}

// FUNCTION: Plays the provided song
func PlaySong(songPath string) bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			logger.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}

	_, err := player.LoadMediaFromPath(songPath)
	if err != nil {
		logger.Log.Errorf("Error loading song: %v", err)
		return false
	}

	err = player.Play()
	if err = player.Play(); err != nil {
		logger.Log.Errorf("Error playing song: %v", err)
		return false
	}
	return true
}

// FUNCTION: Pauses song
func PauseSong() {
	if err := player.SetPause(player.IsPlaying()); err != nil {
		logger.Log.Errorf("Error toggling pause: %v", err)
	}
}

// FUNCTION: Returns the current state of the song
func IsPlaying() bool {
	return player.IsPlaying()
}

// FUNCTION: Gets the current Song Position
func CurrentSongPosition() float32 {
	position, err := player.MediaPosition()
	if err != nil {
		logger.Log.Errorf("Error getting media position: %v", err)
		return 0.0
	}
	return position
}

// FUNCTION: Sets the volume
func SetVolume(volume int) {
	err := player.SetVolume(volume)
	if err != nil {
		logger.Log.Errorf("Error setting volume to %d: %v", volume, err)
	}
}

// FUNCTION: Sets Media Position
func SetMediaPosition(position float32) bool {
	err := player.SetMediaPosition(position)
	if err != nil {
		logger.Log.Errorf("Error setting media position: %v", err)
		return false
	}
	return true
}

func Sleep(duration time.Duration) {
	log.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)
	log.Println("Sleep complete. Stopping playback.")
	Stop()
}

// FUNCTION: Stops the player
func Stop() bool {
	if player.IsPlaying() {
		if err := player.Stop(); err != nil {
			logger.Log.Errorf("Error stopping player: %v", err)
			return false
		}
	}
	return true
}

// FUNCTION: Releases all resources
func Cleanup() {
	if player != nil {
		if err := player.Release(); err != nil {
			logger.Log.Errorf("Error releasing player: %v", err)
		}
	}

	if media != nil {
		if err := vlc.Release(); err != nil {
			logger.Log.Errorf("Error releasing VLC: %v", err)
		}
	}
}

// FUNCTION: Generates the song list for the music list element
func SongList(playlistPath string) []table.Row {
	playlistName := filepath.Base(playlistPath)

	rows, err := loadCache(playlistName)
	if err != nil {
		logger.Log.Errorf("Error loading cache file: %v", err)
		rows = []table.Row{}
	}

	durations, err := getAllDurations(playlistPath)
	if err != nil {
		logger.Log.Errorf("Error getting song durations: %v", err)
		return rows
	}

	if err = saveCache(playlistName, durations); err != nil {
		logger.Log.Warnf("Warning: Could not save playlist cache: %v", err)
	}

	return durations
}

// FUNCTION: Gets cache path
func getCachePath(playlistName string) (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		logger.Log.Errorf("Error getting cache dir: %v", err)
		return "", err
	}

	appCacheDir := filepath.Join(cacheDir, "melodex")
	err = os.MkdirAll(appCacheDir, fs.ModePerm)
	if err != nil {
		logger.Log.Errorf("Error creating cache dir: %v", err)
		return "", err
	}

	return filepath.Join(appCacheDir, playlistName+".cache"), nil
}

// FUNCTION: Loads cached data
func loadCache(playlistName string) ([]table.Row, error) {
	path, err := getCachePath(playlistName)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var rows []table.Row
	if err := json.Unmarshal(data, &rows); err != nil {
		return nil, err
	}

	return rows, nil
}

// FUNCTION: Saves cache
func saveCache(filename string, rows []table.Row) error {
	path, err := getCachePath(filename)
	if err != nil {
		logger.Log.Errorf("Error getting cache dir: %v", err)
		return err
	}

	data, err := json.Marshal(rows)
	if err != nil {
		logger.Log.Errorf("Error marshaling cache rows: %v", err)
		return err
	}

	return os.WriteFile(path, data, 0644)
}

// FUNCTION: Gets duration of a song
func getDuration(songPath string) (string, error) {
	output, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", songPath).Output()
	if err != nil {
		logger.Log.Errorf("Error getting duration: %v", err)
		return "", err
	}

	seconds, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
		logger.Log.Errorf("Error parsing duration: %v", err)
		return "", err
	}

	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60
	seconds = math.Mod(seconds, 60)

	switch {
	case hours > 0:
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, int(seconds)), nil
	case minutes > 0:
		return fmt.Sprintf("%02d:%02d", minutes, int(seconds)), nil
	default:
		return fmt.Sprintf("%02d", int(seconds)), nil
	}
}

// FUNCTION: Gets durations of all songs in a directory
func getAllDurations(directory string) ([]table.Row, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		logger.Log.Errorf("Error reading directory: %v", err)
		return nil, err
	}

	var durationList []table.Row
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		duration, err := getDuration(filepath.Join(directory, file.Name()))
		if err != nil {
			logger.Log.Warnf("Error fetching duration for %s: %v", file.Name(), err)
		} else {
			durationList = append(durationList, table.Row{file.Name(), duration})
		}
	}

	return durationList, nil
}
