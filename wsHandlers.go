package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func setupRoutes() {
	http.HandleFunc("/currentTemp", currentTemp)
}

func currentTemp(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool {
		return strings.HasPrefix(r.Host, "localhost:")
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Client has connected")
	for {
		log.Println("Sending current temp to client")
		temp, unit, err := getCurrentTemp()
		if err != nil {
			if err = ws.Close(); err != nil {
				log.Println("ERROR - closing connection - ", err)
			}
			return
		}
		fmt.Println(temp, unit)
		tempResp := &TemperatureResponse{
			Temp: temp,
			Unit: unit,
		}
		if err = ws.WriteJSON(tempResp); err != nil {
			log.Println("ERROR - writing response - ", err)
			log.Println("Closing connection")
			ws.Close()
			return
		}
		time.Sleep(time.Second * 1)
	}
}
