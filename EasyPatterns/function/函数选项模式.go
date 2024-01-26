package main

import (
	"errors"
	"fmt"
)

/*
https://juejin.cn/post/7234731441073619004
*/

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

func main_o() {
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

type loggerOptions struct {
	OutputFormat string
	LogLevel     string
	LogFilePath  string
}

type logger struct {
	options loggerOptions
}

func (l *logger) info(message string) {
	println(l.options.LogFilePath, l.options.LogLevel, message)
}

func newLogger(opts ...func(options2 *loggerOptions)) *logger {
	log := &logger{
		options: loggerOptions{
			OutputFormat: "text",
			LogLevel:     "info",
			LogFilePath:  "app.log",
		},
	}

	for _, opt := range opts {
		opt(&log.options)
	}

	return log
}

func main() {
	log := newLogger(func(opts *loggerOptions) {
		opts.LogLevel = "debug"
		opts.LogFilePath = "debug.log"
	})

	log.info("hello")
}
