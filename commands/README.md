[![hoarder logo](http://nano-assets.gopagoda.io/readme-headers/hoarder.png)](http://nanobox.io/open-source#hoarder)  
[![Build Status](https://travis-ci.org/nanopack/hoarder.svg)](https://travis-ci.org/nanopack/hoarder)
[![GoDoc](https://godoc.org/github.com/nanopack/hoarder?status.svg)](https://godoc.org/github.com/nanopack/hoarder)

# Hoarder

Hoarder is a simple, api-driven, storage system for storing anything for cloud based deployment services.

## CLI Commands:

```
hoarder - data storage

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

## Server Usage Example:
```
$ ./hoarder --server
```
or
```
$ ./hoarder -c config.json
```

>config.json
>```json
{
  "backend": "file:///var/db/hoarder",
  "host": "127.0.0.1",
  "insecure": false,
  "log-level": "debug",
  "port": "7410",
  "server": true,
  "token": "secret"
}
```

## Common Issues

Error: `Failed to make request - Post https://localhost:7410/blobs/test: http: server gave HTTP response to HTTPS client` means that hoarder is started using `http`, but the client is configured to use `https` to talk to it.  
Solution: Set `-H 'http://localhost:7410'` to use `http`  


## Client Usage Example:

#### add data
```sh
$ hoarder add -k 'small-file' -d 'testfile'
# 'small-file' created!

$ cat test2 | hoarder add -k 'med-file' -d -
# 'med-file' created!

$ hoarder add -k 'med-file' -f 'large-file'
# 'large-file' created!
```

#### show blobs
```sh
$ hoarder list
#[
#  {
#    "Name": "small-file",
#    "Size": 8,
#    "ModTime": "2016-07-25T22:15:35.848159653Z"
#  },
#  {
#    "Name": "med-file",
#    "Size": 10850121,
#    "ModTime": "2016-07-25T23:06:17.578408182Z"
#  },
#  {
#    "Name": "large-file",
#    "Size": 108501210,
#    "ModTime": "2016-07-25T21:50:55.056072975Z"
#  }
#]
```

#### show blob
```sh
$ hoarder show -k small-file
# testfile

$ hoarder show -k large-file -f large-file2
# Finished writing file
$ du -b large-file*
# 108501210 large-file
# 108501210 large-file2
```

#### remove blob
```
$ hoarder remove -k large-file
# 'large-file' destroyed!
```

#### update blob
```
$ hoarder update -k small-file -d 'eliftset'
# 'small-file' created!
```

[![opensource logo](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
