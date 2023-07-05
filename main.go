package main

import (
	"net/http"
)

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	//w.Write([]byte("about page"))

	//request sonucu sunucu tarafından döndürülen ifade
	w.WriteHeader(http.StatusOK)

}

func getData(w http.ResponseWriter, r *http.Request) {
	data := ""
	// "/" işaretinden sonraki tüm ifadeyi alır.
	x := r.URL.Path[1:]
	if len(x) > 0 {
		data = x
	} else {
		data = "Index page"
	}

	w.Write([]byte(data))
}

func main() {
	//HandleFunc() fonksiyonu ; belirlediğimiz adrese girildiğinde hangi fonksiyonun çalıştırılacağınız belirliyor.
	//HandleFunc iki değer alan bir fonksiyondur, ilk değer path ikinci değer de bir fonksiyon’dur
	http.HandleFunc("/input", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Merhaba Golang"))
	})

	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/", getData)

	//listenAndServe() fonksiyonu ise sunucunun ayağa kalkmasını ve istediğimiz bir porttan ulaşılmasını sağlıyor.
	http.ListenAndServe(":4000", nil)
}
