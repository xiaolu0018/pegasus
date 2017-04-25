package dao

//import (
//	"time"
//	"gopkg.in/mgo.v2"
//	"github.com/golang/glog"
//)
//
////用来默认具体的数据库的
//var DataBase,Url string
//func Init(){
//	//初始化DataBase，Url
//	DataBase,Url = "pegasus","127.0.0.1:27017"
//	Session = CreateSession()
//	PingSession()
//}
//
////Session 作为全局变量来应用
//var Session *mgo.Session
//
//func PingSession(){
//
//	go func(){
//		for {
//
//			if Session != nil && Session.Ping() != nil{
//				glog.Errorln(Session.Ping().Error())
//				//需要再次连接数据
//				Session = CreateSession()
//			}
//			// 每5秒进行一次session检查
//			time.Sleep(time.Second*5)
//		}
//	}()
//}
//
//
////创建Session
////没有相关参数时使用默认Url
//func CreateSession(url ...string ) *mgo.Session{
//	var tempUrl string
//	if len(url) >0{
//		tempUrl = url[0]
//	}else{
//		tempUrl = Url
//	}
//
//	//如果没有Session时重新连接
//	var err error
//	if Session,err = mgo.DialWithTimeout(tempUrl,time.Second*15);err != nil{
//		glog.Errorln(err.Error())         //记录错误信息
//		return nil
//	}
//
//	return Session
//
//}
//
//
//
////连接指定的DataBase中的collection
//func ConnectConnection(collection string,sesiong *mgo.Session,databaseName ...string) (*mgo.Collection){
//
//	if len(databaseName) >0{
//		return sesiong.DB(databaseName[0]).C(collection)
//	}
//	return  sesiong.DB(DataBase).C(collection)
//}
