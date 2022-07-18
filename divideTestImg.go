package main

import (
  "fmt"
  "io/ioutil"
  "strconv"
  "os"
)

func main(){
  INPUTDIR := "./input/"
  VALDIR := "./animeFaces/val/"
  TESTDIR := "./animeFaces/test/"
  TRAINDIR := "./animeFaces/train/"


  fmt.Println("divide files for test data")
  files, err := ioutil.ReadDir(INPUTDIR)
  if err != nil{
    panic(err)
  }

  filelen := len(files)
  outdir := TESTDIR
  filecnt := 1
  for i:= 0 ; i < filelen ; i++ {

    fromfile := INPUTDIR +  files[i].Name()
    if  i == 107 {
      filecnt = 1
      outdir = TRAINDIR
    }
    if i ==  507 {
      filecnt = 1
      outdir = VALDIR
    }
    if i > 607 {
      return
    }
    s := strconv.Itoa(filecnt)
    tofile := outdir + s + ".jpg"
    filecnt++

    err := os.Rename(fromfile,tofile)
    if err != nil {
      fmt.Println(err)
    }
    fmt.Println( fromfile + " to " + tofile)
  }

  fmt.Println("image convert end")
}
