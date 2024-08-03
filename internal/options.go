package internal

import (
	"flag"
	"strings"
)

type stringSliceFlag []string

func (s *stringSliceFlag) String() string {
	return strings.Join(*s, ",")
}

func (s *stringSliceFlag) Set(value string) error {
	*s = append(*s, value)
	return nil
}

type Options struct {
	Token   string
	Users   []string
	Verbose bool
}

func ParseOptions() Options {
	var opts Options
	var users stringSliceFlag

	flag.StringVar(&opts.Token, "token", "", "GitHub token")
	flag.StringVar(&opts.Token, "t", "", "GitHub token (shorthand)")
	flag.Var(&users, "users", "Comma separated list of users")
	flag.Var(&users, "u", "Comma separated list of users (shorthand)")
	flag.BoolVar(&opts.Verbose, "verbose", false, "Enable verbose output")
	flag.BoolVar(&opts.Verbose, "v", false, "Enable verbose output (shorthand)")

	flag.Parse()

	for _, u := range users {
		opts.Users = append(opts.Users, strings.Split(u, ",")...)
	}

	return opts
}
