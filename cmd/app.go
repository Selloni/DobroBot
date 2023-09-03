package main

import (
	"DobroBot/model"
	customestore "DobroBot/store/customeStore"
	"DobroBot/transport/rest"
	"DobroBot/transport/telegram"
	"net/http"
)

func main() {
	store := customestore.NewStore()
	ch := make(chan (model.Discont), 10)
	tg := telegram.NewTelegram(store, ch)

	go tg.Run("6548886185:AAH_D2kYxX2GIV5PhuDWKPjwBpidWeeBVx4")

	handler := rest.NewHandler(ch)

	http.ListenAndServe("localhost:8080", handler.Init())
}
