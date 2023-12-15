package events

// import (
// 	"context"
// 	"fmt"
// 	"gameon-twotwentyk-api/graphql"
// 	"gameon-twotwentyk-api/models"
// 	"sync"
// 	"time"
// )

// type EventType string

// const (
// 	EventTypeCreateUser        EventType = "CreateUser"
// 	EventTypeUpdateUser        EventType = "UpdateUser"
// 	EventTypeUpdateUserCredits EventType = "UpdateUserCredits"
// 	EventTypeCreateOrder       EventType = "CreateOrder"
// 	EventTypeUpdateOrder       EventType = "UpdateOrder"
// 	EventTypeRequestPassword   EventType = "RequestPassword"
// )

// type Event struct {
// 	EventType EventType   `json:"event_type"`
// 	Data      interface{} `json:"data,omitempty"`
// }

// type EventHandler struct {
// 	ID                           int64
// 	Name                         string
// 	Ctx                          context.Context
// 	queue                        []Event
// 	mu                           sync.RWMutex
// 	config                       EventHandlerConfig
// 	OnCreateUser                 chan EventDataOnCreateUser
// 	HandleCreateUser             func(EventDataOnCreateUser) error
// 	OnUpdateUser                 chan EventDataOnUpdateUser
// 	HandleUpdateUser             func(EventDataOnUpdateUser) error
// 	OnUpdateUserCredits          chan EventDataOnUpdateUserCredits
// 	HandleUpdateUserCredits      func(EventDataOnUpdateUserCredits) error
// 	OnCreateOrder                chan EventDataOnCreateOrder
// 	HandleCreateOrder            func(EventDataOnCreateOrder) error
// 	OnUpdateOrder                chan EventDataOnUpdateOrder
// 	HandleUpdateOrder            func(EventDataOnUpdateOrder) error
// 	OnRequestPassword            chan EventDataOnRequestPassword
// 	HandleRequestPassword        func(EventDataOnRequestPassword) error
// 	HandleCreateMarketplaceOffer func(EventDataOnCreateMarketplaceOffer) error
// 	OnError                      func(Event, error)
// }

// type EventHandlerConfig struct {
// 	freq       time.Duration
// 	maxThreads int
// }

// type EventDataOnRequestPassword struct {
// 	UserId   int64
// 	Password string
// }

// type EventDataOnCreateUser models.User

// type EventDataOnUpdateUser models.User

// type EventDataOnCreateMarketplaceOffer models.MarketplaceOffer

// type EventDataOnUpdateMarketplaceOffer models.MarketplaceOffer

// type EventDataOnClaimStatusChanged models.Claim

// type EventDataOnUpdateUserCredits struct {
// 	UserId   int64
// 	OldValue int64
// 	NewValue int64
// }

// type EventDataOnUpdateUserPassword struct {
// 	UserId   int64
// 	Password string
// }

// var Global *EventHandler

// func NewEventHandler(ctx context.Context, config *EventHandlerConfig) *EventHandler {
// 	new := EventHandler{
// 		Ctx:   ctx,
// 		queue: []Event{},
// 		mu:    sync.RWMutex{},
// 		config: EventHandlerConfig{
// 			freq:       20 * time.Millisecond,
// 			maxThreads: 1,
// 		},
// 	}

// 	if config != nil {
// 		new.config = *config
// 	}

// 	new.setDefaultHandlers()

// 	for i := 0; i < new.config.maxThreads; i++ {
// 		go new.listen(ctx)
// 		// go new.Listener()
// 	}

// 	return &new
// }

// func (h *EventHandler) setDefaultHandlers() {
// 	h.OnError = func(e Event, err error) {
// 		fmt.Printf("handleEvent:%s:%s\n", e.EventType, err.Error())
// 	}

// 	h.HandleCreateUser = func(d EventDataOnCreateUser) error {
// 		err := emails.SendTemplate(emails.EMAIL_ADDRESS_NOREPLY, *d.Email, emails.POSTMARK_TEMPLATE_ID_CREATE_USER, map[string]interface{}{
// 			"username": d.Username,
// 		})

// 		return err
// 	}

// 	h.HandleRequestPassword = func(d EventDataOnRequestPassword) error {
// 		user, err := graphql.GetUser(context.Background(), d.UserId)
// 		if err != nil {
// 			return err
// 		}

// 		err = emails.SendTemplate(emails.EMAIL_ADDRESS_NOREPLY, *user.Email, emails.POSTMARK_TEMPLATE_ID_PASSWORD_RESET, map[string]interface{}{
// 			"username": *user.Username,
// 		})

// 		return err
// 	}

// 	h.HandleClaimStatusChanged = func(d EventDataOnClaimStatusChanged) error { return nil }

// 	h.HandleCreateOrder = func(d EventDataOnCreateOrder) error {
// 		user, err := graphql.GetUser(context.Background(), *d.UserId)
// 		if err != nil {
// 			return err
// 		}

// 		// data_platform := store

// 		err = emails.SendTemplate(emails.EMAIL_ADDRESS_NOREPLY, *user.Email, emails.POSTMARK_TEMPLATE_ID_CREATE_ORDER, map[string]interface{}{
// 			"username": *user.Username,
// 			"amount":   *d.Amount,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	}

// 	h.HandleUpdateOrder = func(d EventDataOnUpdateOrder) error {
// 		user, err := graphql.GetUser(context.Background(), *d.UserId)
// 		if err != nil {
// 			return err
// 		}

// 		if d.Status == models.ORDER_STATUS_COMPLETE {
// 			err = emails.SendTemplate(emails.EMAIL_ADDRESS_NOREPLY, *user.Email, emails.POSTMARK_TEMPLATE_ID_ORDER_DOWNLOAD, map[string]interface{}{
// 				"username": *user.Username,
// 			})
// 			if err != nil {
// 				return err
// 			}
// 		}

// 		return nil
// 	}

// 	h.HandleCreateMarketplaceOffer = func(d EventDataOnCreateMarketplaceOffer) error {
// 		user, err := graphql.GetUser(context.Background(), *d.MarketplaceListing.OwnerId)
// 		if err != nil {
// 			return err
// 		}

// 		// data_platform := store

// 		err = emails.SendTemplate(emails.EMAIL_ADDRESS_NOREPLY, *user.Email, emails.POSTMARK_TEMPLATE_ID_CREATE_ORDER, map[string]interface{}{
// 			"username": *user.Username,
// 			"amount":   *d.Amount,
// 		})
// 		if err != nil {
// 			return err
// 		}

// 		err = notifications.SendNotification(*user.Id, EventDataOnCreateMarketplaceOffer)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}

// }

// func (h *EventHandler) listen(ctx context.Context) {
// 	for {
// 		h.processQueue()
// 		time.Sleep(h.config.freq)
// 	}
// }

// // func (h *EventHandler) Listener() {
// // 	for {
// // 	sel:
// // 		select {
// // 		case e := <-h.OnCreateUser:
// // 			h.HandleCreateUser(e)
// // 		case e := <-h.OnCreateOrder:
// // 			h.HandleCreateOrder(e)
// // 		case e := <-h.OnRequestPassword:
// // 			h.HandleRequestPassword(e)
// // 		case e := <-h.OnUpdateOrder:
// // 			h.HandleUpdateOrder(e)
// // 		case e := <-h.OnUpdateUserCredits:
// // 			h.HandleUpdateUserCredits(e)
// // 		default:
// // 			break sel
// // 		}
// // 		time.Sleep(h.config.freq)
// // 	}
// // }

// func (h *EventHandler) processQueue() error {

// 	e := h.getNextEvent()
// 	if e != nil {
// 		h.handleEvent(*e)
// 	}

// 	return nil
// }

// func (h *EventHandler) getNextEvent() *Event {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()

// 	qlen := len(h.queue)
// 	switch qlen {
// 	case 0:
// 		return nil
// 	case 1:
// 		e := h.queue[0]
// 		h.queue = []Event{}
// 		return &e
// 	default:
// 		e := h.queue[0]
// 		h.queue = h.queue[1:]
// 		return &e
// 	}
// }

// func (h *EventHandler) addToQueue(e Event) {
// 	h.mu.Lock()
// 	defer h.mu.Unlock()
// 	h.queue = append(h.queue, e)
// }

// func (h *EventHandler) Send(e Event) {
// 	h.addToQueue(e)
// }

// func (h *EventHandler) handleEvent(e Event) {
// 	var err error
// 	fmt.Printf("Handling event: %v\n", e)
// 	switch e.EventType {
// 	// case EventTypeClaimStatusChanged:
// 	// 	err = h.HandleClaimStatusChanged((e.Data.(EventDataOnClaimStatusChanged)))
// 	// case EventTypeRequestPassword:
// 	// 	js, err := json.Marshal(e.Data)
// 	// 	if err != nil {
// 	// 		h.OnError(e, err)
// 	// 		return
// 	// 	}
// 	// 	var ed EventDataOnRequestPassword
// 	// 	err = json.Unmarshal(js, &ed)
// 	// 	if err != nil {
// 	// 		h.OnError(e, err)
// 	// 		return
// 	// 	}
// 	// 	err = h.HandleRequestPassword(ed)
// 	// case EventTypeCreateUser:
// 	// 	var ed EventDataOnCreateUser

// 	// 	copier.Copy(&ed, &e.Data)
// 	// 	err = h.HandleCreateUser(ed)
// 	case EventTypeUpdateUser:
// 		err = h.HandleUpdateUser((e.Data.(EventDataOnUpdateUser)))
// 	// case EventTypeUpdateUserPassword:
// 	// 	err = h.OnUpdateUserPassword(e.Data.(EventDataOnUpdateUserPassword))
// 	// case EventTypeUpdateUserCredits:
// 	// 	err = h.OnUpdateUserCredits(e.Data.(EventDataOnUpdateUserCredits))
// 	// case EventTypeCreateOrder:
// 	// 	err = h.OnCreateOrder(e.Data.(EventDataOnCreateOrder))
// 	case EventTypeUpdateOrder:
// 		err = h.HandleUpdateOrder(e.Data.(EventDataOnUpdateOrder))
// 	}

// 	if err != nil {
// 		h.OnError(e, err)
// 	}
// }
