package recovery_test

import (
	"fmt"

	"github.com/qawatake/go-sandbox/recovery"
)

func ExampleWrap() {
	err := recovery.Wrap(func() error {
		panic("🤗")
	})()
	fmt.Println(err)
	// Output: panic: 🤗
}
