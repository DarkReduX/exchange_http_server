package service

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"main/src/internal/data"
	pb "main/src/internal/protocol"
	"main/src/internal/repository"
)

func ListenPriceUpdates() error {
	conn, err := grpc.Dial("localhost:8081", grpc.WithInsecure(), grpc.WithBlock())
	log.Infof("try connect")
	if err != nil {
		return err
	}
	log.Infof("Connected")
	defer func() {
		if err = conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	client := pb.NewPriceServiceClient(conn)

	res, err := client.SendPrice(context.Background(), &pb.Conn{Message: "connected client"})
	if err != nil {
		log.Error(err)
		return err
	}
	for {
		price, err := res.Recv()

		if err != nil {
			continue
		}
		repository.Repository.Data[price.Symbol] = data.SymbolPrice{
			Uuid:   price.GetUuid(),
			Symbol: price.GetSymbol(),
			Bid:    price.GetBid(),
			Ask:    price.GetAsk(),
		}
		repository.Repository.IsUpdated = true
	}
}
