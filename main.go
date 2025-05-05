package main

import (
	"encoding/csv"
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
	var students []student

	file, err := os.Open(filePath)
	if err != nil {
		log.Panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	data, err := reader.ReadAll()
	if err != nil {
		log.Panic(err)
	}

	for i, row := range data {
		if i == 0 {
			continue
		}

		score1, _ := strconv.Atoi(row[3])
		score2, _ := strconv.Atoi(row[4])
		score3, _ := strconv.Atoi(row[5])
		score4, _ := strconv.Atoi(row[6])

		students = append(students, student{
			firstName:  row[0],
			lastName:   row[1],
			university: row[2],
			test1Score: score1,
			test2Score: score2,
			test3Score: score3,
			test4Score: score4,
		})

	}

	return students
}

func calculateGrade(students []student) []studentStat {
	var results []studentStat

	for _, student := range students {
		totalScore := (student.test1Score + student.test2Score + student.test3Score + student.test4Score)
		finalScore := float32(totalScore) / 4
		var grade Grade

		switch {
		case finalScore >= 70:
			grade = A
		case finalScore >= 50:
			grade = B
		case finalScore >= 35:
			grade = C
		default:
			grade = F
		}

		results = append(results, studentStat{
			student:    student,
			finalScore: finalScore,
			grade:      grade,
		})

	}

	return results
}

func findOverallTopper(gradedStudents []studentStat) studentStat {
	topper := gradedStudents[0]

	for _, s := range gradedStudents {
		if s.finalScore > topper.finalScore {
			topper = s
		}
	}

	return topper
}

func findTopperPerUniversity(gs []studentStat) map[string]studentStat {
	results := make(map[string]studentStat)

	for _, s := range gs {
		topper, present := results[s.student.university]
		if !present || s.finalScore > topper.finalScore {
			results[s.student.university] = s
		}
	}

	return results
}
