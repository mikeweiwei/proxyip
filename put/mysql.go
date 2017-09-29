package put

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
	"log"
	"fmt"
	"github.com/mikeweiwei/proxyip/model"
)

func Init()  {

	creaTable()
}

func DB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/ip?charset=utf8 ")
	if err != nil {
		panic(err)
	}
	log.Println("数据库连接成功！")
	return db
}


func creaTable()  {
	cr_table := "create table if not exists ip(id int(10) not null primary key auto_increment,ip char(200) not null,iptype char(20) not null);"
	stmt1, err := db.Prepare(cr_table)
	checkErr(err)
	_, err = stmt1.Exec()
	defer stmt1.Close()
	log.Println("table sucsess！")
	checkErr(err)

}
func checkErr(err error) {    if err != nil {        panic(err)    } }

//获取所有
func FindAll()  ([]*model.Ip){

	var ips []*model.Ip
	rows, err := db.Query("SELECT * FROM ip")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var ip, iptype string
		rows.Scan(&id, &ip, &iptype)
		ipOne := model.Newip()
		ipOne.Ip = ip
		ipOne.IpType = iptype
		ips = append(ips,ipOne)
	}
	return ips
}
//获取一条
func FindOne(ip string,iptype string) (*model.Ip) {

	rows, err := db.Query("SELECT * FROM ip WHERE ip = ? AND iptype = ?",ip,iptype)
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return nil
	}
	defer rows.Close()
	ipOne := model.Newip()
	for rows.Next() {
		var id int
		var ip, iptype string
		rows.Scan(&id, &ip, &iptype)
		ipOne.Ip = ip
		ipOne.IpType = iptype
	}
	return ipOne
}
//插入一条
func insertOne(db *sql.DB,ip string,iptype string)  {
	exec, error := db.Exec("insert into ip(ip,iptype) select ?,? from dual where not exists(select ip from ip where ip =?)",ip,iptype,ip)
	if error != nil{
		log.Println("insert data failed:", error.Error())
		return
	}
	id, err := exec.LastInsertId()
	if err != nil {
		log.Println("fetch last insert id failed:", err.Error())
		return
	}
	log.Println("insert new record", id)
}
//删除一条
func deleteOne(ip string,iptype string)  {

	exec, error := db.Exec("DELETE FROM ip WHERE ip = ? AND iptype = ?", ip,iptype)
	if error != nil{
		fmt.Println("delete data false",error.Error())
		return
	}
	num, err := exec.RowsAffected()
	if err != nil {
		fmt.Println("fetch row affected failed:", err.Error())
		return
	}
	fmt.Println("delete record number", num)
}
//统计
func count() int {

	rows, err := db.Query("SELECT * FROM ip")
	if err != nil {
		fmt.Println("fetech data failed:", err.Error())
		return 0
	}
	defer rows.Close()
	var num int
	for rows.Next() {
		num ++
	}
	return num
}
//查询具体类型ip
func FindType(iptype string) ([]*model.Ip) {


	var parme string
	if iptype == "http" || iptype == "HTTP"{
		iptype = "%http%"
		parme = "https"
	}else if iptype == "https" || iptype == "HTTPS" {
		iptype = "%https%"
		parme = "http"
	}

	var ips []*model.Ip
	rows, error := db.Query("select * FROM ip WHERE iptype LIKE ? AND iptype != ?", iptype,parme)

	if iptype == "all"  || iptype == "ALL"{
		iptype = "http,https"
		parme = "https,http"
		rows, error = db.Query("select * FROM ip WHERE iptype = ? or iptype = ?", iptype,parme)
	}

	defer rows.Close()
	if error != nil{
		fmt.Println("find type false",error.Error())
	}

	for rows.Next(){
		var id int
		var ip, iptype string
		rows.Scan(&id, &ip, &iptype)
		ipOne := model.Newip()
		ipOne.Ip = ip
		ipOne.IpType = iptype
		ips = append(ips,ipOne)
	}
	return ips
}