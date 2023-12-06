package controller

import (
	"go-clean-arch/app/websocketserver/responder"
	_BookieEntity "go-clean-arch/domain/bookie/entity"
	_logger "go-clean-arch/pkg/logger"

	"fmt"

	"github.com/olahol/melody"
)

type Controller struct {
	logger        _logger.Logger
	bookieUsecase _BookieEntity.BookieUsecase
}

func NewController(logger _logger.Logger, bookieUsecase _BookieEntity.BookieUsecase) *Controller {
	return &Controller{
		logger:        logger,
		bookieUsecase: bookieUsecase,
	}
}

// 收到 ping，回傳 pong
func (c *Controller) Pong(s *melody.Session) {
	resp, err := responder.SuccessResp(
		responder.Pong,
		nil,
	)
	if err != nil {
		c.logger.Error("Pong", err)
	}

	c.sendMessageToSession(s, resp)

	return
}

// 收到訂閱賽事訊息，回傳成功
func (c *Controller) SubscribeEvent(s *melody.Session) {
	// 設定頻道
	s.Set(responder.ChannelEvent, true)

	// 設定回傳
	resp, err := responder.SuccessResp(
		responder.SubscribeEvent,
		nil,
	)
	if err != nil {
		c.logger.Error("SubscribeEvent", err)
	}

	// 發送回傳給該 session client
	c.sendMessageToSession(s, resp)

	return
}

// 收到取消訂閱賽事訊息，回傳成功
func (c *Controller) UnsubscribeEvent(s *melody.Session) {
	// 設定頻道
	s.UnSet(responder.ChannelEvent)

	// 設定回傳
	resp, err := responder.SuccessResp(
		responder.UnsubscribeEvent,
		nil,
	)
	if err != nil {
		c.logger.Error("UnsubscribeEvent", err)
	}

	// 發送回傳給該 session client
	c.sendMessageToSession(s, resp)

	return
}

// 收到訂閱盤口訊息，回傳成功
func (c *Controller) SubscribeMarket(s *melody.Session) {
	// 設定頻道
	s.Set(responder.ChannelMarket, true)

	// 設定回傳
	resp, err := responder.SuccessResp(
		responder.SubscribeMarket,
		nil,
	)
	if err != nil {
		c.logger.Error("SubscribeMarket", err)
	}

	// 發送回傳給該 session client
	c.sendMessageToSession(s, resp)

	return
}

// 收到取消訂閱盤口訊息，回傳成功
func (c *Controller) UnsubscribeMarket(s *melody.Session) {
	// 設定頻道
	s.UnSet(responder.ChannelMarket)

	// 設定回傳
	resp, err := responder.SuccessResp(
		responder.UnsubscribeMarket,
		nil,
	)
	if err != nil {
		c.logger.Error("UnsubscribeMarket", err)
	}

	// 發送回傳給該 session client
	c.sendMessageToSession(s, resp)

	return
}

func (c *Controller) GetSportList(s *melody.Session) {
	var resp []byte

	// 取得球種列表
	sportList, err := c.bookieUsecase.GetSportList()
	if err != nil {
		// 設定錯誤回傳
		resp, err = responder.ErrorResp(
			responder.GetSportList,
			responder.GetErrorCode(err.Error()),
			err.Error(),
		)
		if err != nil {
			c.logger.Error("GetSportList ErrorResp", err)
		}

	} else {
		// 設定正確回傳
		resp, err = responder.SuccessResp(
			responder.GetSportList,
			sportList,
		)
		if err != nil {
			c.logger.Error("GetSportList SuccessResp", err)
		}
	}

	// 發送回傳給該 session client
	c.sendMessageToSession(s, resp)

	return
}

// 廣播給所有 session
func (c *Controller) Broadcast(m *melody.Melody, msg []byte) (err error) {
	return c.sendMessageToAll(m, msg)
}

// 發送訊息給特定 session
func (c *Controller) sendMessageToSession(s *melody.Session, msg []byte) {
	err := s.Write(msg)
	if err != nil {
		c.logger.Error("sendMessageToSession", err, fmt.Sprintf("Message: %s", msg)) // TODO Logger 內應該包含上層 func name
		return
	}
	return
}

// 發送訊息給所有 session
func (c *Controller) sendMessageToAll(m *melody.Melody, msg []byte) (err error) {
	err = m.Broadcast(msg)
	if err != nil {
		c.logger.Error("sendMessageToAll", err, fmt.Sprintf("Message: %s", msg)) // TODO Logger 內應該包含上層 func name
		return
	}
	return
}
