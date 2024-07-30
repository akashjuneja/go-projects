package main
import (
	"fmt"
	"log"
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movies struct {
	id string `json:"id"`
	Isbm string `json:"isbm"`
	title string `json:"title"`
	director *Director `json:"director"`
}

type Director struct{
	firstName string `json:"firstname"`
	lastName string `json:"lastname"`
}

var movies []Movies

func getMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-Type","application/json")
	json.NewEncoder(w).Encode(movies)

}
func deleteMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	params:=mux.Vars(r)
	for index,item :=range movies {
		if item.id==params["id"]{
			movies = append(movies[:index],movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}
 
func getMovieById(w http.ResponseWriter ,r *http.Request){
	w.Header().Set("content-type","application/json")
	params:=mux.Vars(r)
	for _,item:=range movies{
		if item.id==params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}

}

func createMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("application/json","application/json")
	var movie Movies; _= json.NewDecoder(r.Body).Decode(&movie)
	movie.id= strconv.Itoa(rand.Intn(1000000))
	movies=append(movies, movie)
	json.NewEncoder(w).Encode(movies)
}

func updateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("content-type","application/json")
	params:=mux.Vars(r)
	for index, item :=range movies {
		if item.id==params["id"]{
			movies=append(movies[:index],movies[index+1:]...)
			var movie Movies
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.id=params["id"]
			movies=append(movies, movie)
		}
	}
	json.NewEncoder(w).Encode(movies)

}
func main(){
	r:=mux.NewRouter()
	movies = append(movies, Movies{id: "1",Isbm: "438227",title: "Movie One",director: &Director{firstName: "John",lastName: "Doe"}})
	movies = append(movies, Movies{id: "2",Isbm: "45455",title: "Movie Two",director: &Director{firstName: "Steve",lastName: "kaddu" }})
	
	
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/getMovie/{id}",getMovieById).Methods("GET")
	r.HandleFunc("/createMovie",createMovie).Methods("POST")
	r.HandleFunc("/updateMovie/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/deleteMovie/{id}",deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))


}
