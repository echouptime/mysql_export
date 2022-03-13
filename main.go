package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"mysqld_export/auth"
	"mysqld_export/contorller"
	"mysqld_export/logger"
	"net/http"
)

var db *sql.DB

func init() {
	//数据库Dsn信息
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&loc=PRC&parseTime=true", logger.ReadConfig().Mysql.Username, logger.ReadConfig().Mysql.Password, logger.ReadConfig().Mysql.Host, logger.ReadConfig().Mysql.Port, logger.ReadConfig().Mysql.Db)
	db, _ = sql.Open("mysql", dsn)
}

func mysql_up() {
	//数据库健康检查指标采集
	mysql_status := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "Mysql_Up_Status",
			Help: "Mysql Process And Port Status",
		},
		[]string{"ip", "port"},
	)
	if err := db.Ping(); err == nil {
		logrus.Info("mysql status is ok")
		mysql_status.WithLabelValues(logger.ReadConfig().Mysql.Host, logger.ReadConfig().Mysql.Port).Add(1)
	} else {
		logrus.WithFields(logrus.Fields{"metric": "mysql_up"}).Error(err)
		mysql_status.WithLabelValues(logger.ReadConfig().Mysql.Host, logger.ReadConfig().Mysql.Port).Add(0)
	}
	prometheus.MustRegister(mysql_status)
}

func Register() {
	////注册监控指标
	mysql_up()
	prometheus.MustRegister(contorller.NewSlowQueryCollector(db))
	prometheus.MustRegister(contorller.NewQpsCollector(db))
	prometheus.MustRegister(contorller.NewCommandCollector(db))
}

func main() {

	//注册监控指标
	Register()

	webuser, webpass := logger.ReadConfig().Web.Auth.User, logger.ReadConfig().Web.Auth.Password
	http.Handle("/metricinfo", auth.Auth(promhttp.Handler(), auth.AuthSecrets{webuser: webpass}))

	//添加路由,启动服务
	//http.Handle("/metricinfo", auth.Auth(promhttp.Handler(),nil))
	err := http.ListenAndServe(":8000", nil)
	logrus.Error(err)

}
