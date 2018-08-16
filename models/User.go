package models

type User struct {
    Id int
    Nickname string
    Email string
    Password string
    IsLock int8
    CreatedAt int
    UpdatedAt int
    Questions []*Question `orm:"reverse(many)"`
}