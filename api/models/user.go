package models

type User struct {
	Id       uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Name     string `gorm:"size:100; not null" json:"name"`
	Email    string `gorm:"size:100; not null" json:"email"`
	Password string `gorm:"size:256; not null" json:"password"`
	Token    string `gorm:"size:256" json:"token"`
}

func NewUser(user User) error {
	db := Connect()
	defer db.Close()
	return db.Create(&user).Error
}
func GetUser() []User {
	db := Connect()
	defer db.Close()
	var users []User
	db.Find(&users)
	return users
}
func GetUserByToken(token string) (User, error) {
	db := Connect()
	defer db.Close()
	var user User
	db.Where("token = ?", token).Find(&user)
	return user, nil
}
func GetUserById(id uint64) User {
	db := Connect()
	defer db.Close()
	var user User
	db.Where("id = ?", id).Find(&user)
	return user
}
func GetUserByEmail(u string) (User, error) {
	db := Connect()
	defer db.Close()
	var user User
	db.Where("email = ?", u).Find(&user)
	return user, nil
}
func UpdateUserById(user User) (int64, error) {
	db := Connect()
	defer db.Close()
	res := db.Model(&user).Where("id = ?", user.Id).UpdateColumns(
		map[string]interface{}{
			"name":  user.Name,
			"email": user.Email,
		},
	)
	return res.RowsAffected, res.Error
}
func UpdateTokenByEmail(token string, user User) (int64, error) {
	db := Connect()
	defer db.Close()
	res := db.Model(&user).Where("email = ?", user.Email).UpdateColumns(
		map[string]interface{}{
			"token": token,
		},
	)
	return res.RowsAffected, res.Error
}
func DeleteUser(id uint64) (int64, error) {
	db := Connect()
	defer db.Close()
	res := db.Where("id = ?", id).Delete(&User{})
	return res.RowsAffected, res.Error
}
