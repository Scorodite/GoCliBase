package main

import "strings"

type Sequence struct {
	Items []string
}

func (s *Sequence) ParseCli(arg string, source any) error {
	*s = Sequence{Items: strings.Split(arg, ",")}
	return nil
}

func (s Sequence) String() string {
	return strings.Join(s.Items, ",")
}
