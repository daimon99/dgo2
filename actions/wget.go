package actions

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func Wget(c *cli.Context) error {
	if c.NArg() != 1 {
		cli.ShowCommandHelp(c, "wget")
		return nil
	}
	cmd := exec.Command("wget", "--content-disposition", c.Args().First())

	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("执行命令完成：\n%s\n", cmd.String())
	return nil
}
