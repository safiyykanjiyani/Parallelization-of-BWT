package main

import (
  "fmt"
  "sync"
  "sort"
)

func rotate(s string, pid int, wg *sync.WaitGroup, msg *chan string){
  defer wg.Done()
  r:= s[pid:len(s)] + s[0:pid]
  *msg <- r
}



func bwt(s string) string{
  msg := make(chan string)
  s += "$"
  var sl []string
  wg := new(sync.WaitGroup)
  for i:= 0; i < len(s); i++ {
    wg.Add(1)
    go rotate(s, i, wg, &msg)
  }
  go func(wg *sync.WaitGroup, msg chan string) {
		wg.Wait()
		close(msg)
	}(wg, msg)

  count := 0
  for i:= range msg{
    sl = append(sl, i)
    count += 1
  }
  sort.Strings(sl)
  r := ""
  for i:= range sl{
    r += sl[i][len(s) - 1: len(s)]
  }
  return r
}

func main() {
  fmt.Print("Enter character string to compress: ")
  var user string
  fmt.Scanln(&user)
  transform := bwt(user)
  fmt.Println("Burrows-Wheeler Transform:", transform)
}
