package appstoreconnect

import (
	"compress/gzip"
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/google/go-querystring/query"
)

type DownloadFinanceReportsQuery struct {
	FilterRegionCode   []string `url:"filter[regionCode]"`
	FilterReportDate   []string `url:"filter[reportDate]"`
	FilterReportType   []string `url:"filter[reportType]"`
	FilterVendorNumber []string `url:"filter[vendorNumber]"`
}

// https://developer.apple.com/documentation/appstoreconnectapi/download_finance_reports
func (c *Client) DownloadFinanceReports(ctx context.Context, q *DownloadFinanceReportsQuery, out io.Writer) error {
	v, err := query.Values(q)
	if err != nil {
		return err
	}

	u := c.URL.JoinPath("/v1/financeReports")
	u.RawQuery = v.Encode()

	req, err := c.newRequest(ctx, http.MethodGet, u, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Accept", "application/a-gzip")

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return err
	}

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("status code = %v", res.StatusCode)
	}

	gz, err := gzip.NewReader(res.Body)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if _, err = io.Copy(out, gz); err != nil {
		return err
	}

	return nil
}
