package contorller

import (
	"database/sql"
	"github.com/sirupsen/logrus"
)

type mysqlCollector struct {
	db *sql.DB
}

func (c *mysqlCollector) status(name string) float64 {
	sql := "show global status where variable_name=?"
	var (
		vname string
		rs    float64
	)
	err := c.db.QueryRow(sql, name).Scan(&vname, &rs)

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"sql":   sql,
			"name":  name,
			"vaule": &rs,
		}).Error("select status")
	} else {
		logrus.WithFields(logrus.Fields{
			"sql":   sql,
			"name":  name,
			"vaule": &rs,
		}).Debug("select status")
	}
	return rs
}

func (c *mysqlCollector) variables(name string) float64 {
	sql := "show global variable where variable_name=?"
	var (
		vname string
		rs    float64
	)
	err := c.db.QueryRow(sql, name).Scan(&vname, &rs)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"sql":   sql,
			"name":  name,
			"vaule": &rs,
		}).Error("select variable")
	} else {
		logrus.WithFields(logrus.Fields{
			"sql":   sql,
			"name":  name,
			"vaule": &rs,
		}).Debug("select variable")
	}
	return rs
}
