package main

import (
  "fmt"
  "gocv.io/x/gocv"
  "io/ioutil"
)

func main(){
  dir := "./images"
  fmt.Println("start");
  files, err := ioutil.ReadDir(dir)
  if err != nil{
    panic(err)
  }

  fmt.Println(dir);
  edge_thresh := 100
  for _,file := range files{

    filename := dir + "/" + file.Name()

    imageMat := gocv.IMRead(filename, gocv.IMReadAnyColor)

    gray := gocv.NewMat()
    gocv.CvtColor(imageMat,&gray,gocv.ColorBGRToGray)

    edge := gocv.NewMat()

    gocv.Canny(gray,&edge,float32(edge_thresh),float32(edge_thresh*3))
    gocv.IMWrite("./out/" + file.Name(),edge)

    defer gray.Close()
    defer edge.Close()
  }
}
