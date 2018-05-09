package main

import (

	//_ "postgresql/routers"
	"demo-go/private/postgresql/models"
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var (
	TIMELAYOUT = "2006-01-02 15:04:05" // time的layout
)

func init() {
	// PostgreSQL 配置
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default", "postgres", "user=antyiwei password=yw19900702 dbname=pgtest host=127.0.0.1 port=5432 sslmode=disable", 5, 50)
	// 自动建表
	//orm.RunSyncdb("default", false, true)
	/**
	* MySQL 配置
	* 注册驱动
	* orm.RegisterDriver("mysql", orm.DR_MySQL)
	* mysql用户：root ，root的秘密：tom ， 数据库名称：test ， 数据库别名：default
	* orm.RegisterDataBase("default", "mysql", "root:tom@/test?charset=utf8")
	 */

	//MySQL 配置
	//注册驱动
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//mysql用户：root ，root的秘密：tom ， 数据库名称：test ， 数据库别名：default
	mysqlUrl := `root:1234567@tcp(127.0.0.1:3306)/testzcm?charset=utf8&loc=Asia%2FShanghai`
	orm.RegisterDataBase("zcm_read", "mysql", mysqlUrl, 5, 50) //只读库

	/**
	* Sqlite 配置
	* 注册驱动
	* orm.RegisterDriver("sqlite", orm.DR_Sqlite)
	* 数据库存放位置：./datas/test.db ， 数据库别名：default
	* orm.RegisterDataBase("default", "sqlite3", "./datas/test.db")
	 */

}

func main() {
	//orm.Debug = true

	//TestPg()

	beginId := 23151736 // 开始的ID数
Begin:

	beego.Info("beginId:", beginId)
	// 查询数据库的数据
	ums, err := TestMysql(beginId)
	if err != nil {
		println(err.Error())
	}
	if len(ums) <= 0 {
		goto End
	}
	for i := 0; i < len(ums); i++ {
		var state, msg string
		um := ums[i]
		um.MsgId = um.Id
		result, err := AddPgForUsersMessages(um)
		if !result {
			state = "=FAIL="
			msg = " =======err:" + err.Error()
		} else {
			state = "=SUCCESS="
		}
		//beego.Info(utils.ToString(&um))
		beginId = um.Id
		beego.Info("beginId:", beginId, "第", (i + 1), "条:", state, msg)
	}
	goto Begin

End:
	beego.Run()
}

func TestPg() {
	o := orm.NewOrm()
	o.Using("pq_bd_test")
	stu := new(models.Student)
	stu.Name = "tom"
	stu.Age = 25

	fmt.Println(o.Insert(stu))

}

// 查询数据库的数据
func TestMysql(begin int) (ums []models.UsersMessage, err error) {

	o := orm.NewOrm()
	o.Using("zcm_read")
	//uid := 100004704
	//sql := `SELECT * FROM users_message WHERE uid = ?  ORDER BY id asc`
	sql := `SELECT * FROM users_message WHERE id >? ORDER BY id asc LIMIT 10000;`
	if _, err = o.Raw(sql, begin).QueryRows(&ums); err != nil {
		return nil, err
	}
	return
}

// 插入pg
func AddPgForUsersMessages(um models.UsersMessage) (bool, error) {

	o := orm.NewOrm()
	o.Using("pq_bd_test")
	sql := `INSERT INTO "public"."users_message"( 
"title", "content", "uid", "create_time", "msg_type", 
"is_read", "content_activity_id", "isshow_app", "isshow_pc", "isshow_wap", 
"share_counts", "guid", "ope_time", "muuid", "msg_classify","msg_id")
VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)`
	_, err := o.Raw(sql, um.Title, um.Content, um.Uid, um.CreateTime, um.MsgType,
		um.IsRead, um.ContentActivityId, um.IsshowApp, um.IsshowPc, um.IsshowWap,
		um.ShareCounts, um.Guid, um.OpeTime, um.Muuid, um.MsgClassify, um.MsgId).Exec()
	if err == nil {
		return true, nil
	}
	return false, err

}
