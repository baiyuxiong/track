package app

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"time"
)

var Engine *xorm.Engine

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		revel.ActionInvoker,           // Invoke the action.
	}

	// register startup functions with OnAppStart
	// ( order dependent )
	revel.OnAppStart(InitDB)
	// revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	// Add some common security headers
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func InitDB() {
	dns := revel.Config.StringDefault("db.dns", "root:@tcp(127.0.0.1:3306)/track?charset=utf8")

	var err error
	Engine, err = xorm.NewEngine("mysql", dns)
	if err != nil {
		revel.ERROR.Fatalln("InitDB - NewEngine error ", err)
	}
	err = Engine.Ping()
	if err != nil {
		revel.ERROR.Fatalln("InitDB - Ping error ", err)
	}
	Engine.ShowSQL = true
	Engine.ShowErr = true
	Engine.ShowDebug = true
	Engine.ShowWarn = true

	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 5000)
	Engine.SetDefaultCacher(cacher)

	go pingDB()
}

func pingDB() {
	Engine.Ping()
	time.Sleep(time.Minute * 10) //每10分钟ping一下，保持连接有效
}
