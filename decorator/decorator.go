package decorator

import (
	"app/render"
	"errors"
	"fmt"
)

type FieldRules struct {
	fieldPtr interface{}
	rules    []Rule
}

var Need = &Req{}

type Req struct {
}

func (r *Req) Validate(value interface{}) error {
	return errors.New("error from validation ....")
}

type Rule interface {
	Validate(value interface{}) error
}

func Field(fieldPtr interface{}, rules ...Rule) *FieldRules {
	return &FieldRules{
		fieldPtr: fieldPtr,
		rules:    rules,
	}
}

type Data struct {
	Name string
}

func (d *Data) Bind(r *render.Request) error {

	return CreatingError(d, Field(d.Name, Need))
}

func Decorator() {
	r := &render.Request{
		Method: "Post",
	}
	body := &Data{
		Name: "Monte Cristo",
	}
	if err := render.Bind(r, body); err != nil {
		fmt.Println(err)
	}
}

func CreatingError(body *Data, Fields ...*FieldRules) error {
	for _, field := range Fields {
		return Validate(field.fieldPtr, field.rules)
	}
	return nil
}

func Validate(value interface{}, rules []Rule) error {
	for _, rule := range rules {
		return rule.Validate(value)
	}
	return nil
}
