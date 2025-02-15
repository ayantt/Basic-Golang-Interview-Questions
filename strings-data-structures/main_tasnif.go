package main

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go ConcatString(&wg)
	wg.Wait()

	ConvertToString()
	LenOfString()
	Substring()
	MapString()
	Slice()
}

func Slice() {
	num := []int{34, 65, 23, 76, 8}
	sort.Ints(num)
	fmt.Println(num)

	s := []string{"Abc", "Def", "Ghij", "Dab"}
	sort.Strings(s)
	fmt.Println(s)

	num = append(num[:3], num[3:]...)
	fmt.Println(num)
}

func MapString() {
	ages := map[string]int{
		"Ayan":   25,
		"Tasnif": 30,
	}

	fmt.Println("Ayan's age:", ages["Ayan"])

	ages["Taussuk"] = 28

	ages["Tasnif"] = 31

	delete(ages, "Ayan")

	if age, exists := ages["Ayan"]; exists {
		fmt.Println("Ayan's Age:", age)
	} else {
		fmt.Println("Ayan not found!")
	}

	for k, v := range ages {
		fmt.Println("Key:", k, "Value:", v)
	}
}

func Substring() {
	name := "Tasnif Taussuk Turash Ayan"
	firstName := "Tasnif"

	nameB := "তাসনিফ তায়াশশুক তুরাষ অয়ন"
	firstNameB := "তাসনিফ"

	if strings.ContainsAny(name, firstName) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found.")
	}

	if strings.Contains(nameB, firstNameB) {
		fmt.Println("Found")
	} else {
		fmt.Println("Not found.")
	}
}

func LenOfString() {
	firstName := "Tasnif"
	lastName := "Taussuk"

	firstNameB := "তাসনিফ"
	lastNameB := "তায়াশশুক"

	fmt.Println("Using len():", len(firstName))
	fmt.Println("Using utf8.RuneCountInString():", utf8.RuneCountInString(lastName))

	//When expecting Unicode use utf8.RuneCountInString()
	fmt.Println("Using len():", len(firstNameB))
	fmt.Println("Using utf8.RuneCountInString():", utf8.RuneCountInString(firstNameB))

	fmt.Println("Using len():", len(lastNameB))
	fmt.Println("Using utf8.RuneCountInString():", utf8.RuneCountInString(lastNameB))

	fmt.Println("Character length:", len([]rune(firstNameB)))
}

func ConcatString(wg *sync.WaitGroup) {
	defer wg.Done()

	firstName := "Tasnif"
	lastName := "Taussuk"

	var wgI sync.WaitGroup

	wgI.Add(1)
	go func() {
		defer wgI.Done()
		// Efficiency 1: Using strings.Builder
		var sb strings.Builder
		sb.WriteString(firstName)
		sb.WriteString(" ")
		sb.WriteString(lastName)
		fmt.Println("Using strings.Builder:", sb.String())
	}()

	wgI.Add(1)
	go func() {
		defer wgI.Done()
		// Efficiency 2: Using bytes.Buffer
		var sb1 bytes.Buffer
		b := []byte(firstName)
		sb1.Write(b)
		sb1.WriteString(" ")
		sb1.WriteString(lastName)
		fmt.Println("Using bytes.Buffer:", sb1.String())
	}()

	wgI.Add(1)
	go func() {
		defer wgI.Done()
		// Efficiency 3: Using fmt.Sprintf
		result := fmt.Sprintf("%s %s", firstName, lastName)
		fmt.Println("Using fmt.Sprintf:", result)
	}()

	wgI.Wait()
}

func ConvertToString() {
	i := "5"

	s, err := strconv.Atoi(i)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("ConvertToString:", s)
	}
}
