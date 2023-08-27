package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/sergi/go-diff/diffmatchpatch"
)

func Diff() {
	filea := filepath.Join(getAccelerDir(), "c.txt")
	fileb := filepath.Join(getAccelerDir(), "k.txt")
	//result := filepath.Join(getAccelerDir(), "result.txt")

	// // a.txt와 b.txt 파일을 엽니다.
	// a, err := os.Open(filea)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer a.Close()
	// b, err := os.Open(fileb)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer b.Close()

	// // 두 파일의 내용을 문자열로 변환합니다.
	// text1, err := io.ReadAll(a)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// text2, err := io.ReadAll(b)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // // 두 파일의 내용을 문자열로 변환합니다.
	// file1 := string(text1)
	// file2 := string(text2)

	// slice1 := strings.Split(file1, "\n")
	// slice2 := strings.Split(file2, "\n")
	// diff := cmp.Diff(slice1, slice2)
	// // result.txt 파일을 생성하고, 변경사항을 씁니다.
	// f, err := os.Create(result)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer f.Close()
	// log.Println("diff : ", diff[13:len(diff)-2])

	// fmt.Fprint(f, diff[13:len(diff)-2])

	// s := "Hello\nWorld\nThis\nIs\nGo"
	// // \n 으로 쪼개기
	// lines := strings.Split(s, "\n")
	// println(lines)
	// 두 개의 파일을 읽어서 텍스트로 변환합니다.
	file1, err1 := os.ReadFile(filea)
	if err1 != nil {
		log.Fatal(err1)
	}
	text1 := string(file1)

	file2, err2 := os.ReadFile(fileb)
	if err2 != nil {
		log.Fatal(err2)
	}
	text2 := string(file2)

	// diffmatchpatch 객체를 생성합니다.
	dmp := diffmatchpatch.New()

	// 두 텍스트의 diff를 계산합니다.
	diffs := dmp.DiffMain(text1, text2, false)

	results := dmp.DiffCleanupSemantic(diffs)

	// diff 결과를 텍스트로 출력합니다.
	fmt.Println(results)
}
