package emails

const (
	// EMAIL_GREETING    = `Hello, %s\n\n`
	// EMAIL_UNSUBSCRIBE = ``
	// EMAIL_SIGNATURE   = `\n\nAudienceViral.com`

	// EMAIL_CREATE_ORDER        = `Your order has been created: %d`
	// EMAIL_UPDATE_ORDER_STATUS = `The status of your order has changed: %d is now %s.`
	// EMAIL_COMPLETE_ORDER      = `Your order (%d) is now complete. View and download here: %s`

	// EMAIL_CREATE_USER = `Welcome to AudienceViral, your username is: %s`
	// EMAIL_UPDATE_USER_PASSWORD = `Your password has been changed, if this wasn't you then please contact support.`
	// EMAIL_REQUEST_PASSWORD_RESET = ``

	// EMAIL_CREATE_ORDER_CREDITS = `You have received %d credits. Your new balance is %d.`

	POSTMARK_TEMPLATE_ID_CREATE_ORDER        = 30685812
	POSTMARK_TEMPLATE_ID_PASSWORD_RESET      = 30685790
	POSTMARK_TEMPLATE_ID_UPDATE_ORDER_STATUS = 30686017
	POSTMARK_TEMPLATE_ID_CREATE_USER         = 30685811
	POSTMARK_TEMPLATE_ID_ORDER_DOWNLOAD      = 30685813
)

type CreateOrderTemplate struct {
	OrderId   int64
	Cost      int64
	CreatedAt int64
}
