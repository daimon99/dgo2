package main

import (
	"log"
	"os"
	"time"

	"github.com/daimon99/dgo2/actions"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:     "老戴工具箱",
		Version:  "v1.0.0",
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "戴健",
				Email: "daijian1@qq.com",
			},
		},
		Commands: []*cli.Command{
			{
				Name:      "upload",
				Usage:     "上传临时文件",
				ArgsUsage: "file_path",
				Action:    actions.Upload,
			},
			{
				Name:      "wget",
				Usage:     "下载文件。用content-disposition",
				ArgsUsage: "url",
				Action:    actions.Wget,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
