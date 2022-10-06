package gapi

type HealthResponse struct {
	Code     string `json:"code,omitempty"`
	Message  string `json:"message,omitempty"`
	Commit   string `json:"commit,omitempty"`
	Database string `json:"database,omitempty"`
	Version  string `json:"version,omitempty"`
}

func (c *Client) Health() (HealthResponse, error) {
	health := HealthResponse{}
	err := c.request("GET", "/api/health", nil, nil, &health)
	if err != nil {
		return health, err
	}
	return health, nil
}
