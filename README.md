# go-scaleway

[![License Apache 2][badge-license]][LICENSE]
[![travis][badge-travis]][travis]
[![drone][badge-drone]][drone]
[![coveralls][badge-coveralls]][coveralls]

A CLI for the [Scaleway][] cloud.
See https://developer.scaleway.com for documentation.

[Click To View Documentation](http://godoc.org/github.com/nlamirault/go-scaleway)

## Installation

Download binary from [releases][] for your platform.

## Usage

```bash
NAME:
   scaleway - A CLI for Scaleway

USAGE:
   scaleway [global options] command [command options] [arguments...]

VERSION:
   0.8.0

AUTHOR:
  Nicolas Lamirault - <nicolas.lamirault@gmail.com>

COMMANDS:
   servers
   users
   organizations
   tokens
   volumes
   images
   help, h		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level, -l 'info'	Log level (options: debug, info, warn, error, fatal, panic)
   --scaleway-token     Scaleway Token [$SCALEWAY_TOKEN] (required)
   --scaleway-userid 		Scaleway UserID [$SCALEWAY_USERID]
   --scaleway-organization 	Organization identifier [$SCALEWAY_ORGANIZATION]
   --help, -h			show help
   --version, -v		print the version

```

You must have a **tokenid** to use the CLI.
By default, commands requiring a **userid** and/or **organizationid** will use the
user id and primary organization id associated with the token.
You may specify different id's by setting the corresponding command line arguments
or environment variables (see GLOBAL OPTIONS).

For each command, subcommands are availables :

```bash
$ ./scaleway servers
NAME:
   scaleway servers -

USAGE:
   scaleway servers command [command options] [arguments...]

COMMANDS:
   list		List all servers associate with your account
   get		Retrieve a server
   delete	Delete a server
   action	Execute an action on a server
   help, h	Shows a list of commands or help for one command

OPTIONS:
   --help, -h	show help

```

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
Retrieve a server    | [x]
List all actions     | [x]
Create a server      | [x]
Update a server      | [ ]
Remove a server      | [x]
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
List all snapshots        | [ ]
Create a snapshot         | [ ]
Retrieve a snapshot       | [ ]
Update a snapshot         | [ ]
Remove a snapshot         | [ ]

### Images

Action                         | Implementation
-------------------------------|------------------------------
List all images                | [x]
Create a new image             | [ ]
Operation on a single image    | [ ]
Retrieves an image             | [x]
Update an image                | [ ]
Delete an image                | [ ]

### IPs

Action                         | Implementation
-------------------------------|------------------------------
Create a new IP                | [x]
Retrieves all IPs addresses    | [x]
Retrieve an IP address         | [x]
Attach an IP address           | [ ]
Remove an IP address           | [x]

### Metadata

Action                         | Implementation
-------------------------------|------------------------------
Pimouss metadata               | [ ]


## Development

* Checkout the projet and install it into $GOPATH :

        $ mkdir -p $GOPATH/src/github.com/nlamirault
        $ git clone https://github.com/nlamirault/go-scaleway.git $GOPATH/src/github.com/nlamirault/go-scaleway
        $ cd $GOPATH/src/github.com/nlamirault/go-scaleway

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
        $ make covoutput pkg=github.com/nlamirault/go-scaleway/commands

* For a new release, it will run a build which cross-compiles binaries for
  a variety of architectures and operating systems:

        $ make release


## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE][] for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-Apache_2-green.svg?style=flat
[LICENSE]: https://github.com/nlamirault/go-scaleway/blob/master/LICENSE
[travis]: https://travis-ci.org/nlamirault/go-scaleway
[badge-travis]: http://img.shields.io/travis/nlamirault/go-scaleway.svg?style=flat
[badge-drone]: https://drone.io/github.com/nlamirault/go-scaleway/status.png
[drone]: https://drone.io/github.com/nlamirault/go-scaleway/latest
[badge-coveralls]: https://coveralls.io/repos/nlamirault/go-scaleway/badge.svg
[coveralls]: https://coveralls.io/r/nlamirault/go-scaleway

[releases]: https://github.com/nlamirault/go-scaleway/releases

[Scaleway]: https://www.scaleway.com
