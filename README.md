# djay-tools

Command-line companion for [Algoriddim djay Pro](https://www.algoriddim.com/djay-pro-mac).

## Features

- Copy playlist files into a directory with customizable prefix based on the song's position in the playlist and BPM.

## Installation

```bash
go install github.com/ofstudio/djay-tools@latest
```

## Usage

```bash
djay-tools [command]`
```

### Flags

```
-h, --help   help for djay-tools
```

## Available Commands

```
cp          Copy playlist files into a directory
help        Help about any command
```

Use `djay-tools [command] --help` for more information about a command.

## `cp` command

Copy playlist files into a directory.

### Usage

```bash
djay-tools cp [flags] playlist.csv 
```

First export the playlist from djay as a CSV file and use it as input.
Only local files will be copied.  Other songs (e.g. from Apple Music) will be skipped.

By default, the files will be copied to the current directory with the default prefix
based on the song's position in the playlist and BPM.

Example: `001 (128 bpm) Some Song.mp3`

Prefix can be customized using the `--prefix` or `-p` flag.


### Flags

```
  -h, --help            help for cp
  -o, --output string   Output directory (default ".")
  -p, --prefix string   File name prefix. Available values: "full", "pos", "bpm", "none" (default "full")
```

### Examples

Copy playlist files to specified directory:

```bash
djay-tools cp -o /path/to/output /path/to/playlist.csv`
```

Copy playlist files to the current directory:

```bash
djay-tools cp playlist.csv
````

With bpm as prefix:

```bash
djay-tools cp -p bpm playlist.csv
```

## License

Apache License 2.0

## Contributing

Feel free to open an issue or a pull request.

## Author

Oleg Fomin — [@ofstudio](https://t.me/ofstudio), [ofstudio@gmail.com](mailto:ofstudio@gmail.com)
