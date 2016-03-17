[![hoarder logo](http://nano-assets.gopagoda.io/readme-headers/hoarder.png)](http://nanobox.io/open-source#hoarder)  
[![Build Status](https://travis-ci.org/nanopack/hoarder.svg)](https://travis-ci.org/nanopack/hoarder)

Hoarder is a simple, api-driven, storage system for storing anything for cloud based deployment services.

## Usage

#### As a server
To start hoarder as a server run:

`hoarder --server`

An optional config file can also be passed on startup:

`hoarder --server --config /path/to/config`

#### As a CLI

Simply run `hoarder <COMMAND>`

`hoarder` or `hoarder -h` will show usage and a list of commands:

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

To configure hoarder, a config.yml file can be passed with --config. Configuration read in through a file will overwrite the same configuration specified by a flag. If no config file is passed, and no flags are set, reasonable defaults will be used.

```yml
clean_after : 0                           # the age that data is deemed garbage (seconds)
backend     : "file:///var/db/hoarder"    # the pluggable backend the api will use for storage
host        : 127.0.0.1                   # the connection host
insecure    : true                        # connect insecurely
log_type    : "stdout"                    # the type of logging
log_file    : "/var/log/hoarder.log"      # the location of the log file
log_level   : "info"                      # the output log level
port        : "7410"                      # the connection port
token       : ""                          # the secure token used to connect with (no auth by default)
```

## API:

```
| Method |     Route     | Functionality |
------------------------------------------
| GET    | /blobs"      | List all blobs
| GET    | /blobs/{:id} | Retrieve a blob
| POST   | /blobs/{:id} | Publish a new blob
| PUT    | /blobs/{:id} | Update an existing blob
| DELETE | /blobs/{:id} | Remove an existing blob
| HEAD   | /blobs"      | Retrieve file information for all blobs
| HEAD   | /blobs/{:id} | Retrieve file information about a blob
```

#### Examples

##### ping:
```
$ curl -k https://localhost:7410/ping
=> pong
```

##### create:
```
$ curl -k https://localhost:7410/blobs/test -d "data"
=> 'test' created!
```

##### get:
```
$ curl -k https://localhost:7410/blobs/test
=> data
```

##### get head:
```
$ curl -k https://localhost:7410/blobs/test -I
=> HTTP/1.1 200 OK
=> Content-Length: 4
=> Date: Tue, 01 Mar 2016 21:14:28 UTC
=> Last-Modified: Tue, 01 Mar 2016 21:13:57 UTC
=> Content-Type: text/plain; charset=utf-8
```

##### update:
```
$ curl -k https://localhost:7410/blobs/test -d "new data" -X PUT
=> 'test' created!
```

##### list:
```
$ curl -k https://localhost:7410/blobs
=> [{"Name":"test","Size":4,"ModTime":"2016-03-01T21:13:57.534706044Z"}]
```

##### delete:
```
$ curl -k https://localhost:7410/blobs/test -X DELETE
=> 'test' destroyed!
```

*Note*: all examples are run without auth. If auth was enabled when the server was started then an additional header needs to be present:

`-H "x-auth-token: TOKEN"`

## Data

Hoarder simply stores whatever data you give it as a string. So you can literally store whatever you want as long is it can be "stringified".

Some examples of what data could look like when creating a new blob:

##### string
```
$ curl -k https://localhost:7410/blobs/test -d "some string"
```

##### JSON
``` json
$ curl -k https://localhost:7410/blobs/test -d "{\"key\":\"value\"}"
```

When it retrieves data it might look like the following:
```
{
	"Name": "test",
	"Size": 4,
	"ModTime": "2016-03-01T21:13:57.534706044Z"
}
```

## Contributing

Contributions to hoarder are welcome and encouraged. Hoarder is a [Nanobox](https://nanobox.io) project and contributions should follow the [Nanobox Contribution Process & Guidelines](https://docs.nanobox.io/contributing/).

[![open source](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
