package plugins_core

import (
	"fmt"
	"io"
	"os"
	"text/template"
)

const (
	BaseStylePath = "assets/web_interface/style.css"
)

func PageHandler(
	PageHtmlPath string,
	PageStylePath string,
	PageScriptPath string,
	clusterInfo *ClusterInfo,
	out io.Writer,
	data []byte,
) error {
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

	pageInfo := PageInfo{
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
