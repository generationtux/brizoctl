# brizoctl

CLI tool for [Brizo](https://github.com/generationtux/brizo)

## Usage
```
$ brizoctl list apps
$ brizoctl get apps [UUID]
```

## Install
Download the [latest release](https://github.com/generationtux/brizoctl/releases) for your system and move to a location in your $PATH.
```
$ mv ~/Downloads/brizoctl /usr/local/bin/brizoctl
$ chmod +x /usr/loca/bin/brizoctl
```

## Configure
Create a config file at `~/.brizo.json` with the following properties
```json
{
  "endpoint": "https://brizo.example.com",
  "token": "[ACCESS_TOKEN]"
}
```

## Building a release
- Set the version in `main.go`.
- Build binaries and create [Github Release](https://github.com/generationtux/brizoctl/releases)

A helper script for building a new release is provided in the `bin` directory. Run `./bin/buildRelease [VERSION]` and binaries will be compiled for OSX and Linux. The results are placed in the `dist` directory.
