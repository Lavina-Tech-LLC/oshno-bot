package socket

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
	tele "gopkg.in/telebot.v3"
)

func Connection(chatId string, lang string) (*websocket.Conn, error) {

	url := "wss://chatly-ws.lavina.tech/chat/" + chatId
	header := http.Header{}

	header.Add("Origin", "https://oshno.lavina.tech")
	header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	header.Add("Cache-Control", "no-cache")
	header.Add("Accept-Language", lang)
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Host", "chatly-ws.lavina.tech")
	header.Add("Cookie", "user-room="+chatId)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = tlsConfig

	conn, _, err := dialer.Dial(url, header)

	if err != nil {
		log.Fatal("dial:", err)
		return nil, err
	}

	return conn, nil
}

func ConnTest(context tele.Context, chatId string) (*websocket.Conn, error) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	url := "wss://chatly-ws.lavina.tech/chat/" + chatId
	header := http.Header{}

	header.Add("Origin", "https://oshno.lavina.tech")
	header.Add("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36")
	header.Add("Cache-Control", "no-cache")
	header.Add("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	header.Add("Accept-Encoding", "gzip, deflate, br")
	header.Add("Host", "chatly-ws.lavina.tech")
	header.Add("Cookie", "user-room="+chatId)

	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}

	dialer := websocket.DefaultDialer
	dialer.TLSClientConfig = tlsConfig

	c, _, err := dialer.Dial(url, header)

	if err != nil {
		log.Fatal("dial:", err)
		return &websocket.Conn{}, err
	}

	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				return
			}
			fmt.Printf("Received message: %s\n", message)
			context.Send(string(message))
		}
	}()

	for {
		select {
		case <-done:
			return &websocket.Conn{}, nil
		case <-interrupt:
			fmt.Println("Interrupt received, closing connection...")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
				return &websocket.Conn{}, nil
			}
			select {
			case <-done:
			case <-time.After(time.Second):
			}
			return &websocket.Conn{}, nil
		}
	}
}
