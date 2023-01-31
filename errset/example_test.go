package errset_test

import (
	"context"
	"errors"
	"fmt"

	"github.com/qawatake/go-sandbox/errset"
)

func ExampleGroup() {
	g, _ := errset.WithContext(context.Background())

	g.Go(func() error {
		return errors.New("first error")
	})
	g.Go(func() error {
		return errors.New("second error")
	})
	errs := g.Wait()
	for _, err := range errs {
		fmt.Println(err)
	}

	// Unordered output: first error
	// second error
}
