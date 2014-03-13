package main

import (
	"github.com/xrangers/hood"
)
type Users struct {
    Id  hood.Id
    Email string    `sql:"size(50),notnull"`
    Nickname string `sql:"size(32)"`
    Urlname string  `sql:"size(32),notnull"`
    Status bool     `sql:"default(0)"`
    Password string `sql:"size(32),notnull"`
    Salt string     `sql:"size(32),notnull"`
    CreatedAt int    `sql:"size(10)"`
    UpdatedAt int    `sql:"size(10),default(0)"`
}

type Profiles struct {
    User_id int     `sql:"size(10),notnull"`
    Gender bool `sql:"size(2),default(0)"` //性别:0 其他, 1:男. 2:女
    Location string `sql:"size(100)"`
    Summary string `sql:"size(512)"`
    Birthday string `sql:"size(32)"`
    Avatar string `sql:"size(50)"`
}

type Bindings struct {
    Id string    `sql:"size(50),notnull"`
    User_id int     `sql:"size(10),notnull"`
    Service_name string `sql:"size(32),notnull"`
    Auth_type string  `sql:"size(32),notnull"`
    Auth_token string `sql:"size(255),notnull"`
    User_info string
    CreatedAt int   `sql:"size(10)"`
    UpdatedAt int  `sql:"size(10)"`
}

func (m *M) CreateUsersTable_1394700475_Up(hd *hood.Hood) {
    hd.CreateTable(&Users{})
    hd.CreateTable(&Profiles{})
    hd.CreateTable(&Bindings{})

    hd.CreateIndex("users", "uniq_urlname", true, "urlname")
    hd.CreateIndex("users", "uniq_email", true, "email")

    hd.CreateIndex("profiles", "uniq_user", true, "user_id")

    hd.CreateIndex("bindings", "uniq_binding", true, "id")
    hd.CreateIndex("bindings", "user_id", false, "user_id")
}

func (m *M) CreateUsersTable_1394700475_Down(hd *hood.Hood) {
    hd.DropTableIfExists(&Users{})
    hd.DropTableIfExists(&Profiles{})
    hd.DropTableIfExists(&Users{})
}
