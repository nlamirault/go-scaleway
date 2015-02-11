# go-onlinelabs

[![License GPL 3][badge-license]][LICENSE]
[![travis][badge-travis]][travis]
[![drone][badge-drone]][drone]

A CLI for the Online labs cloud. See [https://doc.cloud.online.net/api]
for documentation.

## API

### Tokens

Action               | Implementation
---------------------|-----------------------------
Create a token       | [x]
List all tokens      | [x]
Retrieve a token     | [x]
Update a token       | [ ]
Remove a token       | [x]

### Organizations

Action               | Implementation
---------------------|------------------------------
List organizations   | [x]

### Users

Action               | Implementation
---------------------|------------------------------
List informations    | [x]

### Servers

Action               | Implementation
---------------------|------------------------------
List servers         | [x]
Create a server      | [x]
Retrieve a server    | [x]
Update a server      | [ ]
Remove a server      | [x]
List all actions     | [ ]
Execute an action    | [x]

### Volumes

Action                     | Implementation
---------------------------|------------------------------
List volumes               | [x]
Create a new volume        | [x]
Retrieves informations     | [x]
Delete a volume            | [x]

### Snapshots

Action                    | Implementation
--------------------------|------------------------------
Create a snapshot         | [ ]
List all snapshots        | [ ]
Retrieve a snapshot       | [ ]
Update a snapshot         | [ ]
Remove a snapshot         | [ ]

### Images

Action                         | Implementation
-------------------------------|------------------------------
Create a new image             | [ ]
List all images                | [x]
Operation on a single image    | [ ]
Retrieves an image             | [x]
Update an image                | [ ]
Delete an image                | [ ]

### IPs

Action                         | Implementation
-------------------------------|------------------------------
Create a new IP                | [ ]
Retrieves all IPs addresses    | [ ]
Retrieve an IP address         | [ ]
Attach an IP address           | [ ]
Remove an IP address           | [ ]

### Metadata

Action                         | Implementation
-------------------------------|------------------------------
Pimouss metadata               | [ ]


## Development

* Install requirements :

        $ make init

* Initialize dependencies :

        $ make deps

* Make the binary:

        $ make build

* Launch unit tests :

        $ make test

* Check code coverage for project or specific package :

        $ make coverage
        $ make covoutput pkg=github.com/nlamirault/go-onlinelabs/commands

* For a new release, it will run a build which cross-compiles binaries for
  a variety of architectures and operating systems:

        $ make release


## License

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
