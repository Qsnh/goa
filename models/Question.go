package models

type Question struct {
    Id int
    Title string
    Description string
    View_num int
    Is_ban int8
    Created_at int
    Updated_at int
    User *User `orm:"rel(fk)"`
    Category *Category `orm:"rel(fk)"`
}