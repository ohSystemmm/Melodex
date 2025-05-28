package music

import (
	"encoding/json"
	"fmt"
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

func Init() {
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatalf("Error initializing VLC: %v", err)
	}

	var err error
	player, err = vlc.NewPlayer()
	if err != nil {
		log.Fatalf("Error creating VLC player: %v", err)
	}

	if muted, err := player.IsMuted(); err == nil && muted {
		_ = player.SetMute(false)
	}

	if volume, err := player.Volume(); err == nil && volume == 0 {
		_ = player.SetVolume(100)
	}
}

func PlaySong(song string) {
	if player.IsPlaying() {
		player.Stop()
	}

	_, err := player.LoadMediaFromPath(song)
	if err != nil {
		log.Fatalf("Error loading song: %v", err)
	}

	err = player.Play()
	if err := player.Play(); err != nil {
		log.Fatalf("Error playing song: %v", err)
	}
}
func PauseSong() {
	player.SetPause(IsPlaying())
}
func IsPlaying() bool {
	return player.IsPlaying()
}
func GetSongPosition() (float32, error) {
	return player.MediaPosition()
}
func SetVolume(volume int) {
	_ = player.SetVolume(volume)
}
func SetMediaPosition(pos float32) {
	_ = player.SetMediaPosition(pos)
}
func Sleep(duration time.Duration) {
	log.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)
	log.Println("Sleep complete. Stopping playback.")
	Stop()
}
func Stop() {
	if player != nil {
		player.Stop()
	}
	if media != nil {
		media.Release()
	}
}
func Cleanup() {
	if player != nil {
		player.Release()
	}
	vlc.Release()
}

func Songlist(playlistPath string) ([]table.Row, error) {
	playlistName := filepath.Base(playlistPath)

	if cachedRows, err := loadCache(playlistName); err == nil {
		log.Println("Loaded playlist from cache:", playlistName)
		return cachedRows, nil
	}

	durations, err := getAllDurations(playlistPath)
	if err != nil {
		return nil, err
	}

	if err := saveCache(playlistName, durations); err != nil {
		log.Println("Warning: could not save playlist cache:", err)
	}

	return durations, nil
}
func getCachePath(playlistName string) (string, error) {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		return "", err
	}

	appCacheDir := filepath.Join(cacheDir, "melodex")
	_ = os.MkdirAll(appCacheDir, 0755)

	return filepath.Join(appCacheDir, playlistName+".cache"), nil
}
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
	return rows, json.Unmarshal(data, &rows)
}
func saveCache(filename string, rows []table.Row) error {
	path, err := getCachePath(filename)
	if err != nil {
		return err
	}

	data, err := json.Marshal(rows)
	if err != nil {
		return err
	}

	return os.WriteFile(path, data, 0644)
}
func getDuration(songPath string) (string, error) {
	output, err := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", songPath).Output()
	if err != nil {
		return "", err
	}

	seconds, err := strconv.ParseFloat(strings.TrimSpace(string(output)), 64)
	if err != nil {
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
func getAllDurations(directory string) ([]table.Row, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	var durationList []table.Row
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		duration, err := getDuration(filepath.Join(directory, file.Name()))
		if err == nil {
			durationList = append(durationList, table.Row{file.Name(), duration})
		}
	}

	return durationList, nil
}
