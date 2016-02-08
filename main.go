package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"syscall"

	"github.com/Songmu/wrapcommander"
	"github.com/mattn/go-colorable"
)

func main() {
	defaultFlags := declareFlags("")
	outFlags := declareFlags("out")
	errFlags := declareFlags("err")
	force := flag.Bool("force", false, "Force colorizing.")
	flag.Parse()

	if *outFlags.prefix == "" {
		*outFlags.prefix = *defaultFlags.prefix
	}
	if *errFlags.prefix == "" {
		*errFlags.prefix = *defaultFlags.prefix
	}

	if *defaultFlags.pattern == "" {
		*defaultFlags.pattern = ".*"
	}
	if *outFlags.pattern == "" {
		*outFlags.pattern = *defaultFlags.pattern
	}
	if *errFlags.pattern == "" {
		*errFlags.pattern = *defaultFlags.pattern
	}

	args := flag.Args()
	if len(args) <= 0 {
		fatal(1, "No command specified")
	}

	cmd := exec.Command(args[0], args[1:]...)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fatal(2, err.Error())
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		stdout.Close()
		fatal(2, err.Error())
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
	go func() {
		for {
			if s, ok := (<-sig).(syscall.Signal); ok {
				if cmd.Process == nil {
					os.Exit(int(s) | 0x80)
				}
				syscall.Kill(cmd.Process.Pid, s)
			} else {
				if cmd.Process == nil {
					os.Exit(1)
				}
				cmd.Process.Kill()
			}
		}
	}()

	err = cmd.Start()
	if err != nil {
		stdout.Close()
		stderr.Close()
		fatal(3, err.Error())
	}

	outColor := outFlags.ApplyTo(defaultFlags.ToColor())
	if *force {
		outColor.Force()
	}

	outPipe, err := newPipe(*outFlags.prefix, *outFlags.pattern, outColor)
	if err != nil {
		stdout.Close()
		stderr.Close()
		fatal(4, err.Error())
	}

	errColor := errFlags.ApplyTo(defaultFlags.ToColor())
	if *force {
		errColor.Force()
	}

	errPipe, err := newPipe(*errFlags.prefix, *errFlags.pattern, errColor)
	if err != nil {
		stdout.Close()
		stderr.Close()
		fatal(4, err.Error())
	}

	outDone := make(chan struct{})
	go func() {
		defer stdout.Close()
		output := colorable.NewColorableStdout()
		if err := outPipe.Copy(output, stdout); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		outDone <- struct{}{}
	}()

	errDone := make(chan struct{})
	go func() {
		defer stderr.Close()
		output := colorable.NewColorableStderr()
		if err := errPipe.Copy(output, stderr); err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
		}
		errDone <- struct{}{}
	}()

	<-outDone
	<-errDone
	err = cmd.Wait()

	os.Exit(wrapcommander.ResolveExitCode(err))
}

func fatal(code int, message string) {
	fmt.Fprintln(os.Stderr, message)
	os.Exit(code)
}
