package gtfsdata

import (
	"context"
	"log/slog"
	"os"

	slogctx "github.com/veqryn/slog-context"
)

type newSlogCtxOptions struct {
	WithTrace bool
}

func newSlogCtx(opts newSlogCtxOptions) context.Context {
	h := slogctx.NewHandler(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{AddSource: opts.WithTrace}), nil)
	slog.SetDefault(slog.New(h))

	return slogctx.NewCtx(context.Background(), slog.Default())
}
