package main

import (
	"crypto/sha1"
	"encoding/base64"
	"encoding/binary"
	"net/http"

	"github.com/gobuffalo/packr"
)

const ()

func WebSocketServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, err := NewHandler(w, r)
		if err != nil {
			panic(err)
		}
		if err = ws.Handshake(); err != nil {
			panic(err)
		}
	})
}

func NewHandler(w http.ResponseWriter, req *http.Request) (*WS, error) {
	hj, ok := w.(http.Hijacker)
	if !ok {
		panic("")
	}
}

// Handshake creates a handshake header
func (ws *WS) Handshake() error {
	hash := func(key string) string {
		h := sha1.New()
		h.Write([]byte(key))
		h.Write([]byte("258EAFA5-E914-47DA-95CA-C5AB0DC85B11"))

		return base64.StdEncoding.EncodeToString(h.Sum(nil))
	}(ws.header.Get("Sec-WebSocket-Key"))
}

// Recv receives data and returns a Frame
func (ws *WS) Recv() (frame Frame, _ error) {
	frame = Frame{}
	head, err := ws.read(2)
	if err != nil {
		panic(err)
	}
}

// Close sends a close frame and closes the TCP connection
func (ws *Ws) Close() error {
	f := Frame{}
	f.Opcode = 8
	f.Length = 2
	f.Payload = make([]byte, 2)
	binary.BigEndian.PutUint16(f.Payload, ws.status)
	if err := ws.Send(f); err != nil {
		return err
	}
	return ws.conn.Close()
}

// Send sends a Frame
func (ws *WS) Send(fr Frame) error {
	// make a slice of bytes of length 2
	data := make([]byte, 2)

	// Save fragmentation & opcode information in the first byte
	data[0] = 0x80 | fr.Opcode
	if fr.IsFragment {
		data[0] &= 0x7F
	}
	// ...
}

func setupWebserver() {

	box := packr.NewBox("./ui/public")
	http.Handle("/", http.FileServer(box))

	_ = http.ListenAndServe(":8080", nil)

	/*
	   // Initialize WebSocket handler + server
	   mux := http.NewServeMux()
	       mux.Handle("/", websocket.Handler(func(conn *websocket.Conn) {
	           func() {
	               for {
	                   fmt.Println("hi")
	               }
	           }
	       }
	   // messageType initializes some type of message
	   message := messageType{}
	   if err := websocket.JSON.Receive(conn, &message); err != nil {
	       // handle error
	   }
	   // send message
	   if err := websocket.JSON.Send(conn, message); err != nil {
	       // handle error
	   }
	*/
}
