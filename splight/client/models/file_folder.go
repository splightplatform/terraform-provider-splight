package models

type FileFolderParams struct {
	Name   string `json:"name"`
	Parent string `json:"parent"`
}

type FileFolder struct {
	FileFolderParams
	Id string `json:"id"`
}
