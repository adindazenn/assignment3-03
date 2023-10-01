package model

type Data struct {
    ID    uint `gorm:"primary_key" json:"-"`
    Water int
    Wind  int
}
