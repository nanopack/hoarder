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
  -g, --clean-after="0": Age data is deemed garbage (seconds)
  -H, --host="127.0.0.1": Hoarder hostname/IP
  -i, --insecure[=true]: Disable tls key checking
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
# following are all the available configuration file options (all default values are shown)
clean_after : 0                    # the age that data is deemed garbage (seconds)
collect     : false                # required to enable garbage collector from config file
connection  : "file://"            # the pluggable backend the api will use for storage
host        : 127.0.0.1            # the connection host
insecure    : true                 # connect insecurly
log_level   : "info"               # the output log level
port        : "7410"               # the connection port
token       : "TOKEN"              # the secury token used to connect with
```

## Routes:

```
| Method |     Route     | Functionality |
------------------------------------------
| HEAD   | /blobs/{blob} | Retrieve file information about a blob
| GET    | /blobs/{blob} | Retrieve a blob
| GET    | /blobs"       | List all blobs
| POST   | /blobs/{blob} | Publish a new blob
| PUT    | /blobs/{blob} | Update an existing blob
| DELETE | /blobs/{blob} | Remove an existing blob
```

### Contributing

Contributions to the hoarder project are welcome and encouraged. Hoarder is a [Nanobox](https://nanobox.io) project and contributions should follow the [Nanobox Contribution Process & Guidelines](https://docs.nanobox.io/contributing/).

### Licence

Mozilla Public License Version 2.0

[![open source](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
