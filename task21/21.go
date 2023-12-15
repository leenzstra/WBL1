package main

import (
	"fmt"
)

// Реализовать паттерн «адаптер» на любом примере.

// Надо реализовать этот интерфейс (Target)
type Requester interface {
	SendRequest()
}

// Допустим есть уже реализоанные где-то структуры httpclient и jsonrpcclient
// Которые мы не можем редактировать (или нежелательно)
type HttpClient struct{}

type JSONRPCClient struct{}

// У них есть эти методы, которые нужно унифицировать
func (c *HttpClient) HttpRequest() {
	fmt.Println("HttpClient HttpRequest")
}

func (c *JSONRPCClient) JSONRPCRequest() {
	fmt.Println("JSONRPCClient JSONRPCRequest")
}

// Реализуем адаптеры для интерфейса Requester
type HttpClientAdapter struct {
	client *HttpClient
}

type JSONRPCClientAdapter struct {
	client *JSONRPCClient
}

func (ca *HttpClientAdapter) SendRequest() {
	ca.client.HttpRequest()
}

func (ca *JSONRPCClientAdapter) SendRequest() {
	ca.client.JSONRPCRequest()
}

func main() {
	// Разные структуры, не реализуют интерфейс Requester
	httpClient := HttpClient{}
	jsonrpcClient := JSONRPCClient{}

	// Создаем адапатеры
	httpAdaper := HttpClientAdapter{&httpClient}
	jsonRpcAdaper := JSONRPCClientAdapter{&jsonrpcClient}

	// Адаптеры реализуют Requester (Target)
	httpAdaper.SendRequest()
	jsonRpcAdaper.SendRequest()
}
