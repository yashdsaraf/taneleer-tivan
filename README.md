## Taneleer Tivan

### Description

This is a simple command line program which interacts with a 'password store' to perform basic CRUD operations.
It doesn't make any network calls or store the encryption password anywhere, and can be compiled for almost all
the widely used platforms.

The 'password store' is a CSV (Comma Separated Value) file encrypted with AES-256 bit encryption with data stored
as `key`[,`data`,`data`,...]. Here the `key` is a unique identifier for each entry and `data` is optional data which
can be added corresponding to specific key.

It can additionally encrypt or decrypt files for safe storage.

### Usage

tivan.`os` [OPERATION/ALIAS] [OPTIONS]

`os` can be one of `android, darwin, linux, windows`.

#### OPERATIONS:
* `add` `key` [`data`] - Add a `key` in the store optionally with `data` (update if `key` exists).
* `get` `key` - Get all data corresponding to `key` from the store (Case and spacing insensitive).
* `keys` [`pattern`] - Get all stored keys matching regex `pattern`, return all keys if no pattern supplied. **UNSAFE**
* `dump` - Print the whole decrypted data on the command line. **UNSAFE**
* `rem` `key` - Delete all data corresponding to `key` from the store.
* `enc` `file` - Encrypt `file` (Deletes original copy, use `-k` to ignore).
* `dec` `file` - Decrypt `file` (Deletes encrypted copy, use `-k` to ignore).

#### ALIASES:
###### Shorthand names for above operations
* `a` => `add`
* `g` => `get`
* `r` => `rem`
* `e` => `enc`
* `d` => `dec`

#### OPTIONS:
* --verbose, -v - Print debug info.
* --keep, -k - Keep the source copy in encryption/decryption operations.
* --silent, -s - Do not print or read anything from the command line. Needs the `--password` enabled to work.
* --password, -p - Password to use for decryption/encryption, useful when calling the program from another script.  
**WARNING**: Be careful when supplying plain text password using the above option.
* --out-file, -o - File on which CRUD operations will be performed. Default: taneleer.csv

#### BUILD

* Make sure your `go` version is >= 1.10.4
* Run the bash script `build.sh` in the project root directory.  
When building is done, you can find the binaries for all archs in `bin` directory.
