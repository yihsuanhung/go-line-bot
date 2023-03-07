package conf

type Option func(o *Options)

type Options struct {
	Path string
	Type string
}

func mergeOptions(o ...Option) *Options {
	opts := new(Options)
	for _, opt := range o {
		opt(opts)
	}
	return opts
}

func WithPath(path string) Option {
	return func(o *Options) {
		o.Path = path
	}
}

func WithType(t string) Option {
	return func(o *Options) {
		o.Type = t
	}
}
