package category

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var basepath = "/v1/categories"

type CategoryRoutes struct {
	periodHandler ICategoryHandler
}

func NewCategoryRoutes(periodHandler ICategoryHandler) CategoryRoutes {
	return CategoryRoutes{
		periodHandler: periodHandler,
	}
}

func (categoryRoutes *CategoryRoutes) CategoryRouteMapping() (string, []controllers.Route) {

	return basepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               categoryRoutes.periodHandler.CreateCategory,
			RequiresAuthentication: true,
		},
		{
			URI:                    "",
			Method:                 http.MethodGet,
			Function:               categoryRoutes.periodHandler.ListCategories,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodGet,
			Function:               categoryRoutes.periodHandler.GetCategoryById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               categoryRoutes.periodHandler.UpdateCategory,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               categoryRoutes.periodHandler.DeleteCategory,
			RequiresAuthentication: true,
		},
	}
}
