package main

import (
	"github.com/xrangers/hood"
)

type Posts struct {
    Id  hood.Id
    Urlname string  `sql:"size(32),notnull"`
    User_id int     `sql:"size(10),notnull"`
    Category_id int `sql:"size(10),notnull"`
    Status bool     `sql:"default(0)"`
    Extra string    `sql:"size(512)"`
    Content string
    CreateAt int    `sql:"size(10)"`
    UpdateAt int    `sql:"size(10),default(0)"`
}

func (m *M) CreatePostsTable_1394290039_Up(hd *hood.Hood) {
	// TODO: implement
    hd.CreateTable(&Posts{})
    hd.CreateIndex("posts", "uniq_urlname", true, "urlname")
    hd.CreateIndex("posts", "user_index", false, "User_id")
    hd.CreateIndex("posts", "category_index", false, "Category_id")
}

func (m *M) CreatePostsTable_1394290039_Down(hd *hood.Hood) {
	// TODO: implement
    hd.DropTableIfExists(&Posts{})
}
