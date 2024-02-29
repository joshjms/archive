package fs

type File struct {
	ID       int    `gorm:"primaryKey"`       // File ID
	Name     string `gorm:"unique; not null"` // Unique and non-null file name
	ParentID int    // ID of the parent directory
	Dir      string
	Public   bool
}

func (f *File) Read() {

}
