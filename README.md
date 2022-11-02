# sop

_A command line tool to perform **s**et **op**erations with files_

## Installation

You can install `sop` using `go install` or building it from source.

### Go Install

If you have Go installed in your system, just run:

```bash
> go install github.com/marcosy/sop
```

Remember to add `GOPATH/bin` to your path.

### Build from source

You can also build `sop` from source, just run:

```bash
> git clone git@github.com:marcosy/sop.git
> make build
```

The binary will be saved at `./bin/sop`. For other targets, run `make help`.

## Usage

Given two files A (`fileA.txt`) and B (`fileB.txt`):

`fileA.txt`:

```txt
1
2
3
4
```

`fileB.txt`:

```txt
3
4
5
6
```

`sop` performs set operations with the files.

### Operations

The available operations are: union, intersection and difference.

#### Union

The [union](https://en.wikipedia.org/wiki/Union_(set_theory)) of two sets A and B is the set of elements which are in A, in B, or in both A and B.

```bash
> sop union fileA.txt fileB.txt
1
2
3
4
5
6
```

#### Intersection

The [intersection](https://en.wikipedia.org/wiki/Intersection_(set_theory)) of two sets A and B is the set containing all elements of A that also belong to B or equivalently, all elements of B that also belong to A.

```bash
> sop intersection fileA.txt fileB.txt
3
4
```

#### Difference or Relative Complement

The [difference](https://en.wikipedia.org/wiki/Complement_(set_theory)#Relative_complement) of A and B, is the set of all elements in A that are not in B.

```bash
> sop difference fileA.txt fileB.txt
1
2
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
3
4
```
