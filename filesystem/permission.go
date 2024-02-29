package fs

type FilePermission struct {
	UserID int
	FileID int
	Type   string // "read", "write"
}

type DirectoryPermission struct {
	UserID      int
	DirectoryID int
	Type        string // "read", "write"
}
