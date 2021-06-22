package view

import (
	"fmt"
	"github.com/gin-gonic/gin"
	t "html/template"
	"strings"
)

// DebugPrintLoadTemplate 打印加载的模版名称
func DebugPrintLoadTemplate(temp *t.Template) {
	var buf strings.Builder
	for _, tmpl := range temp.Templates() {
		buf.WriteString("\t- ")
		buf.WriteString(tmpl.Name())
		buf.WriteString("\n")
	}
	format := "Loaded HTML Templates (%d): \n%s\n"
	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}
	fmt.Fprintf(gin.DefaultWriter, "[GIN-debug] "+format, len(temp.Templates()), buf.String())
}