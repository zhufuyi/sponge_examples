package initial

import (
	"context"
	"time"

	"community_gw/internal/config"
	"community_gw/internal/rpcclient"

	"github.com/zhufuyi/sponge/pkg/app"
	"github.com/zhufuyi/sponge/pkg/tracer"
)

// RegisterClose register for released resources
func RegisterClose(servers []app.IServer) []app.Close {
	var closes []app.Close

	// close server
	for _, s := range servers {
		closes = append(closes, s.Stop)
	}

	// close the rpc client connection
	closes = append(closes, func() error {
		_ = rpcclient.CloseUserRPCConn()
		_ = rpcclient.CloseRelationRPCConn()
		_ = rpcclient.CloseCreationRPCConn()
		return nil
	})

	// close tracing
	if config.Get().App.EnableTrace {
		closes = append(closes, func() error {
			ctx, _ := context.WithTimeout(context.Background(), 2*time.Second) //nolint
			return tracer.Close(ctx)
		})
	}

	return closes
}
