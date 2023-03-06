package retry

import "github.com/joomcode/errorx"

var (
	nsRetry              = errorx.NewNamespace("Retry")
	ErrAllAttemptsFailed = nsRetry.NewType("AllAttemptsFailed")
)
