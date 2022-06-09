package main
 
import (
"context"
"math/rand"
"os"
"time"
"github.com/gofiber/fiber/v2"
"go.mongodb.org/mongo-driver/mongo"
"go.mongodb.org/mongo-driver/mongo/options"
)

/*
global structures in src/structures.go
global constants in src/constants.go
network functions in src/network_functions.go
blockchain API in src/API_blockchain.go
*/

func main() {

  // set the random number generator
  rand.Seed(time.Now().UTC().UnixNano())

  // setup the mongodb connection
  ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
  defer cancel()
  mongoClient, mongoClienterror = mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
  if mongoClienterror != nil {
    os.Exit(0)
  }

  defer func() {
    if mongoClienterror = mongoClient.Disconnect(context.TODO()); mongoClienterror != nil {
      os.Exit(0)
    }
  }()
        
  // setup fiber
  app := fiber.New(fiber.Config{
    Prefork: true,
    DisableStartupMessage: true,
  })

  // setup blockchain routes
  app.Get("/v1/xcash/blockchain/unauthorized/stats/",v1_xcash_blockchain_unauthorized_stats)
  app.Get("/v1/xcash/blockchain/unauthorized/blocks/",v1_xcash_blockchain_unauthorized_blocks_blockHeight)
  app.Get("/v1/xcash/blockchain/unauthorized/blocks/:blockHeight/",v1_xcash_blockchain_unauthorized_blocks_blockHeight)
  //app.Get("/v1/xcash/blockchain/unauthorized/tx/:txHash/",v1_xcash_blockchain_unauthorized_tx_txHash)
  app.Post("/v1/xcash/blockchain/unauthorized/tx/prove/",v1_xcash_blockchain_unauthorized_tx_prove)
  app.Post("/v1/xcash/blockchain/unauthorized/address/prove",v1_xcash_blockchain_unauthorized_address_prove)
  app.Post("/v1/xcash/blockchain/unauthorized/address/createIntegrated",v1_xcash_blockchain_unauthorized_address_create_integrated)


  // setup global routes
  app.Get("/*", func(c *fiber.Ctx) error {
    return c.SendString("Invalid API Request")
  })
 
  app.Listen(":9000")
}
