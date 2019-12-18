package io_text

import (
    "io/ioutil"
    "os"
    "log"
)

func ReadTxt(s string) string{
    file, err := os.Open(s)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

  b, err := ioutil.ReadAll(file)
  return string(b)
  // fmt.Print(b)
}