package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

const serverFileDir  ="files/"
func main() {

	http.HandleFunc("/",func(writer http.ResponseWriter, request *http.Request){
		htmlBytes,err:=ioutil.ReadFile("index.html")
		if err != nil {
			log.Println(err)
		}
		_,err= writer.Write(htmlBytes)
		if err != nil {
			log.Fatalf("html error %v\n",err)
		}
	})

	http.HandleFunc("/http.png",func(writer http.ResponseWriter, request *http.Request){
		pngBytes,err:=ioutil.ReadFile(serverFileDir+"go.png")
		if err != nil {
			log.Println(err)
		}

		_,err= writer.Write(pngBytes)
		if err != nil {
			log.Fatalf("png error: %v\n",err)
		}
		_,err= writer.Write([]byte("image/png"))
		if err != nil {
			log.Fatalf("png error: %v\n",err)
		}

	})

	http.HandleFunc("/http.jpg",func(writer http.ResponseWriter, request *http.Request){
		jpgBytes,err:=ioutil.ReadFile(serverFileDir+"gobook.jpg")
		if err != nil {
			log.Println(err)
		}
		_,err= writer.Write(jpgBytes)
		if err != nil {
			log.Fatalf("jpg error: %v\n",err)
		}
		_,err= writer.Write([]byte("image/jpg"))
		if err != nil {
			log.Fatalf("jpg error: %v\n",err)
		}
	})

	http.HandleFunc("/http.pdf",func(writer http.ResponseWriter, request *http.Request){
		pdfBytes,err:=ioutil.ReadFile(serverFileDir+"Go01.pdf")
		if err != nil {
			log.Println(err)
		}
		_,err= writer.Write(pdfBytes)
		if err != nil {
			log.Fatalf("pdf error: %v\n",err)
		}
		_,err= writer.Write([]byte("application/pdf"))
		if err != nil {
			log.Fatalf("pdf error: %v\n",err)
		}
	})

	http.HandleFunc("/http.txt",func(writer http.ResponseWriter, request *http.Request){
		txtBytes,err:=ioutil.ReadFile(serverFileDir+"hello.txt")
		if err != nil {
			log.Println(err)
		}
		_,err= writer.Write(txtBytes)
		if err != nil {
			log.Fatalf("txt error: %v\n",err)
		}
	})

	http.HandleFunc("/files/http.html",func(writer http.ResponseWriter, request *http.Request){
		htmlBytes,err:=ioutil.ReadFile(serverFileDir+"http.html")
		if err != nil {
			log.Println(err)
		}
		_,err= writer.Write(htmlBytes)
		if err != nil {
			log.Fatalf("html error: %v\n",err)
		}
	})

	err := http.ListenAndServe("0.0.0.0:1999",nil)
	if err != nil {
		log.Fatalf("can't ListenAndServe %v\n",err)
	}
}