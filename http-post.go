package main

import "fmt"
import "http"

func main() {
   url := "http://restapi3.apiary.io/notes"
   fmt.Println("posting to ", url)

   var jsonStr = []byte(`{"title":"Buy cheese "}`)
   req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
   req.Header.Set("X-Custom-Header", "myvalue")
   req.Header.Set("Content-Type", "application/json")
   client := &http.Client{}
   resp, err := client.Do(req)
   if err != nil {
      panic(err)
   }
   defer resp.Body.Close()

   fmt.Println("response status : ", resp.Status)


}
