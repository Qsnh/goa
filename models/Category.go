package models

type Category struct {
    Id int
    Name string
    CreatedAt int
    Questions []*Question `orm:"reverse(many)"`
}