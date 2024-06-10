package main

import (
	"api_gateway/handler"
	"api_gateway/proto"
	"context"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	//"github.com/go-micro/generator/cmd/protoc-gen-micro/plugin/micro"
	"github.com/go-micro/plugins/v4/client/grpc"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	// "go-micro.dev/v4/util/grpc"
)

func main() {
	r := gin.Default()

	// CORS middleware configuration
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"http://localhost:5173"},
        AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }))

	addrServiceTransactionOpt := client.WithAddress(":9000")
	clientsrvTransaction := grpc.NewClient()

	srvTransaction := micro.NewService(
		micro.Client(clientsrvTransaction),
	)

	srvTransaction.Init(
		micro.Name("service-transaction"),
		micro.Version("latest"),
	)

	accountRoute := r.Group("/account")
	accountRoute.GET("/get", handler.NewAccount().GetAccount)
	accountRoute.POST("/create", handler.NewAccount().CreateAccount)
	accountRoute.PATCH("/update/:id", handler.NewAccount().UpdateAccount)
	accountRoute.DELETE("/delete/:id", handler.NewAccount().DeleteAccount)
	accountRoute.POST("/getbalance", handler.NewAccount().GetBalance)

	transactionRoute := r.Group("/transaction")
	transactionRoute.POST("/transferbank", handler.NewTransaction().TransferBank)
	transactionRoute.POST("/datatransaksi", handler.NewTransaction().CreateDataTransaksi)
	transactionRoute.GET("/get", func(g *gin.Context) {
		clientResponse, err := proto.NewServiceTransactionService("service-transaction",
			srvTransaction.Client()).
			Login(context.Background(), &proto.LoginRequest{
				Username: "huda",
			}, addrServiceTransactionOpt)

		if err != nil {
			g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		}
		g.JSON(http.StatusOK, gin.H{
			"data": clientResponse,
		})

	})

	authRoute := r.Group("/auth")
	authRoute.POST("/login", handler.NewAuth().Login)

	// r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "pong",
	// 	})
	// })
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
