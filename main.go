package main

import (
  "encoding/json"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
)

// ----------- Init Variable (Ini Global, Bisa Diakses Semua Func)

type Food struct {
  ID int
  Name string
  Category string
}

//  Cata buat Array Kosong berisi Struct
var (
  Foods []Food
)

//  ----------- Function


func CheckHandler(w http.ResponseWriter, r* http.Request){
  w.Header().Set("Content-Type", "application/json")

  res := "API Server is Running"
  response, err := json.Marshal(res) // Marshal = JsonEncode, Unmarshal = JsonDecode

  if err != nil {
    http.Error(w, "Error saat parse json", http.StatusInternalServerError)
  }

  w.Write(response)
}

func GetAllFood(w http.ResponseWriter, r* http.Request){
  w.Header().Set("Content-Type", "application/json")

  if len(Foods) < 1 {
    http.Error(w, "Food Empty", http.StatusNotFound)
    return
  }

  response, err := json.Marshal(Foods)
  if err != nil {
    http.Error(w, "Error saat parse json", http.StatusInternalServerError)
  }

  w.Write(response)
}

//  Kirim Input Pakai Raw String JSON
func CreateFood(w http.ResponseWriter, r* http.Request){
  w.Header().Set("Content-Type", "application/json")

  // ------ Get POST Data, langsung cek err saat parsing

  var temp Food
  if err := json.NewDecoder(r.Body).Decode(&temp); err != nil {
    http.Error(w, "Wrong format input", http.StatusBadRequest)
    return
  }

  temp.ID = len(Foods) + 1
  Foods = append(Foods, temp)

  // -----Cetak Response

  response, err := json.Marshal(temp)
  if err != nil {
    http.Error(w, "Error saat parse json", http.StatusInternalServerError)
  }

  w.WriteHeader(http.StatusCreated)
  w.Write(response)
}

// Pakai Form Data
func CreateFoodAnotherWay(w http.ResponseWriter, r* http.Request){
  w.Header().Set("Content-Type", "application/json")

  // ------ Get POST Data, langsung cek err saat parsing

  var temp Food
  temp.ID = len(Foods) + 1
  temp.Name = r.FormValue("Name")
  temp.Category = r.FormValue("Category")
  Foods = append(Foods, temp)

  // -----Cetak Response

  response, err := json.Marshal(temp)
  if err != nil {
    http.Error(w, "Error saat parse json", http.StatusInternalServerError)
  }

  w.WriteHeader(http.StatusCreated)
  w.Write(response)
}


func main()  {
  fmt.Println("Server Running")

  //  --- Routing

  route := mux.NewRouter()
  route.HandleFunc("/", CheckHandler).Methods("GET")
  route.HandleFunc("/foods", GetAllFood).Methods("GET")
  route.HandleFunc("/foods", CreateFood).Methods("POST")
  route.HandleFunc("/foodsAnotherWay", CreateFoodAnotherWay).Methods("POST")

  // Run server
  // Pakai github.com/labstack/gommon/log jika mau log
  http.ListenAndServe(":8000", route)
}

// Petunjuk
// 1. http itu bawaan
// 2. go get URL_IMPORT
// 3. Framework : Gin/Viper/Echo
