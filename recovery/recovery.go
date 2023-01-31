package recovery

import "fmt"

func Wrap(f func() error) func() error {
	return func() (err error) {
		defer func() {
			if p := recover(); p != nil {
				if err2, ok := p.(error); ok {
					err = fmt.Errorf("panic: %w", err2)
				} else {
					err = fmt.Errorf("panic: %v", p)
				}
			}
		}()
		err = f()
		return
	}
}
