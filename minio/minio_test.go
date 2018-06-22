package main

import (
	"testing"
	"io/ioutil"
	"net/http"
	"bytes"
	"mime/multipart"
	"fmt"
	"os"
	"io"
)

func TestUp(t *testing.T) {
	if err := postFile("./test.txt", "test.txt", "http://127.0.0.1:8080/v1/up"); err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ok")
}

func TestFind(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/v1/find?id=1529647962089361100")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ok")
	fmt.Println(string(respBody))
}

func TestDel(t *testing.T) {
	var client = &http.Client{}
	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8080/v1/remove?id=1529647962089361100", nil)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	resp,err:=client.Do(req)
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println("ok")
	fmt.Println(string(respBody))
}

func postFile(filepath, filename string, targetUrl string) error {
	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)
	//关键的一步操作
	fileWriter, err := bodyWriter.CreateFormFile("up_file", filename)
	if err != nil {
		fmt.Println("error writing to buffer")
		return err
	}
	//打开文件句柄操作
	fh, err := os.Open(filepath)
	if err != nil {
		fmt.Println("error opening file")
		return err
	}
	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := http.Post(targetUrl, contentType, bodyBuf)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(respBody))
	return nil
}
