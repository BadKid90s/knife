module example

go 1.21

replace knife => ../knife

require knife v0.0.0-00010101000000-000000000000

require github.com/julienschmidt/httprouter v1.3.0 // indirect
