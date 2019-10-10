package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/mmcquillan/gaudit/commands"
	"github.com/mmcquillan/gaudit/config"
	"github.com/mmcquillan/matcher"
)

var version string

func main() {

	// options
	var options config.Options
	options.Version = version

	// Github Token
	if os.Getenv("GAUDIT_GITHUB_TOKEN") == "" {
		fmt.Println("ERROR: No GAUDIT_GITHUB_TOKEN defined")
		os.Exit(1)
	}
	options.GithubToken = os.Getenv("GAUDIT_GITHUB_TOKEN")

	// Organization
	if os.Getenv("GAUDIT_ORGANIZATION") == "" {
		fmt.Println("ERROR: No GAUDIT_ORGANIZATION defined")
		os.Exit(1)
	}
	options.Organization = os.Getenv("GAUDIT_ORGANIZATION")

	// Storage
	options.Storage = "audit.state"
	if os.Getenv("GAUDIT_STORAGE") != "" {
		options.Storage = os.Getenv("GAUDIT_STORAGE")
	}

	// Rules
	options.Rules = "rules.yml"
	if os.Getenv("GAUDIT_RULES") != "" {
		options.Rules = os.Getenv("GAUDIT_RULES")
	}

	// Debug
	options.Debug = false
	if os.Getenv("GAUDIT_DEBUG") == "true" {
		options.Debug = true
	}

	match, values := Command("help", os.Args)
	if match {
		options.Args = values
		commands.Help(options)
	}

	match, values = Command("list [filter]", os.Args)
	if match {
		options.Args = values
		commands.List(options)
	}

	match, values = Command("details [filter]", os.Args)
	if match {
		options.Args = values
		commands.Details(options)
	}

	match, values = Command("update", os.Args)
	if match {
		options.Args = values
		commands.Update(options)
	}

	match, values = Command("diff [old] [new] [--verbose]", os.Args)
	if match {
		options.Args = values
		commands.Diff(options)
	}

	match, values = Command("analyze", os.Args)
	if match {
		options.Args = values
		commands.Analyze(options)
	}

	match, values = Command("results [--verbose]", os.Args)
	if match {
		options.Args = values
		commands.Results(options)
	}

	match, values = Command("csv", os.Args)
	if match {
		options.Args = values
		commands.CSV(options)
	}

}

func Command(input string, args []string) (match bool, values map[string]string) {
	match, _, values = matcher.Matcher("<bin> "+input, strings.Join(args, " "))
	return match, values
}