package controllers

import (
	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_formulacion_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/requestresponse"
)

// InversionController operations for Inversion
type InversionController struct {
	beego.Controller
}

// URLMapping ...
func (c *InversionController) URLMapping() {
	c.Mapping("AddProyecto", c.AddProyecto)
	c.Mapping("EditProyecto", c.EditProyecto)
	c.Mapping("GuardarDocumentos", c.GuardarDocumentos)
	c.Mapping("GetProyectoId", c.GetProyectoId)
	c.Mapping("GetMetasProyect", c.GetMetasProyect)
	c.Mapping("GetAllProyectos", c.GetAllProyectos)
	c.Mapping("ActualizarSubgrupoDetalle", c.ActualizarSubgrupoDetalle)
	c.Mapping("ActualizarProyectoGeneral", c.ActualizarProyectoGeneral)
	c.Mapping("CrearPlan", c.CrearPlan)
	c.Mapping("GetPlanId", c.GetPlanId)
	c.Mapping("GetPlan", c.GetPlan)
	c.Mapping("GuardarMeta", c.GuardarMeta)
	c.Mapping("ArmonizarInversion", c.ArmonizarInversion)
	c.Mapping("ActualizarMetaPlan", c.ActualizarMetaPlan)
	c.Mapping("AllMetasPlan", c.AllMetasPlan)
	c.Mapping("ActualizarActividad", c.ActualizarActividad)
	c.Mapping("ActualizarTablaActividad", c.ActualizarTablaActividad)
	c.Mapping("ActualizarPresupuestoMeta", c.ActualizarPresupuestoMeta)
	c.Mapping("VerificarMagnitudesProgramadas", c.VerificarMagnitudesProgramadas)
	c.Mapping("VersionarPlan", c.VersionarPlan)
}

// AddProyecto ...
// @Title AddProyecto
// @Description post AddProyecto
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @router /proyectos [post]
func (c *InversionController) AddProyecto() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.AddProyecto(datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// EditProyecto ...
// @Title EditProyecto
// @Description post EditProyecto
// @Param	body		body 	{}	true		"body for Plan content"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /proyectos/:id [put]
func (c *InversionController) EditProyecto() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody
	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.EditProyecto(id, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GuardarDocumentos ...
// @Title GuardarDocumentos
// @Description post AddProyecto
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /documentos [post]
func (c *InversionController) GuardarDocumentos() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.GuardarDocumentos(datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetProyectoId ...
// @Title GetProyectoId
// @Description get GetProyectoId
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /proyectos/:id [get]
func (c *InversionController) GetProyectoId() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.GetProyectoId(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetMetasProyect ...
// @Title GetMetasProyect
// @Description get GetMetasProyect
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /proyectos/:id/metas [get]
func (c *InversionController) GetMetasProyect() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.GetMetasProyect(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetAllProyectos ...
// @Title GetAllProyectos
// @Description get GetAllProyectos
// @Param	aplicativo_id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :aplicativo_id is empty
// @router /proyectos/:aplicativo_id [get]
func (c *InversionController) GetAllProyectos() {
	defer errorhandler.HandlePanic(&c.Controller)

	tipo_plan_id := c.Ctx.Input.Param(":aplicativo_id")

	if resultado, err := services.GetAllProyectos(tipo_plan_id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ActualizarSubgrupoDetalle ...
// @Title ActualizarSubgrupoDetalle
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /subgrupo/detalle/:id [put]
func (c *InversionController) ActualizarSubgrupoDetalle() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody
	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.ActualizarSubgrupoDetalle(id, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ActualizarProyectoGeneral ...
// @Title ActualizarProyectoGeneral
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /actualizar_proyecto/:id [put]
func (c *InversionController) ActualizarProyectoGeneral() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody
	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.ActualizarProyectoGeneral(id, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// CrearPlan ...
// @Title CrearPlan
// @Description post CrearPlan
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403
// @router /planes [post]
func (c *InversionController) CrearPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.CrearPlan(datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GetPlanId ...
// @Title GetPlanId
// @Description get GetPlanId
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /planes/:id [get]
func (c *InversionController) GetPlanId() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.GetPlanId(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// GuardarMeta ...
// @Title GuardarMeta
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /metas/:id [put]
func (c *InversionController) GuardarMeta() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.GuardarMeta(id, datos); err == nil {
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
// @Success 200
// @Failure 403 :id is empty
// @router /planes/:id/informacion [get]
func (c *InversionController) GetPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.ObtenerPlan(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ArmonizarInversion ...
// @Title ArmonizarInversion
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /armonizar/:id [put]
func (c *InversionController) ArmonizarInversion() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ArmonizarInversion(id, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ActualizarMetaPlan ...
// @Title ActualizarMetaPlan
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /metas/:id/:index [put]
func (c *InversionController) ActualizarMetaPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ActualizarMetaPlan(id, index, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// AllMetasPlan ...
// @Title AllMetasPlan
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /metas/:id [get]
func (c *InversionController) AllMetasPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.AllMetasPlan(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// InactivarMeta ...
// @Title InactivarMeta
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /metas/:id/:index/inactivar [put]
func (c *InversionController) InactivarMeta() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	if resultado, err := services.InactivarMeta(id, index); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ProgMagnitudesPlan ...
// @Title ProgMagnitudesPlan
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /magnitudes/:id/:index [put]
func (c *InversionController) ProgMagnitudesPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ProgMagnitudesPlan(id, index, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// MagnitudesProgramadas ...
// @Title MagnitudesProgramadas
// @Description get MagnitudesProgramadas
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /magnitudes/:id/:indexMeta [get]
func (c *InversionController) MagnitudesProgramadas() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":indexMeta")

	if resultado, err := services.MagnitudesProgramadas(id, index); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// CrearGrupoMeta ...
// @Title CrearPlan
// @Description post CrearPlan
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403
// @router /metas/grupo [post]
func (c *InversionController) CrearGrupoMeta() {
	defer errorhandler.HandlePanic(&c.Controller)

	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.CrearGrupoMeta(datos); err == nil {
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
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /actividad/:id/:index [put]
func (c *InversionController) ActualizarActividad() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ActualizarActividadInv(id, index, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ActualizarTablaActividad ...
// @Title ActualizarTablaActividad
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /actividad/:id/:index/tabla [put]
func (c *InversionController) ActualizarTablaActividad() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ActualizarTablaActividad(id, index, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// ActualizarPresupuestoMeta ...
// @Title ActualizarPresupuestoMeta
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /metas/presupuestos/:id/:index [put]
func (c *InversionController) ActualizarPresupuestoMeta() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	datos := c.Ctx.Input.RequestBody

	if resultado, err := services.ActualizarPresupuestoMeta(id, index, datos); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}

// VerificarMagnitudesProgramadas ...
// @Title VerificarMagnitudesProgramadas
// @Description get VerificarMagnitudesProgramadas
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /magnitudes/:id/verificar [get]
func (c *InversionController) VerificarMagnitudesProgramadas() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.VerificarMagnitudesProgramadas(id); err == nil {
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
// @Description post Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /planes/:id/versionar [post]
func (c *InversionController) VersionarPlan() {
	defer errorhandler.HandlePanic(&c.Controller)

	id := c.Ctx.Input.Param(":id")

	if resultado, err := services.VersionarPlanInv(id); err == nil {
		c.Ctx.Output.SetStatus(200)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 200, resultado)
	} else {
		c.Ctx.Output.SetStatus(404)
		c.Data["json"] = requestresponse.APIResponseDTO(true, 404, nil, err)
	}

	c.ServeJSON()
}
