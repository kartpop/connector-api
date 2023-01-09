package restapi

import (
	"net/url"
	"strconv"

	"github.com/sede-x/gopoc-connector/pkg/models"
)

func GetQueryParams(query url.Values) (models.ConnectorQueryParams, error) {

	var sort bool
	if _, present := query["sort"]; present {
		sort = true // TODO: currently only one type of sort enabled -> based on location_id first, then type
	}
	var limit, page int
	var err error
	if strlimit, present := query["limit"]; present {
		limit, err = strconv.Atoi(strlimit[0])
		if err != nil {
			return models.ConnectorQueryParams{}, err
		}
	}
	if strpage, present := query["page"]; present {
		page, err = strconv.Atoi(strpage[0])
		if err != nil {
			return models.ConnectorQueryParams{}, err
		}
	}
	queryParams := models.ConnectorQueryParams{
		LocationIds: query["location_id"],
		Types:       query["type"],
		Sort:        sort,
		Limit:       limit,
		Page:        page,
	}
	return queryParams, nil
}
