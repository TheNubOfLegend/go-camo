package main

import (
  "fmt"
  "os"
  "os/exec"
)

type Grammar struct {
  translation map[string]string
  whitespace []byte
}

func main() {
  file, err := os.ReadFile("test")
  if err != nil {
    fmt.Println("no file nerd")
  }

  if(err==nil){fmt.Println("nerd")}

  fmt.Println(string(file))

  language := Grammar { translation: make(map[string]string), whitespace: []byte{ ' ', '\n', 0x09, '(', ')', '{', '}', ';' } }
  //TODO: currently, translation could override previous translation
  language.translation["lsa"] = "if"

  for orig, trans := range language.translation {
    file = horspool(file, []byte(orig), []byte(trans))
  }
  fmt.Println(string(file))

  os.WriteFile(".camo/main.camo.go", file, os.FileMode(int(0777)))
  err = exec.Command("go", "run", ".camo/main.camo.go").Run()
  if err != nil {
    fmt.Println("nononononono")
  }
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

//also replaces
func horspool(file, pattern, trans []byte) []byte {
  var badChar [256]int

  for i := range pattern {
    badChar[pattern[i]] = i
  }

  offset := 0
  count := 0
  i := len(pattern) - 1
  for offset + len(pattern) < len(file) {
    if file[i + offset] != pattern[i] {
      offset += max(1, i - badChar[file[i + offset]])
      i = len(pattern) - 1
    } else if i == 0 {
      //TODO: probably inefficient as hell to resize array every time like this instead of doing mass collection and mass replacement
      file = append(file[:offset], append([]byte(trans), file[len(pattern) + offset:]...)...)
      count++
      fmt.Printf("%d found pattern at %d\n", count, offset)

      offset += len(trans)
      i = len(pattern) - 1
    } else {
      i--
    }
  }
  return file
}
