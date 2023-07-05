package main

// queryString:Sayfalar arası veri taşıma yollarından en kolayı olan QueryString yöntemi, URL üzerindeki stringler ile verileri taşımaya yarar.
//https://www.google.com/search?q=havva+nur+durudeniz&rlz=1C1GCEA_enTR971TR971 : ?'den sonraki ifadeler querystring'dir.
//querystring get ile çalışır. sayfalar arası veri tarnsferinde kullanılır.
import "net/http"

func main() {

	http.ListenAndServe(":8080", nil)
}

func search(w http.ResponseWriter, r *http.Request) {
	// r: url üzerindeki ifadeleri temsil eder.

	//https://www.google.com/search?q=havva+nur+durudeniz&rlz=1C1GCEA_enTR971TR971
	//q := r.FormValue("q")
	//lz := r.FormValue("lz")
}
