package fs

type File struct {
	ID       int    `gorm:"primaryKey"` // File ID
	Name     string `gorm:"not null"`   // Unique and non-null file name
	ParentID int    // ID of the parent directory
	Dir      string
	Public   bool
	Slug     string `gorm:"unique"` // Unique slug for the file
}
