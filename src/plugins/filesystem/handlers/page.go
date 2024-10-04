package filesystem_handlers

import (
	"fmt"
	"io"
	"os"
	"text/template"
	plugins_core "vhs/src/plugins/core"
)

const (
	BaseStylePath  = "assets/web_interface/style.css"
	PageHtmlPath   = "assets/web_interface/plugins/filesystem/explorer/explorer.html"
	PageStylePath  = "assets/web_interface/plugins/filesystem/explorer/explorer.css"
	PageScriptPath = "assets/web_interface/plugins/filesystem/explorer/explorer.js"
)

func PageHandler(clusterInfo *plugins_core.ClusterInfo, out io.Writer, data []byte) error {
	pageHtmlData, err := os.ReadFile(PageHtmlPath)
	if err != nil {
		return err
	}

	baseStyleData, err := os.ReadFile(BaseStylePath)
	if err != nil {
		return err
	}

	pageStyleData, err := os.ReadFile(PageStylePath)
	if err != nil {
		return err
	}

	pageScriptData, err := os.ReadFile(PageScriptPath)
	if err != nil {
		return err
	}

	pageInfo := plugins_core.PageInfo{
		BaseStyle: fmt.Sprintf("<style>%s</style>", baseStyleData),
		Style:     fmt.Sprintf("<style>%s</style>", pageStyleData),
		Script:    fmt.Sprintf("<script>%s</script>", pageScriptData),
	}

	pageTemplate, err := template.New("page").Parse(string(pageHtmlData))
	if err != nil {
		return err
	}

	if err := pageTemplate.Execute(out, pageInfo); err != nil {
		return err
	}
	return nil
}
