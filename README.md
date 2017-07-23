# isbn-gen

`isbn-gen` is the single command line tool which generates random and valid [ISBN(International Standard Book Number)](https://en.wikipedia.org/wiki/International_Standard_Book_Number).

## Install

Use `go get` to install.

```sh
$ go get -d github.com/hidapple/isbn-gen
```

## Usage

Generate random ISBN code.
```sh
$ isbn-gen
9784561107828
```

Specify publisher of ISBN with pubcode option.
```sh
$ isbn-gen -p 04
9784047382084
```

## Options

### -h, -help
Desplay a help message

### -v, -version
Display the version of isbn-gen

### -p, -pubcode
The publisher code of ISBN to generate.

