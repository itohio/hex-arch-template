package component

import (
	"context"
	"time"

	"github.com/hexops/vecty"
)

type Effect struct {
	Ctx    context.Context
	Error  error
	Result interface{}
}

func UseEffect(p vecty.Component, timeout time.Duration, f func(context.Context) (interface{}, error)) *Effect {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	ret := &Effect{
		Ctx: ctx,
	}

	go func() {
		defer vecty.Rerender(p)
		defer cancel()
		r, e := f(ret.Ctx)
		ret.Error = e
		ret.Result = r
	}()

	return ret
}

func (e *Effect) InProgress() bool {
	select {
	case _, ok := <-e.Ctx.Done():
		return !ok
	default:
		return true
	}
}

type State struct {
	Component vecty.Component
	State     interface{}
}

func (s *State) Set(state interface{}) {
	s.State = state
	if s.Component != nil {
		vecty.Rerender(s.Component)
	}
}

func (s *State) String() string {
	if v, ok := s.State.(string); ok {
		return v
	}
	return ""
}

func (s *State) Int() int {
	if v, ok := s.State.(int); ok {
		return v
	}
	return 0
}

func (s *State) Float64() float64 {
	if v, ok := s.State.(float64); ok {
		return v
	}
	return 0
}

func (s *State) Bool() bool {
	if v, ok := s.State.(bool); ok {
		return v
	}
	return false
}
