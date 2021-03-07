package main

import (
	"fmt"
	"log"
	"os"
	"time"

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
				ArgsUsage: "[file_path]",
				Action: func(c *cli.Context) error {
					var local_file_path = c.Args().First()

					fmt.Printf("%#v\n", local_file_path)
					return nil
				},
			},
		},
	}
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
