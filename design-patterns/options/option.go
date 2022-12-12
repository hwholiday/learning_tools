package main

type Options struct {
	Name string
	Age  int
}

func defaultOptions() *Options {
	return &Options{
		Name: "",
		Age:  0,
	}
}

type FOptions func(*Options)

func WithName(n string) FOptions {
	return func(options *Options) {
		options.Name = n
	}
}

func WithAge(a int) FOptions {
	return func(options *Options) {
		options.Age = a
	}
}

func NewUserOptions(opt ...FOptions) *Options {
	o := defaultOptions()
	for _, f := range opt {
		f(o)
	}
	return o
}
