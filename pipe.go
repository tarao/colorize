package main

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

type pipe struct {
	pattern *regexp.Regexp
	color   *maybeColor
}

func newPipe(pattern string, color *maybeColor) (*pipe, error) {
	r, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &pipe{r, color}, nil
}

func (p *pipe) Copy(dst io.Writer, src io.Reader) error {
	in := bufio.NewScanner(src)

	wrap := p.color.Wrapper()

	for in.Scan() {
		line := in.Text()
		if p.color.IsColored() {
			line = p.pattern.ReplaceAllStringFunc(line, func(s string) string {
				return wrap(s)
			})
		}
		fmt.Fprintln(dst, line)
	}

	return in.Err()
}
