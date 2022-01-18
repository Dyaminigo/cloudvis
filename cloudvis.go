package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"time"

	requests "github.com/carlmjohnson/requests"
)

type UploadResponse struct {
	Status string `json:"status"`
	Id     string `json:"id"`
}

type RecognitionResult struct {
	UploadResponse
	Text string `json:"text"`
	QR   string `json:"qr"`
}

func uploadImage(url string, params map[string]string, fileData []byte) (*UploadResponse, error) {
	var upResponse UploadResponse

	rbody := new(bytes.Buffer)
	writer := multipart.NewWriter(rbody)
	defer writer.Close()
	filepart, _ := writer.CreateFormFile("file", "img")
	filepart.Write(fileData)
	for pname, pval := range params {
		writer.WriteField(pname, pval)
	}

	err := requests.
		URL(url).
		BodyBytes(rbody.Bytes()).
		ContentType(writer.FormDataContentType()).
		ToJSON(&upResponse).
		Fetch(context.Background())

	return &upResponse, err
}

func getResult(url string, id string) (*RecognitionResult, error) {
	var result RecognitionResult

	err := requests.
		URL(url).
		Param("id", id).
		ToJSON(&result).
		Fetch(context.Background())

	return &result, err
}

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage: %s [opts] [file] \n file is a image file name. png and jpg are officially supported. \n Opts can be one of:\n", os.Args[0])
		flag.PrintDefaults()
	}
	useStdin := flag.Bool("stdin", false, "Get image data from stdin, filename will be ignored")
	lang := flag.String("l", "en", "Set `language` name to use. Language is formed as two letter code, like ru or en")
	target := flag.String("t", "all", "Data `type` to recognize. Can be 'text', 'image' or 'all'")
	translate := flag.Bool("tr", false, "Translate recognition result")
	qr := flag.Bool("qr", false, "Recognize QR/BAR codes")
	flag.Parse()

	var file *os.File
	if *useStdin {
		file = os.Stdin
	} else {
		var filename = flag.Arg(0)
		if filename == "" {
			log.Fatal("File name not specified")
		}
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			log.Fatalf("File %v does not exist", filename)
		}
		var err error
		file, err = os.Open(filename)
		if err != nil {
			log.Fatal(err)
		}
	}

	filedata, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	const apiUrl = "https://visionbot.ru/apiv2"
	params := map[string]string{
		"lang":      *lang,
		"target":    *target,
		"qr":        "0",
		"translate": "0",
	}

	if *qr {
		params["qr"] = "1"
	}
	if *translate {
		params["translate"] = "1"
	}
	fmt.Println("Uploading image...")
	resp, err := uploadImage(apiUrl+"/in.php", params, filedata)
	if err != nil {
		log.Fatal(err)
	}

	if resp.Status != "ok" {
		log.Fatal("Server error. Try again later")
	}

	id := resp.Id
	fmt.Println("Waiting for result...")
	for {
		res, err := getResult(apiUrl+"/res.php", id)
		if err != nil {
			log.Fatal(err)
		}

		if res.Status == "error" {
			log.Fatal("Server error. Try again later")
		} else if res.Status == "ok" {
			if res.Text == "" {
				fmt.Println("No text")
			} else {
				fmt.Println(res.Text)
			}
			if res.QR != "" {
				fmt.Printf("\nQR:\n%v", res.QR)
			}
			break
		}
		time.Sleep(1 * time.Second)
	}
}
