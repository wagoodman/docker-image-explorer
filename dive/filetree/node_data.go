package filetree

var GlobalFileTreeCollapse bool

// NodeData is the payload for a FileNode
type NodeData struct {
	ViewInfo ViewInfo
	FileInfo FileInfo `json:"fileInfo"`
	DiffType DiffType
}

// NewNodeData creates an empty NodeData struct for a FileNode
func NewNodeData() *NodeData {
	return &NodeData{
		ViewInfo: *NewViewInfo(),
		FileInfo: FileInfo{},
		DiffType: Unmodified,
	}
}

// Copy duplicates a NodeData
func (data *NodeData) Copy() *NodeData {
	return &NodeData{
		ViewInfo: *data.ViewInfo.Copy(),
		FileInfo: *data.FileInfo.Copy(),
		DiffType: data.DiffType,
	}
}
