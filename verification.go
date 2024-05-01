package proton

import (
	"context"
	"fmt"

	"github.com/go-resty/resty/v2"
)

func (c *Client) GetVerificationData(ctx context.Context, shareId, linkId, revisionId string) (VerificationData, error) {
	var res struct {
		VerificationData
	}

	if err := c.do(ctx, func(r *resty.Request) (*resty.Response, error) {
		url := fmt.Sprintf("/drive/shares/%s/links/%s/revisions/%s/verification", shareId, linkId, revisionId)
		return r.SetResult(&res).Get(url)
	}); err != nil {
		return VerificationData{}, err
	}

	return res.VerificationData, nil
}
