# Beatiful (and strange) I/O

Lightning Talk, Go and Cloud Native Leipzig
@BasislagerCo, golangleipzig.space, 2019-03-15, 19:00

----

# Go Proverb

* [The bigger the interface, the weaker the abstraction](https://youtu.be/PAAkCSZUG1c?t=5m18s)

----

# Exemplified in package io

Generic I/O with io.Reader and io.Writer. A few other interfaces:

----

|                    | R | W | C | S |
|--------------------|---|---|---|---|
| io.Reader          | x |   |   |   |
| io.Writer          |   | x |   |   |
| io.Closer          |   |   | x |   |
| io.Seeker          |   |   |   | x |
| io.ReadWriter      | x | x |   |   |
| io.ReadCloser      | x |   | x |   |
| io.ReadSeeker      | x |   |   | x |
| io.WriteCloser     |   | x | x |   |
| io.WriteSeeker     |   | x |   | x |
| io.ReadWriteCloser | x | x | x |   |
| io.ReadWriteSeeker | x | x |   | x |

----

# Missing things

Libraries might implement missing pieces, e.g.

* [ReadSeekCloser, ReaderAtCloser](https://github.com/go4org/go4/blob/94abd6928b1da39b1d757b60c93fb2419c409fa1/readerutil/readerutil.go#L33-L43)

----

# IO interface list

* io.ReaderAt
* io.ReaderFrom
* io.WriterAt
* io.WriterTo

----

# Use cases

* io.ReaderAt, io.WriterAt -- (parallel writes) with offset

----

# Use cases

* io.ReaderFrom -- a data structure, that know how to deserialize itself

Example, different JSON API structs, but each of them implements io.ReaderFrom,
so the data fetch can be separated -- [fetchLocation(location string, r
io.ReaderFrom)](https://github.com/miku/span/blob/86aeec55853b795e57ad80978f97caedc4000ea2/cmd/span-amsl-discovery/main.go#L130-L139)

----

# Readers for types

## Rune

* io.RuneReader
* io.RuneScanner

## Byte

* io.ByteReader
* io.ByteScanner
* io.ByteWriter

## String

* io.StringWriter
