package errset

import (
	"context"
	"sync"
)

type Group struct {
	wg     sync.WaitGroup
	cancel func()
	errs   []error
	mux    sync.Mutex
}

func WithContext(ctx context.Context) (*Group, context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	return &Group{cancel: cancel}, ctx
}

func (g *Group) Go(f func() error) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()

		if err := f(); err != nil {
			g.mux.Lock()
			defer g.mux.Unlock()
			g.errs = append(g.errs, err)
		}
	}()
}

func (g *Group) Wait() []error {
	g.wg.Wait()
	if g.cancel != nil {
		g.cancel()
	}
	return g.errs
}
