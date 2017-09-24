package main

import (
	"fmt"
)

// generate img max size
const IMG_MAX_SIZE  = 4294967296
// img RGB level
const IMG_LEVEL  = 3
// img file header (byte)
const IMAGE_HEADER_SIZE  = 4

type ImageFile struct {
	file []byte
	xMax int
	yMax int
}

func NewImageFile(file []byte,xMax int,yMax int) (*ImageFile,error) {
	var result = new(ImageFile)
	if(xMax < 1 || yMax < 1){
		return nil,error("像素必须存在")
	}
	result.file = file
	result.xMax = xMax
	result.yMax = yMax
	return result,nil
}

func (imageFile ImageFile) Count() (int,error)  {
	// use file calc image count
	if(imageFile.file == nil){
		return -1,error("file bytes is nil")
	}
	count := len(imageFile.file)/(imageFile.xMax*imageFile.yMax*IMG_LEVEL-IMAGE_HEADER_SIZE)
	return count,nil
}

func (imageFile ImageFile) covert(fileBytes []byte,i int) ([]byte,error){
	// get []byte size
	count,err := imageFile.Count()
	if(err != nil){
		return nil,nil
	}
	if(count < i){
		return nil,error("i more than image max count")
	}
	fileToImgByteSize := imageFile.xMax*imageFile.yMax*IMG_LEVEL-IMAGE_HEADER_SIZE
	startByteIndex := (i-1)*fileToImgByteSize
	if(count == i){
		imgBytes := imageFile.file[startByteIndex:len(imageFile.file)]
		return imgBytes,nil
	}else{
		imgBytes := imageFile.file[startByteIndex:startByteIndex+fileToImgByteSize]
		return imgBytes,nil
	}
}

func main() {

	fmt.Println("Hello World!")
}
