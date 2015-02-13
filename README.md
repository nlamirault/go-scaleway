# go-onlinelabs

[![License GPL 3][badge-license]][LICENSE]
[![travis][badge-travis]][travis]
[![drone][badge-drone]][drone]

A CLI for the Online labs cloud. See [https://doc.cloud.online.net/api]
for documentation.

## Installation

Download binary from [releases][] for your platform.

## Usage

```bash
$ ./onlinelabs -h
NAME:
   onlinelabs - A CLI for Online Labs

USAGE:
   onlinelabs [global options] command [command options] [arguments...]

VERSION:
   0.4.0

AUTHOR:
  Nicolas Lamirault - <nicolas.lamirault@gmail.com>

COMMANDS:
   server
   user
   organizations
   token
   volume
   image
   help, h		Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --log-level, -l 'info'	Log level (options: debug, info, warn, error, fatal, panic)
   --onlinelabs-userid 		Onlinelabs UserID [$ONLINELABS_USERID]
   --onlinelabs-token 		Onlinelabs Token [$ONLINELABS_TOKEN]
   --onlinelabs-organization 	Organization identifier [$ONLINELABS_ORGANIZATION]
   --help, -h			show help
   --version, -v		print the version

```

You must have a **userid**, **tokenid** and **organizationid** to
use the CLI. You could use command line arguments (in global option), or
environments variables.

For earch commands, subcommands are availables :

```bash
$ ./onlinelabs server
NAME:
   onlinelabs server -

USAGE:
   onlinelabs server command [command options] [arguments...]

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

[releases]: https://github.com/nlamirault/go-onlinelabs/releases
