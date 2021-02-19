package http

import (
	"net/http"

	ms "github.com/fuwensun/goms/pkg/misc"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/unknwon/com"
)

// readLog
func (s *Server) readLog(ctx *gin.Context) {
	log.Debug("start to read log")

	name := com.StrTo(ctx.Param("name")).String()
	if name == "" {
		name = com.StrTo(ctx.Query("name")).String()
	}

	name = "all" //TODO
	log.Debugf("succeeded to create log data, name: %v", name)

	level := ms.GetLogLevel()

	ctx.JSON(http.StatusOK, gin.H{
		"name":  name,
		"level": level,
	})
	log.Debugf("succeeded to get log, name: %v, level: %v", name, level)
	return
}

// upateLog
func (s *Server) updateLog(ctx *gin.Context) {
	log.Debug("start to update log")

	name := com.StrTo(ctx.Param("name")).String()
	if name == "" {
		name = com.StrTo(ctx.PostForm("name")).String()
	}
	level := com.StrTo(ctx.PostForm("level")).String()

	log.Debugf("succeeded to create log data, name: %v, level: %v", name, level)

	ms.SetLogLevel(level)

	ctx.JSON(http.StatusOK, gin.H{})
	log.Debugf("succeeded to set log, name: %v, level: %v", name, level)
	return
}
