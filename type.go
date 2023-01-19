package main

type User []struct {
	Token  string `json:"token"`
	Folder string `json:"folder"`
}
type Content struct {
	LastTime string `json:"last_time"`
	Content  string `json:"content"`
}
type Result struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data Content `json:"data"`
}
type ResultStr struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
type Request struct {
	Token   string `json:"token"`
	Path    string `json:"path"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
