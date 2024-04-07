package integration

import (
	"context"

	match_reporting_main "github.com/ddd/cmd/match_reporting"
)

func Dispatch(ctx context.Context, rawData [][]string) {
	match_reporting_main.Main(ctx, rawData)
}
