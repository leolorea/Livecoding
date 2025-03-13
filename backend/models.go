package main

type MetaData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Data []byte `json:"data"`
}
