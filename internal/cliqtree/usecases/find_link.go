package usecases

import (
	"github.com/NatalNW7/cliqtree/internal/cliqtree/schemas"
	"github.com/NatalNW7/cliqtree/pkg/config"
	"gorm.io/gorm"
)

func FindShortLinkByLinkId(linkID string, db *gorm.DB, logger *config.Logger) (*schemas.Link, error) {
	link := schemas.Link{}
	err := db.Where("link_id = ?", linkID).First(&link).Error
	if err != nil {
		logger.Errorf("Short link with linkId '%s' not found: %v", linkID, err)
		return nil, err
	}

	return &link, nil
}