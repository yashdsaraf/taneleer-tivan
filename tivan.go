package main

import (
	"flag"
	"log"
	"os"
	"regexp"
)

type Options struct {
	Keep      bool
	Password  string
	StoreFile string
	NoPrompt  bool
	Help      bool
}

var Operations = map[string]string{
	"add":  "Add `key [data,...]` in the store (update if `key` exists).",
	"get":  "Get all data corresponding to `key` from the store (Case and spacing insensitive).",
	"keys": "Get all stored keys matching regex `pattern`, return all keys if no pattern supplied. **UNSAFE**",
	"dump": "Print the whole decrypted data on the command line. **UNSAFE**",
	"rem":  "Delete all data corresponding to `key` from the store.",
	"enc":  "Encrypt `file` (Deletes original copy, use `-keep` to ignore).",
	"dec":  "Decrypt `file` (Deletes encrypted copy, use `-keep` to ignore)."}

var Aliases = map[string]string{
	"add": "a",
	"get": "g",
	"rem": "r",
	"enc": "e",
	"dec": "d"}

var OptionUsages = map[string]string{
	"help":      "Print help text for all operations and options",
	"keep":      "Keep the source copy in encryption/decryption operations.",
	"password":  "Password to use for decryption/encryption, program will prompt if empty password is provided. See '-no-prompt'.\n**WARNING**: Be careful when supplying plain text `password` using this option.",
	"store":     "CSV `file` on which CRUD operations will be performed",
	"no-prompt": "Don't prompt for password, will cause the program to error out if not used with '-password'."}

const (
	DEFAULT_PASSWORD_STORE_FILE = "taneleer.csv"
)

func (options *Options) generateFlagOptions() {
	flag.BoolVar(&options.Keep, "keep", false, OptionUsages["keep"])
	flag.StringVar(&options.Password, "password", "", OptionUsages["password"])
	flag.StringVar(&options.StoreFile, "store", DEFAULT_PASSWORD_STORE_FILE, OptionUsages["store"])
	flag.BoolVar(&options.NoPrompt, "no-prompt", false, OptionUsages["no-prompt"])
	flag.BoolVar(&options.Help, "help", false, OptionUsages["help"])
}

func main() {
	log.SetFlags(0) // We only want messages to be printed by logger
	options := Options{}
	options.generateFlagOptions()
	flag.Parse() // Must be called before getCurrentOperation()
	if options.Help {
		printHelp()
		os.Exit(0)
	}
	operation := getCurrentOperation()
	executeOperation(operation, options)
}

func getCurrentOperation() string {
	operation := flag.Arg(0)
	if _, exists := Operations[operation]; ! exists {
		for operationStr, alias := range Aliases {
			if alias == operation {
				operation = operationStr
				exists = true
			}
		}
		if ! exists {
			log.Fatalln("Please supply a valid operation, see -help for more info.")
		}
	}
	return operation
}

func executeOperation(operation string, options Options) {
	switch operation {
	case "add":
		fallthrough
	case "get":
		fallthrough
	case "keys":
		fallthrough
	case "dump":
		fallthrough
	case "rem":
		fallthrough
	case "enc":
		fallthrough
	case "dec":
		fallthrough
	default:
		log.Panicln("Unsupported operation!")
	}
}

func printHelp() {
	log.Printf("Usage: %s [<OPTIONS>] <OPERATION> [<ARGUMENTS>]\n", os.Args[0])
	log.Print("\nOPTIONS:\n\n")
	flag.PrintDefaults()
	log.Print("\nOPERATIONS:\n\n")
	for operation, usage := range Operations {
		var alias string
		if value, exists := Aliases[operation]; exists {
			alias = " | " + value
		}
		log.Printf("%s%s %s\n\t%s\n", operation, alias, getInitialQuotedText(usage), usage)
	}
}

func getInitialQuotedText(data string) string {
	pattern := regexp.MustCompile("`.*?`")
	return pattern.FindString(data)
}
