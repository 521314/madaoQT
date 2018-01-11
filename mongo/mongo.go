package mongo

import "github.com/kataras/golog"

const MongoURL = "mongodb://localhost"
// const MongoURL = "mongodb://192.168.0.102"
const Database = "madaoQT"

const ChartCollectin = "Chart"
const TradeRecordCollection = "TradeRecord"
const ExchangeCollection = "Exchanges"
const OrderCollection = "Orders"
const FundCollection = "Funds"
const OkexDiffHistory = "OkexDiffHistory"

type DBConfig struct {
	CollectionName string
}

var Logger *golog.Logger

func init() {
	logger := golog.New()
	Logger = logger
	Logger.SetLevel("debug")
	Logger.SetTimeFormat("2006-01-02 06:04:05")
	Logger.SetPrefix("[MONG]")
}
