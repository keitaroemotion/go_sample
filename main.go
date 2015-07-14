package main

import (
  "io/ioutil"
  "os"
)

func main(){
  lines := []byte("This is a pen.\nhoahoami")
  ioutil.WriteFile("test2", lines, os.ModePerm)
}
