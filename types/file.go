package types

//go:generate ffjson $GOFILE

type File struct {
	ID             FileID         `json:"id"`
	FileName       String         `json:"file_name"`
	FileOwner      AccountID      `json:"file_owner"`
	CreateTime     Time           `json:"create_time"`
	FileContent    string         `json:"file_content"`
	RelatedAccount AccountIDArray `json:"related_account"`
	Signature      SignaturesType `json:"signature"`
	ParentFile     *FileID        `json:"parent_file"`
	SubFile        FileIDArray    `json:"sub_file"`
}

