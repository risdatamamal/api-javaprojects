// [package main

// import (
// 	"fmt"

// 	"github.com/risdatamamal/final-project/config"
// 	"github.com/risdatamamal/final-project/database"
// 	"github.com/risdatamamal/final-project/router"
// )

// func main() {
// 	r := router.StartApp()
// 	err := database.StartDB()
// 	if err != nil {
// 		fmt.Println("Error starting database: ", err)
// 		return
// 	}
// 	r.Run(config.SERVER_PORT)
// }]

package main

import (
    "github.com/gin-gonic/gin"
)

func main() {
    // Membuat instance router Gin
    router := gin.Default()

    // Route sederhana untuk halaman utama (root)
    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Welcome to Gin Framework!",
        })
    })

    // Route dengan parameter
    router.GET("/hello/:name", func(c *gin.Context) {
        name := c.Param("name")
        c.JSON(200, gin.H{
            "message": "Hello " + name,
        })
    })

    // Jalankan server pada port 8080
    router.Run(":8080")
}
