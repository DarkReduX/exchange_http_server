package service

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"main/src/internal/data"
	pb "main/src/internal/protocol"
	"main/src/internal/repository"
	"time"
)

func connectPositionService(URI string) (*grpc.ClientConn, pb.PositionServiceClient) {
	conn, err := grpc.Dial(URI, grpc.WithInsecure(), grpc.WithBlock())
	log.Infof("try connect")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("Connected")
	client := pb.NewPositionServiceClient(conn)

	return conn, client
}

func OpenPosition(price data.SymbolPrice, isBay bool, token string) (*data.Position, error) {
	conn, client := connectPositionService("localhost:8082")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	res, err := client.SendOpenPositionRequest(ctx,
		&pb.PositionOpenReq{
			IsBay:  isBay,
			Symbol: price.Symbol,
			Uuid:   price.Uuid,
			Ask:    price.Ask,
			Bid:    price.Bid,
			Token:  token,
		})
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if res.Message == "FAIL" {
		return nil, errors.New("Couldn't open position because this data is not actual ")
	}
	var openCost float32
	if isBay {
		openCost = price.Ask
	} else {
		openCost = price.Bid
	}

	openedPosition := &data.Position{
		UUID:   price.Uuid,
		Symbol: price.Symbol,
		Open:   openCost,
		Close:  0,
		IsBay:  isBay,
	}

	return openedPosition, nil
}

func ClosePosition(position data.Position, token string) (string, error) {
	conn, client := connectPositionService("localhost:8082")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}()

	var closePrice float32
	if position.IsBay {
		closePrice = repository.Repository.Data[position.Symbol].Bid
	} else {
		closePrice = repository.Repository.Data[position.Symbol].Ask
	}
	res, err := client.SendClosePositionRequest(ctx,
		&pb.PositionCloseReq{
			Symbol:     position.Symbol,
			Id:         position.UUID,
			Token:      token,
			Isbay:      position.IsBay,
			ClosePrice: closePrice,
		})
	if err != nil {
		return "", err
	}
	if res.Message == "OK" {
		return "", err
	}
	return "", errors.New("Couldn't close position ")
}

func Login(username string, password string) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	conn, client := connectPositionService("localhost:8082")
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	res, err := client.Login(ctx,
		&pb.LoginReq{
			Login:    username,
			Password: password,
		})
	if err != nil {
		return "", err
	}
	return res.Token, nil
}

func Logout() error {
	conn, client := connectPositionService("localhost:8082")
	ctx, cancel := context.WithTimeout(context.TODO(), time.Second*5)
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	_, err := client.Logout(ctx,
		&pb.Token{
			Token: repository.Repository.UserToken,
		})
	return err
}

func Donate(val float32) error {
	conn, client := connectPositionService("localhost:8082")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Fatal(err)
		}
	}()
	res, err := client.Donate(ctx, &pb.DonateValue{Val: val, Token: repository.Repository.UserToken})
	if err != nil {
		return err
	}
	if res.Message != "OK" {
		return errors.New("Some error while use donate option ")
	}
	repository.Repository.Balance += val
	return err
}

func GetUserData(token string) ([]data.Position, error) {
	conn, client := connectPositionService("localhost:8082")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error()
		}
	}()
	res, err := client.GetUserData(ctx, &pb.Token{Token: token})
	if err != nil {
		return nil, err
	}

	positions := make(map[string]data.Position)
	err = json.Unmarshal(res.Positions, &positions)
	if err != nil {
		return nil, err
	}

	positionsResult := make([]data.Position, 0, 10)
	for _, v := range positions {
		positionsResult = append(positionsResult, v)
	}

	return positionsResult, err
}

func GetUserBalance(token string) (float32, error) {
	conn, client := connectPositionService("localhost:8082")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()
	defer func() {
		if err := conn.Close(); err != nil {
			log.Error()
		}
	}()
	res, err := client.GetUserBalance(ctx, &pb.Token{Token: token})
	if err != nil {
		return 0, err
	}
	return res.Balance, err
}
