package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type Grade string

const (
	A Grade = "A"
	B Grade = "B"
	C Grade = "C"
	F Grade = "F"
)

type student struct {
	firstName, lastName, university                string
	test1Score, test2Score, test3Score, test4Score int
}

type studentStat struct {
	student
	finalScore float32
	grade      Grade
}

func parseCSV(filePath string) []student {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open the file: %s", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	var students []student

	_, err = reader.Read()
	if err != nil {
		log.Fatalf("error reading row: %s", err)
	}

	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("error reading row: %s", err)
		}
		if len(row) < 7 {
			log.Printf("Skipping malformed row: %s", row)
			continue
		}

		t1, _ := strconv.Atoi(row[3])
		t2, _ := strconv.Atoi(row[4])
		t3, _ := strconv.Atoi(row[5])
		t4, _ := strconv.Atoi(row[6])

		stu := student{
			firstName:  row[0],
			lastName:   row[1],
			university: row[2],
			test1Score: t1,
			test2Score: t2,
			test3Score: t3,
			test4Score: t4,
		}

		students = append(students, stu)
	}
	return students
}

func calculateGrade(students []student) []studentStat {
	return nil
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	return studentStat{}
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	return nil
}
