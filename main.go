package main

import (
	"context"
	"log"
	"net/http"
	"openccapi/rpc/opencc/kitex_gen/demo/opencc_service"
	"openccapi/rpc/opencc/kitex_gen/demo/opencc_service/openccservice"

	"github.com/cloudwego/kitex/client"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	h := server.New(server.WithHostPorts("localhost:8080"))
	h.GET("/", func(ctx context.Context, c *app.RequestContext) {
		c.JSON(http.StatusOK, "ok")
	})
	h.POST("/convert", func(ctx context.Context, c *app.RequestContext) {
		r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
		if err != nil {
			log.Fatal(err)
		}
		cli, err := openccservice.NewClient("opencc.convertor", client.WithResolver(r))
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		req := &opencc_service.ConvertRequest{
			Type: c.PostForm("type"),
			Text: c.PostForm("text"),
		}
		resp, err := cli.Convert(ctx, req)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		c.JSON(http.StatusOK, resp)
	})

	h.Spin()
}
