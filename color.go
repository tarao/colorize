package main

import (
	"fmt"

	"github.com/fatih/color"
)

type maybeColor struct {
	c *color.Color
}

func newColor(a color.Attribute) *maybeColor {
	if a == color.Reset {
		return &maybeColor{nil}
	}
	return &maybeColor{color.New(a)}
}

func (c *maybeColor) IsColored() bool {
	return c.c != nil
}

func (c *maybeColor) Add(a color.Attribute) *maybeColor {
	if !c.IsColored() {
		return newColor(a)
	}
	if a != color.Reset {
		c.c.Add(a)
	}
	return c
}

func (c *maybeColor) Force() *maybeColor {
	if c.IsColored() {
		c.c.EnableColor()
	}
	return c
}

func (c *maybeColor) Wrapper() func(a ...interface{}) string {
	if c.IsColored() {
		return c.c.SprintFunc()
	}
	return func(a ...interface{}) string {
		return fmt.Sprint(a...)
	}
}
