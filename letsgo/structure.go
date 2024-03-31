package letsgo

type OptFunc func(*Opts)

type Opts struct {
	maxConn int
	id      string
	tls     bool
}

type ServerT struct {
	Opts
}

func defaultOpts() Opts {
	return Opts{
		maxConn: 10,
		id:      "default",
		tls:     false,
	}
}

func WithTLS(opts *Opts) {
	opts.tls = true
}

func WithMaxConn(n int) OptFunc {
	return func(opts *Opts) {
		opts.maxConn = n
	}
}

func WithID(id string) OptFunc {
	return func(opts *Opts) {
		opts.id = id
	}
}

func NewServer(opts ...OptFunc) *ServerT {
	opt := defaultOpts()
	for _, fn := range opts {
		fn(&opt)
	}
	return &ServerT{
		opt,
	}
}
