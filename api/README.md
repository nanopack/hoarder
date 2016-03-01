[![hoarder logo](http://nano-assets.gopagoda.io/readme-headers/hoarder.png)](http://nanobox.io/open-source#hoarder)  
[![Build Status](https://travis-ci.org/nanopack/hoarder.svg)](https://travis-ci.org/nanopack/hoarder)

Hoarder is a simple, api-driven, storage system for storing anything for cloud
based deployment services.

## Usage Example
ping
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/ping
pong
```

add blob
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test -d "data"
'test' created!
```

list blobs
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs
[{"Name":"test","Size":4,"ModTime":"2016-03-01T21:13:57.534706044Z"}]
```

get blob
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test
data
```

get blob info
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test -I
HTTP/1.1 200 OK
Content-Length: 4
Date: Tue, 01 Mar 2016 21:14:28 UTC
Last-Modified: Tue, 01 Mar 2016 21:13:57 UTC
Content-Type: text/plain; charset=utf-8
```

update blob
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test -d "datas" -X PUT
'test' created!
```

get blob
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test
datas
```

delete blob
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test -X DELETE
'test' destroyed!
```

get blob info
```
$ curl -k -H "X-NANOBOX-TOKEN: TOKEN" https://localhost:7410/blobs/test -I
HTTP/1.1 404 Not Found
Date: Tue, 01 Mar 2016 21:16:35 GMT
Content-Length: 50
Content-Type: text/plain; charset=utf-8
```

### Contributing

Contributions to the hoarder project are welcome and encouraged. Hoarder is a [Nanobox](https://nanobox.io) project and contributions should follow the [Nanobox Contribution Process & Guidelines](https://docs.nanobox.io/contributing/).

### Licence

Mozilla Public License Version 2.0

[![open source](http://nano-assets.gopagoda.io/open-src/nanobox-open-src.png)](http://nanobox.io/open-source)
