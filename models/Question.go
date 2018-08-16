package models

type Question struct {
    Id int
    Title string
    Description string
    ViewNum int
    IsBan int8
    CreatedAt int
    UpdatedAt int
    User *User `orm:"rel(fk)"`
    Category *Category `orm:"rel(fk)"`
}