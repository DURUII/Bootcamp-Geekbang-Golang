package ch00

import (
	"context"
	"fmt"
	"testing"

	"go.uber.org/fx"
)

func TestHook(t *testing.T) {
	fx.New(
		fx.Invoke(func(lc fx.Lifecycle) {
			lc.Append(fx.Hook{
				OnStart: func(context.Context) error {
					fmt.Println("Before main logic")
					return nil
				},
				OnStop: func(context.Context) error {
					fmt.Println("After main logic")
					return nil
				},
			})
		}),
	).Run()
}
