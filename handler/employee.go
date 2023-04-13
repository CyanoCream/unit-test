package handler

//
//// delivery, controller
//
//import (
//	"sesi_8/helper"
//	"sesi_8/model"
//	"strconv"
//
//	"github.com/gin-gonic/gin"
//)
//
//// CreateEmployee godoc
//// @Summary Create a new employee
//// @Description Create a new employee
//// @Tags employees
//// @Accept  json
//// @Produce  json
//// @Param employee body model.EmployeeReq true "Create a new employee"
//// @Success 200 {object} helper.Response“
//// @Failure 400 {object} helper.Response
//// @Failure 500 {object} helper.Response
//// @Router /employee [post]
//func (h HttpServer) CreateEmployee(c *gin.Context) {
//	in := model.Employee{}
//
//	err := c.BindJSON(&in)
//	if err != nil {
//		helper.BadRequest(c, err.Error())
//		return
//	}
//
//	err = in.Validation()
//	if err != nil {
//		helper.BadRequest(c, err.Error())
//		return
//	}
//
//	// call service
//	res, err := h.app.CreateEmployee(in)
//	if err != nil {
//		helper.InternalServerError(c, err.Error())
//		return
//	}
//
//	helper.OkWithMessage(c, "Success", res)
//}
//
//// GetEmployeeByID godoc
//// @Summary Get details for a given id
//// @Description Get details of employee corresponding to the input id
//// @Tags employees
//// @Accept  json
//// @Produce  json
//// @Param id path int true "ID of the employee"
//// @Security BasicAuth
//// @Success 200 {object} helper.Response
//// @Failure 401 {object} helper.Response
//// @Failure 404 {object} helper.Response
//// @Failure 500 {object} helper.Response
//// @Router /employee/{id} [get]
//func (h HttpServer) GetEmployeeByID(c *gin.Context) {
//	id := c.Param("id")
//
//	idInt, err := strconv.Atoi(id)
//	if err != nil {
//		helper.BadRequest(c, err.Error())
//		return
//	}
//
//	// call service
//	res, err := h.app.GetEmployeeByID(int64(idInt))
//	if err != nil {
//		if err.Error() == helper.ErrNotFound {
//			helper.NotFound(c, err.Error())
//			return
//		}
//		helper.InternalServerError(c, err.Error())
//		return
//	}
//
//	helper.OkWithMessage(c, "Success", res)
//}
//
//// UpdateEmployee godoc
//// @Summary Update an employee
//// @Description Update an employee
//// @Tags employees
//// @Accept  json
//// @Produce  json
//// @Param id path int true "ID of the employee"
//// @Param employee body model.EmployeeReq true "Update an employee"
//// @Success 200 {object} helper.Response
//// @Failure 400 {object} helper.Response
//// @Failure 500 {object} helper.Response
//// @Router /employee/{id} [put]
//func (h HttpServer) UpdateEmployee(c *gin.Context) {
//	id := c.Param("id")
//
//	idInt, err := strconv.Atoi(id)
//	if err != nil {
//		helper.BadRequest(c, err.Error())
//		return
//	}
//
//	in := model.Employee{}
//
//	err = c.BindJSON(&in)
//	if err != nil {
//		helper.BadRequest(c, err.Error())
//		return
//	}
//
//	in.ID = idInt
//	// call service
//	res, err := h.app.UpdateEmployee(in)
//	if err != nil {
//		helper.InternalServerError(c, err.Error())
//		return
//	}
//
//	helper.OkWithMessage(c, "Success", res)
//}
//
//// DeleteEmployee godoc
//// @Summary Delete an employee
//// @Description Delete an employee
//// @Tags employees
//// @Accept  json
//// @Produce  json
//// @Param id path int true "ID of the employee"
//// @Success 200 {object} helper.Response
//// @Failure 500 {object} helper.Response
//// @Router /employee/{id} [delete]
//func (h HttpServer) DeleteEmployee(c *gin.Context) {
//	id := c.Param("id")
//
//	idInt, err := strconv.Atoi(id)
//	if err != nil {
//		helper.BadRequest(c, err.Error())
//		return
//	}
//
//	// call service
//	err = h.app.DeleteEmployee(int64(idInt))
//	if err != nil {
//		helper.InternalServerError(c, err.Error())
//		return
//	}
//
//	helper.OkWithMessage(c, "Successfully deleted employee", nil)
//}
//
//// func (h HttpServer) GetEmployeeByID(c *gin.Context) {
//// 	id := c.Param("id")
//
//// 	// call service
//// 	res, err := h.app.GetEmployeeByID(id)
//// 	if err != nil {
//// 		helper.InternalServerError(c, err.Error())
//// 		return
//// 	}
//
//// 	helper.Ok(c, res)
//// }
