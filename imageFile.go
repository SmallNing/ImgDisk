package main

import (
	"fmt"
	"os"
	"image"
)

// 最大文件
const FILEMAXSIZE  = 4294967296-32

type ImageFile struct {
	File os.File
	XMax int32
	YMax int32
}

func NwImageFile(file os.File,xMax int32,yMax int32) (*ImageFile,error) {
	var result = new(ImageFile)
	if(xMax < 1 || yMax < 1){
		return nil,error("像素必须存在")
	}
	result.File = file
	result.XMax = xMax
	result.YMax = yMax
	return result,nil
}

func (imageFile ImageFile) Count() (int32,error)  {
	result,err := imageFile.File.Stat()
	if(err != nil){
		return -1,err
	}else{
		return int32(result.Size())/(imageFile.YMax*imageFile.XMax*3-32),nil
	}
	return -1,nil
}

func (imageFile ImageFile) CovertToImg(i int32) (image.Config,error) {
	count,err := imageFile.Count()
	if(err != nil){

	}
	if(i>count){
		return image.Config{},error("数字大于长度")
	}
	img,_,err := image.DecodeConfig(&imageFile.File)
	if(err != nil){
		return image.Config{},err
	}
	return img,nil
}


func main() {

	fmt.Println("Hello World!")
}
