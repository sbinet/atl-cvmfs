atl-cvmfs
=========

``atl-cvmfs`` is a simple command to create, install and remove athena package tarballs off a ``CVMFS`` repository.

## Installation

```sh
$ go get github.com/sbinet/atl-cvmfs
```

## Usage

### ``atl-cvmfs pkg-create``
This command creates a tarball from a release installed on ``CVMFS``

```sh
$ ll /cvmfs/atlas.cern.ch/repo/sw/software
[...]
lrwxrwxrwx. 1 cvmfs cvmfs   26 Apr 17  2012 17.2.1 -> i686-slc5-gcc43-opt/17.2.1/
lrwxrwxrwx. 1 cvmfs cvmfs   27 Mar 31 20:00 17.2.10 -> i686-slc5-gcc43-opt/17.2.10/
[...]
$ atl-cvmfs pkg-create -rel 17.2.10 -outdir=/some/where
$ ll /some/where
-rw-r--r--. 1 binet zp 2.3G Apr  4 18:39 athena-17.2.10-i686-slc5-gcc43-opt.tar.gz
```

/Note/ that by default, ``atl-cvmfs pkg-create`` will package up the
rather biggish ``DBRelease`` directory, to disable it:

```sh
$ atl-cvmfs pkg-create -rel 17.2.10 -outdir=/some/where -with-dbrelease=0
```

