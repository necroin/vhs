package plugins_core

type ServiceInput struct {
	ClusterInfo ClusterInfo
	Data        []byte
}

func (input *ServiceInput) String() string {
	return string(input.Data)
}
