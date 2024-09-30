package filesystem_handlers

type CollectFilesytemRequest struct {
	Path   string `json:"path"`
	Search string `json:"search"`
}

type CreateOperation struct {
	Type string `json:"type"`
	Path string `json:"path"`
}

type DeleteOperation struct {
	Path string `json:"path"`
}

type SelectOperation struct {
	Path string `json:"path"`
}

type CopyOperation struct {
	SrcUrl  string `json:"src_url"`
	SrcPath string `json:"src_path"`
	DstPath string `json:"dst_path"`
}

type MoveOperation CopyOperation

type RenameOperation struct {
	SrcPath string `json:"src_path"`
	DstPath string `json:"dst_path"`
}
