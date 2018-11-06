package main

import (
    "fmt"
    "log"
    "net/http"
    "encoding/json"
    "io/ioutil"
    "html/template"
    "regexp"
    "strconv"
    "sort"
)

type Result struct {
    Base string `json:"base"`
    Rates map[string]float64 `json:"rates"`
}

var currencies []string

type Page struct {
    Currencies []string
    CurrFromText string
    CurrFromAmount float64
    CurrToText string
    CurrToAmount float64
}

var result Result
var t = template.Must(template.ParseFiles("pages/index.html"))

func importJSON() {

    response, err := http.Get("http://data.fixer.io/api/latest?access_key=6402b15466cb5e9c639dbe397ccc4a7d")

    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        byteValue, _ := ioutil.ReadAll(response.Body)
    
        json.Unmarshal(byteValue, &result)

        for key, _ := range result.Rates {
            currencies = append(currencies, key)
        }

        sort.Strings(currencies)
    }
}

func convertValues (r *http.Request, page Page) Page {
    var validPath = regexp.MustCompile("^/convert/([A-Z]{3})/([0-9]+)/([A-Z]{3})/([0-9]+)$")
    m := validPath.FindStringSubmatch(r.URL.Path)

    page.CurrFromText = m[1]
    page.CurrFromAmount, _ = strconv.ParseFloat(m[2], 4)
    page.CurrToText = m[3]
    page.CurrToAmount, _ = strconv.ParseFloat(m[4], 4)

    if page.CurrFromText != "" && page.CurrToText != "" {

        if page.CurrFromAmount > 0 {
            page.CurrToAmount = page.CurrFromAmount * result.Rates[page.CurrToText] / result.Rates[page.CurrFromText]
        } else if page.CurrToAmount > 0 {
            page.CurrFromAmount = page.CurrToAmount / result.Rates[page.CurrToText] * result.Rates[page.CurrFromText]
        } else {
                fmt.Println("Currencies not selected!")
        }
    } else {
        fmt.Println("Currency types not selected!")
    }

    return page
}

func renderTemplate(w http.ResponseWriter, r *http.Request) {
    var page Page
    page.Currencies = currencies

    t.Execute(w, page)
}

func renderConvertTemplate(w http.ResponseWriter, r *http.Request) {
    var page Page
    page.Currencies = currencies

    page = convertValues(r, page)

    t.ExecuteTemplate(w, "convert", page)
}

func handlerequests () {
    http.HandleFunc("/", renderTemplate)
    http.HandleFunc("/convert/", renderConvertTemplate)

    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    importJSON()
    handlerequests()
}