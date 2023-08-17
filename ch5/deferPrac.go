package main

import (
	"context"
	"database/sql"
	"fmt"
	"io"
	"log"
	"os"
)

func deferPrac() {
	fmt.Println("defer prac")
	fmt.Println(noAnswerDefer)
}

func readFileOnArgs() {
	if len(os.Args) < 2 {
		log.Fatal("no file specified")
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	data := make([]byte, 2048)
	for {
		count, err := f.Read(data)
		os.Stdout.Write(data[:count])
		if err != nil {
			if err != io.EOF {
				log.Fatal(err)
			}
		}
		break
	}
}

func noAnswerDefer() {
	defer func() int {
		return 2
	}()
}

func doSomeInserts(ctx context.Context, db *sql.DB, value1, value2 string) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err == nil {
			err = tx.Commit()
		}
		if err != nil {
			tx.Rollback()
		}
	}()
	_, err = tx.ExecContext(ctx, "insert into foo (val) values $1", value1)
	if err != nil {
		return err
	}
	return nil
}

func useGetFile() {
	_, closer, err := getFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	closer()

}

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, err
}
