package actions

import (
	"fmt"
	"log"
	"os"
	"path"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/urfave/cli/v2"
)

/*
pip_conf_file_content = """
[global]
index-url=https://mirrors.aliyun.com/pypi/simple/
trusted-host=
    mirrors.daimon.cc
    mirrors.cloud.tencent.com
    mirrors.aliyun.com
"""
    with open('/etc/pip.conf', 'w') as fout:
        fout.write(pip_conf_file_content)
    click.secho(u'/etc/pip.conf 文件创建成功', fg='green')
*/

const pip_conf_file_content = `[global]
index-url=https://mirrors.aliyun.com/pypi/simple/
trusted-host=
    mirrors.daimon.cc
    mirrors.cloud.tencent.com
    mirrors.aliyun.com
`

func Pipconf(c *cli.Context) error {
	// global pip conf
	global_pip_conf := "/etc/pip.conf"
	f, err := os.OpenFile(global_pip_conf, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		user_pip_conf_dir, _ := homedir.Expand("~/.pip")
		user_pip_conf := path.Join(user_pip_conf_dir, "pip.conf")
		os.Mkdir(user_pip_conf_dir, 0655)
		f, err = os.OpenFile(user_pip_conf, os.O_RDWR|os.O_CREATE, 0644)
		if err != nil {
			log.Fatal(err)
		} else {
			defer f.Close()
			fmt.Println("用户pip.conf生成成功。", user_pip_conf)
		}
	} else {
		f.Write([]byte(pip_conf_file_content))
		defer f.Close()
		fmt.Println("全局pip.conf生成成功：", global_pip_conf)
	}
	return nil
}
