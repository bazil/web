---
title: Installation
github: https://github.com/bazillion/bazil
---

## Installing from source

Building Bazil generally requires the latest *1.x* release of Go.

If you want to install from source, first
[install Go](http://golang.org/doc/install), then
[set up GOPATH](http://golang.org/doc/code.html) and then run

``` console
go get bazil.org/bazil
```

You'll find an executable in `$GOPATH/bin/bazil`. That file is all
that is needed for Bazil itself to run.

## Dependencies

To work right, Bazil needs the FUSE utilities installed, specifically
command `fusermount` from `fuse.deb` or similar for Linux, and
[OSXFUSE](http://osxfuse.github.io/) for Mac.

``` console
sudo apt-get install fuse
```

Your Linux distribution may require you to add your user account to
group `fuse`, and re-login, to be allowed to use FUSE.
