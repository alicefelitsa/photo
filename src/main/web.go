package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"photo/src/function"
	"photo/src/middleware"
	"photo/src/model/admin"
	"photo/src/model/agent"
	"photo/src/model/common"
	"photo/src/model/h5"
	"photo/src/model/notify"
	"time"
)

const qiniuUrl = "https://7niu.trumall.cn/" //"http://qiniu.xf6699.cn/"

func main() {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.NoRoute(NotPage)
	manager := r.Group("/", HandleAdminToken)
	{
		manager.POST("/checkadminlogin", checkadminlogin)
		manager.POST("/adminlogin", adminlogin)
		manager.POST("/adminlogout", adminlogout)
		manager.GET("/getAgentList", getAgentList)
		manager.GET("/getLevelAgent", getLevelAgent)
		manager.POST("/addAgent", addAgent)
		manager.POST("/editAgent", editAgent)
		manager.GET("/delAgent", delAgent)
		manager.GET("/delLevelAgent", delLevelAgent)
		manager.GET("/getModelList", getModelList)
		manager.GET("/getModel", getModel)
		manager.GET("/getWorks", getWorks)
		manager.POST("/addWorks", addWorks)
		manager.POST("/editWorks", editWorks)
		manager.GET("/delWorks", delWorks)
		manager.GET("/delVideo", delVideo)
		manager.GET("/stateWorks", stateWorks)
		manager.GET("/stateVideo", stateVideo)
		manager.GET("/getWorksData", getWorksData)
		manager.POST("/updateWorksImage", updateWorksImage)
		manager.POST("/editModel", editModel)
		manager.POST("/addModel", addModel)
		manager.GET("/delModel", delModel)
		manager.GET("/getUptoken", getUptoken)
		manager.POST("/uploadPhoto", uploadPhoto)
		manager.POST("/uploadFile", uploadFile)
		manager.GET("/delMovie", delMovie)
		manager.GET("/getVideoList", getVideoList)
		manager.POST("/addVideo", addVideo)
		manager.POST("/editVideo", editVideo)
		manager.GET("/getOrderList", getOrderList)
		manager.GET("/getDrawingLisst", getDrawingLisst)
		manager.GET("/getLuodiDomainList", getLuodiDomainList)
		manager.GET("/getDoorDomainList", getDoorDomainList)
		manager.GET("/getJumpDomainList", getJumpDomainList)
		manager.POST("/addLuodiDomain", addLuodiDomain)
		manager.POST("/addDoorDomain", addDoorDomain)
		manager.POST("/addJumpDomain", addJumpDomain)
		manager.GET("/delLuodiDomain", delLuodiDomain)
		manager.GET("/delDoorDomain", delDoorDomain)
		manager.GET("/delJumpDomain", delJumpDomain)
		manager.GET("/stateDrawing", stateDrawing)
		manager.POST("/aliSettlement", aliSettlement)
		manager.GET("/get_order_sales", get_order_sales)
		manager.GET("/getVisit", getVisit)
		manager.GET("/get_adminloginlog", get_adminloginlog)
		manager.GET("/getWxPayList", getWxPayList)
		manager.POST("/addWxPay", addWxPay)
		manager.POST("/editWxPay", editWxPay)
		manager.GET("/delWxPay", delWxPay)
		manager.GET("/stateWxPay", stateWxPay)
		manager.POST("/editWorkTypes", editWorkTypes)
	}
	agent := r.Group("agent", HandleUserOrigin)
	{
		agent.POST("/agentlogin", agentlogin)
		agent.POST("/addAgentLevel", addAgentLevel)
		agent.POST("/editAgentLevel", editAgentLevel)
		agent.GET("/getAgentLevel", getAgentLevel)
		agent.GET("/delAgentLevel", delAgentLevel)
		agent.POST("/agentlogout", agentlogout)
		agent.POST("/checkAgentlogin", checkAgentlogin)
		agent.GET("/getAgentOrder", getAgentOrder)
		agent.GET("/getAgentOrderLevel", getAgentOrderLevel)
		agent.GET("/getAgentDrawing", getAgentDrawing)
		agent.GET("/getAgentInfo", getAgentInfo)
		agent.POST("/addAgentDrawing", addAgentDrawing)
		agent.GET("/queryAgentSales", queryAgentSales)
		agent.GET("/getAgentLoginlog", getAgentLoginlog)
		agent.GET("/getVisitAgent", getVisitAgent)
	}
	notifys := r.Group("notify")
	{
		notifys.POST("/wxNotify/:mchid", wxNotify)
		notifys.POST("/wxNotifyH5", wxNotifyH5)
	}
	h5 := r.Group("h5")
	{
		h5.GET("/getWorksList", getWorksList)
		h5.POST("/getVideoInfo", getVideoInfo)
		h5.POST("/requestRegister", requestRegister)
		h5.POST("/requestLogin", requestLogin)
		h5.GET("/getYigouVideoList", getYigouVideoList)
		h5.POST("/getWorksInfo", getWorksInfo)
		h5.POST("/getAgentPrice", getAgentPrice)
		h5.POST("/goOrder", goOrder)
		h5.POST("/goExOrder", goExOrder)
		h5.POST("/getOrderInfo", getOrderInfo)
		h5.POST("/requestAliPay", requestAliPay)
		h5.POST("/requestWxJsapiPay", requestWxJsapiPay)
		h5.POST("/requestWxH5Pay", requestWxH5Pay)
		h5.GET("/checkOrder", checkOrder)
		h5.GET("/checkOrderData", checkOrderData)
		h5.POST("/getUserPackage", getUserPackage)
		h5.POST("/checkLoginStatus", checkLoginStatus)
		h5.POST("/insertLog", insertLog)
		h5.POST("/updateUserPasswd", updateUserPasswd)
		h5.POST("/updateUserRegister", updateUserRegister)
		h5.GET("/getDomain", getDomain)
		h5.GET("/getDoorDomain", getDoorDomain)
		h5.GET("/getJumpDomain", getJumpDomain)
		h5.GET("/getOpenid", getOpenid)
		h5.GET("/getUserInfo", getUserInfo)
		h5.GET("/addVisit", addVisit)
		h5.GET("/getIp", getIp)
		h5.GET("/getWxPayInfo", getWxPayInfo)
	}
	_ = r.Run(":9100")
	//r.RunTLS(":9100", "./zhengshu/server.pem", "./zhengshu/server.key")
}

// NotPage ????????????????????????????????????
func NotPage(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "???????????????????????????????????????????????????",
	})
}

// HandleAdminToken ????????????????????????????????????token?????????????????????????????????????????????
func HandleAdminToken(c *gin.Context) {
	origin := c.Request.Header.Get("Origin") //????????????
	if origin != "http://boss.felitsa.cn" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "???????????????????????????????????????????????????????????????",
		})
		c.Abort() //???????????????????????????
	} else {
		path := c.Request.URL.Path
		admincode := c.Query("admincode")
		admintoken := c.Query("admintoken")
		redisdb, ctx, _ := function.CreateRedisClient()
		if path == "/adminlogin" || path == "/checkadminlogin" || path == "/adminlogout" {
			c.Next() //???????????????????????????
		} else {
			if admincode != "" && admintoken != "" {
				val, _ := redisdb.Get(ctx, admincode).Result() //??????admintoken
				if val == "" || val != admintoken {
					c.JSON(http.StatusOK, gin.H{
						"code": 400,
						"msg":  "???????????????????????????????????????????????????????????????",
					})
					c.Abort() //???????????????????????????
				}
			} else {
				c.JSON(http.StatusOK, gin.H{
					"code": 403,
					"msg":  "???????????????????????????????????????????????????????????????",
				})
				c.Abort() //???????????????????????????
			}
		}
	}
}

func HandleUserOrigin(c *gin.Context) {
	origin := c.Request.Header.Get("Origin") //????????????
	if origin != "http://agent.felitsa.cn" {
		c.JSON(http.StatusOK, gin.H{
			"code": 400,
			"msg":  "???????????????????????????????????????????????????????????????",
		})
		c.Abort() //???????????????????????????
	}
}

// HandleVisit ??????????????????
func HandleVisit(c *gin.Context) {
	var id int64
	ip := c.ClientIP()
	code := c.Query("code")
	fmt.Println("code???", code)
	ctime := time.Now().Format("2006-01-02")
	where := fmt.Sprintf("where code='%v' && ip='%v' && ctime='%v'", code, ip, ctime)
	_ = common.Db.QueryRow("select id from `visit` " + where).Scan(&id)
	if id == 0 {
		_, _ = common.Db.Exec("insert into visit (`code`,`ctime`,`ip`)values(?,?,?)", code, ctime, ip)
	}
}

//????????????Jsapi????????????
func wxNotify(c *gin.Context) {
	code := notify.WxNotify(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

//????????????H5????????????
func wxNotifyH5(c *gin.Context) {
	code := notify.WxNotifyH5(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
	})
}

/**
???????????????????????????
*/
func checkadminlogin(c *gin.Context) {
	msg, code := admin.Checkadminlogin(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

/**
???????????????
*/
func adminlogin(c *gin.Context) {
	msg, code, admincode, admintoken, admin_id, account := admin.Adminlogin(c)
	c.JSON(http.StatusOK, gin.H{
		"code":       code,
		"msg":        msg,
		"admincode":  admincode,
		"admintoken": admintoken,
		"admin_id":   admin_id,
		"account":    account,
	})
}

//???????????????
func adminlogout(c *gin.Context) {
	msg, code := admin.Adminlogout(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????????????????
func checkUserlogin(c *gin.Context) {
	msg, code := agent.CheckAgentlogin(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func getAgentLevel(c *gin.Context) {
	data, count := agent.GetAgentLevel(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????
func agentlogin(c *gin.Context) {
	msg, code, agentcodeRes, agenttokenRes, aid, account, coder, pid := agent.Agentlogin(c)
	c.JSON(http.StatusOK, gin.H{
		"code":       code,
		"msg":        msg,
		"agentcode":  agentcodeRes,
		"agenttoken": agenttokenRes,
		"aid":        aid,
		"account":    account,
		"coder":      coder,
		"pid":        pid,
	})
}

//???????????????
func agentlogout(c *gin.Context) {
	msg, code := agent.Agentlogout(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????????????????
func checkAgentlogin(c *gin.Context) {
	msg, code := agent.CheckAgentlogin(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????????????????
func getAgentLoginlog(c *gin.Context) {
	data, count := agent.GetAgentLoginlog(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????????????????
func getAgentOrder(c *gin.Context) {
	data, count := agent.GetAgentOrder(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????????????????
func getAgentOrderLevel(c *gin.Context) {
	data, count := agent.GetAgentOrderLevel(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????????????????
func getAgentDrawing(c *gin.Context) {
	data, count := agent.GetAgentDrawing(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????????????????
func getAgentInfo(c *gin.Context) {
	data := agent.GetAgentInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
		"msg":  "",
	})
}

//????????????????????????
func addAgentDrawing(c *gin.Context) {
	msg, code := agent.AddAgentDrawing(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????????????????
func queryAgentSales(c *gin.Context) {
	msg, code, data := agent.QueryAgentSales(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//??????????????????
func getAgentList(c *gin.Context) {
	data, count := admin.GetAgentList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????????????????
func getLevelAgent(c *gin.Context) {
	data, count := admin.GetLevelAgent(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????
func addAgent(c *gin.Context) {
	msg, code := admin.AddAgent(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func addAgentLevel(c *gin.Context) {
	msg, code := agent.AddAgentLevel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func editAgentLevel(c *gin.Context) {
	msg, code := agent.EditAgentLevel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func delAgentLevel(c *gin.Context) {
	msg, code := agent.DelAgentLevel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func editAgent(c *gin.Context) {
	msg, code := admin.EditAgent(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func delAgent(c *gin.Context) {
	msg, code := admin.DelAgent(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func delLevelAgent(c *gin.Context) {
	msg, code := admin.DelLevelAgent(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func getModelList(c *gin.Context) {
	data, count := admin.GetModelList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func getModel(c *gin.Context) {
	data := admin.GetModel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "",
		"data": data,
	})
}

//??????????????????
func getWorks(c *gin.Context) {
	data, count := admin.GetWorks(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????
func addWorks(c *gin.Context) {
	code, msg := admin.AddWorks(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func addVideo(c *gin.Context) {
	code, msg := admin.AddVideo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func editWorks(c *gin.Context) {
	code, msg := admin.EditWorks(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????????????????
func editWorkTypes(c *gin.Context) {
	code, msg := admin.EditWorkTypes(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func editVideo(c *gin.Context) {
	code, msg := admin.EditVideo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func delWorks(c *gin.Context) {
	msg, code := admin.DelWorks(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func delVideo(c *gin.Context) {
	msg, code := admin.DelVideo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func stateWorks(c *gin.Context) {
	msg, code := admin.StateWorks(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func stateVideo(c *gin.Context) {
	msg, code := admin.StateVideo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????????????????
func getWorksData(c *gin.Context) {
	code, msg, data := admin.GetWorksData(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//??????????????????
func updateWorksImage(c *gin.Context) {
	code, msg := admin.UpdateWorksImage(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func addModel(c *gin.Context) {
	msg, code := admin.AddModel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func editModel(c *gin.Context) {
	msg, code := admin.EditModel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func delModel(c *gin.Context) {
	msg, code := admin.DelModel(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//???????????????????????????
func getUptoken(c *gin.Context) {
	uptoken := function.GetUpToken()
	c.JSON(http.StatusOK, gin.H{
		"uptoken":  uptoken,
		"qiniuUrl": qiniuUrl,
	})
}

//????????????
func uploadPhoto(c *gin.Context) {
	var code int
	var msg string
	var image string
	file, err := c.FormFile("imageFile")
	if err != nil {
		code = 400
		msg = "????????????"
	} else {
		ok, imageName, localFile := function.SaveImageFile(c, "./images", file)
		if ok {
			imgName := function.UploadFileToQny(imageName, localFile)
			if imgName != "" {
				code = 200
				msg = "????????????"
				image = qiniuUrl + imgName
			}
		} else {
			code = 400
			msg = "????????????"
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"image": image,
	})
}

//????????????
func uploadFile(c *gin.Context) {
	var code int
	var msg string
	var fileName string
	var filePath string
	file, err := c.FormFile("imageFile")
	if err != nil {
		code = 400
		msg = "????????????"
	} else {
		ok, imageName, localFile := function.SaveFile(c, "./upload", file)
		if ok {
			code = 200
			msg = "????????????"
			fileName = imageName
			filePath = localFile
		} else {
			code = 400
			msg = "????????????"
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code":     code,
		"msg":      msg,
		"fileName": fileName,
		"filePath": filePath,
	})
}

//????????????????????????
func get_order_sales(c *gin.Context) {
	msg, code, data := admin.QueryOrderSales(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//????????????????????????
func getVisit(c *gin.Context) {
	msg, code, data := admin.GetVisit(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//????????????????????????
func getVisitAgent(c *gin.Context) {
	msg, code, data := agent.GetVisitAgent(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"data": data,
		"msg":  msg,
	})
}

//???????????????????????????
func get_adminloginlog(c *gin.Context) {
	data, count := admin.QueryAdminLoginLog(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func aliSettlement(c *gin.Context) {
	msg, code := admin.AliSettlement(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????
func delMovie(c *gin.Context) {
	msg, code := admin.DelMovie(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func getVideoList(c *gin.Context) {
	data, count := admin.GetVideoList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"msg":   "",
		"count": count,
		"data":  data,
	})
}

//??????????????????
func getOrderList(c *gin.Context) {
	data, count := admin.GetOrderList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func getDrawingLisst(c *gin.Context) {
	data, count := admin.GetDrawingLisst(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func getDoorDomainList(c *gin.Context) {
	data, count := admin.GetDoorDomainList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//????????????????????????
func getWxPayList(c *gin.Context) {
	data, count := admin.GetWxPayList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func getJumpDomainList(c *gin.Context) {
	data, count := admin.GetJumpDomainList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func getLuodiDomainList(c *gin.Context) {
	data, count := admin.GetLuodiDomainList(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  0,
		"count": count,
		"data":  data,
		"msg":   "",
	})
}

//??????????????????
func addLuodiDomain(c *gin.Context) {
	msg, code := admin.AddLuodiDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func addDoorDomain(c *gin.Context) {
	msg, code := admin.AddDoorDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func addWxPay(c *gin.Context) {
	msg, code := admin.AddWxPay(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

// ??????????????????
func editWxPay(c *gin.Context) {
	msg, code := admin.EditWxPay(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func delWxPay(c *gin.Context) {
	msg, code := admin.DelWxPay(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func stateWxPay(c *gin.Context) {
	msg, code := admin.StateWxPay(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func addJumpDomain(c *gin.Context) {
	msg, code := admin.AddJumpDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func delDoorDomain(c *gin.Context) {
	msg, code := admin.DelDoorDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func delJumpDomain(c *gin.Context) {
	msg, code := admin.DelJumpDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func delLuodiDomain(c *gin.Context) {
	msg, code := admin.DelLuodiDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func stateDrawing(c *gin.Context) {
	msg, code := admin.StateDrawing(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func getWorksList(c *gin.Context) {
	data := h5.GetWorksList(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"list": data,
	})
}

//??????????????????
func getVideoInfo(c *gin.Context) {
	code, msg, playerUri, title := h5.GetVideoInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code":      code,
		"msg":       msg,
		"playerUri": playerUri,
		"title":     title,
	})
}

//????????????????????????
func getYigouVideoList(c *gin.Context) {
	data := h5.GetYigouVideoList(c)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"list": data,
	})
}

//????????????
func requestRegister(c *gin.Context) {
	code, msg, uid, userName, loginTime := h5.RequestRegister(c)
	c.JSON(http.StatusOK, gin.H{
		"code":      code,
		"msg":       msg,
		"uid":       uid,
		"username":  userName,
		"loginTime": loginTime,
	})
}

//????????????
func requestLogin(c *gin.Context) {
	code, msg, uid, userName, loginTime := h5.RequestLogin(c)
	c.JSON(http.StatusOK, gin.H{
		"code":      code,
		"msg":       msg,
		"uid":       uid,
		"username":  userName,
		"loginTime": loginTime,
	})
}

//??????????????????
func getWorksInfo(c *gin.Context) {
	code, msg, cover, atlas := h5.GetWorksInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"cover": cover,
		"list":  atlas,
	})
}

//????????????????????????
func getAgentPrice(c *gin.Context) {
	code, msg, data := h5.GetAgentPrice(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"price": data,
	})
}

//????????????????????????
func goOrder(c *gin.Context) {
	code, msg, orderNo := h5.GoOrder(c)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     msg,
		"orderNo": orderNo,
	})
}

//?????????????????????
func goExOrder(c *gin.Context) {
	code, msg, orderNo, mchid := h5.GoExOrder(c)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     msg,
		"orderNo": orderNo,
		"mchid":   mchid,
	})
}

//??????????????????
func getOrderInfo(c *gin.Context) {
	code, msg, data := h5.GetOrderInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//?????????????????????
func requestAliPay(c *gin.Context) {
	code, msg, payurl := h5.RequestAliPay(c)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"payurl": payurl,
	})
}

//????????????JSAPI??????
func requestWxJsapiPay(c *gin.Context) {
	code, msg, data := h5.RequestWxJsapiPay(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//????????????H5??????
func requestWxH5Pay(c *gin.Context) {
	code, msg, h5Url := h5.RequestWxH5Pay(c)
	c.JSON(http.StatusOK, gin.H{
		"code":  code,
		"msg":   msg,
		"h5Url": h5Url,
	})
}

//??????????????????
func checkOrder(c *gin.Context) {
	code, msg := h5.CheckOrder(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????????????????
func checkOrderData(c *gin.Context) {
	code, msg, orderNo := h5.CheckOrderData(c)
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"msg":     msg,
		"orderNo": orderNo,
	})
}

//????????????????????????
func getUserPackage(c *gin.Context) {
	code, msg := h5.GetUserPackage(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//????????????????????????
func checkLoginStatus(c *gin.Context) {
	code, msg := h5.CheckLoginStatus(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func insertLog(c *gin.Context) {
	code, msg := h5.InsertLog(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????
func updateUserPasswd(c *gin.Context) {
	code, msg := h5.UpdateUserPasswd(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}

//??????????????????????????????
func updateUserRegister(c *gin.Context) {
	code, msg, username, uid, loginTime := h5.UpdateUserRegister(c)
	c.JSON(http.StatusOK, gin.H{
		"code":      code,
		"msg":       msg,
		"username":  username,
		"uid":       uid,
		"loginTime": loginTime,
	})
}

//??????????????????
func getDoorDomain(c *gin.Context) {
	code, msg, domain := h5.GetDoorDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"domain": domain,
	})
}

//??????????????????
func getJumpDomain(c *gin.Context) {
	code, msg, domain := h5.GetJumpDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"domain": domain,
	})
}

//??????????????????
func getDomain(c *gin.Context) {
	code, msg, domain := h5.GetDomain(c)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"domain": domain,
	})
}

//????????????????????????
func getWxPayInfo(c *gin.Context) {
	code, msg, data := h5.GetWxPayInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//????????????IP
func getIp(c *gin.Context) {
	var code int
	var msg string
	ip := c.ClientIP()
	if ip != "" {
		code = 200
		msg = "??????IP??????"
	} else {
		code = 400
		msg = "??????IP??????"
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"ip":   ip,
	})
}

//??????????????????openid
func getOpenid(c *gin.Context) {
	code, msg, openid := h5.GetOpenid(c)
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"openid": openid,
	})
}

//??????????????????
func getUserInfo(c *gin.Context) {
	code, msg, data := h5.GetUserInfo(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}

//??????????????????
func addVisit(c *gin.Context) {
	code, msg := h5.AddVisit(c)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
	})
}
