package cmd

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/spf13/cobra"

	"github.com/ofstudio/djay-tools/djay"
)

// cpCmd represents the cp command
var cpCmd = &cobra.Command{
	Use:   "cp playlist.csv",
	Args:  cobra.ExactArgs(1),
	Short: "Copy playlist files into a directory",
	Long: `Copy playlist files into a directory.

First export the playlist from djay as a CSV file and use it as input.

By default, the files will be copied to the current directory with the default prefix
based on the song's position in the playlist and BPM.

Example: "001 (128 bpm) Some Song.mp3"

Prefix can be customized using the "--prefix" or "-p" flag.
`,
	Example: `
Copy playlist files to specified directory:
  djay-tools cp -o /path/to/output /path/to/playlist.csv

Copy playlist files to the current directory:
  djay-tools cp playlist.csv

With bpm as prefix:
  djay-tools cp -p bpm playlist.csv
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// Read the command line flags
		outDir, _ := cmd.Flags().GetString("output")
		prefix := cmd.Flag("prefix").Value.String()
		prefixFn, ok := cpPrefixMapFn[prefix]
		if !ok {
			return fmt.Errorf("invalid prefix option: %q", prefix)
		}

		// Read the playlist
		cmd.Printf("Reading playlist from %q\n", args[0])
		f, err := os.Open(args[0])
		if err != nil {
			return fmt.Errorf("failed to open %q: %w", args[0], err)
		}
		playlist, err := djay.ParsePlaylist(f)
		if err != nil {
			return fmt.Errorf("failed to parse %q: %w", args[0], err)
		}

		// Copy playlist files
		cmd.Printf("Copying files to %q\n\n", outDir)
		n := 0
		for i, song := range playlist {
			// Skip non-local file sources
			if song.Source != djay.FileSource {
				cmd.Printf("- Skipping %s song: %s - %s\n", song.Source, song.Artist, song.Title)
				continue
			}

			// Copy the file
			n++
			dst := path.Join(outDir, prefixFn(i+1, song))
			if err = copyFile(song.FullPath, dst); err != nil {
				return fmt.Errorf("failed to copy from %q to %q: %w", song.FullPath, dst, err)
			}
			fmt.Printf("+ Copied %q\n", dst)
		}
		fmt.Printf("\nCopied %d files\n", n)

		return nil
	},
}

func init() {
	rootCmd.AddCommand(cpCmd)
	cpCmd.Flags().StringP(
		"output",
		"o",
		".",
		"Output directory",
	)
	cpCmd.Flags().StringP(
		"prefix",
		"p",
		"full",
		`File name prefix. 
Available values: "full", "pos", "bpm", "none"`,
	)
}

// copyFile copies a file from src to dst.
func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	//goland:noinspection ALL
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	//goland:noinspection ALL
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	return nil
}

// cpPrefixFn is a function that returns a file name prefix for a song.
type cpPrefixFn func(int, djay.Song) string

// cpPrefixMapFn is a map of available file name prefix functions.
var cpPrefixMapFn = map[string]cpPrefixFn{
	"full": func(i int, song djay.Song) string {
		return fmt.Sprintf("%03d (%03d bpm) %s", i, int(song.BPM), song.Filename)
	},
	"pos": func(i int, song djay.Song) string {
		return fmt.Sprintf("%03d %s", i, song.Filename)
	},
	"bpm": func(_ int, song djay.Song) string {
		return fmt.Sprintf("(%03d bpm) %s", int(song.BPM), song.Filename)
	},
	"none": func(_ int, song djay.Song) string {
		return song.Filename
	},
}
