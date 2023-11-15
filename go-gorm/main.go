package main

import (
	"fmt"
	local_sharding "go-gorm/gorm_sharding"
	"gorm.io/driver/mysql"
	// "gorm_sharding.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/sharding"
)

type Order struct {
	ID 			int64 `gorm_sharding:"primarykey"`
	UserID 		int64
	ProductID	int64
}

func main() {

	// mysql自定义分片
	local_sharding.CustomerShardingWithMysql()

	//shardingDemo()
}

func shardingDemo(){
	//dsn := "host=localhost user=postgresUser password=postgresPW dbname=sharding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//db,err := gorm_sharding.Open(postgres.New(postgres.Config{DSN: dsn}))

	dsn := "usercenter:User@SmartOps@tcp(192.168.101.6:3306)/rundeck?charset=utf8mb4&parseTime=True&loc=Local"
	db,err := gorm.Open(mysql.Open(dsn),&gorm.Config{})

	if err != nil{
		fmt.Print(err)
	}

	for i := 0; i<3; i+=1 {
		table := fmt.Sprintf("orders_%02d", i)
		db.Exec(`DROP TABLE IF EXISTS ` + table)
		//db.Exec(`CREATE TABLE ` + table + ` (
		//	id BIGSERIAL PRIMARY KEY,
		//	user_id bigint,
		//	product_id bigint
		//)`)

		db.Exec(`CREATE TABLE ` + table + ` (
			id bigint PRIMARY KEY,
			user_id bigint,
			product_id bigint
		)`)
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey: "user_id",
		NumberOfShards: 10,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	},"orders")

	db.Use(middleware)
	//insertData(db)

	err = db.Create(&Order{UserID: 2}).Error
	if err != nil{
		fmt.Println(err)
	}

	//queryData(db)
}

func insertData(db *gorm.DB){
	err := db.Create(&Order{UserID: 2}).Error
	if err != nil{
		fmt.Println(err)
	}

	//err = db.Exec("INSERT INTO orders(user_id) VALUES(?)", int64(3)).Error
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//err = db.Exec("INSERT INTO orders(product_id) VALUES(1)").Error //会报错因为没有分片键
	//fmt.Println(err)
}

func queryData(db *gorm.DB){
	var orders []Order
	err := db.Model(&Order{}).Where("user_id", int64(2)).Find(&orders).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", orders)
}