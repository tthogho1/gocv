package main

import (
  "fmt"
  "gocv.io/x/gocv"
  "io/ioutil"
  "image"
//  "image/draw"
)

func main(){
  inputdir := "./images/"
  resize := "./resize/"
  convert := "./convert/"

  imgSize := 256
  //outImgWidth := 512
  //outImgHeight := 256

  fmt.Println("image convert start");
  files, err := ioutil.ReadDir(inputdir)
  if err != nil{
    panic(err)
  }

  edge_thresh := 100
  for _,file := range files{

    filename := inputdir +  file.Name()

    imageMat := gocv.IMRead(filename, gocv.IMReadAnyColor)

    gocv.Resize(imageMat,&imageMat,image.Point{imgSize,imgSize},0,0,gocv.InterpolationDefault)

    gocv.IMWrite(resize + file.Name(),imageMat)

    gray := gocv.NewMat()
    gocv.CvtColor(imageMat,&gray,gocv.ColorBGRToGray)

    edge := gocv.NewMat()

    gocv.Canny(gray,&edge,float32(edge_thresh),float32(edge_thresh*3))

    gocv.IMWrite(convert + file.Name(),edge)


    defer gray.Close()
    defer edge.Close()
  }

  fmt.Println("image convert end");
}
