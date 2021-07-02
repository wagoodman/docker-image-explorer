package export

import (
	"github.com/wagoodman/dive/dive/filetree"
)

type layer struct {
	Index     int                 `json:"index"`
	ID        string              `json:"id"`
	DigestID  string              `json:"digestId"`
	SizeBytes uint64              `json:"sizeBytes"`
	Command   string              `json:"command"`
	FileList  []filetree.NodeData `json:"fileList"`
}
