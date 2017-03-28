/*
Package gpshash implements a series of general purpose string hashing algorithms.

Usage:

    package main

    import (
        "fmt"

        hash "github.com/umahmood/gpshash"
    )

    func main() {
        h := hash.RS("abcdefghijklmnopqrstuvwxyz1234567890")
        fmt.Println(h)
    }
*/
package gpshash
