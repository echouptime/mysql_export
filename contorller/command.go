package contorller

import (
	"database/sql"
	"fmt"
	"github.com/prometheus/client_golang/prometheus"
)

type CommandCollector struct {
	mysqlCollector
	desc *prometheus.Desc
}

func NewCommandCollector(db *sql.DB) *CommandCollector {
	return &CommandCollector{
		mysqlCollector: mysqlCollector{db: db},
		desc: prometheus.NewDesc(
			"Mysql_Global_Status_Com_select",
			"Mysql Global Status Com_select",
			[]string{"command"},
			nil,
		),
	}
}

func (c *CommandCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}

func (c CommandCollector) Collect(metric chan<- prometheus.Metric) {

	names := []string{
		"insert",
		"select",
		"update",
		"delete",
	}
	for _, name := range names {
		metric <- prometheus.MustNewConstMetric(
			c.desc,
			prometheus.CounterValue,
			c.status(fmt.Sprintf("Com_%s", name)),
			name,
		)
	}

}
