package djay

import (
	"fmt"
	"net/url"
	"path"
	"strconv"
	"strings"
	"time"
)

// Song represents a song in a playlist
type Song struct {
	Title    string
	Artist   string
	Album    string
	Duration time.Duration
	BPM      float64
	Key      string
	URL      string
	Source   string // file or apple-music
	FullPath string // full path to the file
	Filename string // file name
}

const (
	FileSource       = "file"
	AppleMusicSource = "apple-music"
)

// ParseSong parses a song from a djay CSV record
func ParseSong(record []string) (*Song, error) {
	if len(record) != 7 {
		return nil, fmt.Errorf("want 7 fields, got %d", len(record))
	}

	song := &Song{
		Title:  record[0],
		Artist: record[1],
		Album:  record[2],
		Key:    record[5],
		URL:    record[6],
	}

	var err error
	if song.Duration, err = parseDuration(record[3]); err != nil {
		return nil, fmt.Errorf("failed to parse duration %q: %w", record[3], err)
	}

	if record[4] == "" {
		record[4] = "0"
	}
	if song.BPM, err = strconv.ParseFloat(record[4], 64); err != nil {
		return nil, fmt.Errorf("failed to parse bpm %q: %w", record[4], err)
	}

	u, err := url.Parse(song.URL)
	if err != nil {
		return nil, fmt.Errorf("failed to parse url %q: %w", song.URL, err)
	}

	song.Source = u.Scheme
	if song.Source == FileSource {
		song.FullPath = u.Path
		_, song.Filename = path.Split(song.FullPath)
	}

	return song, nil
}

func parseDuration(str string) (time.Duration, error) {
	parts := strings.Split(str, ":")
	if len(parts) != 2 {
		return 0, fmt.Errorf("invalid time format: %s", str)
	}

	m, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("failed to parse miutes: %s", str)
	}

	s, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, fmt.Errorf("failed to parse seconds: %s", str)
	}

	return time.Duration(m)*time.Minute + time.Duration(s)*time.Second, nil
}
