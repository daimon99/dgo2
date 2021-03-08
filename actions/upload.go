package actions

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/urfave/cli/v2"
)

func Upload(c *cli.Context) error {
	if c.Args().Len() != 1 {
		cli.ShowCommandHelp(c, "upload")
		return nil
	}

	local_file_path := c.Args().First()

	bodyBuf := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyBuf)
	fileWriter, err := writer.CreateFormFile("file", local_file_path)
	f, err := os.Open(local_file_path)
	if err != nil {
		log.Fatal("文件打开失败")
		return nil
	}
	io.Copy(fileWriter, f)
	writer.Close()

	r, err := http.Post(
		"http://tmp.daimon.cc:10080/upload",
		writer.FormDataContentType(),
		bodyBuf,
	)
	body, _ := ioutil.ReadAll(r.Body)
	fields := strings.Fields(string(body))
	if len(fields) != 2 {
		log.Fatal("服务异常：%s", body)
		return fmt.Errorf("服务异常: %s", body)
	}

	fmt.Printf("wget %s\n", fields[1])
	return nil
}
