package plugins_core

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type ServiceInput struct {
	ClusterInfo ClusterInfo
	Data        []byte
}

func (input *ServiceInput) String() string {
	return string(input.Data)
}

func MainPipeline(call map[string]func(clusterInfo *ClusterInfo, out io.Writer, data []byte) error) {
	flags := ParseFlags()

	serviceInput := &ServiceInput{}
	json.NewDecoder(os.Stdin).Decode(serviceInput)

	callFunc, ok := call[flags.Call]
	if !ok {
		fmt.Fprintf(os.Stderr, "wrong service: %s", flags.Call)
		return
	}

	if err := callFunc(&serviceInput.ClusterInfo, os.Stdout, serviceInput.Data); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
