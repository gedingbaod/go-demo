package main

import (
	"mysql/dbs" //这个是相对路径
)

func main() {

	dbs.StructInsert()
	dbs.StructUpdate()
	dbs.StructQueryField()
	dbs.StructQueryAllField()
	dbs.StructDel()
	dbs.StructTx()
	dbs.RawQueryField()
	dbs.RawQueryAllField()
}
