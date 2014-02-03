Key Value Store
=============
kvstore is yet another implementation of a Key Value store that is implemented on a Client Server Model assuming there to be many of them running concurrently. 
Code maturity is considered experimental.

The functions of SET, GET and DEL (Delete) have been successfully implemented

Installation
------------

Use `go get github.com/gargprateek26/kvstore`.  Or alternatively, download or clone the repository.

Usage
-----

The Demo Directory contains the files corresponding to the Client and Server, both of which are contained in package main.

Testing
-------
The `kvstore_test.go` file shows usage examples and also, performs testing of the implementation. The following test cases have been considered keeping in view that several clients may concurrently bombarding the same server with different requests: 

1. Single SET followed by single GET, both by different clients. 
2. Several SET by various clients, followed by a GET. 
3. SET followed by SET (of the same key) followed by GET. 
4. SET followed by DEL followed by GET.


Maintainer
----------
Prateek Garg ( garg_prateek26[AT]yahoo[DOT]com )

Coding Style
------------
The source code is automatically formatted to follow `go fmt` by the [IDE]
(https://code.google.com/p/liteide/).  And where pragmatic, the source code
follows this general [coding style]
(http://slamet.neocities.org/coding-style.html).

