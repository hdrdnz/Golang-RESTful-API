package main

import (
	"fmt"
	"log"
	"net/http"
)

// go da kullanılan struct yapısını javadaki class mantığına benzetebilirsin.

// ironman adında struct oluşturulur.
type ironman int

type messageHandler struct {
	message string
}

func (a *messageHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, a.message)
}

func (x ironman) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// ioutil genellikle dosya işlemleri için kullanılır.
	w.Write([]byte("Mr.Ironi!"))
}

// nesneler type ile oluşturulur.
func main() {
	var i ironman
	m := &messageHandler{"bu bir mesaj"}

	//bir servis örneği oluşturuldu.
	mux := http.NewServeMux()
	//handle iki değer alan bir fonksiyondur, ilk değer path ikinci değer de bir http.Handler interface’ine uygun ifadedir.
	mux.Handle("/ironman", i)
	//  handle de ifadenin çalışabilmesi için ; ironman struct ı içerisinde ServeHTTP adında bir fonksiyon bulunmalı.

	mux.Handle("/bir", m)

	mux.HandleFunc("/", getData)

	log.Println("Listening..")

	// new server oluşturduğun zaman port içerisinde çalıştırırken buraya yazmalısın.
	http.ListenAndServe(":4000", mux)
}

func getData(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ana server "))
}
