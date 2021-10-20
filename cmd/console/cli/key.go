package cli

import (
	"bufio"
	"bytes"
	"fagin/config"
	"fagin/utils"
	"github.com/spf13/cobra"
	"log"
	"os"
	"strings"
)

const tag string = "APP_KEY="

func createAppKey(filePath string) {
	bs, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// O_TRUNC 打开并清空文件
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_TRUNC, 666)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	s := bufio.NewScanner(bytes.NewBuffer(bs))
	w := bufio.NewWriter(file)

	for s.Scan() {
		if strings.HasPrefix(s.Text(), tag) {
			str := utils.RandString(32)
			if _, err = w.WriteString(tag + str + "\n"); err != nil {
				log.Fatal(err)
			}
			continue
		}

		if _, err = w.WriteString(s.Text() + "\n"); err != nil {
			log.Fatal(err)
		}
	}

	if err = w.Flush(); err != nil {
		log.Fatal(err)
	}
}

var KeyCmd = &cobra.Command{
	Use:   "key",
	Short: "生成 key",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		createAppKey(config.App().RootPath + `/.env`)
		log.Printf("key generate run successfully")
	},
}
