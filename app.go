// package main

// import "fmt"

// func main() {
// 	//1. Exercise Segitiga
// 	bintang := 5

// 	// segitiga 1
// 	fmt.Println("Segitiga 1")
// 	for i := 1; i <= bintang; i++ {
// 		for j := 1; j <= i; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}

// 	//segitiga 2
// 	fmt.Println("\nSegitiga 2")
// 	for i := 1; i <= bintang; i++ {
// 		for j := i; j <= bintang; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}

// 	//segitiga 3
// 	fmt.Println("\nSegitiga 3")
// 	for i := 1; i <= bintang; i++ {
// 		for j := bintang; j >= 1; j-- {
// 			if i <= (bintang+1)-j {
// 				fmt.Print("*")
// 			} else {
// 				fmt.Print(" ")
// 			}
// 		}
// 		fmt.Println()
// 	}

// 	//segitiga 4
// 	fmt.Println("\nSegitiga 4")
// 	for k := 1; k <= bintang; k++ {
// 		for i := bintang; i >= k; i-- {
// 			fmt.Print(" ")
// 		}
// 		for j := 1; j <= k; j++ {
// 			fmt.Print("*")
// 		}
// 		fmt.Println()
// 	}
// }

package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	log.Println("Running")
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
