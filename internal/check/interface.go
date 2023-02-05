package check

import "context"

type Checker interface {
	Check(ctx context.Context, url string, cfg map[string]interface{}) error
}
