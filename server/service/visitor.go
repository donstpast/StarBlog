package service

//// 日志写入数据库
//func DbLogger(data *model.Visitor) int {
//	err := model.DB.Create(data).Error
//	if err != nil {
//		return errmsg.ERROR //500
//	}
//	return errmsg.SUCCESS //200
//}

////构建存入数据库的数据
//var data *model.Visitor
//data.RequestPath = reqPath
//data.ClientIP = clientIp
//data.Agent = userAgent
//data.Referer = referer
//data.Date = startTime
//err = model.DB.Create(data).Error
//if err != nil {
//	fmt.Println(err)
//}
//err = model.DB.Model(&user).Where("id = ?", id).Updates(maps).Error
