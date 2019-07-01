# isbn-gen

[![Build Status](https://travis-ci.org/hidapple/isbn-gen.svg?branch=master)](https://travis-ci.org/hidapple/isbn-gen)
[![codecov](https://codecov.io/gh/hidapple/isbn-gen/branch/master/graph/badge.svg)](https://codecov.io/gh/hidapple/isbn-gen)

`isbn-gen` is the command line tool which generates random and valid [ISBN(International Standard Book Number)](https://en.wikipedia.org/wiki/International_Standard_Book_Number).

## Install

Get repository by `go get` and use `make install` to make command available.

```sh
$ go get -d github.com/hidapple/isbn-gen
$ cd $GOPATH/src/github.com/hidapple/isbn-gen
$ make install
```

## Usage

Generate random ISBN code.
```sh
$ isbn-gen
9784561107828
```
### Option Flags

#### -help, -h
Print a help message.

#### -version, -v
Print the version of isbn-gen.

#### -list, -l
Print supported ISBN registration group identifier list. Currently `isbn-gen` supports following group identifiers.
```
+-------------------+--------------+--------+------------+
| IDENTIFYING GROUP | ABBREVIATION | PREFIX | IDENTIFIER |
+-------------------+--------------+--------+------------+
| English           | en           |    978 |          0 |
| English2          | en2          |    978 |          1 |
| French            | fr           |    978 |          2 |
| German            | de           |    978 |          3 |
| Japan             | jp           |    978 |          4 |
| Russia            | ru           |    978 |          5 |
| China             | cn           |    978 |          7 |
| Brazil            | br           |    978 |         65 |
| Czech             | cz           |    978 |         80 |
| India             | in           |    978 |         81 |
| Norge             | no           |    978 |         82 |
| Poland            | pl           |    978 |         83 |
| Spain             | es           |    978 |         84 |
| Brazil2           | br2          |    978 |         85 |
| Serbia            | rs           |    978 |         86 |
| Denmark           | dk           |    978 |         87 |
| Italy             | it           |    978 |         88 |
| SouthKorea        | kr           |    978 |         89 |
| Netherlands       | nl           |    978 |         90 |
| Sweden            | se           |    978 |         91 |
| NGO               | ngo          |    978 |         92 |
| India2            | in2          |    978 |         93 |
| Netherlands2      | nl2          |    978 |         94 |
| French2           | fr2          |    979 |         10 |
| SouthKorea2       | kr2          |    979 |         11 |
| Italy2            | it2          |    979 |         12 |
+-------------------+--------------+--------+------------+
```

#### -id-group, -id
The ISBN registration group identifier name or abbreviation. You can specify the registration identifier group from the table listed above then `isbn-gen` generates ISBN which belongs to the given identifier group.

```sh
$ isbn-gen -id-group cn
9787644999709
```

#### -code, -c
The code part of ISBN which consist of publisher code and book code. You can specify the code part prefix of ISBN.
```sh
$ isbn-gen -code 123
9784123372770
```

## License
MIT
