# wassit
Proxy http calls to a given host with extended options.

## Why
It's useful for people who want to proxy traffic easily.

### Possible Scenarios
- Tunnel IPTV traffic (and follow redirects - some IPTV apps are not able to do so)
- Connect to server and "fake `localhost` hostname"

## Install

### Binary
Download latest binary from releases page.

### Go Style
```sh
go get -u github.com/sbani/wassit
```

## How
```
A fast and simple request http proxy
                with easy to use configuration options
                created by sbani in Go.
Usage:
  wassit target [flags]
Flags:
  -f, --follow-redirect   Follow the first redirect (if present) and proxies content
  -h, --help              help for wassit
  -l, --listen string     Host and port that the wassit server is listening to (default ":9001")
  -q, --quiet             Do not log output to sdtout. Silent mode
  -s, --socks5 string     Use a socks5 socks for connections to the target
  -t, --tor               Enable tor socks5 proxy usage. Please don't forget to start tor
```

## Examples
Easiest way to start wassit
```
$ wassit https://www.google.com
Starting reverse host
Server running on :9001
Pushing request to https://www.google.com
```

Use tor and follow the first redirect
```
$ wassit https://google.com --tor -f
Starting reverse host
Server running on :9001
Pushing request to https://google.com
Using socks proxy 127.0.0.1:9150
```

Specify another socks
```
$ wassit https://www.google.com --socks5 127.0.0.1:9000
Starting reverse host
Server running on :9001
Pushing request to https://www.google.com
Using socks proxy 127.0.0.1:9000
```

## Contribute
Contributions are very welcome. Feel free to file a bug, create feature requests or ask questions through [Github Issues](https://github.com/sbani/wassit/issues). I also appreciate any [Pull Request](https://github.com/sbani/wassit/issues) coming from the community.

## License
See [LICENSE](https://github.com/sbani/wassit/blob/master/LICENSE) file.

