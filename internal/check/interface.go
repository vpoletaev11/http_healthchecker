package check

import (
	"context"
	"fmt"
)

type Checker interface {
	Check(ctx context.Context, url string, cfg map[string]interface{}) error
}

var CheckFailed = fmt.Errorf("check failed")
