package main

import (
  "fmt"
  "os"
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

  fmt.Println(string(file))

  language := Grammar { translation: make(map[string]string), whitespace: []byte{ ' ', '\n', 0x09 } }
  language.translation["lsa"] = "if"

  for orig, trans := range language.translation {
    for i := range file {
      if i + len(orig) < len(file) && orig[0] == file[i] {
        match := true
        for j := range len(orig) - 1 {
          if orig[j] != file[i+j] {
            match = false
          }
        }
        if match {
          for _, char := range language.whitespace {
            if file[i+len(orig)] == char {
              file = append(file[:i], append([]byte(trans), file[i+len(orig):]...)...)
            }
          }
        }
      }
    }
  }
  fmt.Println(string(file))
}
