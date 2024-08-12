package library

import (
	"context"
	"io"

	"github.com/volcengine/ve-tos-golang-sdk/v2/tos"
)

type Tos struct {
	client *tos.ClientV2
}

func NewTos() (*Tos, error) {
	// credential := tos.NewStaticCredentials("AKLTZjI3ZjM2MTgxNGQwNDJkMDk1OWJkODg3NThlNmYyZjk", "WkRNNE1USmxaREl5WmpCak5EQmpPVGxtTnp***********")
	credential := tos.NewStaticCredentials("AKLTOGRlNjNmNDQ5ZTU3NDM5MGFhY2Q4MTI3YTJmOWMxMzc", "WmpCbU1HTm1Oamt3WVRFMU5HRmpZVGd5WlRVM056YzBaak01TTJOaU9UVQ==")
	client, err := tos.NewClientV2("https://tos-cn-beijing.volces.com", tos.WithCredentials(credential), tos.WithRegion("cn-beijing"))
	return &Tos{client}, err
}

func (tos_ *Tos) Upload(file io.Reader, object string) error {
	_, err := tos_.client.PutObjectV2(context.Background(), &tos.PutObjectV2Input{
		PutObjectBasicInput: tos.PutObjectBasicInput{
			Bucket: "cms-bj-tos-cdn",
			Key:    object,
		},
		Content: file,
	})
	return err
}
