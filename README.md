# WLED Backup (wled-backup)

[![standard-readme compliant](https://img.shields.io/badge/readme%20style-standard-brightgreen.svg)](https://github.com/RichardLitt/standard-readme)
[![.github/workflows/ci.yaml](https://github.com/thibmaek/wled-backup/actions/workflows/ci.yaml/badge.svg)](https://github.com/thibmaek/wled-backup/actions/workflows/ci.yaml)

<!-- Short Description: 
- Must not have its own title.
- Must be less than 120 characters.
- Must not start with `> `
- Must be on its own line.
- Must match the description in the packager manager's `description` field.
- Must match GitHub's description (if on GitHub).
 -->
Simple CLI tool to backup presets & configuration from a WLED device.

## Table of Contents
<!-- - Must link to all Markdown sections in the file.
- Must start with the next section; do not include the title or Table of Contents headings.
- Must be at least one-depth: must capture all `##` headings. -->
- [Background](#background)
- [Install](#install)
- [Usage](#usage)
- [License](#license)

## Background

I put a lot of effort in getting my presets right for all of my WLED devices. The cheap ESP8266's I buy from AliExpress aren't resistant to failure however and reflashing WLED on a new ESP means that I would have to recreate my presets from scratch.
Luckily WLED's JSON API exposes an endpoint to retrieve your presets & configuration as JSON files.

This tool does nothing more than calling that endpoint and writing it to a file on the system in parallel. This way you can backup your files via e.g a cronjob.

It's also possible to reupload these files across host and syncing the presets if you'd want that. This tool however does not currently support that.

## Install

Download the latest release binary for your platform from the [releases page](https://github.com/thibmaek/wled-backup/releases). Optionally rename it if you want but this is not required.

## Usage

Run the `export` command and pass a host or a list of (comma separated) hosts.

```console
# A single host
$ ./wled-backup_linux_x64 export --hosts=192.168.1.12

# Multiple hosts
$ ./wled-backup_linux_x64 export --hosts=192.168.1.12,192.168.1.177

# mDNS works too, you can mix and match
$ ./wled-backup_linux_x64 export --hosts=wled-tv.local,192.168.1.177
```

By default this will output the backup files in the current folder. You can optionally specify an output directory with the `--outputDir` flag:

```console
$ ./wled-backup_linux_x64 export --hosts=192.168.1.12,192.168.1.177  --outputDir=/home/user/wled_backups
```

## License

Unlicense

For more info, see [license file](./LICENSE)
