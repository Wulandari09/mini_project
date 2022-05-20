package config


var DB *gorm.DB

func init() {
	InitDB()
	// InitialMigration()
}

type Config struct {
	DB_perpustakaan string
	DB_Password   string
	DB_Port       string
	DB_Host       string
	DB_Name       string
}

func InitDB() {

	config := Config{
		DB_perpustakaan: "root",
		DB_Password:   "WuLaNdArI09",
		DB_Port:       "3306",
		DB_Host:       "localhost",
		DB_Name:       "perpustakaan",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_matkulname,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}
