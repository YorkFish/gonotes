package models

type FormA struct {
	Foo string `form:"foo" json:"foo" binding:"required"`
}

type FormB struct {
	Bar string `form:"bar" json:"bar" binding:"required"`
}
