[![hoarder logo](http://nano-assets.gopagoda.io/readme-headers/hoarder.png)](http://nanobox.io/open-source#hoarder)
[![Build Status](https://travis-ci.org/nanopack/hoarder.svg)](https://travis-ci.org/nanopack/hoarder)

Hoarder is a simple, api-driven, storage system for storing anything for cloud
based deployment services.

## Status

Nearly Complete/Experimental

## Usage

Hoarder can be used in 2 ways:

##### As a server
To start hoarder as a server run `hoarder --server`. An optional config file can
be passed with `--config /path/to/config`

##### As a CLI
Simply run `hoarder <COMMAND>`

`hoarder` or `hoarder -h` will provide help information including usage and a
list of commands:

```
Usage:
   [flags]
   [command]

Available Commands:
  add         Add file to hoarder storage
  list        List all files in hoarder storage
  remove      Remove a file from hoarder storage
  show        Display a file from the hoarder storage
  update      Update a file in hoarder

Flags:
  -b, --backend="filesystem": Hoarder backend driver
      --config="": Path to config options
  -H, --host="0.0.0.0": Hoarder hostname/IP
  -i, --insecure[=false]: Disable tls key checking
      --log-level="info": Hoarder output log level
  -p, --port="7410": Hoarder port
      --server[=false]: Run hoader as a server
  -t, --token="TOKEN": Hoarder auth token
  -v, --version[=false]: Display the current version of this CLI

Use " [command] --help" for more information about a command.
```

## Configuration

To configure hoarder, a config file can be passed with --config. If no config file
is passed reasonable defaults will be used.

```
# following are all the available configuration options (all default values are shown)
Backend    : "filesystem"         # the pluggable backend the api will use for storage
Driver     : backends.Filesystem  # the actual backend driver
GCInterval : 0                    # the interval between clearning out old storage (disabled by default)
GCAmount   : 0                    # the amount of storage to clear at interval (disabled by default)
Host       : 0.0.0.0              # the connection host
Insecure   : false                # connect insecurly
LogLevel   : "info"               # the output log level
Port       : "7410"               # the connection port
Token      : "TOKEN"              # the secury token used to connect with
```

## Routes:

```
| Method |     Route     | Functionality |
------------------------------------------
| HEAD   | /blobs/{blob} | Retrieve file information about a blob
| GET    | /blobs/{blob} | Retrieve a blob
| GET    | /blobs"       | List all blobs
| POST   | /blobs/{blob} | Publish a New blob
| PUT    | /blobs/{blob} | Update an existing blob
| Delete | /blobs/{blob} | Remove a existing blob
```

## Todo

- tests

### Contributing

Contributions to the hoarder project are welcome and encouraged. Hoarder is a [Nanobox](https://nanobox.io) project and contributions should follow the [Nanobox Contribution Process & Guidelines](https://docs.nanobox.io/contributing/).

### Licence

Mozilla Public License Version 2.0

[![open source](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
