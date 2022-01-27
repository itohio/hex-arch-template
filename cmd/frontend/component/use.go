package component

import (
	"context"
	"reflect"
	"sync/atomic"
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

type Watcher struct {
	Ctx       context.Context
	Vars      []interface{}
	Timeout   time.Duration
	Component vecty.Component
	ch        chan struct{}
	Err       error
	count     int32
	prev      int32
}

func UseWatcher(p vecty.Component, timeout time.Duration, f func(context.Context) error, watch ...interface{}) *Watcher {
	ret := &Watcher{
		Ctx:       context.Background(),
		Component: p,
		Timeout:   timeout,
		Vars:      watch,
		ch:        make(chan struct{}, 1),
		count:     1,
	}

	go ret.Run(f)

	return ret
}

func (w *Watcher) Run(f func(context.Context) error) {
	for _ = range w.ch {
		ctx, cancel := context.WithTimeout(w.Ctx, w.Timeout)
		func() {
			defer cancel()
			w.Err = f(ctx)
		}()
		if w.Err != nil {
			break
		}
		atomic.AddInt32(&w.count, -1)
		vecty.Rerender(w.Component)
	}
}

func (w *Watcher) trigger() {
	select {
	case w.ch <- struct{}{}:
	default:
	}
}

func (w *Watcher) Watch(watch ...interface{}) {
	defer atomic.AddInt32(&w.count, 1)

	if len(watch) != len(w.Vars) {
		return
	}
	if len(watch) == 0 {
		w.trigger()
	}

	trigger := false

	c := atomic.LoadInt32(&w.count)
	if w.prev != c {
		trigger = true
	}
	w.prev = c

	for i := range watch {
		if !reflect.DeepEqual(watch[i], w.Vars[i]) {
			trigger = true
			w.Vars[i] = watch[i]
		}
	}

	if trigger {
		w.trigger()
	}
}

func (w *Watcher) Close() {
	close(w.ch)
}
