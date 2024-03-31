package render

var Decode = DefaultDecoder

type Request struct {
	Method string
}

type Binder interface {
	Bind(r *Request) error
}

func Bind(r *Request, v Binder) error {
	if err := Decode(r, v); err != nil {
		return err
	}
	return binder(r, v)
}

func DefaultDecoder(r *Request, v interface{}) error {
	r.Method = "Get"
	return nil
}

func binder(r *Request, v Binder) error {
	return v.Bind(r)
}
