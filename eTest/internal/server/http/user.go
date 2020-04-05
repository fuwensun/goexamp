package http

import (
	"log"
	"net/http"
	"strconv"

	"github.com/fuwensun/goms/eTest/internal/model"
	"github.com/gin-gonic/gin"
)

func (srv *Server) createUser(c *gin.Context) {
	svc := srv.svc
	var err error
	user := model.User{}

	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("http sex err: %v", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "sex error!",
			"sex":   sexstr,
		})
		return
	}
	user.Name = namestr

	if err = svc.CreateUser(c, &user); err != nil {
		log.Printf("http create user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
			"uid":   user.Uid,
			"name":  user.Name,
			"sex":   user.Sex,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ // create ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http create user=%v", user)
}

func (srv *Server) updateUser(c *gin.Context) {
	svc := srv.svc
	var err error
	user := model.User{}
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.PostForm("uid")
	}
	namestr := c.PostForm("name")
	sexstr := c.PostForm("sex")

	user.Uid, err = strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid error!",
			"uid":   uidstr,
		})
		return
	}
	user.Sex, err = strconv.ParseInt(sexstr, 10, 64)
	if sexstr == "" || err != nil {
		log.Printf("http sex err: %v", sexstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "sex error!",
			"uid":   sexstr,
		})
		return
	}
	user.Name = namestr
	err = svc.UpdateUser(c, &user)
	log.Printf("http update user: %v", err)
	if err == model.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found!",
			"uid":   user.Uid,
			"name":  user.Name,
			"sex":   user.Sex,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
			"uid":   user.Uid,
			"name":  user.Name,
			"sex":   user.Sex,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{ //update ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http update user=%v", user)
}

func (srv *Server) readUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	if uidstr == "" {
		uidstr = c.Query("uid")
	}
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid error!",
			"uid":   uidstr,
		})
		return
	}
	user, err := svc.ReadUser(c, uid)
	log.Printf("http read user: %v", err)
	if err == model.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found!",
			"uid":   uidstr,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
			"uid":   uidstr,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{ //read ok
		"uid":  user.Uid,
		"name": user.Name,
		"sex":  user.Sex,
	})
	log.Printf("http read user=%v", user)
}

func (srv *Server) deleteUser(c *gin.Context) {
	svc := srv.svc
	uidstr := c.Param("uid")
	uid, err := strconv.ParseInt(uidstr, 10, 64)
	if uidstr == "" || err != nil {
		log.Printf("http uid err: %v", uidstr)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "uid error!",
			"uid":   uidstr,
		})
		return
	}
	err = svc.DeleteUser(c, uid)
	log.Printf("http delete user: %v", err)
	if err == model.ErrNotFound {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "data not found!",
			"uid":   uidstr,
		})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "internal error!",
			"uid":   uidstr,
		})
		return
	}
	c.JSON(http.StatusNoContent, gin.H{"uid": uid}) //delete ok
	log.Printf("http delete user uid=%v", uid)
}
