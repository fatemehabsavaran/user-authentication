package pgsql

import (
	"github.com/fatemehabsavaran/user-authentication.git/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PgsqlConfig struct {
	Dsn string `json:"dsn"`
}

type PgsqlService struct {
	db *gorm.DB
}

func NewPgsqlFromConfig(config PgsqlConfig) (*PgsqlService, error) {
	db, err := gorm.Open(postgres.Open(config.Dsn))
	if err != nil {
		return nil, err
	}
	if err = db.AutoMigrate(&User{}, &Token{}); err != nil {
		return nil, err
	}

	return &PgsqlService{
		db: db,
	}, nil

}

func (p *PgsqlService) CreateUser(user *models.User) error {
	dbUser := User{
		Email: user.Email,
		Name:  user.Name,
		Pass:  user.Pass,
	}
	if err := p.db.Model(&User{}).Create(&dbUser).Error; err != nil {
		return err
	}
	user.ID = dbUser.ID
	user.CreatedAt = dbUser.CreatedAt

	return nil
}

func (p *PgsqlService) GetUserByEmail(email string) (models.User, error) {
	getUser := User{}
	if err := p.db.Model(&User{}).Where("email = ?", email).Limit(1).Find(&getUser).Error; err != nil {
		return models.User{}, err
	}

	//if getUser.ID == 0 {
	//	return models.User{}, fmt.Errorf("user not found")
	//}

	return models.User{
		ID:        getUser.ID,
		Email:     getUser.Email,
		Name:      getUser.Name,
		Pass:      getUser.Pass,
		CreatedAt: getUser.CreatedAt,
	}, nil
}

func (p *PgsqlService) AddToken(userID uint, token *models.Token) error {
	getToken := Token{
		ExpireAt: token.ExpireAt,
		Token:    token.Token,
		UserID:   userID,
	}
	if err := p.db.Model(&Token{}).Create(&getToken).Error; err != nil {
		return err
	}
	return nil

}

func (p *PgsqlService) RemoveToken(token string) error {
	if err := p.db.Model(&Token{}).Where("token = ?", token).Delete(&Token{}).Error; err != nil {
		return err
	}
	return nil

}
