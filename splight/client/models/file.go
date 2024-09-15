package models

type FileParams struct {
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Parent        string         `json:"parent"`
	RelatedAssets []RelatedAsset `json:"assets"`
}

type FileURL struct {
	URL string `json:"url"`
}

type FileDetails struct {
	Checksum     string `json:"checksum"`
	LastModified string `json:"last_modified"`
	Size         string `json:"size"`
}

type File struct {
	FileParams
	ID string `json:"id"`
}
