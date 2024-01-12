package service

import (
	"fmt"

	"github.com/kdrkrgz/stockme/client"
	"github.com/kdrkrgz/stockme/model"
	webParser "github.com/kdrkrgz/stockme/parser"
)

type IStockService interface {
	GetStockPrice() (model.Stock, error)
}

type StockService struct {
	client client.IStockClient
}

func NewStockService(client client.IStockClient) StockService {
	return StockService{
		client: client,
	}

}

func (svc *StockService) GetStockPrice() (*model.Stock, error) {
	stock := &model.Stock{}
	respBody, err := svc.client.GetLastPrice()
	if err != nil {
		return nil, fmt.Errorf("client error")
	}
	stock, err = webParser.BorsaDovizWebParser(respBody)
	if err != nil {
		return nil, fmt.Errorf("parse error")
	}
	return stock, err
}
