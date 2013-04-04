atl-cvmfs
=========

``atl-cvmfs`` is a simple command to create, install and remove athena package tarballs off a CVMFs repository

## Installation

```sh
$ go get github.com/sbinet/atl-cvmfs
```

## Usage

### ``atl-cvmfs pkg-create``
This command creates a tarball from a release installed on ``CVMFs``

```sh
$ atl-cvmfs pkg-create -rel 17.2.10 -outdir=/some/where
```
