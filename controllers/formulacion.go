package controllers

import (
	"encoding/json"
	"math"
	"net/url"
	"sort"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_formulacion_mid/helpers"
	"github.com/udistrital/planeacion_formulacion_mid/models"
	"github.com/udistrital/utils_oas/request"
)

// FormulacionController operations for Formulacion
type FormulacionController struct {
	beego.Controller
}

// URLMapping ...
func (c *FormulacionController) URLMapping() {
	c.Mapping("ClonarFormato", c.ClonarFormato)
	c.Mapping("GuardarActividad", c.GuardarActividad)
	c.Mapping("ConsultarPlan", c.ConsultarPlan)
	c.Mapping("ActualizarActividad", c.ActualizarActividad)
	c.Mapping("DesactivarActividad", c.DesactivarActividad)
	c.Mapping("ConsultarTodasActividades", c.ConsultarTodasActividades)
	c.Mapping("ConsultarArbolArmonizacion", c.ConsultarArbolArmonizacion)
	c.Mapping("GuardarIdentificacion", c.GuardarIdentificacion)
	c.Mapping("ConsultarIdentificaciones", c.ConsultarIdentificaciones)
	c.Mapping("VersionarPlan", c.VersionarPlan)
	c.Mapping("DesactivarIdentificacion", c.DesactivarIdentificacion)
	c.Mapping("ConsultarPlanVersiones", c.ConsultarPlanVersiones)
	c.Mapping("PonderacionActividades", c.PonderacionActividades)
	c.Mapping("ConsultarRubros", c.ConsultarRubros)
	c.Mapping("ConsultarUnidades", c.ConsultarUnidades)
	c.Mapping("VinculacionTercero", c.VinculacionTercero)
	c.Mapping("Planes", c.Planes)
	c.Mapping("VerificarIdentificaciones", c.VerificarIdentificaciones)
}

// ClonarFormato ...
// @Title ClonarFormato
// @Description post Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /clonar_formato/:id [post]
func (c *FormulacionController) ClonarFormato() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	plan := make(map[string]interface{})
	var respuesta map[string]interface{}
	var respuestaHijos map[string]interface{}
	var hijos []map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &planFormato)
		json.Unmarshal(c.Ctx.Input.RequestBody, &parametros)
		plan["nombre"] = "" + planFormato["nombre"].(string)
		plan["descripcion"] = planFormato["descripcion"].(string)
		plan["tipo_plan_id"] = planFormato["tipo_plan_id"].(string)
		plan["aplicativo_id"] = planFormato["aplicativo_id"].(string)
		plan["activo"] = planFormato["activo"]
		plan["formato"] = false
		plan["vigencia"] = parametros["vigencia"].(string)
		plan["dependencia_id"] = parametros["dependencia_id"].(string)
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"

		var respuestaPost map[string]interface{}
		var respuestaLimpia map[string]interface{}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/", "POST", &respuestaPost, &plan); err == nil {
			respuestaLimpia = respuestaPost["Data"].(map[string]interface{})
			padre := respuestaLimpia["_id"].(string)
			c.Data["json"] = respuestaPost
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
				helpers.ClonarHijos(hijos, padre)
			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
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
// @router /guardar_actividad/:id [put]
func (c *FormulacionController) GuardarActividad() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	id := c.Ctx.Input.Param(":id")
	var body map[string]interface{}
	var respuesta map[string]interface{}
	var entrada map[string]interface{}
	var respuestaPlan map[string]interface{}
	var plan map[string]interface{}
	var armonizacionEjecutada bool = false

	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	armonizacion := body["armo"]
	armonizacionPI := body["armoPI"]
	maximoIndice := helpers.ConsultarIndexActividad(entrada)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &respuestaPlan); err != nil {
		panic(map[string]interface{}{"funcion": "GuardarActividad", "err": "Error get Plan \"id\"", "status": "400", "log": err})
	}
	request.LimpiezaRespuestaRefactor(respuestaPlan, &plan)
	if plan["estado_plan_id"] != "614d3ad301c7a200482fabfd" {
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, "PUT", &respuesta, plan); err != nil {
			panic(map[string]interface{}{"funcion": "GuardarActividad", "err": "Error actualizacion estado \"id\"", "status": "400", "log": err})
		}
	}

	for llave, elemento := range entrada {
		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})

		if elemento != "" {
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+llave, &respuesta); err != nil {
				panic(map[string]interface{}{"funcion": "GuardarActividad", "err": "Error get subgrupo-detalle \"llave\"", "status": "400", "log": err})
			}
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			actividad := make(map[string]interface{})

			if subgrupo_detalle["dato_plan"] == nil {
				actividad["index"] = 1
				actividad["dato"] = elemento
				actividad["activo"] = true
				numeroIndice := strconv.Itoa(actividad["index"].(int))
				dato_plan[numeroIndice] = actividad
				dato2, _ := json.Marshal(dato_plan)
				string_dato2 := string(dato2)
				subgrupo_detalle["dato_plan"] = string_dato2

				if !armonizacionEjecutada {
					armonizacion_dato := make(map[string]interface{})
					variable_auxiliar := make(map[string]interface{})
					variable_auxiliar["armonizacionPED"] = armonizacion
					variable_auxiliar["armonizacionPI"] = armonizacionPI
					armonizacion_dato[numeroIndice] = variable_auxiliar
					dato, _ := json.Marshal(armonizacion_dato)
					stringArmonizacion := string(dato)
					subgrupo_detalle["armonizacion_dato"] = stringArmonizacion
					armonizacionEjecutada = true
				}
			} else {
				dato_plan_string := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_string), &dato_plan)
				actividad["index"] = maximoIndice + 1
				actividad["dato"] = elemento
				actividad["activo"] = true
				numeroIndice := strconv.Itoa(actividad["index"].(int))
				dato_plan[numeroIndice] = actividad
				dato2, _ := json.Marshal(dato_plan)
				string_dato2 := string(dato2)
				subgrupo_detalle["dato_plan"] = string_dato2

				if !armonizacionEjecutada {
					armonizacion_dato := make(map[string]interface{})

					if subgrupo_detalle["armonizacion_dato"] != nil {
						armonizacion_dato_string := subgrupo_detalle["armonizacion_dato"].(string)
						json.Unmarshal([]byte(armonizacion_dato_string), &armonizacion_dato)
					}
					variable_auxiliar := make(map[string]interface{})
					variable_auxiliar["armonizacionPED"] = armonizacion
					variable_auxiliar["armonizacionPI"] = armonizacionPI
					armonizacion_dato[numeroIndice] = variable_auxiliar
					dato, _ := json.Marshal(armonizacion_dato)
					stringArmonizacion := string(dato)
					subgrupo_detalle["armonizacion_dato"] = stringArmonizacion
					armonizacionEjecutada = true
				}
			}
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
				panic(map[string]interface{}{"funcion": "GuardarActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
			}
		}
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": entrada}
	c.ServeJSON()
}

// GetPlan ...
// @Title GetPlan
// @Description get Plan by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_plan/:id/:index [get]
func (c *FormulacionController) ConsultarPlan() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
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

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)
		helpers.Limpia()
		arbol := helpers.ConstruirArbolFormulacion(hijos, indice)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": arbol}
	} else {
		panic(err)
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
// @router /actualizar_actividad/:id/:index [put]
func (c *FormulacionController) ActualizarActividad() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
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
	var respuesta1 map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	_ = identificador
	json.Unmarshal(c.Ctx.Input.RequestBody, &body)
	entrada = body["entrada"].(map[string]interface{})
	armonizacion := body["armo"]
	armonizacionPI := body["armoPI"]

	for llave, elemento := range entrada {
		var respuesta2 map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})
		var armonizacion_dato map[string]interface{}
		var identificador_subgrupoDetalle string
		llaveString := strings.Split(llave, "_")

		if len(llaveString) > 1 && llaveString[1] == "o" {
			identificador_subgrupoDetalle = llaveString[0]

			if elemento != "" {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuesta2); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
				}
				request.LimpiezaRespuestaRefactor(respuesta2, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]

				if subgrupo_detalle["dato_plan"] != nil {
					actividad := make(map[string]interface{})
					dato_plan_string := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_string), &dato_plan)

					for indice_actividad := range dato_plan {
						if indice_actividad == indice {
							auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
							actividad["index"] = indice_actividad
							actividad["dato"] = auxiliar_actividad["dato"]
							actividad["activo"] = auxiliar_actividad["activo"]
							actividad["observacion"] = elemento
							dato_plan[indice_actividad] = actividad
						}
					}
					auxiliar, _ := json.Marshal(dato_plan)
					str := string(auxiliar)
					subgrupo_detalle["dato_plan"] = str
				}
				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta1, subgrupo_detalle); err != nil {
					panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			}
			continue
		}
		identificador_subgrupoDetalle = llave
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador_subgrupoDetalle, &respuesta2); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
		}
		request.LimpiezaRespuestaRefactor(respuesta2, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["armonizacion_dato"] != nil {
			dato_armonizacion_string := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(dato_armonizacion_string), &armonizacion_dato)

			if armonizacion_dato[indice] != nil {
				auxiliar := make(map[string]interface{})
				auxiliar["armonizacionPED"] = armonizacion
				auxiliar["armonizacionPI"] = armonizacionPI
				armonizacion_dato[indice] = auxiliar
			}
			auxiliar, _ := json.Marshal(armonizacion_dato)
			stringArmonizacion := string(auxiliar)
			subgrupo_detalle["armonizacion_dato"] = stringArmonizacion
		}
		nuevoDato := true
		actividad := make(map[string]interface{})

		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_string := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_string), &dato_plan)

			for indice_actividad := range dato_plan {
				if indice_actividad == indice {
					nuevoDato = false
					auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
					actividad["index"] = indice_actividad
					actividad["dato"] = elemento
					actividad["activo"] = auxiliar_actividad["activo"]

					if auxiliar_actividad["observacion"] != nil {
						actividad["observacion"] = auxiliar_actividad["observacion"]
					}
					dato_plan[indice_actividad] = actividad
				}
			}
		}

		if nuevoDato {
			actividad["index"] = indice
			actividad["dato"] = elemento
			actividad["activo"] = true
			dato_plan[indice] = actividad
		}
		auxiliar, _ := json.Marshal(dato_plan)
		str := string(auxiliar)
		subgrupo_detalle["dato_plan"] = str

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta1, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "ActualizarActividad", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": entrada}
	c.ServeJSON()
}

// DesactivarActividad ...
// @Title DesactivarActividad
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /desactivar_actividad/:id/:index [put]
func (c *FormulacionController) DesactivarActividad() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
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

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)
		helpers.RecorrerHijos(hijos, indice)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Actividades Inactivas"}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// ConsultarTodasActividades ...
// @Title ConsultarTodasActividades
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_todas_actividades/:id/ [get]
func (c *FormulacionController) ConsultarTodasActividades() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
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
	var auxiliarHijos []interface{}

	helpers.LimpiaTabla()

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)

		for i := 0; i < len(hijos); i++ {
			auxiliarHijos = append(auxiliarHijos, hijos[i]["_id"])
		}
		tabla = helpers.ConsultarTabla(auxiliarHijos)
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": tabla}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// ConsultarArbolArmonizacion ...
// @Title ConsultarArbolArmonizacion
// @Description post Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	body		body 	{}	true		"body for Plan content"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_arbol_armonizacion/:id/ [post]
func (c *FormulacionController) ConsultarArbolArmonizacion() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var entrada map[string][]string
	var arregloIdentificador []string
	var armonizacion []map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &entrada)
	arregloIdentificador = entrada["Data"]

	for i := 0; i < len(arregloIdentificador); i++ {
		armonizacion = append(armonizacion, helpers.ConsultarArmonizacion(arregloIdentificador[i]))
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": armonizacion}
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
// @router /guardar_identificacion/:id/:idTipo [put]
func (c *FormulacionController) GuardarIdentificacion() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	identificador_plan := c.Ctx.Input.Param(":id")
	tipoIdentificacion := c.Ctx.Input.Param(":idTipo")
	var entrada map[string]interface{}
	var res map[string]interface{}
	var resJ map[string]interface{}
	var respuesta []map[string]interface{}
	var idStr string
	var identificacion map[string]interface{}
	json.Unmarshal(c.Ctx.Input.RequestBody, &entrada)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+identificador_plan+",tipo_identificacion_id:"+tipoIdentificacion, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuesta)

		if tipoIdentificacion == "61897518f6fc97091727c3c3" { // ? Recurso docente unicamente
			if len(respuesta) > 0 {
				if strings.Contains(respuesta[0]["dato"].(string), "ids_detalle") {
					identificacion = respuesta[0]
					dato_json := map[string]interface{}{}
					json.Unmarshal([]byte(identificacion["dato"].(string)), &dato_json)
					iddetail := ""
					identificacionDetalle := map[string]interface{}{}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhf"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rhf"]
						identificacionDetalle = map[string]interface{}{}

						if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut); err != nil {
							panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhf\"", "status": "400", "log": err})
						}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhf\"", "status": "400", "log": err})
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pre"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rhv_pre"]
						identificacionDetalle = map[string]interface{}{}

						if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut); err != nil {
							panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhv_pre\"", "status": "400", "log": err})
						}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhv_pre\"", "status": "400", "log": err})
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pos"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rhv_pos"]
						identificacionDetalle = map[string]interface{}{}

						if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut); err != nil {
							panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhv_pos\"", "status": "400", "log": err})
						}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhv_pos\"", "status": "400", "log": err})
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rubros"]
						identificacionDetalle = map[string]interface{}{}

						if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut); err != nil {
							panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rubros\"", "status": "400", "log": err})
						}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rubros\"", "status": "400", "log": err})
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros_pos"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						identificacionDetallePut := identificacionDetalle["Data"].(map[string]interface{})
						identificacionDetallePut["dato"] = entrada["rubros_pos"]
						identificacionDetalle = map[string]interface{}{}

						if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, "PUT", &identificacionDetalle, identificacionDetallePut); err != nil {
							panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rubros_pos\"", "status": "400", "log": err})
						}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error consultando detalle identificacion \"rubros_pos\"", "status": "400", "log": err})
					}
					c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Registro de identificación"}
				} else {
					// ? -- Transicional mientras migran datos --
					identificacion = respuesta[0]
					detalles := map[string]interface{}{
						"rhf":        "",
						"rhv_pre":    "",
						"rhv_pos":    "",
						"rubros":     "",
						"rubros_pos": "",
					}
					identificacionDetalle := map[string]interface{}{}

					data := map[string]interface{}{
						"dato": entrada["rhf"],
					}

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data); err == nil {
						detalles["rhf"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error guardando detalle identificacion \"rhf\"", "status": "400", "log": err})
					}

					data = map[string]interface{}{
						"dato": entrada["rhv_pre"],
					}

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data); err == nil {
						detalles["rhv_pre"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error guardando detalle identificacion \"rhv_pre\"", "status": "400", "log": err})
					}

					data = map[string]interface{}{
						"dato": entrada["rhv_pos"],
					}

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data); err == nil {
						detalles["rhv_pos"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error guardando detalle identificacion \"rhv_pos\"", "status": "400", "log": err})
					}

					data = map[string]interface{}{
						"dato": entrada["rubros"],
					}

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data); err == nil {
						detalles["rubros"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error guardando detalle identificacion \"rubros\"", "status": "400", "log": err})
					}

					data = map[string]interface{}{
						"dato": entrada["rubros_pos"],
					}

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/", "POST", &identificacionDetalle, data); err == nil {
						detalles["rubros_pos"] = identificacionDetalle["Data"].(map[string]interface{})["_id"].(string)
						identificacionDetalle = map[string]interface{}{}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error guardando detalle identificacion \"rubros_pos\"", "status": "400", "log": err})
					}
					bt, _ := json.Marshal(map[string]interface{}{"ids_detalle": detalles})
					identificacion["dato"] = string(bt)
					identificacionAns := map[string]interface{}{}

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion/"+identificacion["_id"].(string), "PUT", &identificacionAns, identificacion); err == nil {
						c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Registro de identificación"}
					} else {
						panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando identificacion", "status": "400", "log": err})
					}
				}
			} else {
				panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error sin dato identificacion", "status": "400", "log": err})
			}
		} else {
			jsonString, _ := json.Marshal(respuesta[0]["_id"])
			json.Unmarshal(jsonString, &idStr)
			identificacion = respuesta[0]
			b, _ := json.Marshal(entrada)
			str := string(b)
			identificacion["dato"] = str

			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion/"+idStr, "PUT", &resJ, identificacion); err != nil {
				panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando identificacion \"idStr\"", "status": "400", "log": err})
			}
			c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Registro de identificación"}
		}
	} else {
		panic(map[string]interface{}{"funcion": "GuardarIdentificacion", "err": "Error actualizando detalle identificacion \"rhv_pre\"", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarIdentificaciones ...
// @Title ConsultarIdentificaciones
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	idTipo		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_identificaciones/:id/:idTipo [get]
func (c *FormulacionController) ConsultarIdentificaciones() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	identificador_plan := c.Ctx.Input.Param(":id")
	tipoIdentificacion := c.Ctx.Input.Param(":idTipo")
	var respuesta []map[string]interface{}
	var res map[string]interface{}
	var identificacion map[string]interface{}
	var dato map[string]interface{}
	var data_identificacion []map[string]interface{}

	if tipoIdentificacion == "61897518f6fc97091727c3c3" { // ? Recurso docente unicamente
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+identificador_plan+",tipo_identificacion_id:"+tipoIdentificacion, &res); err == nil {
			request.LimpiezaRespuestaRefactor(res, &respuesta)
			if len(respuesta) > 0 {
				if strings.Contains(respuesta[0]["dato"].(string), "ids_detalle") {
					identificacion = respuesta[0]
					dato_json := map[string]interface{}{}
					json.Unmarshal([]byte(identificacion["dato"].(string)), &dato_json)
					result := dato_json["ids_detalle"].(map[string]interface{})
					iddetail := ""
					identificacionDetalle := map[string]interface{}{}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhf"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						datos = helpers.FiltrarIdentificaciones(dato)

						if len(datos) > 0 {
							result["rhf"] = datos
						} else {
							result["rhf"] = "{}"
						}
					} else {
						result["rhf"] = "{}"
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pre"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						datos = helpers.FiltrarIdentificaciones(dato)

						if len(datos) > 0 {
							result["rhv_pre"] = datos
						} else {
							result["rhv_pre"] = "{}"
						}
					} else {
						result["rhf"] = "{}"
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rhv_pos"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						datos = helpers.FiltrarIdentificaciones(dato)

						if len(datos) > 0 {
							result["rhv_pos"] = datos
						} else {
							result["rhv_pos"] = "{}"
						}
					} else {
						result["rhv_pos"] = "{}"
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						datos = helpers.FiltrarIdentificaciones(dato)

						if len(datos) > 0 {
							result["rubros"] = datos
						} else {
							result["rubros"] = "{}"
						}
					} else {
						result["rubros"] = "{}"
					}
					iddetail = dato_json["ids_detalle"].(map[string]interface{})["rubros_pos"].(string)

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion-detalle/"+iddetail, &identificacionDetalle); err == nil {
						dato_str := identificacionDetalle["Data"].(map[string]interface{})["dato"].(string)
						dato := map[string]interface{}{}
						datos := []map[string]interface{}{}
						json.Unmarshal([]byte(dato_str), &dato)
						datos = helpers.FiltrarIdentificaciones(dato)

						if len(datos) > 0 {
							result["rubros_pos"] = datos
						} else {
							result["rubros_pos"] = "{}"
						}
					} else {
						result["rubros_pos"] = "{}"
					}
					c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": result}
				} else {
					identificacion = respuesta[0]
					if identificacion["dato"] != nil && identificacion["dato"] != "{}" { // ? Antiguo método unificado
						result := make(map[string]interface{})
						dato_str := identificacion["dato"].(string)
						json.Unmarshal([]byte(dato_str), &dato)
						var identi map[string]interface{} = nil
						dato_aux := ""
						dato_aux = dato["rhf"].(string)

						if dato_aux == "{}" {
							result["rhf"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identificacion = append(data_identificacion, element)
								}
							}
							result["rhf"] = data_identificacion
						}
						identi = nil
						data_identificacion = nil
						dato_aux = dato["rhv_pre"].(string)

						if dato_aux == "{}" {
							result["rhv_pre"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identificacion = append(data_identificacion, element)
								}
							}
							result["rhv_pre"] = data_identificacion
						}

						identi = nil
						data_identificacion = nil
						dato_aux = dato["rhv_pos"].(string)

						if dato_aux == "{}" {
							result["rhv_pos"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identificacion = append(data_identificacion, element)
								}
							}
							result["rhv_pos"] = data_identificacion
						}
						identi = nil
						data_identificacion = nil
						dato_aux = dato["rubros"].(string)

						if dato_aux == "{}" {
							result["rubros"] = "{}"
						} else {
							json.Unmarshal([]byte(dato_aux), &identi)
							for key := range identi {
								element := identi[key].(map[string]interface{})
								if element["activo"] == true {
									data_identificacion = append(data_identificacion, element)
								}
							}
							result["rubros"] = data_identificacion
						}
						identi = nil
						data_identificacion = nil

						if dato["rubros_pos"] != nil {
							dato_aux = dato["rubros_pos"].(string)
							if dato_aux == "{}" {
								result["rubros_pos"] = "{}"
							} else {
								json.Unmarshal([]byte(dato_aux), &identi)
								for key := range identi {
									element := identi[key].(map[string]interface{})
									if element["activo"] == true {
										data_identificacion = append(data_identificacion, element)
									}
								}
								result["rubros_pos"] = data_identificacion
							}
						}
						c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": result}
					} else {
						c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ""}
					}
				}
			} else {
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ""}
			}
		} else {
			panic(err)
		}
	} else {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+identificador_plan+",tipo_identificacion_id:"+tipoIdentificacion, &res); err == nil {
			request.LimpiezaRespuestaRefactor(res, &respuesta)
			identificacion = respuesta[0]

			if identificacion["dato"] != nil {
				dato_str := identificacion["dato"].(string)
				json.Unmarshal([]byte(dato_str), &dato)

				for key := range dato {
					element := dato[key].(map[string]interface{})
					if element["activo"] == true {
						data_identificacion = append(data_identificacion, element)
					}
				}

				sort.SliceStable(data_identificacion, func(i, j int) bool {
					a, _ := strconv.Atoi(data_identificacion[i]["index"].(string))
					b, _ := strconv.Atoi(data_identificacion[j]["index"].(string))
					return a < b
				})
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": data_identificacion}
			} else {
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ""}
			}
		} else {
			panic(err)
		}
	}
	c.ServeJSON()
}

// DesactivarIdentificacion ...
// @Title DesactivarIdentificacion
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Param	idTipo		path 	string	true		"The key for staticblock"
// @Param	index		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /desactivar_identificacion/:id/:idTipo/:index [put]
func (c *FormulacionController) DesactivarIdentificacion() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	identificador_plan := c.Ctx.Input.Param(":id")
	indice := c.Ctx.Input.Param(":index")
	identificacionTipo := c.Ctx.Input.Param(":idTipo")
	var identificadorString string
	var res map[string]interface{}
	var respuesta []map[string]interface{}
	var identificacion map[string]interface{}
	var dato map[string]interface{}
	var respuestaJ map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+identificador_plan+",tipo_identificacion_id:"+identificacionTipo, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuesta)
		identificacion = respuesta[0]
		jsonString, _ := json.Marshal(respuesta[0]["_id"])
		json.Unmarshal(jsonString, &identificadorString)

		if identificacion["dato"] != nil {
			dato_str := identificacion["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)

			for llave := range dato {
				intVar, _ := strconv.Atoi(llave)
				intVar = intVar + 1
				strr := strconv.Itoa(intVar)
				if strr == indice {
					elemento := dato[llave].(map[string]interface{})
					elemento["activo"] = false
					dato[llave] = elemento
				}
			}

			b, _ := json.Marshal(dato)
			str := string(b)
			identificacion["dato"] = str
		}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion/"+identificadorString, "PUT", &respuestaJ, identificacion); err != nil {
			panic(map[string]interface{}{"funcion": "DeleteIdentificacion", "err": "Error eliminando identificacion \"idStr\"", "status": "400", "log": err})
		}
		c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": "Identificación Inactiva"}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// VersionarPlan ...
// @Title VersionarPlan
// @Description post Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /versionar_plan/:id [post]
func (c *FormulacionController) VersionarPlan() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	identificador_plan := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	var respuestaHijos map[string]interface{}
	var respuestaIdentificacion map[string]interface{}
	var hijos []map[string]interface{}
	var identificaciones []map[string]interface{}
	var planPadre map[string]interface{}
	var respuestaPost map[string]interface{}
	var planVersionado map[string]interface{}
	plan := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificador_plan, &respuesta); err == nil {
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
		plan["padre_plan_id"] = identificador_plan

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuestaPost, plan); err != nil {
			panic(map[string]interface{}{"funcion": "VersionarPlan", "err": "Error versionando plan \"plan[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		planVersionado = respuestaPost["Data"].(map[string]interface{})
		c.Data["json"] = respuestaPost

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+identificador_plan, &respuestaIdentificacion); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaIdentificacion, &identificaciones)
			if len(identificaciones) != 0 {
				helpers.VersionarIdentificaciones(identificaciones, planVersionado["_id"].(string))
			}
		} else {
			panic(err)
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+identificador_plan, &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
			helpers.VersionarHijos(hijos, planVersionado["_id"].(string))
		} else {
			panic(err)
		}

		var respuestaPadres map[string]interface{}
		var planesPadre []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=dependencia_id:"+plan["dependencia_id"].(string)+",vigencia:"+plan["vigencia"].(string)+",formato:false,nombre:"+url.QueryEscape(plan["nombre"].(string)), &respuestaPadres); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaPadres, &planesPadre)

			for _, padre := range planesPadre {
				var respuestaActualizacion map[string]interface{}

				if padre["_id"].(string) != planVersionado["_id"].(string) {
					padre["activo"] = false

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+padre["_id"].(string), "PUT", &respuestaActualizacion, padre); err != nil {
						panic(err)
					}
				}
			}
		} else {
			panic(err)
		}
	}
	c.ServeJSON()
}

// ConsultarPlanVersiones ...
// @Title ConsultarPlanVersiones
// @Description get Formulacion by id
// @Param	unidad		path 	string	true		"The key for staticblock"
// @Param	vigencia		path 	string	true		"The key for staticblock"
// @Param	nombre		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_plan_versiones/:unidad/:vigencia/:nombre [get]
func (c *FormulacionController) ConsultarPlanVersiones() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	unidad := c.Ctx.Input.Param(":unidad")
	vigencia := c.Ctx.Input.Param(":vigencia")
	nombre := c.Ctx.Input.Param(":nombre")
	var respuesta map[string]interface{}
	var versiones []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=dependencia_id:"+unidad+",vigencia:"+vigencia+",formato:false,nombre:"+nombre, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &versiones)
		versionesOrdenadas := helpers.OrdenarVersiones(versiones)
		c.Data["json"] = versionesOrdenadas
	}
	c.ServeJSON()
}

// GetPonderacionActividades ...
// @Title GetPonderacionActividades
// @Description get Formulacion by id
// @Param	plan		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /ponderacion_actividades/:plan [get]
func (c *FormulacionController) PonderacionActividades() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	plan := c.Ctx.Input.Param(":plan")
	var respuesta map[string]interface{}
	var respuestaDetalle map[string]interface{}
	var respuestaLimpiaDetalle []map[string]interface{}
	var subgrupoDetalle map[string]interface{}
	var hijos []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+plan, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &hijos)

		for i := 0; i < len(hijos); i++ {
			if strings.Contains(strings.ToUpper(hijos[i]["nombre"].(string)), "PONDERACIÓN") && strings.Contains(strings.ToUpper(hijos[i]["nombre"].(string)), "ACTIVIDAD") || strings.Contains(strings.ToUpper(hijos[i]["nombre"].(string)), "PONDERACIÓN") {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+hijos[i]["_id"].(string), &respuestaDetalle); err == nil {
					request.LimpiezaRespuestaRefactor(respuestaDetalle, &respuestaLimpiaDetalle)
					subgrupoDetalle = respuestaLimpiaDetalle[0]

					if subgrupoDetalle["dato_plan"] != nil {
						var suma float64 = 0
						datoPlan := make(map[string]map[string]interface{})
						json.Unmarshal([]byte(subgrupoDetalle["dato_plan"].(string)), &datoPlan)
						ponderacionActividades := make(map[string]interface{})

						for j, dato := range datoPlan {
							if dato["activo"] != false && len(dato) != 0 {
								ponderacionActividades["Actividad "+(j)] = dato["dato"]
								suma += dato["dato"].(float64)
								suma = math.Round(suma*100) / 100
							}
						}
						ponderacionActividades["Total"] = suma
						c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ponderacionActividades}
					}
				} else {
					panic(map[string]interface{}{"funcion": "PonderacionActividades", "err": "Error subgrupo_detalle plan \"plan\"", "status": "400", "log": err})
				}
			}
		}
	} else {
		panic(map[string]interface{}{"funcion": "PonderacionActividades", "err": "Error subgrupo_hijos plan \"plan\"", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarRubros ...
// @Title ConsultarRubros
// @Description get Rubros
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_rubros [get]
func (c *FormulacionController) ConsultarRubros() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var respuesta map[string]interface{}
	var rubros []interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanCuentasService")+"/arbol_rubro", &respuesta); err == nil {
		rubros = respuesta["Body"].([]interface{})

		for i := 0; i < len(rubros); i++ {
			if strings.ToUpper(rubros[i].(map[string]interface{})["Nombre"].(string)) == "GASTOS" {
				aux := rubros[i].(map[string]interface{})
				hojas := helpers.ConsultarHijosRubro(aux["Hijos"].([]interface{}))
				if len(hojas) != 0 {
					c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": hojas}
				} else {
					c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ""}
				}
				break
			}
		}
	} else {
		panic(map[string]interface{}{"funcion": "ConsultarRubros", "err": "Error arbol_rubros", "status": "400", "log": err})
	}
	c.ServeJSON()
}

// ConsultarUnidades ...
// @Title ConsultarUnidades
// @Description get Unidades
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /consultar_unidades [get]
func (c *FormulacionController) ConsultarUnidades() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var respuestaTipoDependencia []models.TipoDependencia
	var respuestaDependencia []models.Dependencia
	var respuesta []models.DependenciaTipoDependencia
	var Dependencias []models.Dependencia
	var TiposDependencias []models.TipoDependencia
	var unidad models.Dependencia
	var unidades []models.Dependencia

	TiposDependenciasUsadas := map[string]bool{
		"FACULTAD":             true,
		"DECANATURA":           true,
		"VICERRECTORIA":        true,
		"OFICINA ASESORA":      true,
		"DIVISIÓN":             true,
		"SECCIÓN":              true,
		"SECRETARÍA ACADÉMICA": true,
		"INSTITUTO":            true,
		"ENTIDAD":              true,
		"ASOCIACIÓN":           true,
		"PREGRADO":             true,
		"POSGRADO":             true,
		"UNIDAD EJECUTORA":     true,
		"OFICINA":              true,
	}

	DependenciasUsadas := map[string]bool{
		"PROYECTO PLANESTIC-UD": true,
		"SECCION BIBLIOTECA":    true,
		"SUBSISTEMA DE LA SEGURIDAD Y SALUD EN EL TRABAJO SG-SST":   true,
		"SECCION DE ACTAS, ARCHIVO Y MICROFILMACION":                true,
		"INSTITUTO DE INVESTIGACIÓN E INNOVACIÓN EN INGENERÍA -I3+": true,
		"SECCION DE PUBLICACIONES":                                  true,
		"CATEDRA UNESCO EN DESARROLLO DEL NINO":                     true,
	}

	if err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/tipo_dependencia/?&limit=0", &respuestaTipoDependencia); err == nil {
		for i := 0; i < len(respuestaTipoDependencia); i++ {
			if TiposDependenciasUsadas[respuestaTipoDependencia[i].Nombre] == true {
				TiposDependencias = append(TiposDependencias, respuestaTipoDependencia[i])
			}
		}
	} else {
		panic(err)
	}

	if err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia/?limit=0", &respuestaDependencia); err == nil {
		for i := 0; i < len(respuestaDependencia); i++ {
			if DependenciasUsadas[respuestaDependencia[i].Nombre] {
				Dependencias = append(Dependencias, respuestaDependencia[i])
			}
		}
	} else {
		panic(err)
	}

	for i := 0; i < len(TiposDependencias); i++ {
		if TiposDependencias[i].Nombre == "INSTITUTO" || TiposDependencias[i].Nombre == "PREGRADO" || TiposDependencias[i].Nombre == "OFICINA" {
			for j := 0; j < len(Dependencias); j++ {
				if Dependencias[j].DependenciaTipoDependencia[0].TipoDependenciaId.Nombre == TiposDependencias[i].Nombre {
					if err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia_tipo_dependencia?query=TipoDependenciaId:"+strconv.Itoa(TiposDependencias[i].Id)+",DependenciaId:"+strconv.Itoa(Dependencias[j].Id)+"&limit=0", &respuesta); err == nil {
						unidad = *respuesta[0].DependenciaId
						unidad.TipoDependencia = respuesta[0].TipoDependenciaId
						unidades = append(unidades, unidad)
					} else {
						panic(err)
					}
				}
			}
		} else {
			if err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia_tipo_dependencia?query=TipoDependenciaId:"+strconv.Itoa(TiposDependencias[i].Id)+"&limit=0", &respuesta); err == nil {
				for j := 0; j < len(respuesta); j++ {
					unidad = *respuesta[j].DependenciaId
					unidad.TipoDependencia = respuesta[j].TipoDependenciaId
					unidades = append(unidades, unidad)
				}
			} else {
				panic(err)
			}
		}
		respuesta = nil
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": unidades}
	c.ServeJSON()
}

// VinculacionTercero ...
// @Title VinculacionTercero
// @Description get VinculacionTercero
// @Param	tercero_id	path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /vinculacion_tercero/:tercero_id [get]
func (c *FormulacionController) VinculacionTercero() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	terceroIdentificacion := c.Ctx.Input.Param(":tercero_id")
	var vinculaciones []models.Vinculacion

	if err := request.GetJson("http://"+beego.AppConfig.String("TercerosService")+"/vinculacion?query=Activo:true,TerceroPrincipalId:"+terceroIdentificacion, &vinculaciones); err == nil {
		for i := 0; i < len(vinculaciones); i++ {
			if vinculaciones[i].CargoId == 319 || vinculaciones[i].CargoId == 312 || vinculaciones[i].CargoId == 320 {
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": vinculaciones[i]}
				break
			} else {
				c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ""}
			}
		}
	} else {
		panic(map[string]interface{}{"funcion": "VinculacionTercero", "err": "Error get vinculacion", "status": "400", "log": err})
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
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	var respuesta map[string]interface{}
	var planes []map[string]interface{}
	var planesPED []map[string]interface{}
	var planesPI []map[string]interface{}
	var tipoPlanes []map[string]interface{}
	var plan map[string]interface{}
	var arregloPlanes []map[string]interface{}
	var auxArregloPlanes []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=formato:true", &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &planes)

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=tipo_plan_id:6239117116511e20405d408b", &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &planesPI)
		} else {
			panic(err)
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=tipo_plan_id:616513b91634adfaffed52bf", &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &planesPED)
		} else {
			panic(err)
		}

		auxArregloPlanes = append(auxArregloPlanes, planes...)
		auxArregloPlanes = append(auxArregloPlanes, planesPI...)
		auxArregloPlanes = append(auxArregloPlanes, planesPED...)

		for i := 0; i < len(auxArregloPlanes); i++ {
			plan = auxArregloPlanes[i]
			tipoPlanId := plan["tipo_plan_id"].(string)

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/tipo-plan?query=_id:"+tipoPlanId, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &tipoPlanes)
				tipoPlan := tipoPlanes[0]
				nombreTipoPlan := tipoPlan["nombre"]
				planesTipo := make(map[string]interface{})
				planesTipo["_id"] = plan["_id"]
				planesTipo["nombre"] = plan["nombre"]
				planesTipo["descripcion"] = plan["descripcion"]
				planesTipo["tipo_plan_id"] = tipoPlanId
				planesTipo["formato"] = plan["formato"]
				planesTipo["vigencia"] = plan["vigencia"]
				planesTipo["dependencia_id"] = plan["dependencia_id"]
				planesTipo["aplicativo_id"] = plan["aplicativo_id"]
				planesTipo["activo"] = plan["activo"]
				planesTipo["nombre_tipo_plan"] = nombreTipoPlan

				arregloPlanes = append(arregloPlanes, planesTipo)

				if arregloPlanes != nil {
					c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": arregloPlanes}
				} else {
					c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": ""}
				}
			} else {
				panic(err)
			}
		}
	} else {
		panic(err)
	}
	c.ServeJSON()
}

// VerificarIdentificaciones ...
// @Title VerificarIdentificaciones
// @Description put Formulacion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Formulacion
// @Failure 403 :id is empty
// @router /verificar_identificaciones/:id [get]
func (c *FormulacionController) VerificarIdentificaciones() {
	defer func() {
		if err := recover(); err != nil {
			localError := err.(map[string]interface{})
			c.Data["mesaage"] = (beego.AppConfig.String("appname") + "/" + "FormulacionController" + "/" + (localError["funcion"]).(string))
			c.Data["data"] = (localError["err"])
			if status, ok := localError["status"]; ok {
				c.Abort(status.(string))
			} else {
				c.Abort("404")
			}
		}
	}()

	identificacion := c.Ctx.Input.Param(":id")
	var respuesta map[string]interface{}
	var respuestaPlan map[string]interface{}
	var respuestaDependencia []map[string]interface{}
	var dependencia map[string]interface{}
	var plan map[string]interface{}
	var identificaciones []map[string]interface{}
	var bandera bool

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+identificacion, &respuestaPlan); err == nil {
		request.LimpiezaRespuestaRefactor(respuestaPlan, &plan)

		if err := request.GetJson("http://"+beego.AppConfig.String("OikosService")+"/dependencia_tipo_dependencia?query=DependenciaId:"+plan["dependencia_id"].(string), &respuestaDependencia); err == nil {
			dependencia = respuestaDependencia[0]

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion?query=plan_id:"+identificacion, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &identificaciones)

				tipoDependencia := dependencia["TipoDependenciaId"].(map[string]interface{})
				identificacion := dependencia["DependenciaId"].(map[string]interface{})["Id"]
				if (tipoDependencia["Id"] == 2.00 || identificacion == 67.00) && identificacion != 8.0 {
					bandera = helpers.VerificarDataIdentificaciones(identificaciones, "facultad")
				} else {
					bandera = helpers.VerificarDataIdentificaciones(identificaciones, "unidad")
				}

			} else {
				panic(err)
			}
		} else {
			panic(err)
		}
	} else {
		panic(err)
	}
	c.Data["json"] = map[string]interface{}{"Success": true, "Status": "200", "Message": "Successful", "Data": bandera}
	c.ServeJSON()
}
