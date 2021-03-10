package actions

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/daimon99/dgo2/utils"
	"github.com/urfave/cli/v2"
)

func Enc(c *cli.Context) error {
	var filename string
	var extra []string
	if c.NArg() > 1 {
		cli.ShowCommandHelp(c, "dec")
		return nil
	} else if c.NArg() == 1 {
		filename = c.Args().First()
	}
	if filename != "" {
		extra = []string{"-in", filename, "-out", filename + ".enc"}
	}
	salt_key := utils.Prompt("请输入加密Key", "")
	var oneline_flag string

	if c.Bool("oneline") {
		oneline_flag = "-A"
	} else {
		oneline_flag = "-a"
	}
	cmds := []string{"aes-256-cbc", "-k", salt_key, oneline_flag, "-md", "md5", "-base64"}
	cmds = append(cmds, extra...)

	cmd := exec.Command("openssl", cmds...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	if filename == "" {
		fmt.Println("请输入要加密的内容, ctrl+d结束输入")
	}
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	if filename != "" {
		fmt.Println("加密完成，见文件", filename+".enc")
	}
	return nil
}
