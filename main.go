package main

import (
	"fmt"
	"html/template"
	"log"
	"math"
	"net/http"
	"strconv"
)

var tpl *template.Template

func init (){
	tpl=template.Must(template.ParseGlob("templates/*.gohtml"))
}


func main()  {
	http.HandleFunc("/", index)
	http.HandleFunc("/apply", apply)
	err := http.ListenAndServe(":8181", nil)
	if err !=nil {
		fmt.Println(err)
	}

}

func index(w http.ResponseWriter, r *http.Request)  {
	err := tpl.ExecuteTemplate(w,"index.gohtml",nil)
	if err !=nil {
		log.Println(err)
		http.Error(w,"internal error server", http.StatusInternalServerError)
	}
	fmt.Println("we got index page")

}

func apply(w http.ResponseWriter, r *http.Request)  {
	var number_input string
		number_input="0"

	if r.Method==http.MethodPost {
		number_input = r.FormValue("txtprime")

	}
	var number_int int
	number_int,error1:= strconv.Atoi(number_input)
	if error1 !=nil {
		log.Println(error1)
	}
	number_input=getPrimeNumber(number_int)
	fmt.Println(number_input)
	err := tpl.ExecuteTemplate(w, "apply.gohtml",number_input)
	if err != nil {
			log.Println(err)
			http.Error(w, "internal error server", http.StatusInternalServerError)
		}
		fmt.Println("we got apply page")

}
func  getPrimeNumber(value int) string {
	var prime_number string
	f := make([]bool, value)
	for i := 2; i <= int(math.Sqrt(float64(value))); i++ {
		if f[i] == false {
			for j := i * i; j < value; j += i {
				f[j] = true
			}
		}
	}
	for i := value-1; i > 1; i--{
		if f[i] == false {
			fmt.Printf("%v ", i)
			prime_number= strconv.Itoa(i)
			break
			}

	}
	return prime_number
}