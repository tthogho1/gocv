package main

import (
  "os"
  "fmt"
  "io/ioutil"
  "image"
  "image/draw"
  "image/jpeg"
)

func main(){
  inputdir := "./convert/"
  resize := "./resize/"
  out := "./testdata/"

  imgSize := 256
  outImgWidth := 512
  outImgHeight := 256

  fmt.Println("create image start");
  files, err := ioutil.ReadDir(inputdir)
  if err != nil{
    fmt.Fprintln(os.Stderr,err)
  }

  for _,file := range files{

    filename := inputdir +  file.Name()
    src,err := os.Open(filename)
    if err != nil{
      fmt.Println("error 1")
      fmt.Fprintln(os.Stderr,err)
    }

    resizename := resize + file.Name()
    resizesrc,err2 := os.Open(resizename)
    if err2 != nil{
      fmt.Println("error 2")
      fmt.Fprintln(os.Stderr,err2)
    }

    imgSrc , _, err3 :=  image.Decode(src)
    if err3 != nil{
      fmt.Println("error 3")
      fmt.Fprintln(os.Stderr,err3)
    }

    resizeSrc ,_, err4 := image.Decode(resizesrc)
    if err4 != nil{
      fmt.Println("error 4")
      fmt.Fprintln(os.Stderr,err4)
    }

    outImg := image.NewRGBA(image.Rect(0,0,outImgWidth,outImgHeight))

    rect1 := image.Rect(0,0,imgSize,imgSize)
    draw.Draw(outImg,rect1,imgSrc,image.Point{0,0},draw.Over)

    rect2 := image.Rect(imgSize,0,imgSize + imgSize,imgSize)
    draw.Draw(outImg,rect2,resizeSrc,image.Point{0,0},draw.Over)

    outname ,err5 := os.Create(out + file.Name())
    if err5 != nil {
      fmt.Println("error 5")
      fmt.Fprintln(os.Stderr,err5)
      return
    }

    qt := jpeg.Options{
      Quality:100,
    }
    err6 := jpeg.Encode(outname,outImg,&qt)
    if err6 != nil {
      fmt.Println("error 6")
      fmt.Fprintln(os.Stderr,err6)
    }

    resizesrc.Close()
    src.Close()
    outname.Close()

  }

  fmt.Println("image convert end");
}
