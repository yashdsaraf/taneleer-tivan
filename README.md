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

```
tivan.<os> [<OPTIONS>] <OPERATION>
e.g ./tivan.linux -store mypasswordfile add hello.com password@1234 "Some extra data"
This will add the key (hello.com) with data (password@1234 "Some extra data") in an encrypted file called mypasswordfile.aes
OR
e.g ./tivan.linux -keep enc my_ssh_keys.txt
This will encrypt my_ssh_keys.txt and create my_ssh_keys.txt.aes
```

`<os>` can be one of `android, darwin, linux, windows`.

#### OPERATIONS:
* add `key` [`data`] - Add a `key` in the store optionally with `data` (update if `key` exists).
* get `key` - Get all data corresponding to `key` from the store (Case and spacing insensitive).
* keys [`pattern`] - Get all stored keys matching regex `pattern`, return all keys if no pattern supplied. **UNSAFE**
* dump - Print the whole decrypted data on the command line. **UNSAFE**
* rem `key` - Delete all data corresponding to `key` from the store.
* enc `file` - Encrypt `file` (Deletes original copy, use `-keep` to ignore).
* dec `file` - Decrypt `file` (Deletes encrypted copy, use `-keep` to ignore).

#### ALIASES:
###### Shorthand names for above operations
* `a` => `add`
* `g` => `get`
* `r` => `rem`
* `e` => `enc`
* `d` => `dec`

#### OPTIONS:
* help - Print help text for all operations and options.
* keep - Keep the source copy in encryption/decryption operations.
* password - Password to use for decryption/encryption, program will prompt if empty password is provided. See '-no-prompt'.  
**WARNING**: Be careful when supplying plain text password using this option.
* store - File on which CRUD operations will be performed. Default: taneleer.csv
* no-prompt: Don't prompt for password, will cause the program to error out if not used with '-password'.

#### BUILD

* Make sure your `go` version is >= 1.10.4
* Run the bash script `build.sh` in the project root directory.  
When building is done, you can find the binaries for all archs in `bin` directory.
