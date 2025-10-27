package db

import (
	"fmt"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"wxcloudrun-golang/db/model"
)

var dbInstance *gorm.DB
var sqliteDB *gorm.DB

// Init 初始化数据库
func Init() error {

	source := "%s:%s@tcp(%s)/%s?readTimeout=1500ms&writeTimeout=1500ms&charset=utf8&loc=Local&&parseTime=true"
	user := os.Getenv("MYSQL_USERNAME")
	pwd := os.Getenv("MYSQL_PASSWORD")
	addr := os.Getenv("MYSQL_ADDRESS")
	dataBase := os.Getenv("MYSQL_DATABASE")
	
	// 设置默认值
	if user == "" {
		user = "root"
	}
	if addr == "" {
		addr = "localhost:3306"
	}
	if dataBase == "" {
		dataBase = "golang_demo"
	}
	
	source = fmt.Sprintf(source, user, pwd, addr, dataBase)
	fmt.Println("start init mysql with ", source)
	fmt.Printf("MySQL connection details: user=%s, addr=%s, database=%s\n", user, addr, dataBase)

	// Retry connection with exponential backoff
	var db *gorm.DB
	var err error
	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = gorm.Open(mysql.Open(source), &gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			}})
		if err == nil {
			fmt.Println("Successfully connected to MySQL")
			break
		}
		
		fmt.Printf("DB Open error (attempt %d/%d): %s\n", i+1, maxRetries, err.Error())
		if i < maxRetries-1 {
			waitTime := time.Duration(i+1) * 2 * time.Second
			fmt.Printf("Waiting %v before retry...\n", waitTime)
			time.Sleep(waitTime)
		}
	}
	
	if err != nil {
		return fmt.Errorf("failed to connect to MySQL after %d attempts: %w", maxRetries, err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("DB Init error,err=", err.Error())
		return err
	}

	// 用于设置连接池中空闲连接的最大数量
	sqlDB.SetMaxIdleConns(100)
	// 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(200)
	// 设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(time.Hour)

	dbInstance = db

	fmt.Println("finish init mysql with ", source)
	return nil
}

// Get ...
func Get() *gorm.DB {
	// 如果在测试模式下，返回SQLite实例
	if os.Getenv("TEST_MODE") == "true" && sqliteDB != nil {
		return sqliteDB
	}
	return dbInstance
}

// InitSQLite 初始化SQLite数据库（用于测试）
func InitSQLite() error {
	dbPath := os.Getenv("SQLITE_PATH")
	if dbPath == "" {
		dbPath = "test.db"
	}

	fmt.Println("初始化SQLite数据库:", dbPath)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Println("SQLite打开失败:", err.Error())
		return err
	}

	sqliteDB = db

	// 自动迁移表结构
	err = autoMigrateSQLite(db)
	if err != nil {
		return err
	}

	fmt.Println("SQLite初始化成功")
	return nil
}

// autoMigrateSQLite 自动创建SQLite表结构
func autoMigrateSQLite(db *gorm.DB) error {
	// 使用GORM的自动迁移
	models := []interface{}{
		&model.CounterModel{},
		&model.UserModel{},
		&model.CategoryModel{},
		&model.ProductModel{},
		&model.MemberModel{},
		&model.MemberLevelModel{},
		&model.DistributorLevelModel{},
		&model.OrderModel{},
		&model.PaymentModel{},
	}

	for _, m := range models {
		if err := db.AutoMigrate(m); err != nil {
			return fmt.Errorf("failed to migrate %T: %v", m, err)
		}
	}

	// 创建默认管理员
	var count int64
	db.Model(&model.UserModel{}).Where("username = ?", "admin").Count(&count)
	if count == 0 {
		admin := &model.UserModel{
			ID:       "admin-001",
			Username: "admin",
			Password: "0192023a7bbd73250516f069df18b500", // admin123的MD5
			Role:     "admin",
		}
		db.Create(admin)
	}
	
	return nil
}
