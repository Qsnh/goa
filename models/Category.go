package models

type Category struct {
    Id int
    Name string
    Created_at int
    Questions []*Question `orm:"reverse(many)"`
}