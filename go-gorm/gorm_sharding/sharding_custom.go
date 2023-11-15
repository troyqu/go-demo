package gorm_sharding

import (
	"errors"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
	"strings"
)

var mysqlDsn 		= "usercenter:User@SmartOps@tcp(192.168.101.6:3306)/rundeck?charset=utf8mb4&parseTime=True&loc=Local"
var postgresSqlDsn 	= "host=localhost user=postgresUser password=postgresPW dbname=sharding port=5432 sslmode=disable TimeZone=Asia/Shanghai"
var shardings		= []string{"2022-08","2022-11","2023-02","2023-04","2023-05"} //默契按照日期批次进行分片

type Order struct {
	ID 			int64 `gorm_sharding:"primarykey"`
	UserID 		int64
	ProductID	int64
	ExamDate	string
}

func CustomerShardingWithMysql(){
	db,err := gorm.Open(mysql.Open(mysqlDsn),&gorm.Config{})
	if err != nil{
		fmt.Print(err)
	}

	fmt.Print(db)

	for _, value := range shardings{
		table := fmt.Sprintf("orders_%s", convDateToSuffix(value))
		db.Exec(`DROP TABLE IF EXISTS ` + table)

		db.Exec(`CREATE TABLE ` + table + ` (
			id bigint PRIMARY KEY,
			user_id bigint,
			product_id bigint,
			exam_date varchar(10)
		)`)
	}

	middleware := sharding.Register(sharding.Config{
		ShardingKey			: "exam_date",
		//NumberOfShards		: 5,

		//PrimaryKeyGenerator	: sharding.PKSnowflake,

		PrimaryKeyGenerator	: sharding.PKCustom,
		PrimaryKeyGeneratorFn	: pkGeneratorFn,

		ShardingAlgorithm	: shardingAlgorithm,
		ShardingSuffixs		: shardingSuffixs,
	},"orders")

	db.Use(middleware)
	insertData(db)
	queryData(db)
}


// 主键生成策略
func pkGeneratorFn(idx int64) int64 {
	var i = 0
	return int64(i)
}

// 分片算法
func shardingAlgorithm(columnValue interface{}) (suffix string, err error){
	if val, ok := columnValue.(string); ok{
		newVal := convDateToSuffix(val)
		return fmt.Sprintf("_%s", newVal), nil
	}
	return "", errors.New("invalid user_id")
}

// 分片后缀
func shardingSuffixs() (suffixs []string) {
	for _, value := range shardings{
		suffixs = append(suffixs, convDateToSuffix(value))
	}
	return suffixs
}

func convDateToSuffix(columnValue string) string{
	return strings.Replace(columnValue, "-", "", 1)
}

func insertData(db *gorm.DB){
	err := db.Create(&Order{ExamDate: "2022-08"}).Error
	if err != nil{
		fmt.Println(err)
	}
}

// 数据库查询操作
func queryData(db *gorm.DB){
	var orders []Order
	err := db.Model(&Order{}).Where("exam_date", "2022-08").Find(&orders).Error
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", orders)
}



func CustomShardingWithPostgres(db *gorm.DB){

}

