package socket

import (
	"crypto/tls"
	"log"
	"net/http"
	"oshno/config"

	"github.com/gorilla/websocket"
)

func Connection(chatId string, lang string) (*websocket.Conn, error) {
	cfg := config.Config()
	url := "wss://chatly-ws.lavina.tech/chat?chatId=" + chatId + "&serviceKey=" + cfg.Chatly.Key
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
