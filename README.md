[![hoarder logo](http://nano-assets.gopagoda.io/readme-headers/hoarder.png)](http://nanobox.io/open-source#hoarder)
 [![Build Status](https://travis-ci.org/nanopack/hoarder.svg)](https://travis-ci.org/nanopack/hoarder)

# Hoarder

Hoarder is a simple, api-driven storage system for storing code builds and cached libraries for cloud-based deployment services.

It automatically removes the oldest builds once the limit is reached (currently set at 5 and will be adjustable through the API int he future).

## Status

Nearly Complete/Experimental

## Configuration:

Start hoarder by passing a config file (hoarder /path/to/config). If the config file is not passed a default set of options will be used.

```
# Hoarder config file

# The address you want me to listen to
# ip and port are important
listenAddr 0.0.0.0:1234

# show a specific amount of logs
# default value is info
logLevel info

# the authentication token
token supersecrettoken

#the folder you want to store the files
dataDir /tmp/hoarder/
```

## Routes:

| Method | Route | Functionality |
| --- | --- | --- |
| GET | /builds/{file} | Retrieve the build by the file name |
| HEAD | /builds/{file} | Retrieve just the head of the file which includes NAME and SIZE |
| POST | /builds/{file} | Publish a New file, if the file exists replace it |
| PUT | /builds/{file} | Publish a New file, if the file exists replace it |
| Delete | /builds/{file} | Remove a existing file |
| GET | /builds" | List all the build files |
| GET | /libs" | Retrieve the Libs |
| POST | /libs" | Publish a new libs file, replace anything that already exists |
| PUT | /libs" | Publish a new libs file, replace anything that already exists |

## Todo

- refactor architecture for cleaner pattern
- add authentication layer
- tests

### Contributing

Contributions to the hoarder project are welcome and encouraged. Hoarder is a [Nanobox](https://nanobox.io) project and contributions should follow the [Nanobox Contribution Process & Guidelines](https://docs.nanobox.io/contributing/).

### Licence

Mozilla Public License Version 2.0

[![open source](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
