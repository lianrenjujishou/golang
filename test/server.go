package main
 import (
   "log"
   "net/http"
   "io"
 )


//   req.ParseForm()
//   log.Print(req.Form.Get("p1"), " ",req.Form.Get("p2"), " ",req.Form.Get("c"), " ", req.Form.Get("p3"), " ",req.Form.Get("p4"))
//   io.WriteString(res, "<person><name>ZhangSan</name><email>ZhangSan@test.com</email></person> <person><name>Adam</name><email>adam@test.com</email></person> <person><name>Eve</name><email>eve@test.com</email></person>")
func main() {
  http.HandleFunc("/red", func(res http.ResponseWriter, req *http.Request){
    io.WriteString(res, "RED")
  })
  http.HandleFunc("/a", func(res http.ResponseWriter, req *http.Request){
    io.WriteString(res, "A")
  })
//  http.HandleFunc("/b", func(res http.ResponseWriter, req *http.Request){
//  http.Redirect(res, req, "/red", 301)
//  })
  h := http.RedirectHandler("/red", 304)
  http.Handle("/b", h)
   log.Fatal(http.ListenAndServe(":8090", nil))
 }
 
