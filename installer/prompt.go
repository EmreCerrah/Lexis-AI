package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type prompter struct {
	in  *bufio.Reader
	out io.Writer
}

func newPrompter(in io.Reader, out io.Writer) *prompter {
	return &prompter{in: bufio.NewReader(in), out: out}
}

func (p *prompter) line() string {
	s, _ := p.in.ReadString('\n')
	return strings.TrimSpace(s)
}

// choose prints a numbered menu and returns the 0-based index; def is used on empty input.
func (p *prompter) choose(question string, options []string, def int) int {
	for {
		fmt.Fprintln(p.out, question)
		for i, o := range options {
			fmt.Fprintf(p.out, "%d) %s\n", i+1, o)
		}
		fmt.Fprintf(p.out, "Seçim [1-%d] (varsayılan %d): ", len(options), def+1)
		s := p.line()
		if s == "" {
			fmt.Fprintln(p.out)
			return def
		}
		if n, err := strconv.Atoi(s); err == nil && n >= 1 && n <= len(options) {
			fmt.Fprintln(p.out)
			return n - 1
		}
		fmt.Fprintf(p.out, "Geçersiz seçim: %q\n\n", s)
	}
}

func (p *prompter) ask(question, def string) string {
	fmt.Fprintf(p.out, "%s [%s]: ", question, def)
	if s := p.line(); s != "" {
		return s
	}
	return def
}

func (p *prompter) confirm(question string, def bool) bool {
	hint := "E/h"
	if !def {
		hint = "e/H"
	}
	fmt.Fprintf(p.out, "%s [%s]: ", question, hint)
	switch strings.ToLower(p.line()) {
	case "":
		return def
	case "e", "evet", "y", "yes":
		return true
	default:
		return false
	}
}
