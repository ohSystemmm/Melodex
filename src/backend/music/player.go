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

// Initialize VLC with no video output and quiet mode
func Init() {
	if err := vlc.Init("--no-video", "--quiet"); err != nil {
		log.Fatalf("Error initializing VLC: %v", err)
	}

	var err error
	player, err = vlc.NewPlayer()
	if err != nil {
		log.Fatalf("Error creating VLC player: %v", err)
	}

	if muted, err := player.IsMuted(); err != nil {
		log.Printf("Error checking mute state: %v", err)
	} else if muted {
		if err := player.SetMute(false); err != nil {
			log.Printf("Error unmuting player: %v", err)
		}
	}

	volume, err := player.Volume()
	if err != nil {
		log.Printf("Error getting volume: %v", err)
		return
	}

	if volume == 0 {
		if err := player.SetVolume(100); err != nil {
			log.Printf("Error setting volume: %v", err)
		}
	}
}

// loads and plays songs
func PlaySong(song string) {
	if player.IsPlaying() {
		player.Stop()
	}

	var err error
	media, err = player.LoadMediaFromPath(song)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error loading song:", err)
	}

	err = player.Play()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error playing song:", err)
	}
}

// toggles pause and play
func PauseSong() {
	if IsPlaying() {
		player.SetPause(true)
	} else {
		player.SetPause(false)
	}
}

func IsPlaying() bool {
	return player.IsPlaying()
}

func GetSongPosition() (position float32, err error) {
	position, err = player.MediaPosition()
	if err != nil {
		log.Println("Error getting position:", err)
		return
	}
	return position, nil
}

// sets the volume (0–100)
func SetVolume(volume int) {
	if err := player.SetVolume(volume); err != nil {
		log.Println("Failed to set volume:", err)
	}
}

// toggles mute
func SetMute() {
	isMuted, err := player.IsMuted()
	if err != nil {
		log.Println("Failed to get mute status:", err)
		return
	}

	err = player.SetMute(!isMuted)
	if err != nil {
		log.Println("Failed to toggle mute:", err)
	}
}

// sets playback position (0.0–1.0) - float32 btw
func SetMediaPosition(pos float32) {
	if err := player.SetMediaPosition(pos); err != nil {
		log.Println("Failed to set position:", err)
	}
}

// Sleeps and stops playback after time is up
func Sleep(duration time.Duration) {
	log.Printf("Sleeping for %s...\n", duration)
	time.Sleep(duration)

	log.Println("Sleep complete. Stopping playback.")
	Stop()
}

// stops playback
func Stop() {
	if player != nil {
		player.Stop()
	}
	if media != nil {
		media.Release()
	}
}

// cleanup to release resources
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
		log.Println("Error getting durations:", err)
		return nil, err
	}

	err = saveCache(playlistName, durations)
	if err != nil {
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

	if err := os.MkdirAll(appCacheDir, 0755); err != nil {
		return "", err
	}

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
	err = json.Unmarshal(data, &rows)
	if err != nil {
		return nil, err
	}

	return rows, nil
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
	cmd := exec.Command("ffprobe", "-v", "error", "-show_entries", "format=duration", "-of", "default=noprint_wrappers=1:nokey=1", songPath)
	output, err := cmd.Output()
	if err != nil {
		log.Printf("Error running ffprobe for %s: %v", songPath, err)
		return "", err
	}

	durationStr := strings.TrimSpace(string(output))
	seconds, err := strconv.ParseFloat(durationStr, 64)
	if err != nil {
		log.Printf("Error parsing duration for %s: %v", songPath, err)
		return "", err
	}

	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60
	seconds = math.Mod(seconds, 60)

	switch {
	case hours > 0:
		return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, int(seconds)), nil
	case minutes > 0:
		return fmt.Sprintf("     %02d:%02d", minutes, int(seconds)), nil
	default:
		return fmt.Sprintf("          %02d", int(seconds)), nil
	}
}

func getAllDurations(directory string) ([]table.Row, error) {
	files, err := os.ReadDir(directory)
	if err != nil {
		log.Printf("Error reading directory %s: %v", directory, err)
		return nil, err
	}

	durationList := make([]table.Row, 0)

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		songPath := fmt.Sprintf("%s/%s", directory, file.Name())

		duration, err := getDuration(songPath)
		if err != nil {
			log.Printf("Error getting duration for %s: %v", songPath, err)
			continue
		}

		row := table.Row{
			file.Name(),
			duration,
		}
		durationList = append(durationList, row)
	}

	return durationList, nil
}
