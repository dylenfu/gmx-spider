package dao

import (
	"github.com/dylenfu/gmx-spider/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type GmxDao struct {
	cfg *config.DBConfig
	db  *gorm.DB
}

func NewGmxDao(cfg *config.DBConfig) *GmxDao {
	dao := &GmxDao{
		cfg: cfg,
	}
	Logger := logger.Default
	if cfg.Debug == true {
		Logger = Logger.LogMode(logger.Info)
	}
	db, err := gorm.Open(mysql.Open(cfg.User+":"+cfg.Password+"@tcp("+cfg.URL+")/"+
		cfg.Scheme+"?charset=utf8"), &gorm.Config{Logger: Logger})
	if err != nil {
		panic(err)
	}

	db.Migrator().CreateTable(&Token{})
	dao.db = db

	// create tables
	return dao
}

// func (dao *SwapDao) GetFees() ([]*models.ChainFee, error) {
// 	fees := make([]*models.ChainFee, 0)
// 	res := dao.db.Find(&fees)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	if res.RowsAffected == 0 {
// 		return nil, fmt.Errorf("no record!")
// 	}
// 	return fees, nil
// }
// func (dao *SwapDao) SaveFees(fees []*models.ChainFee) error {
// 	if fees != nil && len(fees) > 0 {
// 		res := dao.db.Save(fees)
// 		if res.Error != nil {
// 			return res.Error
// 		}
// 	}
// 	return nil
// }

// func (dao *ExplorerDao) RemoveEvents(srcHashes []string, polyHashes []string, dstHashes []string) error {
// 	dao.db.Where("`txhash` in ?", srcHashes).Delete(&SrcTransfer{})
// 	dao.db.Where("`txhash` in ?", srcHashes).Delete(&SrcTransaction{})

// 	dao.db.Where("`txhash` in ?", polyHashes).Delete(&PolyTransaction{})

// 	dao.db.Where("`txhash` in ?", dstHashes).Delete(&DstTransfer{})
// 	dao.db.Where("`txhash` in ?", dstHashes).Delete(&DstTransaction{})
// 	return nil
// }

// func (dao *ExplorerDao) GetChain(chainId uint64) (*models.Chain, error) {
// 	chain := new(Chain)
// 	res := dao.db.Where("id = ?", chainId).First(chain)
// 	if res.Error != nil {
// 		return nil, res.Error
// 	}
// 	if res.RowsAffected == 0 {
// 		return nil, fmt.Errorf("no record!")
// 	}
// 	chainJson, err := json.Marshal(chain)
// 	if err != nil {
// 		return nil, err
// 	}
// 	newChain := new(models.Chain)
// 	err = json.Unmarshal(chainJson, newChain)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return newChain, nil
// }

// func (dao *ExplorerDao) UpdateChain(chain *models.Chain) error {
// 	chainJson, err := json.Marshal(chain)
// 	if err != nil {
// 		return err
// 	}
// 	newChain := new(Chain)
// 	err = json.Unmarshal(chainJson, newChain)
// 	if err != nil {
// 		return err
// 	}
// 	res := dao.db.Model(newChain).Updates(map[string]interface{}{
// 		"height": gorm.Expr("?", newChain.Height)})
// 	if res.Error != nil {
// 		return res.Error
// 	}
// 	if res.RowsAffected == 0 {
// 		return fmt.Errorf("update chain failed!")
// 	}
// 	return nil
// }

// func (dao *ExplorerDao) AddTokens(tokens []*models.TokenBasic, tokenMaps []*models.TokenMap) error {
// 	explorerTokens, explorerTokenMaps := dao.BuildTokens(tokens)
// 	if explorerTokens != nil && len(explorerTokens) > 0 {
// 		res := dao.db.Save(explorerTokens)
// 		if res.Error != nil {
// 			return res.Error
// 		}
// 		if res.RowsAffected == 0 {
// 			return fmt.Errorf("update explorer tokens failed!")
// 		}
// 	}
// 	for _, tokenMap := range tokenMaps {
// 		explorerTokenMaps = append(explorerTokenMaps, &TokenBind{
// 			SrcHash: tokenMap.SrcTokenHash,
// 			DstHash: tokenMap.DstTokenHash,
// 		})
// 	}
// 	if explorerTokenMaps != nil && len(explorerTokenMaps) > 0 {
// 		res := dao.db.Save(explorerTokenMaps)
// 		if res.Error != nil {
// 			return res.Error
// 		}
// 		if res.RowsAffected == 0 {
// 			return fmt.Errorf("update explorer tokens map failed!")
// 		}
// 	}
// 	return nil
// }
