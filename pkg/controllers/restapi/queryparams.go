package restapi

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/kartpop/connector-api/pkg/models"
)

var validQueryParams = map[string]bool{
	"location_name": true,
	"customer_name": true,
	"type":          true,
	"sort":          true,
	"limit":         true,
	"page":          true,
}

func ValidateAndGetQueryParams(url *url.URL) (models.ConnectorQueryParams, error) {
	// validate query parameters
	query, found, err := validateQueryParams(url.String())
	if err != nil {
		return models.ConnectorQueryParams{}, err
	}
	if !found {
		return models.ConnectorQueryParams{}, nil
	}

	// process query parameters
	var sort bool
	if _, present := query["sort"]; present {
		sort = true // TODO: currently only one type of sort enabled -> based on location_id first, then type
	}
	var limit, page int
	if strlimit, present := query["limit"]; present {
		if len(strlimit) > 1 {
			return models.ConnectorQueryParams{}, fmt.Errorf("expected 1 value for query parameter: limit, received %d", len(strlimit))
		}
		limit, err = strconv.Atoi(strlimit[0])
		if err != nil {
			return models.ConnectorQueryParams{}, err
		}
	}
	if strpage, present := query["page"]; present {
		if len(strpage) > 1 {
			return models.ConnectorQueryParams{}, fmt.Errorf("expected 1 value for query parameter: page, received %d", len(strpage))
		}
		page, err = strconv.Atoi(strpage[0])
		if err != nil {
			return models.ConnectorQueryParams{}, err
		}
	}
	queryParams := models.ConnectorQueryParams{
		CustomerNames: query["customer_name"],
		LocationNames: query["location_name"],
		Types:         query["type"],
		Sort:          sort,
		Limit:         limit,
		Page:          page,
	}

	return queryParams, nil
}

func validateQueryParams(urlstr string) (url.Values, bool, error) {
	_, after, found := strings.Cut(urlstr, "?")
	if !found {
		return nil, found, nil
	}

	queryMap, err := url.ParseQuery(after)
	if err != nil {
		return nil, found, err
	}

	// validate if queryMap has valid query params
	for k := range queryMap {
		if _, ok := validQueryParams[k]; !ok {
			return nil, found, fmt.Errorf("%s is not a valid query parameter", k)
		}
	}

	return queryMap, found, nil
}
