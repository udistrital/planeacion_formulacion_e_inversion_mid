package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_formulacion_mid/helpers"
	"github.com/udistrital/utils_oas/request"
)

// InversionController operations for Inversion
type InversionController struct {
	beego.Controller
}

// URLMapping ...
func (c *InversionController) URLMapping() {
	c.Mapping("AgregarProyecto", c.AgregarProyecto)
	c.Mapping("EditarProyecto", c.EditarProyecto)
	c.Mapping("GuardarDocumentos", c.GuardarDocumentos)
	c.Mapping("ConsultarProyectoId", c.ConsultarProyectoId)
	c.Mapping("ConsultarMetasProyecto", c.ConsultarMetasProyecto)
	c.Mapping("ConsultarTodosProyectos", c.ConsultarTodosProyectos)
	c.Mapping("ActualizarSubgrupoDetalle", c.ActualizarSubgrupoDetalle)
	c.Mapping("ActualizarProyectoGeneral", c.ActualizarProyectoGeneral)
	c.Mapping("CrearPlan", c.CrearPlan)
	c.Mapping("ConsultarPlanIdentificador", c.ConsultarPlanIdentificador)
	c.Mapping("GuardarMeta", c.GuardarMeta)
	c.Mapping("ConsultarPlan", c.ConsultarPlan)
	c.Mapping("ArmonizarInversion", c.ArmonizarInversion)
	c.Mapping("ActualizarMetaPlan", c.ActualizarMetaPlan)
	c.Mapping("ConsultarTodasMetasPlan", c.ConsultarTodasMetasPlan)
	c.Mapping("InactivarMeta", c.InactivarMeta)
	c.Mapping("ProgramarMagnitudesPlan", c.ProgramarMagnitudesPlan)
	c.Mapping("ConsultarMagnitudesProgramadas", c.ConsultarMagnitudesProgramadas)
	c.Mapping("CrearGrupoMeta", c.CrearGrupoMeta)
	c.Mapping("ActualizarActividad", c.ActualizarActividad)
	c.Mapping("ActualizarTablaActividad", c.ActualizarTablaActividad)
	c.Mapping("ActualizarPresupuestoMeta", c.ActualizarPresupuestoMeta)
	c.Mapping("VerificarMagnitudesProgramadas", c.VerificarMagnitudesProgramadas)
	c.Mapping("VersionarPlan", c.VersionarPlan)
}

// AgregarProyecto ...
// @Title AgregarProyecto
// @Description post AgregarProyecto
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @router /proyecto [post]
func (c *InversionController) AgregarProyecto() {
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

	var registroProyecto map[string]interface{}
	var identificadorProyecto string
	var respuestaPlan map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroProyecto); err == nil {
		var respuesta map[string]interface{}
		plan := map[string]interface{}{
			"activo":        true,
			"nombre":        registroProyecto["nombre_proyecto"],
			"descripcion":   registroProyecto["codigo_proyecto"],
			"tipo_plan_id":  "63ca86f1b6c0e5725a977dae",
			"aplicativo_id": " ",
		}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuesta, plan); err == nil {
			respuestaPlan = respuesta["Data"].(map[string]interface{})
			identificadorProyecto = respuestaPlan["_id"].(string)
			soportes := map[string]interface{}{"codigo_proyecto": registroProyecto["codigo_proyecto"], "data": registroProyecto["soportes"]}

			if err := helpers.RegistrarInformacionComplementaria(identificadorProyecto, soportes, "soportes"); err == nil {
				fuentes := map[string]interface{}{"codigo_proyecto": registroProyecto["codigo_proyecto"], "data": registroProyecto["fuentes"]}

				if err := helpers.RegistrarInformacionComplementaria(identificadorProyecto, fuentes, "fuentes apropiacion"); err == nil {
					helpers.ActualizarPresupuestoDisponible(registroProyecto["fuentes"].([]interface{}))

					metas := map[string]interface{}{"codigo_proyecto": registroProyecto["codigo_proyecto"], "data": registroProyecto["metas"]}

					if err := helpers.RegistrarInformacionComplementaria(identificadorProyecto, metas, "metas asociadas al proyecto de inversion"); err == nil {
						c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": respuestaPlan}
					} else {
						c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": err}
						c.Abort("400")
					}

				} else {
					c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": err}
					c.Abort("400")
				}
			} else {
				c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": err}
				c.Abort("400")
			}
		} else {
			panic(map[string]interface{}{"funcion": "AgregarProyecto", "err": "Error Registrando Proyecto", "status": "400", "log": err})
		}
	} else {
		panic(map[string]interface{}{"funcion": "AgregarProyecto", "err": "Error Registrando Proyecto", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// EditarProyecto ...
// @Title EditarProyecto
// @Description post EditarProyecto
// @Param	body		body 	{}	true		"body for Plan content"
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /proyecto/:id [put]
func (c *InversionController) EditarProyecto() {
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

	identificador := c.Ctx.Input.Param(":id")
	var registroProyecto map[string]interface{}
	var respuesta map[string]interface{}
	var informacionProyecto map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &registroProyecto); err == nil {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &informacionProyecto)
			informacionProyecto["nombre"] = registroProyecto["nombre_proyecto"]
			informacionProyecto["descripcion"] = registroProyecto["codigo_proyecto"]
			var respuesta2 map[string]interface{}

			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, "PUT", &respuesta2, informacionProyecto); err == nil {
				if err := helpers.ActualizarInformacionComplementariaDetalle(registroProyecto["id_detalle_soportes"].(string), registroProyecto["soportes"].([]interface{})); err == nil {
					if err := helpers.ActualizarInformacionComplementariaDetalle(registroProyecto["id_detalle_fuentes"].(string), registroProyecto["fuentes"].([]interface{})); err == nil {
						helpers.ActualizarPresupuestoDisponible(registroProyecto["fuentes"].([]interface{}))

						if err := helpers.ActualizarInformacionComplementariaDetalle(registroProyecto["id_detalle_metas"].(string), registroProyecto["metas"].([]interface{})); err == nil {
							c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": informacionProyecto}
						} else {
							c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": err}
							c.Abort("400")
						}
					} else {
						c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": err}
						c.Abort("400")
					}
				} else {
					c.Data["json"] = map[string]interface{}{"Success": false, "Status": "400", "Message": "error", "Data": err}
					c.Abort("400")
				}
			} else {
				panic(map[string]interface{}{"funcion": "EditarProyecto", "err": "Error Editando Proyecto", "status": "400", "log": err})
			}
		} else {
			panic(map[string]interface{}{"funcion": "EditarProyecto", "err": "Error Editando Proyecto", "status": "400", "log": err})
		}
	} else {
		panic(map[string]interface{}{"funcion": "EditarProyecto", "err": "Error Editando Proyecto", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// GuardarDocumentos ...
// @Title GuardarDocumentos
// @Description post GuardarDocumentos
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /guardar_documentos [post]
func (c *InversionController) GuardarDocumentos() {
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

	var body map[string]interface{}
	var evidencias []map[string]interface{}

	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &body); err == nil {
		if body["documento"] != nil {
			respuestaDocumentos := helpers.GuardarDocumento(body["documento"].([]interface{}))

			for _, doc := range respuestaDocumentos {
				evidencias = append(evidencias, map[string]interface{}{
					"Id":     doc.(map[string]interface{})["Id"],
					"Enlace": doc.(map[string]interface{})["Enlace"],
					"nombre": doc.(map[string]interface{})["Nombre"],
					"TipoDocumento": map[string]interface{}{
						"id":                doc.(map[string]interface{})["TipoDocumento"].(map[string]interface{})["Id"],
						"codigoAbreviacion": doc.(map[string]interface{})["TipoDocumento"].(map[string]interface{})["CodigoAbreviacion"],
					},
					"Observacion": "",
					"Activo":      true,
				})
			}
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": evidencias}
		} else {
			c.Data["json"] = map[string]interface{}{"Code": "400", "Body": err, "Type": "error"}
			c.Abort("400")
		}
	} else {
		panic(map[string]interface{}{"funcion": "GuardarDocumentos", "err": "Error Guardando Documentos", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarProyectoId ...
// @Title ConsultarProyectoId
// @Description get ConsultarProyectoId
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /proyecto/:id [get]
func (c *InversionController) ConsultarProyectoId() {
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	consultaProyecto := make(map[string]interface{})
	var informacionProyecto map[string]interface{}
	var subgruposData map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &informacionProyecto)
		consultaProyecto["nombre_proyecto"] = informacionProyecto["nombre"]
		consultaProyecto["codigo_proyecto"] = informacionProyecto["descripcion"]
		consultaProyecto["fecha_creacion"] = informacionProyecto["fecha_creacion"]
		padreIdentificador := informacionProyecto["_id"].(string)
		var informacionSubgrupos []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+padreIdentificador, &subgruposData); err == nil {
			request.LimpiezaRespuestaRefactor(subgruposData, &informacionSubgrupos)

			for i := range informacionSubgrupos {
				var subgrupoDetalle map[string]interface{}
				var detalleSubgrupos []map[string]interface{}

				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=subgrupo_id:"+informacionSubgrupos[i]["_id"].(string), &subgrupoDetalle); err == nil {
					request.LimpiezaRespuestaRefactor(subgrupoDetalle, &detalleSubgrupos)
					armonizacion_dato_str := detalleSubgrupos[0]["dato"].(string)
					var subgrupo_dato []map[string]interface{}
					json.Unmarshal([]byte(armonizacion_dato_str), &subgrupo_dato)

					if strings.Contains(strings.ToLower(informacionSubgrupos[i]["nombre"].(string)), "soporte") {
						consultaProyecto["soportes"] = subgrupo_dato
						consultaProyecto["id_detalle_soportes"] = detalleSubgrupos[0]["_id"]
					}
					if strings.Contains(strings.ToLower(informacionSubgrupos[i]["nombre"].(string)), "metas") {
						consultaProyecto["metas"] = subgrupo_dato
						consultaProyecto["id_detalle_metas"] = detalleSubgrupos[0]["_id"]
					}
					if strings.Contains(strings.ToLower(informacionSubgrupos[i]["nombre"].(string)), "fuentes") {
						consultaProyecto["fuentes"] = subgrupo_dato
						consultaProyecto["id_detalle_fuentes"] = detalleSubgrupos[0]["_id"]
					}
				} else {
					panic(err)
				}
			}
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": consultaProyecto}
		} else {
			panic(map[string]interface{}{"funcion": "ConsultarProyectoId", "err": "Error obteniendo información plan", "status": "400", "log": err})
		}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarProyectoId", "err": "Error obteniendo información plan", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarMetasProyecto ...
// @Title ConsultarMetasProyecto
// @Description get ConsultarMetasProyecto
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /consultar_metas_proyecto/:id [get]
func (c *InversionController) ConsultarMetasProyecto() {
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	consultaProyecto := make(map[string]interface{})
	var informacionProyecto map[string]interface{}
	var subgruposData map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &informacionProyecto)
		var informacionSubgrupos []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+identificador, &subgruposData); err == nil {
			request.LimpiezaRespuestaRefactor(subgruposData, &informacionSubgrupos)

			for i := range informacionSubgrupos {
				var subgrupoDetalle map[string]interface{}
				var detalleSubgrupos []map[string]interface{}

				if strings.Contains(strings.ToLower(informacionSubgrupos[i]["nombre"].(string)), "metas") {
					consultaProyecto["subgrupo_id_metas"] = informacionSubgrupos[i]["_id"]

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=subgrupo_id:"+informacionSubgrupos[i]["_id"].(string), &subgrupoDetalle); err == nil {
						request.LimpiezaRespuestaRefactor(subgrupoDetalle, &detalleSubgrupos)
						var dato_metas []map[string]interface{}
						consultaProyecto["id_detalle_meta"] = detalleSubgrupos[0]["_id"]
						datoMeta_str := detalleSubgrupos[0]["dato"].(string)
						json.Unmarshal([]byte(datoMeta_str), &dato_metas)
						consultaProyecto["metas"] = dato_metas
					}
				}
			}
		} else {
			panic(map[string]interface{}{"funcion": "ConsultarMetasProyecto", "err": "Error obteniendo información Metas Proyecto Inversión", "status": "400", "log": err})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": consultaProyecto}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarMetasProyecto", "err": "Error obteniendo información Metas Proyecto Inversión", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarTodosProyectos ...
// @Title ConsultarTodosProyectos
// @Description get ConsultarTodosProyectos
// @Param	aplicativo_id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :aplicativo_id is empty
// @router /proyectos/:aplicativo_id [get]
func (c *InversionController) ConsultarTodosProyectos() {
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

	tipo_plan_identificador := c.Ctx.Input.Param(":aplicativo_id")
	var respuesta map[string]interface{}
	var consultaProyecto []map[string]interface{}
	var proyecto map[string]interface{}
	var dataProyectos []map[string]interface{}
	proyecto_data := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=activo:true,tipo_plan_id:"+tipo_plan_identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &dataProyectos)

		for i := range dataProyectos {
			if dataProyectos[i]["activo"] == true {
				proyecto_data["id"] = dataProyectos[i]["_id"]
				proyecto = helpers.ConsultarDataProyectos(dataProyectos[i])
			}
			consultaProyecto = append(consultaProyecto, proyecto)
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": consultaProyecto}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarTodosProyectos", "err": "Error obteniendo información plan", "status": "400", "log": err})
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
// @router /actualizar_subgrupo_detalle/:id [put]
func (c *InversionController) ActualizarSubgrupoDetalle() {
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

	var subDetalle map[string]interface{}
	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &subDetalle)

	if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+identificador, "PUT", &respuesta, subDetalle); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update Subgrupo Detalle Successful", "Data": respuesta}
		c.ServeJSON()
	} else {
		panic(map[string]interface{}{"funcion": "ActualizarSubgrupoDetalle", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
	}

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

	var informacionProyecto map[string]interface{}
	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &informacionProyecto)

	if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, "PUT", &respuesta, informacionProyecto); err == nil {
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update Plan Successful", "Data": respuesta}
		c.ServeJSON()
	} else {
		panic(map[string]interface{}{"funcion": "ActualizarProyectoGeneral", "err": "Error actualizando plan \"id\"", "status": "400", "log": err})
	}

}

// CrearPlan ...
// @Title CrearPlan
// @Description post CrearPlan
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403
// @router /crearplan [post]
func (c *InversionController) CrearPlan() {
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

	var respuesta map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}
	var respuestaPost map[string]interface{}
	var planSubgrupo map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &parametros)
	plan := make(map[string]interface{})
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	identificador := parametros["id"].(string)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &planFormato)

		plan["nombre"] = "" + planFormato["nombre"].(string)
		plan["descripcion"] = planFormato["descripcion"].(string)
		plan["tipo_plan_id"] = planFormato["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planFormato["aplicativo_id"].(string)
		plan["activo"] = planFormato["activo"]
		plan["formato"] = false
		plan["vigencia"] = parametros["vigencia"].(string)
		plan["dependencia_id"] = parametros["dependencia_id"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/", "POST", &respuestaPost, plan); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaPost, &planSubgrupo)
			padre := planSubgrupo["_id"].(string)

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuestaHijos); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
				helpers.ClonarHijos(hijos, padre)
			} else {
				panic(map[string]interface{}{"funcion": "CrearPlan", "err": "Error creando plan", "status": "400", "log": err})
			}
		} else {
			panic(map[string]interface{}{"funcion": "CrearPlan", "err": "Error creando plan", "status": "400", "log": err})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update Plan Successful", "Data": planSubgrupo}
	} else {
		panic(map[string]interface{}{"funcion": "CrearPlan", "err": "Error consultando datos Plan Formato", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarPlanIdentificador ...
// @Title ConsultarPlanIdentificador
// @Description get ConsultarPlanIdentificador
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /consultar_plan/:id [get]
func (c *InversionController) ConsultarPlanIdentificador() {
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	consultaProyecto := make(map[string]interface{})
	var informacionProyecto map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &informacionProyecto)

		consultaProyecto["nombre_proyecto"] = informacionProyecto["nombre"]
		consultaProyecto["codigo_proyecto"] = informacionProyecto["descripcion"]
		consultaProyecto["fecha_creacion"] = informacionProyecto["fecha_creacion"]

		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": consultaProyecto}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarProyectoIdentificador", "err": "Error obteniendo información plan", "status": "400", "log": err})
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
// @router /guardar_meta/:id [put]
func (c *InversionController) GuardarMeta() {
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

	identificador := c.Ctx.Input.Param(":id")
	var body map[string]interface{}
	var respuesta map[string]interface{}
	var entrada map[string]interface{}
	var respuestaPlan map[string]interface{}
	var plan map[string]interface{}
	var respuestaGuardado map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	identificadorSubDetalleProI := body["idSubDetalle"]
	indiceMetaSubProI := body["indexMetaSubPro"]
	maximoIndice := helpers.ConsultarIndexActividad(entrada)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuestaPlan); err != nil {
		panic(map[string]interface{}{"funcion": "GuardarMeta", "err": "Error get Plan \"id\"", "status": "400", "log": err})
	}

	request.LimpiezaRespuestaRefactor(respuestaPlan, &plan)
	if plan["estado_plan_id"] != "614d3ad301c7a200482fabfd" {
		var respuesta map[string]interface{}
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, "PUT", &respuesta, plan); err != nil {
			panic(map[string]interface{}{"funcion": "GuardarMeta", "err": "Error actualizacion estado \"id\"", "status": "400", "log": err})
		}
	}

	for llave, elemento := range entrada {
		var respuestaJ map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})

		if elemento != "" {
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+llave, &respuestaJ); err != nil {
				panic(map[string]interface{}{"funcion": "GuardarMeta", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
			}

			request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			actividad := make(map[string]interface{})

			if subgrupo_detalle["dato_plan"] == nil {
				actividad["index"] = 1
				actividad["dato"] = elemento
				actividad["activo"] = true
				i := strconv.Itoa(actividad["index"].(int))
				dato_plan[i] = actividad

				b, _ := json.Marshal(dato_plan)
				str := string(b)
				subgrupo_detalle["dato_plan"] = str

				armonizacion_dato := make(map[string]interface{})
				auxiliar := make(map[string]interface{})
				auxiliar["idSubDetalleProI"] = identificadorSubDetalleProI
				auxiliar["indexMetaSubProI"] = indiceMetaSubProI
				auxiliar["indexMetaPlan"] = 1
				armonizacion_dato[i] = auxiliar
				c, _ := json.Marshal(armonizacion_dato)
				strArmonizacion := string(c)
				subgrupo_detalle["armonizacion_dato"] = strArmonizacion
			} else {
				dato_plan_str := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_str), &dato_plan)

				actividad["index"] = maximoIndice + 1
				actividad["dato"] = elemento
				actividad["activo"] = true
				i := strconv.Itoa(actividad["index"].(int))
				dato_plan[i] = actividad
				b, _ := json.Marshal(dato_plan)
				str := string(b)
				subgrupo_detalle["dato_plan"] = str
				armonizacion_dato := make(map[string]interface{})

				if subgrupo_detalle["armonizacion_dato"] != nil {
					armonizacion_dato_str := subgrupo_detalle["armonizacion_dato"].(string)
					json.Unmarshal([]byte(armonizacion_dato_str), &armonizacion_dato)
					auxiliar := make(map[string]interface{})
					auxiliar["idSubDetalleProI"] = identificadorSubDetalleProI
					auxiliar["indexMetaSubProI"] = indiceMetaSubProI
					auxiliar["indexMetaPlan"] = i
					armonizacion_dato[i] = auxiliar
					c, _ := json.Marshal(armonizacion_dato)
					strArmonizacion := string(c)
					subgrupo_detalle["armonizacion_dato"] = strArmonizacion
				}
			}
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
				panic(map[string]interface{}{"funcion": "GuardarMeta", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
			}
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaGuardado)
		}
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": respuestaGuardado}
	c.ServeJSON()
}

// ConsultarPlan ...
// @Title ConsultarPlan
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /consultar_informacion_plan/:id [get]
func (c *InversionController) ConsultarPlan() {
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

	identificador := c.Ctx.Input.Param(":id")
	var subgrupo []map[string]interface{}
	var respuesta map[string]interface{}
	var identificador_subgrupoDetalle string
	var respuestaJ map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	armo_dato := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Armonizar,activo:true,padre:"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)
		identificador_subgrupoDetalle = subgrupo[0]["_id"].(string)

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+identificador_subgrupoDetalle, &respuestaJ); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			armonizacion_dato_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(armonizacion_dato_str), &armo_dato)
		} else {
			panic(map[string]interface{}{"funcion": "ConsultarPlan", "err": "Error get subgrupo-detalle", "status": "400", "log": err})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": armo_dato}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarPlan", "err": "Error get subgrupo", "status": "400", "log": err})
	}
	c.ServeJSON()
}

func (c *InversionController) ArmonizarInversion() {
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	var body map[string]interface{}
	var subgrupo []map[string]interface{}
	var identificador_subgrupoDetalle string
	var respuestaJ map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var armonizacionUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	armonizacion_data, _ := json.Marshal(body)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Armonizar,activo:true,padre:"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)
		subgrupoPost := make(map[string]interface{})
		subDetallePost := make(map[string]interface{})

		if len(subgrupo) != 0 {
			identificador_subgrupoDetalle = subgrupo[0]["_id"].(string)

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+identificador_subgrupoDetalle, &respuestaJ); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)

				if len(respuestaLimpia) > 0 {
					subgrupo_detalle = respuestaLimpia[0]

					subDetallePost["subgrupo_id"] = identificador_subgrupoDetalle
					subDetallePost["fecha_creacion"] = subgrupo_detalle["fecha_creacion"]
					subDetallePost["nombre"] = "Detalle Información Armonización"
					subDetallePost["descripcion"] = "Armonizar"
					subDetallePost["dato"] = " "
					subDetallePost["activo"] = true
					subDetallePost["armonizacion_dato"] = string(armonizacion_data)

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(respuesta, &armonizacionUpdate)
					} else {
						panic(map[string]interface{}{"funcion": "GuardarPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
					}
				} else {
					subDetallePost["subgrupo_id"] = identificador_subgrupoDetalle
					subDetallePost["nombre"] = "Detalle Información Armonización"
					subDetallePost["descripcion"] = "Armonizar"
					subDetallePost["dato"] = " "
					subDetallePost["activo"] = true
					subDetallePost["armonizacion_dato"] = string(armonizacion_data)

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &respuesta, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(respuesta, &armonizacionUpdate)
					} else {
						panic(map[string]interface{}{"funcion": "GuardarPlan", "err": "Error registrando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
					}
				}
			} else {
				panic(map[string]interface{}{"funcion": "ArmonizarInversion", "err": "Error get subgrupo-detalle", "status": "400", "log": err})
			}
		} else {
			subgrupoPost["nombre"] = "Armonización Plan Inversión"
			subgrupoPost["descripcion"] = "Armonizar"
			subgrupoPost["padre"] = identificador
			subgrupoPost["activo"] = true

			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/", "POST", &respuestaJ, subgrupoPost); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				subDetallePost["subgrupo_id"] = subgrupo_detalle["_id"]
				subDetallePost["nombre"] = "Detalle Información Armonización"
				subDetallePost["descripcion"] = "Armonizar"
				subDetallePost["dato"] = " "
				subDetallePost["activo"] = true
				subDetallePost["armonizacion_dato"] = string(armonizacion_data)

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &respuesta, subDetallePost); err == nil {
					request.LimpiezaRespuestaRefactor(respuesta, &armonizacionUpdate)
				} else {
					panic(map[string]interface{}{"funcion": "ArmonizarInversion", "err": "Error registrando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			} else {
				panic(map[string]interface{}{"funcion": "ArmonizarInversion", "err": "Error registrando subgrupo", "status": "400", "log": err})
			}
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": respuesta}
	} else {
		panic(map[string]interface{}{"funcion": "ArmonizarInversion", "err": "Error", "status": "400", "log": err})
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
// @router /actualizar_meta/:id/:index [put]
func (c *InversionController) ActualizarMetaPlan() {
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

	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	var respuesta map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	_ = identificador
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	indiceMetaSubProI := body["indexMetaSubPro"]

	for llave, elemento := range entrada {
		var respuestaJ map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var identificador_subgrupoDetalle string
		llaveStr := strings.Split(llave, "_")

		if len(llaveStr) > 1 && llaveStr[1] == "o" {
			identificador_subgrupoDetalle = llaveStr[0]

			if elemento != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]

				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)

					for indice_actividad := range dato_plan {
						if indice_actividad == indice {
							auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
							meta["index"] = indice_actividad
							meta["dato"] = auxiliar_actividad["dato"]
							meta["activo"] = auxiliar_actividad["activo"]
							meta["observacion"] = elemento
							dato_plan[indice_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}
				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			}
			continue
		}
		identificador_subgrupoDetalle = llave

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["armonizacion_dato"] != nil {
			dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)

			if armonizacion_dato[indice] != nil {
				auxiliar_armonizacion := armonizacion_dato[indice].(map[string]interface{})
				auxiliar := make(map[string]interface{})
				auxiliar["idSubDetalleProI"] = auxiliar_armonizacion["idSubDetalleProI"]
				auxiliar["indexMetaSubProI"] = indiceMetaSubProI
				auxiliar["presupuesto_programado"] = auxiliar_armonizacion["presupuesto_programado"]
				armonizacion_dato[indice] = auxiliar
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

			for indice_actividad := range dato_plan {
				if indice_actividad == indice {
					nuevoDato = false
					auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
					meta["index"] = indice_actividad
					meta["dato"] = elemento
					meta["activo"] = auxiliar_actividad["activo"]
					if auxiliar_actividad["observacion"] != nil {
						meta["observacion"] = auxiliar_actividad["observacion"]
					}
					dato_plan[indice_actividad] = meta
				}
			}
		}

		if nuevoDato {
			meta["index"] = indice
			meta["dato"] = elemento
			meta["activo"] = true
			dato_plan[indice] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarMetaPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": entrada}
	c.ServeJSON()
}

// ConsultarTodasMetasPlan ...
// @Title ConsultarTodasMetasPlan
// @Description put Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /consulta_todas_metas/:id [get]
func (c *InversionController) ConsultarTodasMetasPlan() {
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	var hijos []map[string]interface{}
	var tabla map[string]interface{}
	var metas []map[string]interface{}
	var auxiliarHijos []interface{}
	var data_fuente []map[string]interface{}

	helpers.LimpiaInversion()

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)

		for i := 0; i < len(hijos); i++ {
			auxiliarHijos = append(auxiliarHijos, hijos[i]["_id"])
		}

		tabla = helpers.ConsultarTabla(auxiliarHijos)
		metas = tabla["data_source"].([]map[string]interface{})

		for indiceMeta := range metas {
			if metas[indiceMeta]["activo"] == true {
				data_fuente = append(data_fuente, metas[indiceMeta])
			}
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": data_fuente}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarTodasMetasPlan", "err": "Error al consultar metas del plan \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
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
// @router /inactivar_meta/:id/:index [put]
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

	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	var respuesta map[string]interface{}
	var hijos []map[string]interface{}
	var tabla map[string]interface{}
	var auxiliarHijos []interface{}
	helpers.LimpiaInversion()

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)

		for i := 0; i < len(hijos); i++ {
			auxiliarHijos = append(auxiliarHijos, hijos[i]["_id"])
		}

		helpers.ConsultarHijosInversion(auxiliarHijos, indice)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": tabla}
	} else {
		panic(map[string]interface{}{"funcion": "InactivarMeta", "err": "Error al consultar metas del plan \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ProgramarMagnitudesPlan ...
// @Title ProgramarMagnitudesPlan
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403 :id is empty
// @router /magnitudes/:id/:index [put]
func (c *InversionController) ProgramarMagnitudesPlan() {
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

	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	var respuesta map[string]interface{}
	var body map[string]interface{}
	var subgrupo []map[string]interface{}
	var identificador_subgrupoDetalle string
	var respuestaJ map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var magnitudesUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)
		subgrupoPost := make(map[string]interface{})
		subDetallePost := make(map[string]interface{})

		if len(subgrupo) > 0 {
			identificador_subgrupoDetalle = subgrupo[0]["_id"].(string)

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+identificador_subgrupoDetalle, &respuestaJ); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)

				if len(respuestaLimpia) > 0 {
					subgrupo_detalle = respuestaLimpia[0]
					magnitud := make(map[string]interface{})

					if subgrupo_detalle["dato"] == nil {
						magnitud["index"] = indice
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
						magnitud["index"] = indice
						magnitud["dato"] = body
						magnitud["activo"] = true
						dato[indice] = magnitud
						b, _ := json.Marshal(dato)
						str := string(b)
						subgrupo_detalle["dato"] = str
					}

					subDetallePost["dato"] = subgrupo_detalle["dato"]
					subDetallePost["subgrupo_id"] = identificador_subgrupoDetalle
					subDetallePost["fecha_creacion"] = subgrupo_detalle["fecha_creacion"]
					subDetallePost["nombre"] = "Detalle Información Programación de Magnitudes y Presupuesto"
					subDetallePost["descripcion"] = "Magnitudes"
					subDetallePost["activo"] = true

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(respuesta, &magnitudesUpdate)
					} else {
						panic(map[string]interface{}{"funcion": "ProgramarMagnitudesPlan", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
					}
				} else {
					subgrupo_detalle := make(map[string]interface{})
					magnitud := make(map[string]interface{})
					magnitud["index"] = indice
					magnitud["dato"] = body
					magnitud["activo"] = true
					dato[indice] = magnitud
					b, _ := json.Marshal(dato)
					str := string(b)
					subgrupo_detalle["dato"] = str
					subDetallePost["subgrupo_id"] = identificador_subgrupoDetalle
					subDetallePost["nombre"] = "Detalle Información Programación de Magnitudes y Presupuesto"
					subDetallePost["descripcion"] = "Magnitudes"
					subDetallePost["dato"] = subgrupo_detalle["dato"]
					subDetallePost["activo"] = true

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &respuesta, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(respuesta, &magnitudesUpdate)
					} else {
						panic(map[string]interface{}{"funcion": "ProgramarMagnitudesPlan", "err": "Error registrando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
					}
				}
			} else {
				panic(map[string]interface{}{"funcion": "ProgramarMagnitudesPlan", "err": "Error get subgrupo-detalle", "status": "400", "log": err})
			}
		} else {
			subgrupoPost["nombre"] = "Programación Magnitudes y Prespuesto Plan de Inversión"
			subgrupoPost["descripcion"] = "Magnitudes"
			subgrupoPost["padre"] = identificador
			subgrupoPost["activo"] = true

			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/", "POST", &respuestaJ, subgrupoPost); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]
				magnitud := make(map[string]interface{})
				magnitud["index"] = indice
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

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &respuesta, subDetallePost); err == nil {
					request.LimpiezaRespuestaRefactor(respuesta, &magnitudesUpdate)
				} else {
					panic(map[string]interface{}{"funcion": "ProgramarMagnitudesPlan", "err": "Error registrando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			} else {
				panic(map[string]interface{}{"funcion": "ProgramarMagnitudesPlan", "err": "Error registrando subgrupo", "status": "400", "log": err})
			}
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": respuesta}
	} else {
		panic(map[string]interface{}{"funcion": "ProgramarMagnitudesPlan", "err": "Error", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarMagnitudesProgramadas ...
// @Title ConsultarMagnitudesProgramadas
// @Description get ConsultarMagnitudesProgramadas
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /magnitudes/:id/:indexMeta [get]
func (c *InversionController) ConsultarMagnitudesProgramadas() {
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

	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":indexMeta")
	var respuesta map[string]interface{}
	var subgrupo map[string]interface{}
	var respuestaJ map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})
	var magnitud map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo = respuestaLimpia[0]

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+subgrupo["_id"].(string), &respuestaJ); err != nil {
			panic(map[string]interface{}{"funcion": "ConsultarMagnitudesProgramadas", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["dato"] != nil {
			dato_str := subgrupo_detalle["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)

			for indice_actividad := range dato {
				if indice_actividad == indice {
					auxiliar_actividad := dato[indice_actividad].(map[string]interface{})
					magnitud = auxiliar_actividad
				}
			}
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": magnitud}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarMagnitudesProgramadas", "err": "Error consultando subgrupo", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// CrearGrupoMeta ...
// @Title CrearGrupoMeta
// @Description post CrearGrupoMeta
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200
// @Failure 403
// @router /crear_grupo_meta [post]
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

	var respuesta map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}
	var respuestaPost map[string]interface{}
	var planSubgrupo map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &parametros)
	plan := make(map[string]interface{})
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	identificador := parametros["id"].(string)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuesta); err == nil {
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

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuestaHijos); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
				helpers.ClonarHijos(hijos, padre)
			} else {
				panic(err)
			}
		} else {
			panic(map[string]interface{}{"funcion": "CrearGrupoMeta", "err": "Error", "status": "400", "log": err})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Update Plan Successful", "Data": planSubgrupo}
	} else {
		panic(map[string]interface{}{"funcion": "CrearGrupoMeta", "err": "Error", "status": "400", "log": err})
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
// @router /actualizar_actividad/:id/:index [put]
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

	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	var respuesta map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	var respuestaSubgrupo map[string]interface{}
	var subgrupo map[string]interface{}
	_ = identificador
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	indiceMetaSubProI := body["indexMetaSubPro"]
	identificadorDetalleFuentesPro := body["idDetalleFuentesPro"].(string)
	fuentesActividad := body["fuentesActividad"]
	ponderacionH := body["ponderacionH"]
	var dato_fuente []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+identificadorDetalleFuentesPro, &respuestaSubgrupo); err == nil {
		request.LimpiezaRespuestaRefactor(respuestaSubgrupo, &subgrupo)

		if subgrupo["dato"] != nil {
			dato_str := subgrupo["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato_fuente)

			for llave := range dato_fuente {
				fuenteActividad := body["fuentesActividad"].([]interface{})

				for llave2 := range fuenteActividad {
					if dato_fuente[llave]["_id"] == fuenteActividad[llave2].(map[string]interface{})["id"] {
						fuente := make(map[string]interface{})
						fuente["_id"] = dato_fuente[llave]["_id"]
						fuente["activo"] = dato_fuente[llave]["activo"]
						fuente["descripcion"] = dato_fuente[llave]["descripcion"]
						fuente["fecha_creacion"] = dato_fuente[llave]["fecha_creacion"]
						fuente["nombre"] = dato_fuente[llave]["nombre"]
						fuente["posicion"] = dato_fuente[llave]["posicion"]
						fuente["presupuesto"] = dato_fuente[llave]["presupuesto"]
						fuente["presupuestoDisponible"] = dato_fuente[llave]["presupuestoDisponible"]
						fuente["presupuestoProyecto"] = dato_fuente[llave]["presupuestoProyecto"]
						fuente["presupuestoDisponiblePlanes"] = fuenteActividad[llave2].(map[string]interface{})["presupuestoDisponible"]
						dato_fuente[llave] = fuente
					}
				}
			}
			b, _ := json.Marshal(dato_fuente)
			str := string(b)
			subgrupo["dato"] = str
		}
		var respuestaDetalle map[string]interface{}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+identificadorDetalleFuentesPro, "PUT", &respuestaDetalle, subgrupo); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error put subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
	} else {
		panic(err)
	}

	for llave, elemento := range entrada {
		var respuestaJ map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var identificador_subgrupoDetalle string
		llaveStr := strings.Split(llave, "_")

		if len(llaveStr) > 1 && llaveStr[1] == "o" {
			identificador_subgrupoDetalle = llaveStr[0]

			if elemento != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]

				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)

					for indice_actividad := range dato_plan {
						if indice_actividad == indice {
							auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
							meta["index"] = indice_actividad
							meta["dato"] = auxiliar_actividad["dato"]
							meta["activo"] = auxiliar_actividad["activo"]
							meta["observacion"] = elemento
							dato_plan[indice_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}
				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			}
			continue
		}
		identificador_subgrupoDetalle = llave

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if fuentesActividad != nil {
			if subgrupo_detalle["armonizacion_dato"] != nil {
				dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
				json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)

				if armonizacion_dato[indice] != nil {
					auxiliar := make(map[string]interface{})
					auxiliar["fuentesActividad"] = fuentesActividad
					auxiliar["indexMetaSubProI"] = indiceMetaSubProI
					auxiliar["ponderacionH"] = ponderacionH
					auxiliar["presupuesto_programado"] = body["presupuesto_programado"]
					armonizacion_dato[indice] = auxiliar
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

			for indice_actividad := range dato_plan {
				if indice_actividad == indice {
					nuevoDato = false
					auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
					meta["index"] = indice_actividad
					meta["dato"] = elemento
					meta["activo"] = auxiliar_actividad["activo"]
					if auxiliar_actividad["observacion"] != nil {
						meta["observacion"] = auxiliar_actividad["observacion"]
					}
					dato_plan[indice_actividad] = meta
				}
			}
		}

		if nuevoDato {
			meta["index"] = indice
			meta["dato"] = elemento
			meta["activo"] = true
			dato_plan[indice] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
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
// @router /actualizar_tabla_actividad/:id/:index [put]
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

	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	var respuesta map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	_ = identificador
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})

	for llave, elemento := range entrada {
		var respuestaJ map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var identificador_subgrupoDetalle string
		llaveStr := strings.Split(llave, "_")

		if len(llaveStr) > 1 && llaveStr[1] == "o" {
			identificador_subgrupoDetalle = llaveStr[0]

			if elemento != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarTablaActividad", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]

				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)

					for indice_actividad := range dato_plan {
						if indice_actividad == indice {
							auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
							meta["index"] = indice_actividad
							meta["dato"] = auxiliar_actividad["dato"]
							meta["activo"] = auxiliar_actividad["activo"]
							meta["observacion"] = elemento
							dato_plan[indice_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}
				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarTablaActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			}
			continue
		}
		identificador_subgrupoDetalle = llave
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarTablaActividad", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]
		nuevoDato := true
		meta := make(map[string]interface{})

		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_str := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_str), &dato_plan)

			for indice_actividad := range dato_plan {
				if indice_actividad == indice {
					nuevoDato = false
					auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
					meta["index"] = indice_actividad
					meta["dato"] = elemento
					meta["activo"] = auxiliar_actividad["activo"]
					if auxiliar_actividad["observacion"] != nil {
						meta["observacion"] = auxiliar_actividad["observacion"]
					}
					dato_plan[indice_actividad] = meta
				}
			}
		}
		if nuevoDato {
			meta["index"] = indice
			meta["dato"] = elemento
			meta["activo"] = true
			dato_plan[indice] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarTablaActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
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
// @router /actualizar_presupuesto_meta/:id/:index [put]
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
	identificador := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	var respuesta map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	_ = identificador
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})

	for llave, elemento := range entrada {
		var respuestaJ map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var identificador_subgrupoDetalle string
		llaveStr := strings.Split(llave, "_")

		if len(llaveStr) > 1 && llaveStr[1] == "o" {
			identificador_subgrupoDetalle = llaveStr[0]

			if elemento != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarPresupuestoMeta", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]

				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)

					for indice_actividad := range dato_plan {
						if indice_actividad == indice {
							auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
							meta["index"] = indice_actividad
							meta["dato"] = auxiliar_actividad["dato"]
							meta["activo"] = auxiliar_actividad["activo"]
							meta["observacion"] = elemento
							dato_plan[indice_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}
				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarPresupuestoMeta", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			}
			continue
		}
		identificador_subgrupoDetalle = llave
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuestaJ); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarPresupuestoMeta", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["armonizacion_dato"] != nil {
			dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)

			if armonizacion_dato[indice] != nil {
				for indice_armonizacion := range armonizacion_dato {
					if indice_armonizacion == indice {
						auxiliar_armonizacion := armonizacion_dato[indice_armonizacion].(map[string]interface{})
						auxiliar := make(map[string]interface{})
						auxiliar["idSubDetalleProI"] = auxiliar_armonizacion["idSubDetalleProI"]
						auxiliar["indexMetaSubProI"] = auxiliar_armonizacion["indexMetaSubProI"]
						auxiliar["indexMetaPlan"] = auxiliar_armonizacion["indexMetaPlan"]
						auxiliar["presupuesto_programado"] = body["presupuesto_programado"]
						armonizacion_dato[indice] = auxiliar
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

			for indice_actividad := range dato_plan {
				if indice_actividad == indice {
					nuevoDato = false
					auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
					meta["index"] = indice_actividad
					meta["dato"] = elemento
					meta["activo"] = auxiliar_actividad["activo"]
					if auxiliar_actividad["observacion"] != nil {
						meta["observacion"] = auxiliar_actividad["observacion"]
					}
					dato_plan[indice_actividad] = meta
				}
			}
		}
		if nuevoDato {
			meta["index"] = indice
			meta["dato"] = elemento
			meta["activo"] = true
			dato_plan[indice] = meta
		}

		b, _ := json.Marshal(dato_plan)
		str := string(b)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarPresupuestoMeta", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
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
// @router /verificar_magnitudes/:id [get]
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	var subgrupo map[string]interface{}
	var respuestaJ map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})
	var magnitud int

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo = respuestaLimpia[0]

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+subgrupo["_id"].(string), &respuestaJ); err != nil {
			panic(map[string]interface{}{"funcion": "VerificarMagnitudesProgramadas", "err": "Error \"key\"", "status": "400", "log": err})
		}

		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["dato"] != nil {
			dato_str := subgrupo_detalle["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)
			magnitud = len(dato)
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": magnitud}
	} else {
		panic(map[string]interface{}{"funcion": "VerificarMagnitudesProgramadas", "err": "Error consultando subgrupo", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// VersionarPlan ...
// @Title VersionarPlan
// @Description post Inversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200
// @Failure 403 :id is empty
// @router /versionar_plan/:id [post]
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

	identificador := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	var respuestaJ map[string]interface{}
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

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador, &respuestaJ); err == nil {
		request.LimpiezaRespuestaRefactor(respuestaJ, &planPadre)
		plan["nombre"] = planPadre["nombre"].(string)
		plan["descripcion"] = planPadre["descripcion"].(string)
		plan["tipo_plan_id"] = planPadre["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planPadre["aplicativo_id"].(string)
		plan["activo"] = planPadre["activo"]
		plan["formato"] = false
		plan["vigencia"] = planPadre["vigencia"].(string)
		plan["dependencia_id"] = planPadre["dependencia_id"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		plan["padre_plan_id"] = identificador

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuestaPost, plan); err != nil {
			panic(map[string]interface{}{"funcion": "VersionarPlan", "err": "Error versionando plan \"plan[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		planVersionado = respuestaPost["Data"].(map[string]interface{})
		c.Data["json"] = respuestaPost

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
			helpers.VersionarHijos(hijos, planVersionado["_id"].(string))
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=dependencia_id:"+planPadre["dependencia_id"].(string)+",vigencia:"+planPadre["vigencia"].(string)+",formato:false,arbol_padre_id:"+identificador, &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &actividadesPadre)

			if len(actividadesPadre) > 0 {
				for llave := range actividadesPadre {
					actividadPadre := actividadesPadre[llave]
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
						helpers.VersionarHijos(hijos, actividadVersionada["_id"].(string))
					}
				}
			}
		}
	}
	c.ServeJSON()
}
