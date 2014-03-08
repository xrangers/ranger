package main

import (
	"github.com/eaigner/hood"
)

type Article struct {
    Id  hood.Id
    Urlname string  `sql:"size(128),notnull"`
    Content string
    CreateAt int    `sql:"size(10)"`
    UpdateAt int    `sql:"size(10),default(0)"`
}

func (m *M) CreateArticleTable_1394290039_Up(hd *hood.Hood) {
	// TODO: implement
    hd.CreateTable(&Article{})
    hd.CreateIndex("article", "uniq_urlname", true, "urlname")
}

func (m *M) CreateArticleTable_1394290039_Down(hd *hood.Hood) {
	// TODO: implement
    hd.DropTableIfExists(&Article{})
}
