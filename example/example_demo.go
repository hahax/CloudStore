package main

import (
	"fmt"
	"github.com/hahax/CloudStore"
	"io/ioutil"
	"os"
)

var (
	Minio       *CloudStore.MinIO
	objectMinio = "minio.go"
	err         error

	objectSVG      = "../test_data/test.svg"             //未经过gzip压缩的svg图片
	objectSVGGzip  = "../test_data/test.gzip.svg"        //gzip压缩后的svg图片
	objectHtml     = "../test_data/helloworld.html"      //未经gzip压缩的HTML
	objectHtmlGzip = "../test_data/helloworld.gzip.html" //gzip压缩后的HTML
	objectNotExist = "not exist object"
	objectPrefix   = "test_data"
	objectDownload = "../test_data/download.svg"
	headerGzip     = map[string]string{"Content-Encoding": "gzip"}
	headerSVG      = map[string]string{"Content-Type": "image/svg+xml"}
	headerHtml     = map[string]string{"Content-Type": "text/html; charset=UTF-8"}
)

func init() {
	key := "your minio ak"
	secret := "your minio sk"
	bucket := "your bucket"
	domain := "http://you domain:9000"
	endpoint := "your domain:9000"
	Minio, err = CloudStore.NewMinIO(key, secret, bucket, endpoint, domain)
	if err != nil {
		panic(err)
	}
}

func main() {
	//upload
	err = Minio.Upload(objectSVG, objectSVG, headerSVG)
	if err != nil {
		fmt.Println(err)
	}
	err = Minio.Upload(objectSVGGzip, objectSVGGzip, headerGzip, headerSVG)
	if err != nil {
		fmt.Println(err)
	}
	//("=====IsExist=====")
	fmt.Println(objectSVG, "is exist?(Y):", Minio.IsExist(objectSVG) == nil)
	fmt.Println(objectNotExist, "is exist?(N):", Minio.IsExist(objectNotExist) == nil)
	//("=====Lists=====")
	if files, err := Minio.Lists(objectPrefix); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(files)
	}
	//("=====GetInfo=====")
	if info, err := Minio.GetInfo(objectSVG); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(info)
	}
	//("=====Download=====")
	if err := Minio.Download(objectSVG, objectDownload); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("download success")
		b, _ := ioutil.ReadFile(objectDownload)
		fmt.Println("Content:", string(b))
		os.Remove(objectDownload)
	}
	//("====GetSignURL====")
	fmt.Println(Minio.GetSignURL(objectSVG, 120))
	fmt.Println(Minio.GetSignURL(objectSVGGzip, 120))
	//("========Finished========")
}
