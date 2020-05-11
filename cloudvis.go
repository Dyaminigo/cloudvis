package main
import "time"
import "fmt"
import "net/http"
import "encoding/json"
import "mime/multipart"
import "io/ioutil"
import "os"
import "log"
import "flag"
import "bytes"
var ApiUrl string = "https://visionbot.ru/apiv2"
func multipartuploadreq(url string, params map[string]string, filedata []byte) *http.Request {
rbody := new(bytes.Buffer)
writer := multipart.NewWriter(rbody)
filepart,_ := writer.CreateFormFile("file", "img")
filepart.Write(filedata)
for parname,parval := range params {
writer.WriteField(parname, parval)
}
err := writer.Close()
if err != nil {
log.Fatal(err)
}
req,_ := http.NewRequest("POST", "https://visionbot.ru/apiv2/in.php", rbody)
req.Header.Set("Content-Type", writer.FormDataContentType())
return req
}
func main() {
flag.Usage = func() {
fmt.Fprintf(os.Stderr, "Usage: %s [opts] [file] \n file is a image file name. png and jpg are officially supported. \n Opts can be one of:\n", os.Args[0])
flag.PrintDefaults()
}
usestdin := flag.Bool("stdin", false, "get image data from stdin, filename will be ignored")
lang := flag.String("l", "en", "set `language` name to use. language is formed as two letter code, like ru or en, default is en.")
target := flag.String("t", "all", "data `type` to recognize. can be 'text', 'image' or 'all'.")
flag.Parse()
var file *os.File
var err error
if *usestdin {
file = os.Stdin
} else {
file,err = os.Open(flag.Arg(0))
}
if err != nil {
log.Fatal(err)
}
filedata,err := ioutil.ReadAll(file)
if err != nil {
log.Fatal(err)
}
file.Close()
params := map[string]string{"lang": *lang, "target": *target}
req := multipartuploadreq(ApiUrl + "/in.php", params, filedata)
os.Stderr.WriteString("uploading image\n")
resp,err := http.DefaultClient.Do(req)
if err != nil {
log.Fatal(err)
}
bodydata,err := ioutil.ReadAll(resp.Body)
if err != nil {
log.Fatal(err)
}
resp.Body.Close()
var bodyresult map[string]string
err = json.Unmarshal(bodydata, &bodyresult)
if err != nil {
log.Fatal(err)
}
if bodyresult["status"] != "ok" {
log.Fatal(fmt.Errorf("server error"))
}
id := bodyresult["id"]
os.Stderr.WriteString("waiting for result\n")
for {
time.Sleep(500 * time.Millisecond)
gresp,err := http.Get(ApiUrl + "/res.php?id=" + id)
if err != nil {
log.Fatal(err)
}
bodydata,err = ioutil.ReadAll(gresp.Body)
if err != nil {
log.Fatal(err)
}
bodyresult = nil
err = json.Unmarshal(bodydata, &bodyresult)
if err != nil {
log.Fatal(err)
}
if bodyresult["status"] == "error" {
log.Fatal(fmt.Errorf("server error"))
}
if bodyresult["status"] == "ok" {
fmt.Println(bodyresult["text"])
break
}
}
}