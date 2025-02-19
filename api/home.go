package api

import (
	"net/http"
   
	p "vercel-pos/api/_pkg"
   )
   
   func Hello(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
   
	model := p.NewModel()
   
	resp := "Hello "+ model.Username
	w.Write([]byte(resp))
   }