package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name         string
	SecondName   string
	LastName     string
	Birthday     string
	PlaceOfBirth string
	Height       string
	Weight       string
}

type Product struct {
	Material1     string
	Material2     string
	Material3     string
	Material4     string
	Manufacturer1 string
	Manufacturer2 string
	Manufacturer3 string
	Manufacturer4 string
	Price1        string
	Price2        string
	Price3        string
	Price4        string
	Articul1      string
	Articul2      string
	Articul3      string
	Articul4      string
	Clas1         string
	Clas2         string
	Clas3         string
	Clas4         string
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")

	})
	http.HandleFunc("/usertable", func(w http.ResponseWriter, r *http.Request) {
		data := User{
			Name:         r.FormValue("name"),
			SecondName:   r.FormValue("second_name"),
			LastName:     r.FormValue("last_name"),
			Birthday:     r.FormValue("birthday"),
			PlaceOfBirth: r.FormValue("place_of_birth"),
			Height:       r.FormValue("height"),
			Weight:       r.FormValue("weight"),
		}

		tmpl, _ := template.ParseFiles("usertable.html")
		tmpl.Execute(w, data)
	})
	http.HandleFunc("/map", func(w http.ResponseWriter, r *http.Request) {
		var arrayData [20]string
		m := map[string]string{
			"Бетон":         "Материал",
			"Шпаклевка":     "Материал",
			"Грунтовка":     "Материал",
			"Клей":          "Материал",
			"Гора":          "Производитель",
			"NOVOL":         "Производитель",
			"Сагус":         "Производитель",
			"Henkel":        "Производитель",
			"600р":          "Цена",
			"500р":          "Цена",
			"700р":          "Цена",
			"001":           "Артикул",
			"002":           "Артикул",
			"003":           "Артикул",
			"004":           "Артикул",
			"B30":           "Класс",
			"Универсальный": "Класс",
			"4 класс":       "Класс",
			"Усиленный":     "Класс",
		}
		i := 0
		for key, _ := range m {
			arrayData[i] = key
			i++
		}

		data := Product{
			Material1:     arrayData[0],
			Material2:     arrayData[1],
			Material3:     arrayData[2],
			Material4:     arrayData[3],
			Manufacturer1: arrayData[4],
			Manufacturer2: arrayData[5],
			Manufacturer3: arrayData[6],
			Manufacturer4: arrayData[7],
			Price1:        arrayData[8],
			Price2:        arrayData[9],
			Price3:        arrayData[10],
			Price4:        arrayData[11],
			Articul1:      arrayData[12],
			Articul2:      arrayData[13],
			Articul3:      arrayData[14],
			Articul4:      arrayData[15],
			Clas1:         arrayData[16],
			Clas2:         arrayData[17],
			Clas3:         arrayData[18],
			Clas4:         arrayData[19],
		}

		//newData := make(map[string]map[string]string)

		tmpl, _ := template.ParseFiles("map.html")
		tmpl.Execute(w, data)
	})

	fmt.Println("Server is listening...")
	http.ListenAndServe(":80", nil)

}
