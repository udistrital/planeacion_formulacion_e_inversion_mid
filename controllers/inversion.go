package controllers

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	formulacionhelper "github.com/udistrital/planeacion_formulacion_mid/helpers/formulacionHelper"
	inversionhelper "github.com/udistrital/planeacion_formulacion_mid/helpers/inversionHelper"
	"github.com/udistrital/planeacion_formulacion_mid/services"
	"github.com/udistrital/utils_oas/errorhandler"
	"github.com/udistrital/utils_oas/request"
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	// var res map[string]interface{}
	// var hijos []map[string]interface{}
	// inversionhelper.Limpia()

	// if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
	// 	request.LimpiezaRespuestaRefactor(res, &hijos)
	// 	inversionhelper.GetSons(hijos, index)
	// 	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Meta Inactivada"}
	// } else {
	// 	panic(err)
	// }

	// c.ServeJSON()
	var res map[string]interface{}
	var hijos []map[string]interface{}
	var tabla map[string]interface{}
	var auxHijos []interface{}
	inversionhelper.Limpia()
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &hijos)
		for i := 0; i < len(hijos); i++ {
			auxHijos = append(auxHijos, hijos[i]["_id"])
		}
		fmt.Println(auxHijos, "auxhijos")
		inversionhelper.GetSons(auxHijos, index)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": tabla}
	} else {
		panic(map[string]interface{}{"funcion": "AllMetasPlan", "err": "Error al consultar metas del plan \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ProgMagnitudesPlan ...
// @Title ProgMagnitudesPlan
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /magnitudes/:id/:index [put]
func (c *InversionController) ProgMagnitudesPlan() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")
	var res map[string]interface{}
	var body map[string]interface{}
	var subgrupo []map[string]interface{}
	var id_subgrupoDetalle string
	//var index string
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var magnitudesUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})

	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	//index = body["indiceMetaProyecto"].(string)
	//magnitudes_data, _ := json.Marshal(body)
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		fmt.Println(subgrupo, "subgrupo")
		//armo_dato := make(map[string]interface{})
		subgrupoPost := make(map[string]interface{})
		subDetallePost := make(map[string]interface{})

		if len(subgrupo) > 0 {
			id_subgrupoDetalle = subgrupo[0]["_id"].(string)
			fmt.Println(id_subgrupoDetalle, "id_subgrupoDetalle")
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+id_subgrupoDetalle, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				if len(respuestaLimpia) > 0 {
					fmt.Println("ingresa a PUT")
					subgrupo_detalle = respuestaLimpia[0]
					fmt.Println(subgrupo_detalle["dato"], "subgrupo detalle put")
					magnitud := make(map[string]interface{})

					if subgrupo_detalle["dato"] == nil {
						magnitud["index"] = index
						magnitud["dato"] = body
						magnitud["activo"] = true
						i := strconv.Itoa(magnitud["index"].(int))
						dato[i] = magnitud
						b, _ := json.Marshal(dato)
						str := string(b)
						subgrupo_detalle["dato"] = str
					} else {
						dato_str := subgrupo_detalle["dato"].(string)
						json.Unmarshal([]byte(dato_str), &dato)
						magnitud["index"] = index
						magnitud["dato"] = body
						magnitud["activo"] = true
						//i := strconv.Itoa(magnitud["index"].(int))
						dato[index] = magnitud
						b, _ := json.Marshal(dato)
						str := string(b)
						subgrupo_detalle["dato"] = str
					}
					subDetallePost["dato"] = subgrupo_detalle["dato"]
					subDetallePost["subgrupo_id"] = id_subgrupoDetalle
					subDetallePost["fecha_creacion"] = subgrupo_detalle["fecha_creacion"]
					subDetallePost["nombre"] = "Detalle Información Programación de Magnitudes y Presupuesto"
					subDetallePost["descripcion"] = "Magnitudes"
					subDetallePost["activo"] = true

					fmt.Println(subDetallePost, "dataJSON")

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &magnitudesUpdate)
						fmt.Println(magnitudesUpdate, "update1290")
					} else {
						panic(map[string]interface{}{"funcion": "ProgMagnitudesPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
					}
				} else {
					subgrupo_detalle := make(map[string]interface{})
					magnitud := make(map[string]interface{})
					magnitud["index"] = index
					magnitud["dato"] = body
					magnitud["activo"] = true
					fmt.Println(magnitud["index"], "index")
					//i := strconv.Itoa(magnitud["index"].(int))
					//fmt.Println(i, "index")
					dato[index] = magnitud
					b, _ := json.Marshal(dato)
					str := string(b)
					subgrupo_detalle["dato"] = str

					subDetallePost["subgrupo_id"] = id_subgrupoDetalle
					subDetallePost["nombre"] = "Detalle Información Programación de Magnitudes y Presupuesto"
					subDetallePost["descripcion"] = "Magnitudes"
					subDetallePost["dato"] = subgrupo_detalle["dato"]
					subDetallePost["activo"] = true
					fmt.Println(subDetallePost, "dataJSON")

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &magnitudesUpdate)
						fmt.Println(res, "update1304")
					} else {
						panic(map[string]interface{}{"funcion": "ProgMagnitudesPlan", "err": "Error registrando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
					}
				}

			} else {
				panic(map[string]interface{}{"funcion": "ProgMagnitudesPlan", "err": "Error get subgrupo-detalle", "status": "400", "log": err})
			}
		} else {
			subgrupoPost["nombre"] = "Programación Magnitudes y Prespuesto Plan de Inversión"
			subgrupoPost["descripcion"] = "Magnitudes"
			subgrupoPost["padre"] = id
			subgrupoPost["activo"] = true
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/", "POST", &respuesta, subgrupoPost); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
				fmt.Println(respuestaLimpia, "respuesta subgrupo POST")
				subgrupo_detalle = respuestaLimpia[0]
				magnitud := make(map[string]interface{})
				magnitud["index"] = index
				magnitud["dato"] = body
				magnitud["activo"] = true
				i := strconv.Itoa(magnitud["index"].(int))
				dato[i] = magnitud
				b, _ := json.Marshal(dato)
				str := string(b)
				subgrupo_detalle["dato"] = str
				subDetallePost["subgrupo_id"] = subgrupo_detalle["_id"]
				subDetallePost["nombre"] = "Detalle Información Programación de Magnitudes y Presupuesto"
				subDetallePost["descripcion"] = "Magnitudes"
				subDetallePost["dato"] = subgrupo_detalle["dato"]
				subDetallePost["activo"] = true
				fmt.Println(subDetallePost, "dataJSON")

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
					request.LimpiezaRespuestaRefactor(res, &magnitudesUpdate)
					fmt.Println(magnitudesUpdate, "update1331")
				} else {
					panic(map[string]interface{}{"funcion": "ProgMagnitudesPlan", "err": "Error registrando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}

			} else {
				panic(map[string]interface{}{"funcion": "ProgMagnitudesPlan", "err": "Error registrando subgrupo", "status": "400", "log": err})
			}
		}
		//formulacionhelper.Limpia()
		//tree := formulacionhelper.BuildTreeFa(hijos, index)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": res}
	} else {
		panic(err)

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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":indexMeta")
	var res map[string]interface{}
	//var body map[string]interface{}
	var subgrupo map[string]interface{}
	//var id_subgrupoDetalle string
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	//var armonizacionUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	//json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	//armonizacion_data, _ := json.Marshal(body)
	dato := make(map[string]interface{})
	var magnitud map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuestaLimpia)
		subgrupo = respuestaLimpia[0]
		fmt.Println(res, "subgrupo")
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+subgrupo["_id"].(string), &respuesta); err != nil {
			panic(map[string]interface{}{"funcion": "MagnitudesProgramadas", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["dato"] != nil {

			dato_str := subgrupo_detalle["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)
			for index_actividad := range dato {
				if index_actividad == index {
					aux_actividad := dato[index_actividad].(map[string]interface{})
					magnitud = aux_actividad
				}
			}

		}
		fmt.Println(magnitud)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": magnitud}
	} else {
		panic(map[string]interface{}{"funcion": "MagnitudesProgramadas", "err": "Error consultando subgrupo", "status": "400", "log": err})
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

	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	//id := c.Ctx.Input.Param(":id")

	var respuesta map[string]interface{}
	//var respuestaHijos map[string]interface{}
	//var hijos []map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}
	var respuestaPost map[string]interface{}
	var planSubgrupo map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &parametros)
	plan := make(map[string]interface{})
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	id := parametros["id"].(string)
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuesta); err == nil {

		request.LimpiezaRespuestaRefactor(respuesta, &planFormato)

		plan["nombre"] = "" + planFormato["nombre"].(string)
		plan["descripcion"] = planFormato["descripcion"].(string)
		plan["tipo_plan_id"] = planFormato["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planFormato["aplicativo_id"].(string)
		plan["activo"] = planFormato["activo"]
		plan["formato"] = false
		plan["vigencia"] = parametros["vigencia"].(string)
		plan["dependencia_id"] = parametros["dependencia_id"].(string)
		plan["documento_id"] = parametros["indexMeta"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		plan["arbol_padre_id"] = parametros["arbol_padre_id"]

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/", "POST", &respuestaPost, plan); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaPost, &planSubgrupo)
			padre := planSubgrupo["_id"].(string)
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
				formulacionhelper.ClonarHijos(hijos, padre)
			}

		} else {
			panic(map[string]interface{}{"funcion": "CrearPlan", "err": "Error creando plan", "status": "400", "log": err})
		}

		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update Plan Successful", "Data": planSubgrupo}
		c.ServeJSON()
	} else {
		panic(map[string]interface{}{"funcion": "CrearPlan", "err": "Error consultando datos Plan Formato", "status": "400", "log": err})
	}
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	var resSubgrupo map[string]interface{}
	var subgrupo map[string]interface{}
	//var fuentesActividad []map[string]interface{}
	//var subGrupo map[string]interface{}

	_ = id
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	//idSubDetalleProI := body["idSubDetalle"]
	indexMetaSubProI := body["indexMetaSubPro"]
	//fuentesPro := body["fuentesPro"]
	idDetalleFuentesPro := body["idDetalleFuentesPro"].(string)
	fmt.Println(body["fuentesActividad"], "fuentesActividad")
	fuentesActividad := body["fuentesActividad"]
	ponderacionH := body["ponderacionH"]
	var dato_fuente []map[string]interface{}

	errGet := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+idDetalleFuentesPro, &resSubgrupo)
	if errGet == nil {
		request.LimpiezaRespuestaRefactor(resSubgrupo, &subgrupo)
		//subgrupo = resSubgrupo["Data"].(map[string]interface{})
		fmt.Println(subgrupo, "subgrupo")
		if subgrupo["dato"] != nil {

			dato_str := subgrupo["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato_fuente)
			for key := range dato_fuente {
				//for _, evidencia := range body["evidencia"].([]interface{})
				fuenteActividad := body["fuentesActividad"].([]interface{})
				for key2 := range fuenteActividad {
					fmt.Println(dato_fuente[key]["_id"], key, dato_fuente[key]["nombre"], fuenteActividad[key2].(map[string]interface{})["id"], key2, fuenteActividad[key2].(map[string]interface{})["presupuestoDisponible"], "idsAntes")
					if dato_fuente[key]["_id"] == fuenteActividad[key2].(map[string]interface{})["id"] {
						fmt.Println(dato_fuente[key]["_id"], key, dato_fuente[key]["nombre"], fuenteActividad[key2].(map[string]interface{})["id"], key, fuenteActividad[key2].(map[string]interface{})["presupuestoDisponible"], "ids")
						//aux_dato_fuente := dato_fuente[key]
						fuente := make(map[string]interface{})
						fuente["_id"] = dato_fuente[key]["_id"]
						fuente["activo"] = dato_fuente[key]["activo"]
						fuente["descripcion"] = dato_fuente[key]["descripcion"]
						fuente["fecha_creacion"] = dato_fuente[key]["fecha_creacion"]
						fuente["nombre"] = dato_fuente[key]["nombre"]
						fuente["posicion"] = dato_fuente[key]["posicion"]
						fuente["presupuesto"] = dato_fuente[key]["presupuesto"]
						fuente["presupuestoDisponible"] = dato_fuente[key]["presupuestoDisponible"]
						fuente["presupuestoProyecto"] = dato_fuente[key]["presupuestoProyecto"]
						fuente["presupuestoDisponiblePlanes"] = fuenteActividad[key2].(map[string]interface{})["presupuestoDisponible"]
						dato_fuente[key] = fuente
					}
				}
			}

			fmt.Println(dato_fuente, "dato")
			b, _ := json.Marshal(dato_fuente)
			str := string(b)
			subgrupo["dato"] = str
		}
		var resDetalle map[string]interface{}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+idDetalleFuentesPro, "PUT", &resDetalle, subgrupo); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}

		fmt.Println(resDetalle, "resDetalle")
		//return errDetalle
	}
	//return errGet

	// errFuentes := inversionhelper.ActualizarInfoComplDetalle(idDetalleFuentesPro.(string), fuentesPro.([]interface{}))

	// if errFuentes != nil {
	// 	c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": errFuentes}
	// 	c.Abort("400")
	// }
	// for i := range fuentesPro {

	// }

	for key, element := range entrada {
		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var id_subgrupoDetalle string
		keyStr := strings.Split(key, "_")

		if len(keyStr) > 1 && keyStr[1] == "o" {
			id_subgrupoDetalle = keyStr[0]
			if element != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					//fmt.Println(dato_plan, "dato_plan")
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta

							// aux := make(map[string]interface{})
							// aux["idSubDetalleProI"] = idSubDetalleProI
							// aux["indexMetaSubProI"] = indexMetaSubProI
							// armonizacion_dato[index] = aux
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
				fmt.Println(res, "res 1058")

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]
		if fuentesActividad != nil {
			if subgrupo_detalle["armonizacion_dato"] != nil {
				dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
				json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)
				if armonizacion_dato[index] != nil {
					aux := make(map[string]interface{})
					aux["fuentesActividad"] = fuentesActividad
					aux["indexMetaSubProI"] = indexMetaSubProI
					aux["ponderacionH"] = ponderacionH
					aux["presupuesto_programado"] = body["presupuesto_programado"]
					armonizacion_dato[index] = aux
					fmt.Println(armonizacion_dato, "armonizacion_dato")
				}
				c, _ := json.Marshal(armonizacion_dato)
				strArmonizacion := string(c)
				subgrupo_detalle["armonizacion_dato"] = strArmonizacion

			}
		}

		nuevoDato := true
		meta := make(map[string]interface{})

		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_str := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_str), &dato_plan)

			for index_actividad := range dato_plan {
				if index_actividad == index {
					nuevoDato = false
					aux_actividad := dato_plan[index_actividad].(map[string]interface{})
					meta["index"] = index_actividad
					meta["dato"] = element
					meta["activo"] = aux_actividad["activo"]
					if aux_actividad["observacion"] != nil {
						meta["observacion"] = aux_actividad["observacion"]
					}
					dato_plan[index_actividad] = meta
				}
			}
		}

		if nuevoDato {
			meta["index"] = index
			meta["dato"] = element
			meta["activo"] = true
			dato_plan[index] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		fmt.Println(res, "res 1121")
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": entrada}
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}

	_ = id
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})

	for key, element := range entrada {
		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var id_subgrupoDetalle string
		keyStr := strings.Split(key, "_")

		if len(keyStr) > 1 && keyStr[1] == "o" {
			id_subgrupoDetalle = keyStr[0]
			if element != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					//fmt.Println(dato_plan, "dato_plan")
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta

							// aux := make(map[string]interface{})
							// aux["idSubDetalleProI"] = idSubDetalleProI
							// aux["indexMetaSubProI"] = indexMetaSubProI
							// armonizacion_dato[index] = aux
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
				fmt.Println(res, "res 1058")

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]
		nuevoDato := true
		meta := make(map[string]interface{})
		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_str := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_str), &dato_plan)

			for index_actividad := range dato_plan {
				if index_actividad == index {
					nuevoDato = false
					aux_actividad := dato_plan[index_actividad].(map[string]interface{})
					meta["index"] = index_actividad
					meta["dato"] = element
					meta["activo"] = aux_actividad["activo"]
					if aux_actividad["observacion"] != nil {
						meta["observacion"] = aux_actividad["observacion"]
					}
					dato_plan[index_actividad] = meta
				}
			}
		}

		if nuevoDato {
			meta["index"] = index
			meta["dato"] = element
			meta["activo"] = true
			dato_plan[index] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		fmt.Println(res, "res 1121")
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": entrada}
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()
	id := c.Ctx.Input.Param(":id")
	index := c.Ctx.Input.Param(":index")

	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}

	_ = id
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	//idSubDetalleProI := body["idSubDetalle"]
	//indexMetaSubProI := body["indexMetaSubPro"]
	for key, element := range entrada {
		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var id_subgrupoDetalle string
		keyStr := strings.Split(key, "_")

		if len(keyStr) > 1 && keyStr[1] == "o" {
			id_subgrupoDetalle = keyStr[0]
			if element != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					fmt.Println(dato_plan, "dato_plan")
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta

							// aux := make(map[string]interface{})
							// aux["idSubDetalleProI"] = idSubDetalleProI
							// aux["indexMetaSubProI"] = indexMetaSubProI
							// armonizacion_dato[index] = aux
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
				fmt.Println(res, "res 2002")

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

		subgrupo_detalle = respuestaLimpia[0]
		if subgrupo_detalle["armonizacion_dato"] != nil {
			dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)
			if armonizacion_dato[index] != nil {
				for index_armo := range armonizacion_dato {
					if index_armo == index {
						aux_armonizacion := armonizacion_dato[index_armo].(map[string]interface{})
						aux := make(map[string]interface{})
						aux["idSubDetalleProI"] = aux_armonizacion["idSubDetalleProI"]
						aux["indexMetaSubProI"] = aux_armonizacion["indexMetaSubProI"]
						aux["indexMetaPlan"] = aux_armonizacion["indexMetaPlan"]
						aux["presupuesto_programado"] = body["presupuesto_programado"]
						armonizacion_dato[index] = aux
						fmt.Println(armonizacion_dato, "armonizacion_dato")
					}
				}

			}
			c, _ := json.Marshal(armonizacion_dato)
			strArmonizacion := string(c)
			subgrupo_detalle["armonizacion_dato"] = strArmonizacion

		}

		nuevoDato := true
		meta := make(map[string]interface{})

		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_str := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_str), &dato_plan)

			for index_actividad := range dato_plan {
				if index_actividad == index {
					nuevoDato = false
					aux_actividad := dato_plan[index_actividad].(map[string]interface{})
					meta["index"] = index_actividad
					meta["dato"] = element
					meta["activo"] = aux_actividad["activo"]
					if aux_actividad["observacion"] != nil {
						meta["observacion"] = aux_actividad["observacion"]
					}
					dato_plan[index_actividad] = meta
				}
			}
		}

		if nuevoDato {
			meta["index"] = index
			meta["dato"] = element
			meta["activo"] = true
			dato_plan[index] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		fmt.Println(res, "res 2074")
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": entrada}
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "InversionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	//index := c.Ctx.Input.Param(":indexMeta")
	var res map[string]interface{}
	//var body map[string]interface{}
	var subgrupo map[string]interface{}
	//var id_subgrupoDetalle string
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	//var armonizacionUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	//json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	//armonizacion_data, _ := json.Marshal(body)
	dato := make(map[string]interface{})
	var magnitud int

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuestaLimpia)
		subgrupo = respuestaLimpia[0]
		fmt.Println(res, "subgrupo")
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+subgrupo["_id"].(string), &respuesta); err != nil {
			panic(map[string]interface{}{"funcion": "MagnitudesProgramadas", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["dato"] != nil {

			dato_str := subgrupo_detalle["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)
			// for index_actividad := range dato {
			// 	if index_actividad == index {
			// 		aux_actividad := dato[index_actividad].(map[string]interface{})
			// 		magnitud = aux_actividad
			// 	}
			// }
			magnitud = len(dato)
			fmt.Println(dato, "dato")
		}
		fmt.Println(magnitud)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": magnitud}
	} else {
		panic(map[string]interface{}{"funcion": "MagnitudesProgramadas", "err": "Error consultando subgrupo", "status": "400", "log": err})
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["message"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")

	var respuesta map[string]interface{}
	var res map[string]interface{}
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	var actividadesPadre []map[string]interface{}
	var planPadre map[string]interface{}
	var respuestaPost map[string]interface{}
	var respuestaPostActividad map[string]interface{}
	var planVersionado map[string]interface{}
	var actividadVersionada map[string]interface{}
	plan := make(map[string]interface{})
	grupoActividad := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuesta); err == nil {

		request.LimpiezaRespuestaRefactor(respuesta, &planPadre)

		plan["nombre"] = planPadre["nombre"].(string)
		plan["descripcion"] = planPadre["descripcion"].(string)
		plan["tipo_plan_id"] = planPadre["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planPadre["aplicativo_id"].(string)
		plan["activo"] = planPadre["activo"]
		plan["formato"] = false
		plan["vigencia"] = planPadre["vigencia"].(string)
		plan["dependencia_id"] = planPadre["dependencia_id"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		plan["padre_plan_id"] = id

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuestaPost, plan); err != nil {
			panic(map[string]interface{}{"funcion": "VersionarPlan", "err": "Error versionando plan \"plan[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		planVersionado = respuestaPost["Data"].(map[string]interface{})
		c.Data["json"] = respuestaPost

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
			formulacionhelper.VersionarHijos(hijos, planVersionado["_id"].(string))
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=dependencia_id:"+planPadre["dependencia_id"].(string)+",vigencia:"+planPadre["vigencia"].(string)+",formato:false,arbol_padre_id:"+id, &res); err == nil {
			request.LimpiezaRespuestaRefactor(res, &actividadesPadre)
			if len(actividadesPadre) > 0 {
				for key := range actividadesPadre {
					actividadPadre := actividadesPadre[key]

					grupoActividad["nombre"] = actividadPadre["nombre"].(string)
					grupoActividad["descripcion"] = actividadPadre["descripcion"].(string)
					grupoActividad["tipo_plan_id"] = actividadPadre["tipo_plan_id"].(string)
					grupoActividad["aplicativo_id"] = actividadPadre["aplicativo_id"].(string)
					grupoActividad["activo"] = actividadPadre["activo"]
					grupoActividad["formato"] = false
					grupoActividad["vigencia"] = actividadPadre["vigencia"].(string)
					grupoActividad["dependencia_id"] = actividadPadre["dependencia_id"].(string)
					grupoActividad["estado_plan_id"] = "614d3ad301c7a200482fabfd"
					grupoActividad["padre_plan_id"] = actividadPadre["_id"].(string)
					grupoActividad["arbol_padre_id"] = planVersionado["_id"].(string)
					grupoActividad["documento_id"] = actividadPadre["documento_id"].(string)

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuestaPostActividad, grupoActividad); err != nil {
						panic(map[string]interface{}{"funcion": "VersionarPlan", "err": "Error versionando actividades \"actividadPadre[\"_id\"].(string)\"", "status": "400", "log": err})
					}

					actividadVersionada = respuestaPostActividad["Data"].(map[string]interface{})
					c.Data["json"] = respuestaPost

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+actividadPadre["_id"].(string), &respuestaHijos); err == nil {
						request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
						formulacionhelper.VersionarHijos(hijos, actividadVersionada["_id"].(string))
					}

				}
			}
		}
	}
	c.ServeJSON()
}
