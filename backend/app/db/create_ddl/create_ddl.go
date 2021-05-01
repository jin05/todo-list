package main

import (
	"flag"
	ddlmaker "github.com/kayac/ddl-maker"
	"log"
	"todo-list/app/infrastructure/model"
)

func main() {
	var (
		driver      string
		engine      string
		charset     string
		outFilePath string
	)
	flag.StringVar(&driver, "d", "mysql", "set driver")
	flag.StringVar(&driver, "driver", "mysql", "set driver")
	flag.StringVar(&outFilePath, "o", "db/sql/master.sql", "set ddl output file path")
	flag.StringVar(&outFilePath, "outfile", "db/sql/master.sql", "set ddl output file path")
	flag.StringVar(&engine, "e", "InnoDB", "set driver engine")
	flag.StringVar(&engine, "engine", "InnoDB", "set driver engine")
	flag.StringVar(&charset, "c", "utf8mb4", "set driver charset")
	flag.StringVar(&charset, "charset", "utf8mb4", "set driver charset")
	flag.Parse()

	if driver == "" {
		log.Println("Please set driver name. -d or -driver")
		return
	}
	if outFilePath == "" {
		log.Println("Please set outFilePath. -o or -outfile")
		return
	}

	conf := ddlmaker.Config{
		DB: ddlmaker.DBConfig{
			Driver:  driver,
			Engine:  engine,
			Charset: charset,
		},
		OutFilePath: outFilePath,
	}

	dm, err := ddlmaker.New(conf)
	if err != nil {
		log.Println(err.Error())
		return
	}

	structs := model.GetModels()
	dm.AddStruct(structs...)

	err = dm.Generate()
	if err != nil {
		log.Println(err.Error())
		return
	}
}
