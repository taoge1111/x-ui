package service

import (
	"encoding/json"
	"x-ui/database"
	"x-ui/database/model"
	"x-ui/logger"
	"x-ui/xray"

	"gorm.io/gorm"
)

type OutboundService struct {
	xrayApi xray.XrayAPI
}

func (s *OutboundService) AddOutbound(outbound *model.Outbound) (*model.Outbound, bool, error) {
	var err error

	db := database.GetDB()
	tx := db.Begin()
	defer func() {
		if err == nil {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	//todo 去重

	err = tx.Save(outbound).Error
	if err != nil {
		return nil, false, err
	}

	needRestart := false
	outboundJson, err := json.MarshalIndent(outbound.GenXrayOutboundConfig(), "", "  ")
	if err != nil {
		logger.Debug("Unable to marshal outbound config:", err)
	}
	s.xrayApi.Init(p.GetAPIPort())
	err = s.xrayApi.AddOutbound(outboundJson)
	if err == nil {
		logger.Debug("new outbound is add by xray")
	}
	if err != nil {
		logger.Debug("Unable to add outbound by api", err)
		needRestart = true
	}
	defer s.xrayApi.Close()
	//
	return outbound, needRestart, err
}

func (s *OutboundService) GetAllOutbounds() ([]*model.Outbound, error) {
	db := database.GetDB()
	var outbounds []*model.Outbound
	err := db.Model(model.Outbound{}).Find(&outbounds).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return outbounds, nil
}

func (s *OutboundService) DelOutbound(id int) (bool, error) {
	db := database.GetDB()
	var tag string
	needRestart := false
	result := db.Model(model.Outbound{}).Select("tag").Where("id = ?", id, true).First(&tag)
	if result.Error == nil {
		s.xrayApi.Init(p.GetAPIPort())
		err1 := s.xrayApi.DelOutbound(tag)
		if err1 == nil {
			logger.Debug("Outbound deleted by api:", tag)
		} else {
			logger.Debug("Unable to delete outbound by api:", err1)
			needRestart = true
		}
		s.xrayApi.Close()
	} else {
		logger.Debug("No enabled outbound founded to removing by api", tag)
	}
	return needRestart, db.Delete(model.Outbound{}, id).Error
}
