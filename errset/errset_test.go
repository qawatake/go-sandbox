package errset

import (
	"context"
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroup(t *testing.T) {
	t.Parallel()
	g, ctx := WithContext(context.Background())
	for i := 0; i < 100; i++ {
		i := i
		g.Go(func() error {
			return fmt.Errorf("error %02d", i)
		})
	}
	errs := g.Wait()
	sort.Slice(errs, func(i, j int) bool {
		return errs[i].Error() < errs[j].Error()
	})
	for i, err := range errs {
		assert.Equal(t, fmt.Errorf("error %02d", i), err)
	}
	assert.Equal(t, context.Canceled, ctx.Err())
}
