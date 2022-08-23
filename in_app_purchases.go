package appstoreconnect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ListAllInAppPurchasesForAnAppQuery struct {
	FieldsInAppPurchaseAppStoreReviewScreenshots []string `url:"fields[inAppPurchaseAppStoreReviewScreenshots],omitempty"`
	FieldsInAppPurchaseContents                  []string `url:"fields[inAppPurchaseContents],omitempty"`
	FieldsInAppPurchaseLocalizations             []string `url:"fields[inAppPurchaseLocalizations],omitempty"`
	FieldsInAppPurchasePricePoints               []string `url:"fields[inAppPurchasePricePoints],omitempty"`
	FieldsInAppPurchases                         []string `url:"fields[inAppPurchases],omitempty"`
	FieldsPromotedPurchases                      []string `url:"fields[promotedPurchases],omitempty"`
	FilterInAppPurchaseType                      []string `url:"filter[inAppPurchaseType],omitempty"`
	FilterName                                   []string `url:"filter[name],omitempty"`
	FilterProductID                              []string `url:"filter[productId],omitempty"`
	FilterState                                  []string `url:"filter[state],omitempty"`
	Include                                      []string `url:"include,omitempty"`
	Limit                                        int      `url:"limit,omitempty"`
	LimitInAppPurchaseLocalizations              int      `url:"limit[inAppPurchaseLocalizations],omitempty"`
	LimitPricePoints                             int      `url:"limit[pricePoints],omitempty"`
	Sort                                         []string `url:"sort,omitempty"`
	FieldsInAppPurchasePriceSchedules            []string `url:"fields[inAppPurchasePriceSchedules],omitempty"`
}

type ListAllInAppPurchasesForAnAppResponse struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Name                      string      `json:"name"`
			ProductID                 string      `json:"productId"`
			InAppPurchaseType         string      `json:"inAppPurchaseType"`
			State                     string      `json:"state"`
			ReviewNote                interface{} `json:"reviewNote"`
			FamilySharable            bool        `json:"familySharable"`
			ContentHosting            interface{} `json:"contentHosting"`
			AvailableInAllTerritories bool        `json:"availableInAllTerritories"`
		} `json:"attributes"`
		Relationships struct {
			InAppPurchaseLocalizations struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"inAppPurchaseLocalizations"`
			PricePoints struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"pricePoints"`
			Content struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"content"`
			AppStoreReviewScreenshot struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appStoreReviewScreenshot"`
			PromotedPurchase struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"promotedPurchase"`
			IapPriceSchedule struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"iapPriceSchedule"`
		} `json:"relationships"`
		Links struct {
			Self string `json:"self"`
		} `json:"links"`
	} `json:"data"`
	Links struct {
		Self string `json:"self"`
	} `json:"links"`
	Meta struct {
		Paging struct {
			Total int `json:"total"`
			Limit int `json:"limit"`
		} `json:"paging"`
	} `json:"meta"`
}

func (c *Client) ListAllInAppPurchasesForAnApp(ctx context.Context, id string, q *ListAllInAppPurchasesForAnAppQuery) (*ListAllInAppPurchasesForAnAppResponse, error) {
	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	u := c.URL.JoinPath("v1/apps", id, "inAppPurchasesV2")
	u.RawQuery = v.Encode()

	req, err := c.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status code = %v", res.StatusCode)
	}

	body := new(ListAllInAppPurchasesForAnAppResponse)
	if err := decodeBody(res, body); err != nil {
		return nil, err
	}

	return body, nil
}
