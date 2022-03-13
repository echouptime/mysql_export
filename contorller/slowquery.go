package contorller

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/sirupsen/logrus"
)

type SlowQueryCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func NewSlowQueryCollector(db *sql.DB) *SlowQueryCollector {
	return &SlowQueryCollector{
		mysqlCollector: mysqlCollector{db},
		desc: prometheus.NewDesc(
			"mysql_global_status_slow_queries",
			"Mysql Global Status Slow Queries",
			nil,
			nil,
		),
	}
}

func (c *SlowQueryCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc

}

func (c *SlowQueryCollector) Collect(metrics chan<- prometheus.Metric) {

	sample := c.status("slow_queries")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, sample)
	logrus.WithFields(logrus.Fields{
		"sample": sample,
		"metric": "slow_queries",
	}).Debug("command metric")
}

type QpsCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func NewQpsCollector(db *sql.DB) *QpsCollector {
	return &QpsCollector{
		mysqlCollector: mysqlCollector{db: db},
		desc: prometheus.NewDesc(
			"Mysql_Global_Status_queries",
			"Mysql Global Status queries",
			nil,
			nil,
		),
	}
}

func (c *QpsCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}

func (c *QpsCollector) Collect(metric chan<- prometheus.Metric) {
	/*
		优化前版本
		var (
			name string
			vaule float64
		)
		c.db.QueryRow("show global status where variable_name=?","queries").Scan(&name,&vaule)
		metric <- prometheus.MustNewConstMetric(c.desc,prometheus.CounterValue,vaule)
	*/

	sample := c.status("queries")
	metric <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, sample)
	logrus.WithFields(logrus.Fields{
		"sample": sample,
		"metric": "queries",
	}).Debug("command metric")

}
