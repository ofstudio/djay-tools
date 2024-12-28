package djay

import (
	"encoding/csv"
	"fmt"
	"io"
)

// Playlist represents a list of songs
type Playlist []Song

// ParsePlaylist parses a djay CSV playlist from a reader
func ParsePlaylist(r io.Reader) (Playlist, error) {
	var playlist Playlist

	csvReader := csv.NewReader(r)
	csvReader.Comma = ','

	n := 0
	for {
		n++
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		// Skip the header
		if n == 1 {
			continue
		}

		song, err := ParseSong(record)
		if err != nil {
			return nil, fmt.Errorf("failed to parse song at line %d: %w", len(playlist)+1, err)
		}
		playlist = append(playlist, *song)
	}

	return playlist, nil
}
