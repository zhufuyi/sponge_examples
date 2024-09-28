package initial

import (
	"context"
	"time"

	"github.com/zhufuyi/sponge/pkg/app"
	"github.com/zhufuyi/sponge/pkg/tracer"

	"eshop/order_gw/internal/config"
	//"eshop/order_gw/internal/rpcclient"
)

// Close releasing resources after service exit
func Close(servers []app.IServer) []app.Close {
	var closes []app.Close

	// close server
	for _, s := range servers {
		closes = append(closes, s.Stop)
	}

	// close the rpc client connection
	// example:
	//closes = append(closes, func() error {
	//	return rpcclient.CloseOrder_gwRPCConn()
	//})

	// close tracing
	if config.Get().App.EnableTrace {
		closes = append(closes, func() error {
			ctx, _ := context.WithTimeout(context.Background(), 2*time.Second) //nolint
			return tracer.Close(ctx)
		})
	}

	return closes
}
