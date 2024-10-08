package usecases

import (
	"github.com/NatalNW7/cliqtree/internal/cliqtree/schemas"
	"github.com/NatalNW7/cliqtree/internal/cliqtree/util"
	"github.com/NatalNW7/cliqtree/pkg/config"
	"gorm.io/gorm"
)

func CreateShortLink(redirectUrl string, db *gorm.DB, logger *config.Logger) (*schemas.Link, error){
	link := schemas.Link{
		LinkId: util.GenerateLinkId(),
		RedirectUrl: util.ValidadeRedirectUrl(redirectUrl),
	}

	err := db.Create(&link).Error
	if err != nil {
		logger.Errorf("Error to save shortened link on database: %v", err.Error())
		return nil, err
	}

	return &link, nil
}
