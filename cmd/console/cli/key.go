package cli

import (
	"bufio"
	"bytes"
	"fadmin/config"
	"fadmin/utils"
	"io/fs"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

const tag string = "APP_KEY="

func createAppKey(filePath string) {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// O_TRUNC 打开并清空文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, fs.ModePerm)

	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(bytes.NewBuffer(bs))
	w := bufio.NewWriter(file)

	const num = 32

	for s.Scan() {
		if strings.HasPrefix(s.Text(), tag) {
			str := utils.RandString(num)
			if _, err = w.WriteString(tag + str + "\n"); err != nil {
				log.Panic(err)
			}

			continue
		}

		if _, err = w.WriteString(s.Text() + "\n"); err != nil {
			log.Panic(err)
		}
	}

	if err = w.Flush(); err != nil {
		log.Panic(err)
	}
}

var KeyCmd = &cobra.Command{
	Use:   "key",
	Short: "生成 key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		createAppKey(config.App().RootPath() + `/.env`)
		log.Printf("key generate run successfully")
	},
}
