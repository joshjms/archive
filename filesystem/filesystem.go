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

func (fs *FileSystem) CreateFile(file *File) error {
	return fs.DB.Create(file).Error
}

func (fs *FileSystem) GetFile(id int) (*File, error) {
	file := &File{}
	err := fs.DB.First(file, id).Error
	return file, err
}

func (fs *FileSystem) GetFileFromSlug(slug string) (*File, error) {
	file := &File{}
	err := fs.DB.Where("name = ?", slug).First(file).Error
	return file, err
}
