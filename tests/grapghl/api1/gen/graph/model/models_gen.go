// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type EntityA struct {
	Amount *float64 `json:"Amount"`
	ID     *int     `json:"Id"`
	Name   *string  `json:"Name"`
}

type EntityB struct {
	Description *string  `json:"Description"`
	EntityA     *EntityA `json:"EntityA"`
	ID          *int     `json:"Id"`
}

type EntityC struct {
	EntityB *EntityB `json:"EntityB"`
	ID      *int     `json:"Id"`
	Tag     *string  `json:"Tag"`
}

type Person struct {
	FirstName *string `json:"FirstName"`
	ID        *int    `json:"Id"`
	LastName  *string `json:"LastName"`
}
