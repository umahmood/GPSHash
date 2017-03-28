# General Purpose String Hash Algorithms

A Go library which implements a series of general purpose string hash algorithms. 
Implemented hashes:

- AP
- BP
- BKDR
- DJB
- DEK
- ELF
- FNV
- JS
- PJW
- RS
- SDBM

# Installation

> go get github.com/umahmood/gpshash

# Usage

```
package main

import (
    "fmt"

    hash "github.com/umahmood/gpshash"
)

func main() {
    h := hash.RS("abcdefghijklmnopqrstuvwxyz1234567890")
    fmt.Println(h)
}
```
Output:
```
6985544498910593518
```

# Documentation

- [Godoc - gpshash](http://godoc.org/github.com/umahmood/gpshash)

# References

- [General Purpose Hash Function Algorithms](http://www.partow.net/programming/hashfunctions/index.html#AvailableHashFunctions)

# License

See the [LICENSE](LICENSE.md) file for license rights and limitations (MIT).
