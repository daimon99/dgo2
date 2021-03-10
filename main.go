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
		Version:  "v1.1.0",
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
			{
				Name:   "pipconf",
				Usage:  "生成pip.conf。先生成全局的，失败会只生成用户空间的",
				Action: actions.Pipconf,
			},
			{
				Name:   "pypirc",
				Usage:  "生成 .pypirc 文件",
				Action: actions.Pypirc,
			},
			{
				Name:  "enc",
				Usage: "加密",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "oneline", Value: false},
				},
				Action: actions.Enc,
			},
			{
				Name:  "dec",
				Usage: "解密",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "oneline", Value: false},
				},
				Action: actions.Dec,
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
