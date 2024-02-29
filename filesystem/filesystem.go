package fs

import (
	"gorm.io/gorm"
)

type FileSystem struct {
	DB *gorm.DB
}

func New(db *gorm.DB) *FileSystem {
	return &FileSystem{
		DB: db,
	}
}

func (fs *FileSystem) Init() {
	fs.DB.AutoMigrate(&File{})
	fs.DB.AutoMigrate(&Directory{})
	fs.DB.AutoMigrate(&FilePermission{})
	fs.DB.AutoMigrate(&DirectoryPermission{})
}
