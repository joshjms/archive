package fs

type Directory struct {
	ID       int    `gorm:"primaryKey"`
	Name     string // Directory name
	ParentID int    // Optional parent directory ID, allows for root directories
	Public   bool   // Whether the directory is public or not
	Slug     string `gorm:"unique"` // Unique slug for the directory
}
