package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", ping())
	r.GET("/tes1", tes1())
	r.GET("/himasi", himasi())
	r.GET("/tes-parameter1", tesParameter1())
	r.GET("/tes-parameter2", tesParameter2())
	r.GET("/tes-parameter-if", tesParameterIf())
	r.GET("/tes-parameter-jurusan", tesParameterJurusan())
	r.GET("/tes-parameter-himpunan", tesParameterHimpunan())
	r.GET("/tes-parameter-dan", tesParameterDan())
	r.GET("/tes-parameter-slash/:name", tesParameterSlash())
	r.GET("/tes-parameter-slash2/:nama/:email", tesParameterSlash2())
	r.Run(":9999") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func tes1() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "haloo ini tes 1",
		})
	}
}

func ping() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	}
}

func himasi() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ini get himasi",
		})
	}
}

func tesParameter1() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"id":      id,
		})
	}
}

func tesParameter2() gin.HandlerFunc {
	return func(c *gin.Context) {
		nama := c.Query("nama")
		id := c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"nama":    nama,
			"id":      id,
		})
	}
}

func tesParameterIf() gin.HandlerFunc {
	return func(c *gin.Context) {
		nama := c.Query("nama")
		id := c.Query("id")
		nim := c.Query("nim")

		if nama == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "nama tidak boleh kosong"})
			return
		}

		if nim == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "nim tidak boleh kosong"})
			return

		}

		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"nama":    nama,
			"id":      id,
			"nim":     nim,
		})
	}
}

func tesParameterJurusan() gin.HandlerFunc {
	return func(c *gin.Context) {
		jurusan := c.Query("jurusan")

		if jurusan != "sistem informasi" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "jurusan hanya boleh di isi sistem informasi"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"jurusan": jurusan,
		})
	}
}

func tesParameterHimpunan() gin.HandlerFunc {
	return func(c *gin.Context) {
		himpunan := c.Query("himpunan")

		if himpunan != "himasi" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "himpunan hanya boleh di isi himasi"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message":  "hallo ",
			"himpunan": himpunan,
		})
	}
}

func tesParameterDan() gin.HandlerFunc {
	return func(c *gin.Context) {
		nama := c.Query("nama")
		nim := c.Query("nim")

		if nama == "" || nim == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "nama  dan nim tidak boleh kosong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"nama":    nama,
			"nim":     nim,
		})
	}
}

func tesParameterSlash() gin.HandlerFunc {
	return func(c *gin.Context) {
		nama := c.Param("name")

		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"nama":    nama,
		})
	}
}

func tesParameterSlash2() gin.HandlerFunc {
	return func(c *gin.Context) {
		nama := c.Param("nama")
		email := c.Param("email")

		if nama == "" || email == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "nama  dan email tidak boleh kosong"})
			return
		}
		if nama != "himasi" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "nama dan  hanya boleh di isi namanya himasi"})
			return
		}
		if email != "himasi@gmail.com" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "email hanya boleh di isi email himasi@gmail.com"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "hallo ",
			"nama":    nama,
			"email":   email,
		})
	}
}
