package main

type User []struct {
	Token  string `json:"token"`
	Folder string `json:"folder"`
}
type ResultFiles struct {
	Code int     `json:"code"`
	Msg  string  `json:"msg"`
	Data []Files `json:"data"`
}
type ResultInfo struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data string `json:"data"`
}
type RequestUpload struct {
	Token string  `json:"token"`
	Files []Files `json:"files"`
}
type Files struct {
	Path    string `json:"path"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
