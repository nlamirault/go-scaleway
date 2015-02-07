# go-onlinelabs

[![License GPL 3][badge-license]][LICENSE]
[![travis][badge-travis]][travis]
[![drone][badge-drone]][drone]

A CLI for the Online labs cloud


## Development

* Install requirements :

        $ make init

* Initialize dependencies :

        $ make deps

* Make the binary:

        $ make build

* It make the build inside of a Docker container and the compiled binaries will
  appear in the project directory on the host. By default, it will run a build
  which cross-compiles binaries for a variety of architectures and operating
  systems.

        $ make build-all
        $ make build-linux


## License

go-onlinelabs is free software: you can redistribute it and/or modify it under the
terms of the GNU General Public License as published by the Free Software
Foundation, either version 3 of the License, or (at your option) any later
version.

go-onlinelabs is distributed in the hope that it will be useful, but WITHOUT ANY
WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A
PARTICULAR PURPOSE.  See the GNU General Public License for more details.

See [LICENSE][] for the complete license.


## Changelog

A changelog is available [here](ChangeLog.md).


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-GPL_3-green.svg?style=flat
[LICENSE]: https://github.com/nlamirault/go-onlinelabs/blob/master/LICENSE
[travis]: https://travis-ci.org/nlamirault/go-onlinelabs
[badge-travis]: http://img.shields.io/travis/nlamirault/go-onlinelabs.svg?style=flat
[badge-drone]: https://drone.io/github.com/nlamirault/go-onlinelabs/status.png
[drone]: https://drone.io/github.com/nlamirault/go-onlinelabs/latest
