package main 

import (
  "flag"
  "fmt"
  "os"
  "time"
)

func main(){
  var query string
  var pages int
  var delay int

  flag.StringVar(&query, "q","inurl:viewer.html", "Search query")
  flag.IntVar(&pages, "p",1, "Number of pages to search")
  flag.IntVar(&delay, "d",5,"Delay between requests (seconds)")
  flag parse()

  if query == "" {
    fmt.PrintIn("Please provide a search query with -q")
    os.Exit(1)
    }

  fmt.Printf("Searching for: %s\n",query)
  fmt.Printf("Pages: %d, Delay: %ds\n\n",pages,delay)

  simulatedDevices :=[]string{
    "https://192.168.1.100:8080/viewer.html",
    "https://192.168.1.101:80/cgi-bin/viewer",
    "https://10.0.0.50:8080/visul.html",
    "https://203.0.113.10:80/admin",
    "https://198.51.100.23:8080/camera.html",
    }

  for i, device:=range simulatedDevices {
    fmt.Printf("[%d] Found: %s\n",i+1,device)
    if i<len(simulatedDevices)-1 {
      time.Sleep(time.Duration(delay)*time.Second)
      }
    }

  fmt.Printin("\nSearch completed. Use Wireshark to analyze TCP headers.")
  }
    
    
