package bind

import (
	"fmt"
	"path/filepath"

	"github.com/saltbo/zpan/model"
)

type QueryFiles struct {
	QueryPage
	Dir     string `form:"dir"`
	Type    string `form:"type"`
	Search  bool   `form:"search"`
	Keyword string `form:"keyword"`
}

type BodyFile struct {
	Name   string `json:"name" binding:"required"`
	Size   int64  `json:"size" binding:"required"`
	Type   string `json:"type"`
	Dir    string `json:"dir"`
	Public bool   `json:"public"`
}

func (p *BodyFile) ToMatter(uid int64) *model.Matter {
	if p.Type == "" {
		p.Type = "application/octet-stream"
	}
	m := model.NewMatter()
	m.Uid = uid
	m.Name = p.Name
	m.Type = p.Type
	m.Size = p.Size
	m.Parent = p.Dir
	m.Object = fmt.Sprintf("%d/%s%s", uid, m.Alias, filepath.Ext(p.Name))
	if p.Public {
		m.ACL = model.AclPublic
	}

	return m
}

type BodyFileRename struct {
	NewName string `json:"name" binding:"required"`
}

type BodyFileMove struct {
	NewDir string `json:"dir"`
}

type BodyFileCopy struct {
	NewPath string `json:"path" binding:"required"`
}
