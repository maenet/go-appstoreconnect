package appstoreconnect

import (
	"context"
	"fmt"
	"net/http"

	"github.com/google/go-querystring/query"
)

type ListAppsQuery struct {
	FieldsApps                          []string `url:"fields[apps],omitempty"`
	FieldsBetaLicenseAgreements         []string `url:"fields[betaLicenseAgreements],omitempty"`
	FilterBundleID                      []string `url:"filter[bundleId],omitempty"`
	FilterID                            []string `url:"filter[id],omitempty"`
	FilterName                          []string `url:"filter[name],omitempty"`
	FilterSKU                           []string `url:"filter[sku],omitempty"`
	Include                             []string `url:"include,omitempty"`
	Limit                               int      `url:"limit,omitempty"`
	Sort                                []string `url:"sort,omitempty"`
	FieldsPreReleaseVersions            []string `url:"fields[preReleaseVersions],omitempty"`
	LimitPreReleaseVersions             int      `url:"limit[preReleaseVersions],omitempty"`
	FieldsBetaAppReviewDetails          []string `url:"fields[betaAppReviewDetails],omitempty"`
	FieldsBetaAppLocalizations          []string `url:"fields[betaAppLocalizations],omitempty"`
	FieldsBuilds                        []string `url:"fields[builds],omitempty"`
	FieldsBetaGroups                    []string `url:"fields[betaGroups],omitempty"`
	LimitBuilds                         int      `url:"limit[builds],omitempty"`
	LimitBetaGroups                     int      `url:"limit[betaGroups],omitempty"`
	LimitBetaAppLocalizations           int      `url:"limit[betaAppLocalizations],omitempty"`
	LimitPrices                         int      `url:"limit[prices],omitempty"`
	LimitAvailableTerritories           int      `url:"limit[availableTerritories],omitempty"`
	LimitAppStoreVersions               int      `url:"limit[appStoreVersions],omitempty"`
	LimitAppInfos                       int      `url:"limit[appInfos],omitempty"`
	FieldsEndUserLicenseAgreements      []string `url:"fields[endUserLicenseAgreements],omitempty"`
	FieldsAppStoreVersions              []string `url:"fields[appStoreVersions],omitempty"`
	FieldsTerritories                   []string `url:"fields[territories],omitempty"`
	FieldsAppPrices                     []string `url:"fields[appPrices],omitempty"`
	FieldsAppPreOrders                  []string `url:"fields[appPreOrders],omitempty"`
	FieldsAppInfos                      []string `url:"fields[appInfos],omitempty"`
	FieldsPerfPowerMetrics              []string `url:"fields[perfPowerMetrics],omitempty"`
	FilterAppStoreVersions              []string `url:"filter[appStoreVersions],omitempty"`
	FilterAppStoreVersionsPlatform      []string `url:"filter[appStoreVersions.platform],omitempty"`
	FilterAppStoreVersionsAppStoreState []string `url:"filter[appStoreVersions.appStoreState],omitempty"`
	LimitGameCenterEnabledVersions      int      `url:"limit[gameCenterEnabledVersions],omitempty"`
	FieldsGameCenterEnabledVersions     []string `url:"fields[gameCenterEnabledVersions],omitempty"`
	ExistsGameCenterEnabledVersions     bool     `url:"exists[gameCenterEnabledVersions],omitempty"`
	LimitInAppPurchases                 int      `url:"limit[inAppPurchases],omitempty"`
	FieldsInAppPurchases                []string `url:"fields[inAppPurchases],omitempty"`
	FieldsCiProducts                    []string `url:"fields[ciProducts],omitempty"`
	LimitAppClips                       int      `url:"limit[appClips],omitempty"`
	FieldsAppClips                      []string `url:"fields[appClips],omitempty"`
	FieldsReviewSubmissions             []string `url:"fields[reviewSubmissions],omitempty"`
	FieldsAppCustomProductPages         []string `url:"fields[appCustomProductPages],omitempty"`
	FieldsAppEvents                     []string `url:"fields[appEvents],omitempty"`
	LimitAppCustomProductPages          int      `url:"limit[appCustomProductPages],omitempty"`
	LimitAppEvents                      int      `url:"limit[appEvents],omitempty"`
	LimitReviewSubmissions              int      `url:"limit[reviewSubmissions],omitempty"`
	FieldsAppPricePoints                []string `url:"fields[appPricePoints],omitempty"`
	FieldsCustomerReviews               []string `url:"fields[customerReviews],omitempty"`
	FieldsSubscriptionGracePeriods      []string `url:"fields[subscriptionGracePeriods],omitempty"`
	FieldsPromotedPurchases             []string `url:"fields[promotedPurchases],omitempty"`
	FieldsSubscriptionGroups            []string `url:"fields[subscriptionGroups],omitempty"`
	LimitInAppPurchasesV2               int      `url:"limit[inAppPurchasesV2],omitempty"`
	LimitPromotedPurchases              int      `url:"limit[promotedPurchases],omitempty"`
	LimitSubscriptionGroups             int      `url:"limit[subscriptionGroups],omitempty"`
}

type ListAppsResponse struct {
	Data []struct {
		Type       string `json:"type"`
		ID         string `json:"id"`
		Attributes struct {
			Name                                   string      `json:"name"`
			BundleID                               string      `json:"bundleId"`
			Sku                                    string      `json:"sku"`
			PrimaryLocale                          string      `json:"primaryLocale"`
			IsOrEverWasMadeForKids                 bool        `json:"isOrEverWasMadeForKids"`
			SubscriptionStatusURL                  interface{} `json:"subscriptionStatusUrl"`
			SubscriptionStatusURLVersion           interface{} `json:"subscriptionStatusUrlVersion"`
			SubscriptionStatusURLForSandbox        interface{} `json:"subscriptionStatusUrlForSandbox"`
			SubscriptionStatusURLVersionForSandbox interface{} `json:"subscriptionStatusUrlVersionForSandbox"`
			AvailableInNewTerritories              bool        `json:"availableInNewTerritories"`
			ContentRightsDeclaration               string      `json:"contentRightsDeclaration"`
		} `json:"attributes"`
		Relationships struct {
			CiProduct struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"ciProduct"`
			BetaTesters struct {
				Links struct {
					Self string `json:"self"`
				} `json:"links"`
			} `json:"betaTesters"`
			BetaGroups struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"betaGroups"`
			AppStoreVersions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appStoreVersions"`
			PreReleaseVersions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"preReleaseVersions"`
			BetaAppLocalizations struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"betaAppLocalizations"`
			Builds struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"builds"`
			BetaLicenseAgreement struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"betaLicenseAgreement"`
			BetaAppReviewDetail struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"betaAppReviewDetail"`
			AppInfos struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appInfos"`
			AppClips struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appClips"`
			AppPricePoints struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appPricePoints"`
			PricePoints struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"pricePoints"`
			EndUserLicenseAgreement struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"endUserLicenseAgreement"`
			PreOrder struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"preOrder"`
			Prices struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"prices"`
			AppPriceSchedule struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appPriceSchedule"`
			AvailableTerritories struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"availableTerritories"`
			AppAvailability struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appAvailability"`
			InAppPurchases struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"inAppPurchases"`
			SubscriptionGroups struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"subscriptionGroups"`
			GameCenterEnabledVersions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"gameCenterEnabledVersions"`
			PerfPowerMetrics struct {
				Links struct {
					Related string `json:"related"`
				} `json:"links"`
			} `json:"perfPowerMetrics"`
			AppCustomProductPages struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appCustomProductPages"`
			InAppPurchasesV2 struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"inAppPurchasesV2"`
			PromotedPurchases struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"promotedPurchases"`
			AppEvents struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"appEvents"`
			ReviewSubmissions struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"reviewSubmissions"`
			SubscriptionGracePeriod struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"subscriptionGracePeriod"`
			CustomerReviews struct {
				Links struct {
					Self    string `json:"self"`
					Related string `json:"related"`
				} `json:"links"`
			} `json:"customerReviews"`
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

// https://developer.apple.com/documentation/appstoreconnectapi/list_apps
func (c *Client) ListApps(ctx context.Context, q *ListAppsQuery) (*ListAppsResponse, error) {
	v, err := query.Values(q)
	if err != nil {
		return nil, err
	}

	u := c.URL.JoinPath("v1/apps")
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

	body := new(ListAppsResponse)
	if err := decodeBody(res, body); err != nil {
		return nil, err
	}

	return body, nil
}
