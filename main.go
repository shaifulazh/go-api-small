package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// type Payload struct {
// 	Stuff Data
// }
// type Data struct {
// 	Fruit   Fruits
// 	Veggies Vegetables
// }
// type Fruits map[string]int
// type Vegetables map[string]int

// func hello(w http.ResponseWriter, req *http.Request) {
// 	fruits := make(map[string]int)
// 	fruits["Apples"] = 25
// 	fruits["Oranges"] = 10

// 	vegetables := make(map[string]int)
// 	vegetables["Carrats"] = 10
// 	vegetables["Beets"] = 0

// 	d := Data{fruits, vegetables}
// 	p := Payload{d}
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(p)
// }

// func headers(w http.ResponseWriter, req *http.Request) {

// 	for name, headers := range req.Header {
// 		for _, h := range headers {
// 			fmt.Fprintf(w, "%v: %v\n", name, h)
// 		}
// 	}
// }

// func main() {

// 	http.HandleFunc("/", hello)
// 	http.HandleFunc("/headers", headers)

// 	http.ListenAndServe(":8090", nil)
// }
type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)

	router.Run("localhost:8080")
}
