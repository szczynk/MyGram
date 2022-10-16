package delivery

import (
	"strconv"

	"github.com/szczynk/MyGram/models"

	"github.com/gin-gonic/gin"
)

func PaginateFromQuery(c *gin.Context) models.Pagination {
	var limit, page int
	var sort string

	query := c.Request.URL.Query()

	for key, value := range query {
		queryValue := value[len(value)-1]

		switch key {
		case "limit":
			limit, _ = strconv.Atoi(queryValue)
		case "page":
			page, _ = strconv.Atoi(queryValue)
		case "sort":
			sort = queryValue
		}
	}

	return models.Pagination{
		Limit: limit,
		Page:  page,
		Sort:  sort,
	}
}
