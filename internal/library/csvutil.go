package library

import (
	"context"
	"encoding/csv"
	"github.com/Picus-Security-Golang-Backend-Bootcamp/homework-4-bakyazi/internal/db/domain/model"
	"log"
	"os"
	"strconv"
)

// readCsvFile opens and reads CSV file name @filename
func readCsvFile(filename string) [][]string {
	f, err := os.Open(filename)
	if err != nil {
		log.Println("cannot open sample csv file")
		panic(err)
	}

	reader := csv.NewReader(f)
	reader.Comma = ';'
	data, err := reader.ReadAll()
	if err != nil {
		log.Println("cannot read sample csv file")
		panic(err)
	}
	return data[1:]
}

// csvToBookSlice read a csv file and converts it to []book.Book
func csvToBookSlice(filename string) []model.Book {
	data := readCsvFile(filename)
	books := make([]model.Book, len(data))
	for i, row := range data {
		authorId, _ := strconv.Atoi(row[2])
		pageCount, _ := strconv.Atoi(row[5])
		price, _ := strconv.Atoi(row[6])
		stockAmount, _ := strconv.Atoi(row[7])
		books[i] = model.Book{
			Name:        row[1],      // index 1
			AuthorID:    authorId,    // index 2
			StockCode:   row[3],      // index 3
			ISBN:        row[4],      // index 4
			PageCount:   pageCount,   // index 5
			Price:       price,       // index 6
			StockAmount: stockAmount, // index 7
		}
	}
	return books
}

// csvToAuthorMap read a csv file and converts it to []author.Author
func csvToAuthorMap(filename string) map[int]model.Author {
	data := readCsvFile(filename)
	authors := map[int]model.Author{}
	for _, row := range data {
		id, _ := strconv.Atoi(row[0])
		authors[id] = model.Author{
			Name: row[1], // index 1
		}
	}
	return authors
}

// insertSampleData inserts sample data into DB
func insertSampleData() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	savedAuthors, err := authorRepo.FindAll(ctx, true)
	if err != nil {
		log.Printf("cannot get all author data from DB, %v\n", err)
	}
	if len(savedAuthors) == 0 {
		authors := csvToAuthorMap("resources/author.csv")
		books := csvToBookSlice("resources/book.csv")
		for _, book := range books {
			id := book.AuthorID
			author, ok := authors[id]
			if ok {
				book.AuthorID = 0
				author.Books = append(author.Books, book)
				authors[id] = author
			}
		}

		for a := range authors {
			author := authors[a]
			err := authorRepo.Insert(ctx, &author)
			if err != nil {
				panic(err)
			}
		}
	}
	log.Printf("there are already records in DB, sample data not inserted\n")

}
