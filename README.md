# mpd_on_go 0.0.1

Go package + sample client `gmpc` for interacting with Music Player
Daemon (MPD).

It is being developed as a "learning Go" project.
Unless you want to improve it I suggest you to use
[something else](https://github.com/fhs/gompd) because my package
lacks of many MPD features.

## Usage

```
package main

import "github.com/nondv/mpd_on_go/mpd"

func main() {
    client, err := mpd.Connect("localhost", 6600)
    client.SetVolume(30)
}
```

See [sample client gmpc](sample_client.go)

## Development

I use `Ruby` and `rake` for development.
For running tests:

```
rake test
```

Tests are running by interacting with
[fake server](mpd/test/server/fake-mpd-server.rb).
So if you want to add some functionality with tests you need to make
sure that fake server supports command for functionality you're adding.


## Contributing

Bug reports and pull requests are welcome on GitHub at
[https://github.com/Nondv/mpd_on_go](https://github.com/Nondv/mpd_on_go).

## License

The library and `gmpc` are available as open source under the terms
of the [MIT License](https://opensource.org/licenses/MIT).
