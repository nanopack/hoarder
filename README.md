# Hoarder

[![Build Status](https://travis-ci.org/nanopack/hoarder.svg)](https://travis-ci.org/nanopack/hoarder)
[![GoDoc](https://godoc.org/github.com/nanopack/hoarder?status.svg)](https://godoc.org/github.com/nanopack/hoarder)

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
  hoarder [flags]
  hoarder [command]

Available Commands:
  add         Add file to hoarder storage
  list        List all files in hoarder storage
  remove      Remove a file from hoarder storage
  show        Display a file from the hoarder storage
  update      Update a file in hoarder

Flags:
  -b, --backend string       Hoarder backend (default "file:///var/db/hoarder")
  -g, --clean-after uint     Age, in seconds, after which data is deemed garbage (default 0)
  -c, --config string        Path to config file (with extension)
  -H, --listen-addr string   Hoarder listen uri (scheme defaults to https) (default "https://127.0.0.1:7410")
      --log-level string     Output level of logs (TRACE, DEBUG, INFO, WARN, ERROR, FATAL) (default "INFO")
  -s, --server               Run hoarder as a server
  -t, --token string         Auth token used when connecting to a secure Hoarder
  -v, --version              Display the current version of this CLI

Use "hoarder [command] --help" for more information about a command.
```

## Configuration

To configure hoarder, a config.yml file can be passed with --config. Configuration read in through a file will overwrite the same configuration specified by a flag. If no config file is passed, and no flags are set, reasonable defaults will be used.

```yml
backend     : "file:///var/db/hoarder"    # the pluggable backend the api will use for storage
listen-addr : "https://127.0.0.1:7410"    # the connection host uri (scheme defaults to https)
log-level   : "INFO"                      # the output log level (trace, debug, info, warn, error, fatal)
server      : false                       # run as a server
token       : ""                          # the secure token used to connect with (no auth by default)
```

## API:

```
| Method |     Route     | Functionality |
------------------------------------------
| GET    | /blobs/{:id} | Retrieve a blob
| HEAD   | /blobs/{:id} | Retrieve file information about a blob
| POST   | /blobs/{:id} | Publish a new blob
| PUT    | /blobs/{:id} | Update an existing blob
| DELETE | /blobs/{:id} | Remove an existing blob
| GET    | /blobs       | List all blobs
| HEAD   | /blobs       | Retrieve file information for all blobs
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

`-H "X-AUTH-TOKEN: TOKEN"`

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

## Todo

## Contributing

Contributions to hoarder are welcome and encouraged. Hoarder is a [Nanobox](https://nanobox.io) project and contributions should follow the [Nanobox Contribution Process & Guidelines](https://docs.nanobox.io/contributing/).

[![open source](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
