package subcategory

import (
	"marcelofelixsalgado/financial-period-api/api/controllers"
	"net/http"
)

var basepath = "/v1/subcategories"

type SubCategoryRoutes struct {
	periodHandler ISubCategoryHandler
}

func NewSubCategoryRoutes(periodHandler ISubCategoryHandler) SubCategoryRoutes {
	return SubCategoryRoutes{
		periodHandler: periodHandler,
	}
}

func (categoryRoutes *SubCategoryRoutes) SubCategoryRouteMapping() (string, []controllers.Route) {

	return basepath, []controllers.Route{
		{
			URI:                    "",
			Method:                 http.MethodPost,
			Function:               categoryRoutes.periodHandler.CreateSubCategory,
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
			Function:               categoryRoutes.periodHandler.GetSubCategoryById,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodPut,
			Function:               categoryRoutes.periodHandler.UpdateSubCategory,
			RequiresAuthentication: true,
		},
		{
			URI:                    "/:id",
			Method:                 http.MethodDelete,
			Function:               categoryRoutes.periodHandler.DeleteSubCategory,
			RequiresAuthentication: true,
		},
	}
}
