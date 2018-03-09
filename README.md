# isbn-gen

`isbn-gen` is the single command line tool which generates random and valid [ISBN(International Standard Book Number)](https://en.wikipedia.org/wiki/International_Standard_Book_Number).

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

Specify publisher of ISBN with pubcode option.
```sh
$ isbn-gen -p 04
9784047382084
```

## Options

### -help, -h
Desplay a help message

### -version, -v
Display the version of isbn-gen

### -pubcode=code, -p
The publisher code of ISBN to generate.

### -repeat=n, -r
Output ISBN n times

## License
MIT
