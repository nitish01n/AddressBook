package models

import (
	"io"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/nitish-126/phonebook/pkg/config"
)

var db *gorm.DB

type User struct {
	// gorm.Model
	Id       uint   `gorm:"column:id"`
	Name     string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
	Mobileno string `gorm:"column:mobileno"`
	Address  string `gorm:"column:address"`
	// FileName string `gorm:"column:filename"`
	// filedata string `gorm:"column:filedata"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&User{})
}

func (b *User) CreateUser() *User {
	db.NewRecord(b)
	db.Table("contacts").Create(&b)
	db.Table("contacts").Save(&b)
	return b
}

func GetUsers() []User {
	// var r *http.Request
	var users []User
	db.Table("contacts").Debug().Find(&users)
	// for i := range users {
	// 	users[i].filedata = base64.StdEncoding.EncodeToString([]byte(users[i].filedata))
	// }
	// fmt.Println(users)
	// file, handler, _ := r.FormFile("file")

	// jsonUsers, _ := json.Marshal(users)
	return users
}

func GetUserByName(name string) *User {
	var getUser User
	db.LogMode(true)
	db.Table("contacts").Where("name=?", name).Find(&getUser)
	return &getUser
}

func GetUserById(Id string) *User {
	var getUser User
	db.LogMode(true)
	db.Table("contacts").Where("id=?", Id).Find(&getUser)
	return &getUser
}

func GetUserBymobileNo(num string) *User {
	var getUser User
	db.LogMode(true)
	db.Table("contacts").Where("mobileno=?", num).Find(&getUser)
	return &getUser
}

func DeleteUser(name string) User {
	var user User
	db.LogMode(true)
	db.Table("contacts").Where("name=?", name).Delete(&user)
	return user
}

func DeleteUserById(id string) User {
	var user User
	db.LogMode(true)
	db.Table("contacts").Where("id=?", id).Delete(&user)
	return user
}

func (UpdateUser *User) UpdateUser(Id string) {

	userDetails := GetUserById(Id)
	// fmt.Println(Id)

	if UpdateUser.Name != "" {
		userDetails.Name = UpdateUser.Name
	}

	if UpdateUser.Address != "" {
		userDetails.Address = UpdateUser.Address
	}

	if UpdateUser.Email != "" {
		userDetails.Email = UpdateUser.Email
	}

	if UpdateUser.Mobileno != "" {
		userDetails.Mobileno = UpdateUser.Mobileno
	}

	db.Table("contacts").Save(&userDetails)
}

func Addprofile(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	db.LogMode(true)
	file, handler, err := r.FormFile("file")

	if err != nil {
		http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	db.Table("contacts").Create(&file)

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer f.Close()

	io.Copy(f, file)
	w.Write([]byte("Image Uploaded Sucessfully"))

}

// func GetProfile(w http.ResponseWriter, r *http.Request) {

// 	vars := mux.Vars(r)
// 	id := vars["id"]

// 	var image Image
// 	db.First(&image, id)

// 	if image.ID == 0 {
// 		http.Error(w, "Image Not Found", http.StatusNotFound)
// 		return
// 	}

// 	http.ServeFile(w, r, image.FileName)

// }
