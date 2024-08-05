package services

import (
	"context"
	"sync"
	"tinyrdm/backend/types"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

type databaseService struct {
	ctx       context.Context
	ctxCancel context.CancelFunc
	mutex     sync.Mutex
	db        *gorm.DB
}

var database *databaseService
var onceDatabase sync.Once

func Database() *databaseService {
	if database == nil {
		onceDatabase.Do(func() {
			database = &databaseService{
				db: &gorm.DB{},
			}
		})
	}
	return database
}

func (c *databaseService) Start(ctx context.Context) {
	c.ctx, c.ctxCancel = context.WithCancel(ctx)
}

func (c *databaseService) StartDatabase() (resp types.JSResp) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	c.db = db
	resp.Success = true
	return
}

func (b *databaseService) Get() (resp types.JSResp) {
	var product Product
	b.db.First(&product, 1)
	resp.Success = true
	resp.Data = product.Code
	return
}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})
}
