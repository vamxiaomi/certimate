package onepanel

import (
	"fmt"
	"net/http"
)

func (c *Client) UpdateSystemSSL(req *UpdateSystemSSLRequest) (*UpdateSystemSSLResponse, error) {
	resp := &UpdateSystemSSLResponse{}

	path := "/settings/ssl/update"
	if c.version == "v2" {
		path = "/core" + path
	}

	err := c.sendRequestWithResult(http.MethodPost, path, req, resp)
	return resp, err
}

func (c *Client) SearchWebsiteSSL(req *SearchWebsiteSSLRequest) (*SearchWebsiteSSLResponse, error) {
	resp := &SearchWebsiteSSLResponse{}
	err := c.sendRequestWithResult(http.MethodPost, "/websites/ssl/search", req, resp)
	return resp, err
}

func (c *Client) GetWebsiteSSL(req *GetWebsiteSSLRequest) (*GetWebsiteSSLResponse, error) {
	if req.SSLID == 0 {
		return nil, fmt.Errorf("1panel api error: invalid parameter: SSLID")
	}

	resp := &GetWebsiteSSLResponse{}
	err := c.sendRequestWithResult(http.MethodGet, fmt.Sprintf("/websites/ssl/%d", req.SSLID), req, resp)
	return resp, err
}

func (c *Client) UploadWebsiteSSL(req *UploadWebsiteSSLRequest) (*UploadWebsiteSSLResponse, error) {
	resp := &UploadWebsiteSSLResponse{}
	err := c.sendRequestWithResult(http.MethodPost, "/websites/ssl/upload", req, resp)
	return resp, err
}

func (c *Client) GetHttpsConf(req *GetHttpsConfRequest) (*GetHttpsConfResponse, error) {
	if req.WebsiteID == 0 {
		return nil, fmt.Errorf("1panel api error: invalid parameter: WebsiteID")
	}

	resp := &GetHttpsConfResponse{}
	err := c.sendRequestWithResult(http.MethodGet, fmt.Sprintf("/websites/%d/https", req.WebsiteID), req, resp)
	return resp, err
}

func (c *Client) UpdateHttpsConf(req *UpdateHttpsConfRequest) (*UpdateHttpsConfResponse, error) {
	if req.WebsiteID == 0 {
		return nil, fmt.Errorf("1panel api error: invalid parameter: WebsiteID")
	}

	resp := &UpdateHttpsConfResponse{}
	err := c.sendRequestWithResult(http.MethodPost, fmt.Sprintf("/websites/%d/https", req.WebsiteID), req, resp)
	return resp, err
}
