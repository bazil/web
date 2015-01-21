---
title: Architecture
---

##  Filesystem {#filesystem}

Bazil is a [userspace](/doc/related#fuse) file system. This means the
files you see when using Bazil do not have to actually exist as files
on a local file system. This lets you do things like browse terabytes
of archived data on a laptop with a small SSD.

Your local disk will be used as the default storage location and a
cache.

##  Content-Addressed Storage {#cas}

Each file stored in Bazil is divided into chunks,
[hashed](/doc/related#blake2) in a
[Merkle tree](https://en.wikipedia.org/wiki/Merkle_tree), and stored
into a Content-Addressed Store (CAS). Chunk sizes may vary across
files, but each file uses just one chunk size; this is used for fast
random access inside the file.

Using a CAS means that storing the content more than once does not
take up any extra space (if using the same chunk size). Using a Merkle
tree means that small updates to the file only need to store the
updated chunks.

The Merkle tree also ensures data integrity. Knowing the hash at the
root of the tree means none of the content below can have changed,
maliciously or by bitrot.

##  Convergent encryption {#crypto}

[Convergent encryption](https://en.wikipedia.org/wiki/Convergent_encryption)
means data is encrypted with a secret key derived from the data
itself. Bazil uses an extra secret to limit the scope to people you
wish to share and/or deduplicate data with.

Each chunk stored in CAS is encrypted and authenticated with a
[NaCl](/doc/related#nacl) secretbox, with a configured secret and a
nonce derived from a personalized [Blake2](/doc/related#blake2) hash
of the key (which is another personalized Blake2 hash of the content).
The encrypted data is identified by a new key, which is again a
personalized Blake2 hash of the old key, keyed by the configured
secret.

##  Directory storage {#directory}

Trickling every single file content change as Merkle trees of
directories all the way to the top of the root of the volume would be
a lot of extra work. Instead, Bazil takes the hybrid approach:
directories are stored in the CAS only when taking a
[snapshot](/doc/architecture#snapshot); live data lives in a
[key-value database](/doc/related#bolt).

The directory contents are stored as `<dir_inode><basename>` ->
`<file_metadata>`. This makes `readdir+stat` fast, while always
serving directories in alphabetical order.

There is currently no support for multiple directory entries pointing
to the same inode, aka [hard links](/doc/antigoals#limits-hardlink).

##  Snapshot {#snapshot}

Taking a snapshot writes the directory contents from the
[database](/doc/architecture#directory) to the
[CAS](/doc/architecture#cas), thereby archiving a snapshot of the
state of the volume.

Snapshots are created and accessed through the top-level `.snap`
pseudo-directory. A new snapshot can be taken simply with `mkdir`.

``` console
$ echo Hello >greeting
$ mkdir .snap/remember-me
$ rm Hello
$ cat .snap/remember-me/greeting
Hello
```

The inode space is partitioned so that this and other dynamic content
gets distinct inode numbers that never collide with allocated inodes.
