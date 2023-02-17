package server

type GeneralResponse[T any] struct {
	Data T `json:"data"`
}
type PingResponse struct {
	Pong string `json:"pong"`
}
