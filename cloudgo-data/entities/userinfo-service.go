package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	_, err := engine.Insert(u)
	checkErr(err)
	return nil
}

// func (*UserInfoAtomicService) Save(u *UserInfo) error {
// 	tx, err := mydb.Begin()
// 	checkErr(err)

// 	dao := userInfoDao{tx}
// 	err = dao.Save(u)

// 	if err == nil {
// 		tx.Commit()
// 	} else {
// 		tx.Rollback()
// 	}
// 	return nil
// }

// FindAll .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	rows, err := engine.Rows(new(UserInfo))
	defer rows.Close()
	checkErr(err)
	bean := new(UserInfo)
	var uList []UserInfo
	for rows.Next() {
		err = rows.Scan(bean)
		uList = append(uList, *bean)
	}
	return uList
}

// func (*UserInfoAtomicService) FindAll() []UserInfo {
// 	dao := userInfoDao{mydb}
// 	return dao.FindAll()
// }

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	u := new(UserInfo)
	_, err := engine.ID(id).Get(u)
	checkErr(err)
	return u
}

// func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
// 	dao := userInfoDao{mydb}
// 	return dao.FindByID(id)
// }
