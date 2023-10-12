package ctfd

import "fmt"

type Challenge struct {
	Name     string `json:"name"`
	Category string `json:"category"`
	Desc     string `json:"description"`
	Type     string `json:"type"`
	Value    int    `json:"value"`
}

func (c *APIClient) getChallenge(id string) (interface{}, error) {
	obj, err := c.get(fmt.Sprintf("/api/v1/challenges/%v", id))
	if err != nil {
		return nil, err
	}
	return obj["data"], nil
}

func (c *APIClient) getChallenges() (interface{}, error) {
	obj, err := c.get("/api/v1/challenges")
	if err != nil {
		return nil, err
	}
	return obj["data"], nil
}

func (c *APIClient) getChallengeFiles(chalId string) (interface{}, error) {
	obj, err := c.get(fmt.Sprintf("/api/v1/challenges/%v/files", chalId))
	if err != nil {
		return nil, err
	}
	return obj["data"], nil
}
