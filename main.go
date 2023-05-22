package main 

import (
    "net/http"
    "github.com/gin-gonic/gin"

)

//represents data about a flight operated by Stratos
type flight struct {
    ID     string  `json:"id"`
	CallSign  string `json:"callSign"`
    Destination  string  `json:"destination"`
    Origin string  `json:"origin"`
}

// Test data about flights
var flights  = [] flight{
    {ID: "STRATOS1", CallSign: "OO-SHK", Destination: "Brussels", Origin: "Namur"},
	{ID: "STRATOS2", CallSign: "OO-SPK", Destination: "Namur", Origin: "Namur"},
	{ID: "STRATOS3", CallSign: "OO-IJK", Destination: "Namur", Origin: "Brussels"},
    {ID: "STRATOS4", CallSign: "OO-SHK", Destination: "Liege", Origin: "Charleroi"},

}


func  main()  {

	router := gin.Default()
    router.LoadHTMLFiles("index.tmpl")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Stratos Flights Management Platform",
		})
	})

    router.GET("/flights", getAllFlights)

    router.Run() // by default on port 8080 
	
}

// getAllFlights will return all known flights as JSON list (indented for readability)
func getAllFlights(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, flights)	
}

