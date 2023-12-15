package emails

import (
	"github.com/mrz1836/postmark"
)

const (
	EMAIL_ADDRESS_NOREPLY           = "pandoratoolbox@gmail.com"
	POSTMARK_API_TOKEN              = "74d5dba3-e53e-4ef0-86e0-5249d6eb8a99"
	POSTMARK_ACCOUNT_TOKEN          = "a04c0dc0-e985-42c0-af71-f07e7653cfb1"
	TEMPLATE_ID_CREATE_USER         = 1
	TEMPLATE_ID_VERIFY_USER         = 2
	TEMPLATE_ID_RESET_PASSWORD      = 1
	TEMPLATE_ID_CLAIM_STATUS_UPDATE = 1
)

var Postmark *postmark.Client

func Init() error {

	// auth := &http.Client{
	// 	Transport: &postmark.AuthTransport{
	// 		Token: POSTMARK_TOKEN,
	// 	},
	// }

	Postmark = postmark.NewClient(Postmark.ServerToken, Postmark.AccountToken)

	return nil
}

// func SendTemplate(from string, to string, template_id int, template_model map[string]interface{}) error {

// 	res, err := Postmark.SendTemplatedEmail(context.Background(), postmark.TemplatedEmail{
// 		From:          from,
// 		To:            to,
// 		TemplateID:    int64(template_id),
// 		TemplateModel: template_model,
// 	})

// 	return nil
// }
