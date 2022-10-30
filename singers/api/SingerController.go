package api

import(
	"log"
	"github.com/gin-gonic/gin"
	"davidone.it/singers/repository"
)
func Router(){
	router := gin.Default()
	router.GET("/singers/api/albums", getAlbums)
	router.Group("/singers/api").GET("/albums/:name", albumsByArtist)
	router.Group("/singers/api").POST("/albums", addAlbum)
	router.GET("/", notFound)
	router.Run("localhost:8080")
}
func getAlbums(c *gin.Context) {
	log.Println("Get Albums called")
	c.Status(204)
}

func albumsByArtist(c *gin.Context) { // ([]Album,error){
	name := c.Param("name")
	albums , err := repository.AlbumsByArtist(name)
	if err != nil {
		log.Println(err.Error())
		c.Errors.JSON()
	}
	c.JSON(200,albums)
	//c.Status(204)

}

func addAlbum(c *gin.Context) {
	//var alb Album
	var alb repository.Album 
	if err := c.BindJSON(&alb); err != nil {
        c.IndentedJSON(400,err)
    }
	err := repository.Create(alb)
	if err != nil {
		log.Println(err.Error())
		c.JSON(500,err.Error())
	}
	c.Status(201)
}

func notFound(c *gin.Context) {
	c.String(404,"NO SINGER FOUND")
}