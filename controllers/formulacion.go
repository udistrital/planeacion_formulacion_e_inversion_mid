package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_formulacion_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

// FormulacionController operations for Formulacion
type FormulacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *FormulacionController) URLMapping() {
	c.Mapping("ClonarFormato", c.ClonarFormato)
	c.Mapping("GuardarActividad", c.GuardarActividad)
	c.Mapping("GetPlan", c.GetPlan)
	c.Mapping("ActualizarActividad", c.ActualizarActividad)
	c.Mapping("DeleteActividad", c.DeleteActividad)
	c.Mapping("GetAllActividades", c.GetAllActividades)
	c.Mapping("GuardarIdentificacion", c.GuardarIdentificacion)
	c.Mapping("GetAllIdentificacion", c.GetAllIdentificacion)
	c.Mapping("DeleteIdentificacion", c.DeleteIdentificacion)
	c.Mapping("VersionarPlan", c.VersionarPlan)
	c.Mapping("GetPlanVersiones", c.GetPlanVersiones)
	c.Mapping("PonderacionActividades", c.PonderacionActividades)
	c.Mapping("GetRubros", c.GetRubros)
	c.Mapping("GetUnidades", c.GetUnidades)
	c.Mapping("VinculacionTercero", c.VinculacionTercero)
	c.Mapping("Planes", c.Planes)
	c.Mapping("VerificarIdentificaciones", c.VerificarIdentificaciones)
	c.Mapping("PlanesEnFormulacion", c.PlanesEnFormulacion)
	c.Mapping("GetPlanesUnidadesComun", c.GetPlanesUnidadesComun)
	c.Mapping("DefinirFechasFuncionamiento", c.DefinirFechasFuncionamiento)
	c.Mapping("CalculosDocentes", c.CalculosDocentes)
	c.Mapping("EstructuraPlanes", c.EstructuraPlanes)
	c.Mapping("PlanesDeAccion", c.PlanesDeAccion)
	c.Mapping("PlanesDeAccionPorUnidad", c.PlanesDeAccionPorUnidad)
}

// ClonarFormato ...
// @Title ClonarFormato
// @Description post Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /formato/:id/clonar [post]
func (c *FormulacionController) ClonarFormato() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	body := c.Ctx.Input.RequestBody

	if resultado, err := services.ClonarFormato(id, body); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GuardarActividad ...
// @Title GuardarActividad
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /actividad/:id [put]
func (c *FormulacionController) GuardarActividad() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.GuardarActividad(id, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetPlan ...
// @Title GetPlan
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /plan/:id/:index [get]
func (c *FormulacionController) GetPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	if resultado, err := services.GetPlan(id, index); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ActualizarActividad ...
// @Title ActualizarActividad
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /actividad/:id/:index [put]
func (c *FormulacionController) ActualizarActividad() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ActualizarActividad(id, index, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// DeleteActividad ...
// @Title DeleteActividad
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /actividad/:id/:index/desactivar [put]
func (c *FormulacionController) DeleteActividad() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	if resultado, err := services.DeleteActividad(id, index); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetAllActividades ...
// @Title GetAllActividades
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /actividad/:id/ [get]
func (c *FormulacionController) GetAllActividades() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.GetAllActividades(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetArbolArmonizacion ...
// @Title GetArbolArmonizacion
// @Description post Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /get_arbol_armonizacion/:id/ [post]
func (c *FormulacionController) GetArbolArmonizacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.GetArbolArmonizacion(datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GuardarIdentificacion ...
// @Title GuardarIdentificacion
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	idTipo		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /identificacion/:id/:idTipo [put]
func (c *FormulacionController) GuardarIdentificacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	tipoIdenti := c.Ctx.Input.Param(":idTipo")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.GuardarIdentificacion(id, tipoIdenti, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetAllIdentificacion ...
// @Title GetAllIdentificacion
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	idTipo		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /identificacion/:id/:idTipo [get]
func (c *FormulacionController) GetAllIdentificacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	tipoIdenti := c.Ctx.Input.Param(":idTipo")

	if resultado, err := services.GetAllIdentificacion(id, tipoIdenti); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// DeleteIdentificacion ...
// @Title DeleteIdentificacion
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	idTipo		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /identificacion/:id/:idTipo/:index [put]
func (c *FormulacionController) DeleteIdentificacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	idTipo := c.Ctx.Input.Param(":idTipo")
	index := c.Ctx.Input.Param(":index")

	if resultado, err := services.DeleteIdentificacion(id, idTipo, index); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// VersionarPlan ...
// @Title VersionarPlan
// @Description post Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /plan/:id/versionar [post]
func (c *FormulacionController) VersionarPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.VersionarPlan(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetPlanVersiones ...
// @Title GetPlanVersiones
// @Description get Formulacion by id
// @Param	unidad		path 	string	true		"The key for staticblock"
// @Param	vigencia		path 	string	true		"The key for staticblock"
// @Param	nombre		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /plan/versiones/:unidad/:vigencia/:nombre [get]
func (c *FormulacionController) GetPlanVersiones() {
	defer errorhandler.HandlePanic(&c.Controller)

	unidad := c.Ctx.Input.Param(":unidad")
	vigencia := c.Ctx.Input.Param(":vigencia")
	nombre := c.Ctx.Input.Param(":nombre")

	if resultado, err := services.GetPlanVersiones(unidad, vigencia, nombre); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetPonderacionActividades ...
// @Title GetPonderacionActividades
// @Description get Formulacion by id
// @Param	plan		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /actividad/ponderacion/:plan [get]
func (c *FormulacionController) PonderacionActividades() {
	defer errorhandler.HandlePanic(&c.Controller)

	plan := c.Ctx.Input.Param(":plan")

	if resultado, err := services.PonderacionActividades(plan); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetRubros ...
// @Title GetRubros
// @Description get Rubros
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /rubros [get]
func (c *FormulacionController) GetRubros() {
	defer errorhandler.HandlePanic(&c.Controller)

	if resultado, err := services.GetRubros(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetUnidades ...
// @Title GetUnidades
// @Description get Unidades
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /unidades [get]
func (c *FormulacionController) GetUnidades() {
	defer errorhandler.HandlePanic(&c.Controller)

	if resultado, err := services.GetUnidades(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// VinculacionTercero ...
// @Title VinculacionTercero
// @Description get VinculacionTercero
// @Param	tercero_id	path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /tercero/:tercero_id [get]
func (c *FormulacionController) VinculacionTercero() {
	defer errorhandler.HandlePanic(&c.Controller)

	terceroId := c.Ctx.Input.Param(":tercero_id")

	if resultado, err := services.VinculacionTercero(terceroId); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// Planes ...
// @Title Planes
// @Description get Rubros
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /planes [get]
func (c *FormulacionController) Planes() {
	defer errorhandler.HandlePanic(&c.Controller)

	if resultado, err := services.Planes(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// VerificarIdentificaciones ...
// @Title VerificarIdentificaciones
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /identificacion/verificacion/:id [get]
func (c *FormulacionController) VerificarIdentificaciones() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.VerificarIdentificaciones(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// Get Planes En Formulacion ...
// @Title GetPlanesEnFormulacion
// @Description get Planes en formulacion
// @Success 200 {object} models.Formulacion
// @Failure 400 bad response
// @router /planes_formulacion [get]
func (c *FormulacionController) PlanesEnFormulacion() {
	defer errorhandler.HandlePanic(&c.Controller)

	if resultado, err := services.PlanesEnFormulacion(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// CalculosDocentes ...
// @Title CalculosDocentes
// @Description post Formulacion
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /calculos_docentes [post]
func (c *FormulacionController) CalculosDocentes() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.CalculosDocentes(datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// EstructuraPlanes ...
// @Title EstructuraPlanes
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /estructura_planes/:id [put]
func (c *FormulacionController) EstructuraPlanes() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.EstructuraPlanes(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// DefinirFechasFuncionamiento ...
// @Title DefinirFechasFuncionamiento
// @Description Peticion POST para definir fechas en planes de acción de funcionamiento
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 400 bad request
// @router /habilitar_fechas_funcionamiento [post]
func (c *FormulacionController) DefinirFechasFuncionamiento() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.DefinirFechasFuncionamiento(datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetPlanesUnidadesComun ...
// @Title GetPlanesUnidadesComun
// @Description post Get planes en comun con unidades by id periodo-seguimiento
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /get_planes_unidades_comun/:id [post]
func (c *FormulacionController) GetPlanesUnidadesComun() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.GetPlanesUnidadesComun(id, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// PlanesDeAccion ...
// @Title PlanesDeAccion
// @Description get Planes de Acción
// @Success 200 {object} models.Formulacion
// @Failure 400 bad response
// @router /planes_accion [get]
func (c *FormulacionController) PlanesDeAccion() {
	defer errorhandler.HandlePanic(&c.Controller)

	if resultado, err := services.GetPlanesDeAccion(); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}
	c.ServeJSON()
}

// PlanesDeAccionPorUnidad...
// @Title PlanesDeAccionPorUnidad
// @Description get Planes de accion filtrando por el id de la unidad
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /planes_accion/:unidad_id [get]
func (c *FormulacionController) PlanesDeAccionPorUnidad() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":unidad_id")

	if resultado, err := services.GetPlanesDeAccionPorUnidad(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}
	c.ServeJSON()

}