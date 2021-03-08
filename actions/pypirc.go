package actions

import (
	"fmt"
	"log"
	"os"

	"github.com/daimon99/dgo2/utils"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

const pypirc_content = `[distutils]
index-servers =
    pypi

[pypi]
repository: https://upload.pypi.org/legacy/
username: %s
password: %s

`

func Pypirc(c *cli.Context) error {
	pypirc_path, _ := homedir.Expand("~/.pypirc")
	is_exist, err := utils.IsExist(pypirc_path)
	if err != nil {
		log.Fatal(err)
	}
	if is_exist {
		return fmt.Errorf("文件已经存在，不覆盖退出！", pypirc_path)
	}
	fmt.Print("username: ")
	var (
		username string
		password string
	)
	fmt.Scanln(&username)
	fmt.Print("password: ")
	fmt.Scanln(&password)
	content := fmt.Sprintf(pypirc_content, username, password)
	f, _ := os.OpenFile(pypirc_path, os.O_RDWR|os.O_CREATE, 0644)
	defer f.Close()
	f.Write([]byte(content))
	f.Close()
	fmt.Printf("文件生成成功: %s\n", pypirc_path)
	return nil

}
