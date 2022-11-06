package main

import (
     "io"
     "log"
     "net/http"
     "os"
     "fmt"
     "net/url"
     "strings"
)


// parse URL

func main () {

    if len(os.Args) !=2 {
       fmt.Println ("Enter URL: ")
       return
    }

    urll := os.Args[1]
    var _, err = os.Stat(urll)

    u,err := url.Parse(urll)
    if err != nil {
      fmt.Println(err)
    }


    // create output file
       fileName := u.Path
       pars := strings.Split(fileName,"/")
       zz := strings.Join(pars,"")
       file, err := os.Create(zz)
         if err != nil {
           fmt.Println(err,u)
           return
         }
         defer file.Close()


   // HTTP GET request devdungeon.com
   lnk := urll
   response, err := http.Get(lnk)
   defer response.Body.Close()

   // Write bytes from HTTP response to file.
   // response.Body satisfies the reader interface.
   // newFile satisfies the writer interface.
   // That allows us to use io.Copy which accepts
   // any type that implements reader and writer interface

    numBytesWritten, err := io.Copy(file, response.Body)
    if err != nil {
      log.Fatal(err)
    }
    log.Printf("Downloaded %d byte file. \n", numBytesWritten)
}
