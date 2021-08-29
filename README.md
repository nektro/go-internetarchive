# InternetArchive

![loc](https://sloc.xyz/github/nektro/internetarchive)
[![license](https://img.shields.io/github/license/nektro/internetarchive.svg)](https://github.com/nektro/internetarchive/blob/master/LICENSE)
[![discord](https://img.shields.io/discord/551971034593755159.svg?logo=discord)](https://discord.gg/P6Y4zQC)
[![paypal](https://img.shields.io/badge/donate-paypal-009cdf?logo=paypal)](https://paypal.me/nektro)
[![circleci](https://circleci.com/gh/nektro/internetarchive.svg?style=svg)](https://circleci.com/gh/nektro/internetarchive)
[![release](https://img.shields.io/github/v/release/nektro/internetarchive)](https://github.com/nektro/internetarchive/releases/latest)
[![goreportcard](https://goreportcard.com/badge/github.com/nektro/internetarchive)](https://goreportcard.com/report/github.com/nektro/internetarchive)
[![downloads](https://img.shields.io/github/downloads/nektro/internetarchive/total.svg)](https://github.com/nektro/internetarchive/releases)

`ia` is a command-line interface for interacting with https://archive.org/ written in Golang. It is an alternative to the offical Python client available here https://github.com/jjjake/internetarchive.

## Usage

- `ia`
```
Usage:
  ia [command]

Available Commands:
  download    download an item or collection
  help        Help about any command
  metadata    retrieve metadata for items and collections

Flags:
  -h, --help                help for ia
      --mbpp-bar-gradient   Enabling this will make the bar gradient from red/yellow/green.
```

- `ia download`
```
Usage:
  ia download {item_name} [flags]

Flags:
  -c, --concurrency int   number of concurrent download jobs to run at once (default 10)
      --dense             when enabled, stores items based on their creation date
  -h, --help              help for download
      --only-meta         when enabled, only saves _meta.xml files
  -o, --save-dir string    (default "./data")
      --resume            When enabled, performs a deeper check for item completion
```

## Built With
- Golang 1.14
- https://github.com/spf13/cobra
- https://github.com/nektro/go-util
- https://github.com/PuerkitoBio/goquery

## License
MIT
