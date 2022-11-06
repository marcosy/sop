# sop

[![Build](https://github.com/marcosy/sop/actions/workflows/build.yml/badge.svg)](https://github.com/marcosy/sop/actions/workflows/build.yml)
[![coverage](https://img.shields.io/badge/coverage-98.9%25-brightgreen)](https://github.com/marcosy/sop/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/marcosy/sop)](https://goreportcard.com/report/github.com/marcosy/sop)
[![License](https://img.shields.io/github/license/marcosy/sop?color=brightgreen)](./LICENSE)

_A command line tool to perform **s**et **op**erations with files_

## Installation

You can install `sop` using `homebrew`, `go install` or building it from source.

### Homebrew

```bash
> brew install marcosy/tap/sop
```

### Go Install

```bash
> go install github.com/marcosy/sop
```

Remember to add `GOPATH/bin` to your `PATH`:

```bash
> export PATH="$GOPATH/bin:$PATH"
```

### Build from source

You can also build `sop` from source, just run:

```bash
> git clone git@github.com:marcosy/sop.git
> make build
```

The binary will be saved at `./bin/sop`. For other targets, run `make help`.

## Usage

`sop` considers files as sets of elements and performs set operations with those files.

```bash
sop [options] <operation> <filepath A> <filepath B>
```

- **operation** can be one of:
  - `union`: Print elements that exists in file A or file B

  - `intersection`: Print elements that exists in file A and file B
  
  - `difference`: Print elements that exists in file A and do not exist in file B

- **filepath A** and **filepath B** are the filepaths to the files containing
the elements to operate with. Elements are delimited by a separator string which
by default is `"\n"`.

- **options** can be:
  - `-s`: String used as element separator (default `"\n"`)

## Examples

Given two files A (`fileA.txt`) and B (`fileB.txt`):

`fileA.txt`:

```txt
Fox
Duck
Dog
Cat
```

`fileB.txt`:

```txt
Dog
Cat
Cow
Goat
```

`sop` performs set operations with the files.

### Operations

The available operations are: union, intersection and difference.

#### Union

The [union](https://en.wikipedia.org/wiki/Union_(set_theory)) of two sets A and B is the set of elements which are in A, in B, or in both A and B.

```bash
> sop union fileA.txt fileB.txt
Fox
Duck
Dog
Cat
Cow
Goat
```

#### Intersection

The [intersection](https://en.wikipedia.org/wiki/Intersection_(set_theory)) of two sets A and B is the set containing all elements of A that also belong to B or equivalently, all elements of B that also belong to A.

```bash
> sop intersection fileA.txt fileB.txt
Dog
Cat
```

#### Difference

The [difference](https://en.wikipedia.org/wiki/Complement_(set_theory)#Relative_complement) (a.k.a. relative complement) of A and B, is the set of all elements in A that are not in B.

```bash
> sop difference fileA.txt fileB.txt
Fox
Duck
```

### Considerations

#### Separator

The separator character used to delimitate elements is set by default to the new
line character (`\n`) but can also be configured using the flag `-s`:

```bash
> sop -s , union fileA.csv fileB.csv 
```

#### Sorting

The result sets are not ordered by default, so consecutive executions may return
elements in different order. To obtain a consistent order pipe the output of `sop`
to `sort`:

```bash
> sop intersection fileA.txt fileB.txt | sort
```
