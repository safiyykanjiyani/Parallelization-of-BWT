package main

import (
  "fmt"
  "sync"
  "sort"
  //"os"
  //"strconv"
)

/*
0  5: ip
1  2: issip
2  6: p
3  4: sip
4  1: sissip
5  3: ssip
6  0: ssissip

[5, 2, 6, 4, 1, 3 , 0]
*/

func linearbwt(s string) string{
  //get the bwt of string s
  rotations := make([]string, len(s))
  for i := 0; i < len(s) ; i++{
    r := s[i:len(s)] + s[0:i]
    rotations[i] = r
  }
  sort.Strings(rotations)
  r := ""
  for i := range rotations{
    r += rotations[i][len(s) - 1: len(s)]
  }

  suffixes := make([]string, len(s))
  sa := make([]int, len(s))

  for i:=0; i<len(s);i++{
    suffixes[i] = s[i:len(s)]
  }
  sort.Strings(suffixes)
  for i := 0; i < len(s); i++ {
    sa[i] = len(s) - len(suffixes[i])
  }

  return r
  //*saresult <- sa
}


func linearsa(s string) []int{
  suffixes := make([]string, len(s))
  sa := make([]int, len(s))

  for i:=0; i<len(s);i++{
    suffixes[i] = s[i:len(s)]
  }
  sort.Strings(suffixes)
  for i := 0; i < len(s); i++ {
    sa[i] = len(s) - len(suffixes[i])
  }

  return sa
}

func sa(s string, wg *sync.WaitGroup, saresult *chan []int){
  defer wg.Done()
  //get the bwt of string s

  suffixes := make([]string, len(s))
  sa := make([]int, len(s))

  for i:=0; i<len(s);i++{
    suffixes[i] = s[i:len(s)]
  }
  sort.Strings(suffixes)
  for i := 0; i < len(s); i++ {
    sa[i] = len(s) - len(suffixes[i])
  }

  *saresult <- sa
}


func bwt(s string, wg *sync.WaitGroup, bwtresult *chan string){
  defer wg.Done()
  //get the bwt of string s
  rotations := make([]string, len(s))
  for i := 0; i < len(s) ; i++{
    r := s[i:len(s)] + s[0:i]
    rotations[i] = r
  }
  sort.Strings(rotations)
  r := ""
  for i := range rotations{
    r += rotations[i][len(s) - 1: len(s)]
  }

  suffixes := make([]string, len(s))
  sa := make([]int, len(s))

  for i:=0; i<len(s);i++{
    suffixes[i] = s[i:len(s)]
  }
  sort.Strings(suffixes)
  for i := 0; i < len(s); i++ {
    sa[i] = len(s) - len(suffixes[i])
  }

  *bwtresult <- r
  //*saresult <- sa
}

func main() {
  fmt.Print("Enter character string to compress: ")
  var original string
  fmt.Scanln(&original)
  //original = original + "$"
  processors := 2

  sl := ""

  bwtresult := make(chan string)


  wg := new(sync.WaitGroup)

  for i := 0; i < processors; i++ {
    wg.Add(1)
    go bwt(original[i*len(original)/processors:(i+1)*len(original)/processors], wg, &bwtresult)
  }
  go func(wg *sync.WaitGroup, bwtresult chan string) {
		wg.Wait()
		close(bwtresult)
	}(wg, bwtresult)


  count := 0
  for i:= range bwtresult{
    sl = sl + i
  }

  saslice := make([][]int, processors)
  saresult := make(chan []int)
  wg = new(sync.WaitGroup)

  for i := 0; i < processors; i++ {
    wg.Add(1)
    go sa(original[i*len(original)/processors:(i+1)*len(original)/processors], wg, &saresult)
  }
  go func(wg *sync.WaitGroup, saresult chan []int) {
		wg.Wait()
		close(saresult)
	}(wg, saresult)

  for i := range saresult{
    saslice[count] = i
    count++
  }

  total := 0
  elements := 0
  for i:= 0; i < len(saslice); i++{
    for j := 0; j < len(saslice[i]); j++{
      saslice[i][j] += total
      elements += 1
    }
    total += elements
    elements = 0
  }

  single := make([]int, len(original))

  for i := 0; i < len(saslice); i++{
    for j := 0; j < len(saslice[i]); j++{
      single[i*len(saslice[i]) + j] = saslice[i][j]
    }
  }

  for i := 0; i < len(single); i++{
    if single[i] == total/processors{
      single[i] = 0
    } else if single[i] == 0 {
      single[i] = total/processors
    }
  }

  totalsa := linearsa(original)
  bwtoriginal := linearbwt(original)


  final := ""
  for i := 0; i < len(totalsa); i++{
    for j := 0; j < len(single); j++{
      if totalsa[i] == single[j]{
        final += sl[j:j+1]
      }
    }
  }

  fmt.Println(single)
  fmt.Println(sl)
  fmt.Println(totalsa)
  fmt.Println(bwtoriginal)
  fmt.Println(final)
  fmt.Println("Are they equal? ", bwtoriginal == final)


}
