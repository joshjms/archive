package fs

type Directory struct {
	ID       int    `gorm:"primaryKey"`
	Name     string // Directory name
	Owner    int    // Owner of the directory, this references a user ID
	ParentID int    // Optional parent directory ID, allows for root directories
	Public   bool   // Whether the directory is public or not
}
