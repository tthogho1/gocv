package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"io/ioutil"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	inputdir := "./images/"
	resizefolder := "./resize/"
	eddgefolder := "./edge/"
	out := "./outdata/"

	imgSize := 256
	outImgWidth := 512
	outImgHeight := 256

	fmt.Println("image convert start")
	files, err := ioutil.ReadDir(inputdir)
	if err != nil {
		panic(err)
	}

	// エッジ抽出の閾値
	edge_thresh := 100
	for _, file := range files {

		filename := inputdir + file.Name()
		// 画像のリサイズ
		resizeMat := gocv.IMRead(filename, gocv.IMReadAnyColor)
		gocv.Resize(resizeMat, &resizeMat, image.Point{imgSize, imgSize}, 0, 0, gocv.InterpolationDefault)
		gocv.IMWrite(resizefolder+file.Name(), resizeMat)
		resizeReader, _ := os.Open(resizefolder + file.Name())

		// グレースケールからエッジ抽出
		gray := gocv.NewMat()
		gocv.CvtColor(resizeMat, &gray, gocv.ColorBGRToGray)
		edge := gocv.NewMat()
		gocv.Canny(gray, &edge, float32(edge_thresh), float32(edge_thresh*3))
		gocv.IMWrite(eddgefolder+file.Name(), edge)
		edgeReader, _ := os.Open(eddgefolder + file.Name())

		// 画像の結合　左：エッジ抽出　右：リサイズ
		outImg := image.NewRGBA(image.Rect(0, 0, outImgWidth, outImgHeight))
		edgeSrc, _, _ := image.Decode(edgeReader)
		rect1 := image.Rect(0, 0, imgSize, imgSize)
		draw.Draw(outImg, rect1, edgeSrc, image.Point{0, 0}, draw.Over)

		resizeSrc, _, _ := image.Decode(resizeReader)
		rect2 := image.Rect(imgSize, 0, imgSize+imgSize, imgSize)
		draw.Draw(outImg, rect2, resizeSrc, image.Point{0, 0}, draw.Over)

		// 画像の保存
		outname, _ := os.Create(out + file.Name())
		qt := jpeg.Options{
			Quality: 100,
		}
		err6 := jpeg.Encode(outname, outImg, &qt)
		if err6 != nil {
			fmt.Fprintln(os.Stderr, err6)
			return
		}

		resizeMat.Close()
		gray.Close()
		edge.Close()
	}

	fmt.Println("image convert end")
}
