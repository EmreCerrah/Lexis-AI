package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strings"
)

type env struct {
	opts   options
	home   string
	src    string
	meta   pluginMeta
	runner *runner
}

type step struct {
	desc string
	run  func() error
}

type plan struct {
	steps []step
	next  []string
}

type runner struct{ dryRun bool }

// run executes an AI tool's CLI with the user's terminal attached, so any prompt it shows reaches the user.
func (r *runner) run(dir, bin string, args ...string) error {
	cmd := exec.Command(bin, args...)
	cmd.Dir = dir
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
	// `claude plugin …` refuses to start inside a Claude Code terminal; plugin subcommands open no session.
	cmd.Env = withoutEnv(os.Environ(), "CLAUDECODE")
	return cmd.Run()
}

// output runs a command quietly and returns its combined output.
func (r *runner) output(bin string, args ...string) (string, error) {
	cmd := exec.Command(bin, args...)
	cmd.Env = withoutEnv(os.Environ(), "CLAUDECODE")
	b, err := cmd.CombinedOutput()
	return string(b), err
}

func withoutEnv(environ []string, key string) []string {
	out := environ[:0:0]
	for _, kv := range environ {
		if !strings.HasPrefix(strings.ToUpper(kv), key+"=") {
			out = append(out, kv)
		}
	}
	return out
}

// findBin returns the first candidate that exists and, if probe is set, accepts probe args.
func (r *runner) findBin(candidates []string, probe ...string) string {
	for _, c := range candidates {
		p, err := exec.LookPath(c)
		if err != nil {
			continue
		}
		if len(probe) == 0 {
			return p
		}
		if _, err := r.output(p, probe...); err == nil {
			return p
		}
	}
	return ""
}

// codexCandidates: `codex` on PATH may be an old npm build without plugin support;
// the desktop app ships its own CLI under %LOCALAPPDATA%\OpenAI\Codex\bin\<hash>\.
func codexCandidates() []string {
	c := []string{"codex"}
	if la := os.Getenv("LOCALAPPDATA"); la != "" {
		matches, _ := filepath.Glob(filepath.Join(la, "OpenAI", "Codex", "bin", "*", "codex.exe"))
		sort.Slice(matches, func(i, j int) bool { return modTime(matches[i]) > modTime(matches[j]) })
		c = append(c, matches...)
	}
	return c
}

func modTime(p string) int64 {
	if st, err := os.Stat(p); err == nil {
		return st.ModTime().UnixNano()
	}
	return 0
}

func quoteCmd(bin string, args ...string) string {
	parts := []string{quoteArg(bin)}
	for _, a := range args {
		parts = append(parts, quoteArg(a))
	}
	return strings.Join(parts, " ")
}

func quoteArg(a string) string {
	if strings.ContainsAny(a, " \t") {
		return fmt.Sprintf("%q", a)
	}
	return a
}
