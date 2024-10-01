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
    occurrences := horspool(file, []byte(orig))
    for _, val := range occurrences {
      file = append(file[:val], append([]byte(trans), file[val+len(orig):]...)...)
    }

    // for i := range file {
    //   if i + len(orig) < len(file) && orig[0] == file[i] {
    //     match := true
    //     for j := range len(orig) - 1 {
    //       if orig[j] != file[i+j] {
    //         match = false
    //       }
    //     }
    //     if match {
    //       for _, char := range language.whitespace {
    //         if file[i+len(orig)] == char {
    //           file = append(file[:i], append([]byte(trans), file[i+len(orig):]...)...)
    //         }
    //       }
    //     }
    //   }
    // }
  }
  fmt.Println(string(file))
}

func max(x, y int) int {
  if x > y {
    return x
  }
  return y
}

func horspool(file, pattern []byte) []int {
  // file, err := os.ReadFile("test")
  // if err != nil {
  //     fmt.Println("no file nerd")
  // }

  // var pattern []byte = []byte("good")
  var result []int
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
      result = append(result, offset)
      count++
      fmt.Printf("%d found pattern at %d\n", count, offset)

      offset += len(pattern)
      i = len(pattern) - 1
    } else {
      i--
    }
  }
  return result
}
