package main

import (
	"errors"
	"fmt"
)

type options struct {
	option1 string
	option2 int
	option3 func() error
}

type Option func(*options)

func widthOption1(value string) Option {
	return func(o *options) {
		o.option1 = value
	}
}

func widthOption2(value int) func(*options) {
	return func(o *options) {
		o.option2 = value
	}
}

func widthOption3(value func() error) func(*options) {
	return func(o *options) {
		o.option3 = value
	}
}

func doSomething(o *options, opts ...Option) {
	for _, opt := range opts {
		opt(o)
	}
}

func main() {
	opt := options{
		option1: "ok",
		option2: 1,
		option3: func() error {
			return errors.New("hello")
		},
	}
	doSomething(&opt, widthOption1("hello"))

	fmt.Println(opt)
}
