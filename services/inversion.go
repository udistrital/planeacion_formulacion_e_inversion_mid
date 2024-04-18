package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/planeacion_formulacion_mid/helpers"
	formulacionhelper "github.com/udistrital/planeacion_formulacion_mid/helpers/formulacionHelper"
	inversionhelper "github.com/udistrital/planeacion_formulacion_mid/helpers/inversionHelper"
	"github.com/udistrital/utils_oas/request"
)

func AddProyecto(datos []byte) (interface{}, error) {
	var registroProyecto map[string]interface{}
	var idProyecto string
	var resPlan map[string]interface{}
	if err := json.Unmarshal(datos, &registroProyecto); err == nil {
		plan := map[string]interface{}{
			"activo":        true,
			"nombre":        registroProyecto["nombre_proyecto"],
			"descripcion":   registroProyecto["codigo_proyecto"],
			"tipo_plan_id":  "63ca86f1b6c0e5725a977dae",
			"aplicativo_id": " ",
		}
		var respuesta map[string]interface{}

		err1 := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan", "POST", &respuesta, plan)
		if err1 == nil {
			resPlan = respuesta["Data"].(map[string]interface{})
			idProyecto = resPlan["_id"].(string)

			soportes := map[string]interface{}{"codigo_proyecto": registroProyecto["codigo_proyecto"], "data": registroProyecto["soportes"]}
			errSoporte := inversionhelper.RegistrarInfoComplementaria(idProyecto, soportes, "soportes")

			fuentes := map[string]interface{}{"codigo_proyecto": registroProyecto["codigo_proyecto"], "data": registroProyecto["fuentes"]}
			errFuentes := inversionhelper.RegistrarInfoComplementaria(idProyecto, fuentes, "fuentes apropiacion")
			inversionhelper.ActualizarPresupuestoDisponible(registroProyecto["fuentes"].([]interface{}))

			metas := map[string]interface{}{"codigo_proyecto": registroProyecto["codigo_proyecto"], "data": registroProyecto["metas"]}
			errMetas := inversionhelper.RegistrarInfoComplementaria(idProyecto, metas, "metas asociadas al proyecto de inversion")

			if errSoporte != nil || errFuentes != nil || errMetas != nil {
				return nil, errors.New("error del servicio AddProyecto: Error Registrando la información complementaria " + errSoporte.Error())
			}
			return resPlan, nil
		}

	} else {
		return nil, errors.New("error del servicio AddProyecto: Error Registrando Proyecto " + err.Error())
	}

	return nil, errors.New("error del servicio AddProyecto: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido ")
}

func EditProyecto(id string, datos []byte) (interface{}, error) {
	var registroProyecto map[string]interface{}
	var res map[string]interface{}
	var infoProyect map[string]interface{}
	if err := json.Unmarshal(datos, &registroProyecto); err == nil {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
			request.LimpiezaRespuestaRefactor(res, &infoProyect)
			infoProyect["nombre"] = registroProyecto["nombre_proyecto"]
			infoProyect["descripcion"] = registroProyecto["codigo_proyecto"]

			var respuesta map[string]interface{}
			err1 := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, "PUT", &respuesta, infoProyect)
			if err1 == nil {
				errSoporte := inversionhelper.ActualizarInfoComplDetalle(registroProyecto["id_detalle_soportes"].(string), registroProyecto["soportes"].([]interface{}))

				errFuentes := inversionhelper.ActualizarInfoComplDetalle(registroProyecto["id_detalle_fuentes"].(string), registroProyecto["fuentes"].([]interface{}))
				inversionhelper.ActualizarPresupuestoDisponible(registroProyecto["fuentes"].([]interface{}))

				errMetas := inversionhelper.ActualizarInfoComplDetalle(registroProyecto["id_detalle_metas"].(string), registroProyecto["metas"].([]interface{}))

				if errSoporte != nil || errFuentes != nil || errMetas != nil {
					return nil, errors.New("error del servicio EditProyecto: Error Actualizando la información complementaria " + errSoporte.Error())
				}
				return infoProyect, nil
			}
		} else {
			return nil, errors.New("error del servicio GetProyectoId: Error obteniendo información plan " + err.Error())
		}

	} else {
		return nil, errors.New("error del servicio EditProyecto: Error Registrando Proyecto " + err.Error())
	}

	return nil, errors.New("error del servicio EditProyecto: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido")
}

func GuardarDocumentos(datos []byte) (interface{}, error) {
	var body map[string]interface{}
	var evidencias []map[string]interface{}
	if err := json.Unmarshal(datos, &body); err == nil {
		if body["documento"] != nil {
			resDocs := helpers.GuardarDocumento(body["documento"].([]interface{}))

			for _, doc := range resDocs {
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
			return evidencias, nil
		} else {
			return nil, errors.New("error del servicio GuardarDocumentos: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	} else {
		return nil, errors.New("error del servicio GuardarDocumentos: Error Registrando Documentos " + err.Error())
	}
}

func GetProyectoId(id string) (interface{}, error) {
	var res map[string]interface{}
	getProyect := make(map[string]interface{})
	var infoProyect map[string]interface{}
	var subgruposData map[string]interface{}
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &infoProyect)
		getProyect["nombre_proyecto"] = infoProyect["nombre"]
		getProyect["codigo_proyecto"] = infoProyect["descripcion"]
		getProyect["fecha_creacion"] = infoProyect["fecha_creacion"]
		padreId := infoProyect["_id"].(string)

		var infoSubgrupos []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+padreId, &subgruposData); err == nil {
			request.LimpiezaRespuestaRefactor(subgruposData, &infoSubgrupos)
			for i := range infoSubgrupos {
				var subgrupoDetalle map[string]interface{}
				var detalleSubgrupos []map[string]interface{}

				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=subgrupo_id:"+infoSubgrupos[i]["_id"].(string), &subgrupoDetalle); err == nil {
					request.LimpiezaRespuestaRefactor(subgrupoDetalle, &detalleSubgrupos)

					armonizacion_dato_str := detalleSubgrupos[0]["dato"].(string)
					var subgrupo_dato []map[string]interface{}
					json.Unmarshal([]byte(armonizacion_dato_str), &subgrupo_dato)

					if strings.Contains(strings.ToLower(infoSubgrupos[i]["nombre"].(string)), "soporte") {
						getProyect["soportes"] = subgrupo_dato
						getProyect["id_detalle_soportes"] = detalleSubgrupos[0]["_id"]
					}
					if strings.Contains(strings.ToLower(infoSubgrupos[i]["nombre"].(string)), "metas") {
						getProyect["metas"] = subgrupo_dato
						getProyect["id_detalle_metas"] = detalleSubgrupos[0]["_id"]
					}
					if strings.Contains(strings.ToLower(infoSubgrupos[i]["nombre"].(string)), "fuentes") {
						getProyect["fuentes"] = subgrupo_dato
						getProyect["id_detalle_fuentes"] = detalleSubgrupos[0]["_id"]
					}
				}
			}
		} else {
			return nil, errors.New("error del servicio GetProyectoId: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		return getProyect, nil
	} else {
		return nil, errors.New("error del servicio GetProyectoId: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func GetMetasProyect(id string) (interface{}, error) {
	var res map[string]interface{}
	getProyect := make(map[string]interface{})
	var infoProyect map[string]interface{}
	var subgruposData map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &infoProyect)
		var infoSubgrupos []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+id, &subgruposData); err == nil {
			request.LimpiezaRespuestaRefactor(subgruposData, &infoSubgrupos)
			for i := range infoSubgrupos {
				var subgrupoDetalle map[string]interface{}
				var detalleSubgrupos []map[string]interface{}
				if strings.Contains(strings.ToLower(infoSubgrupos[i]["nombre"].(string)), "metas") {
					getProyect["subgrupo_id_metas"] = infoSubgrupos[i]["_id"]
					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=subgrupo_id:"+infoSubgrupos[i]["_id"].(string), &subgrupoDetalle); err == nil {
						request.LimpiezaRespuestaRefactor(subgrupoDetalle, &detalleSubgrupos)
						var dato_metas []map[string]interface{}
						getProyect["id_detalle_meta"] = detalleSubgrupos[0]["_id"]
						datoMeta_str := detalleSubgrupos[0]["dato"].(string)
						json.Unmarshal([]byte(datoMeta_str), &dato_metas)
						getProyect["metas"] = dato_metas
					}
				}
			}
		} else {
			return nil, errors.New("error del servicio GetMetasProyect: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		return getProyect, nil
	} else {
		return nil, errors.New("error del servicio GetMetasProyect: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func GetAllProyectos(tipo_plan_id string) (interface{}, error) {
	var res map[string]interface{}
	var getProyect []map[string]interface{}
	var proyecto map[string]interface{}
	var dataProyects []map[string]interface{}

	proyect := make(map[string]interface{})
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan?query=activo:true,tipo_plan_id:"+tipo_plan_id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &dataProyects)
		for i := range dataProyects {
			if dataProyects[i]["activo"] == true {
				proyect["id"] = dataProyects[i]["_id"]
				proyecto = inversionhelper.GetDataProyects(dataProyects[i])
			}
			getProyect = append(getProyect, proyecto)
		}
		return getProyect, nil
	} else {
		return nil, errors.New("error del servicio GetProyectoId: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func ActualizarSubgrupoDetalle(id string, datos []byte) (interface{}, error) {
	var subDetalle map[string]interface{}
	var res map[string]interface{}
	json.Unmarshal(datos, &subDetalle)
	if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+id, "PUT", &res, subDetalle); err == nil {
		return res, nil
	} else {
		return nil, errors.New("error del servicio ActualizarSubgrupoDetalle: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func ActualizarProyectoGeneral(id string, datos []byte) (interface{}, error) {
	var infoProyecto map[string]interface{}
	var res map[string]interface{}
	json.Unmarshal(datos, &infoProyecto)
	if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, "PUT", &res, infoProyecto); err == nil {
		return res, nil
	} else {
		return nil, errors.New("error del servicio ActualizarProyectoGeneral: La solicitud de actualizar plan \"id\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func CrearPlan(datos []byte) (interface{}, error) {
	var respuesta map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}
	var respuestaPost map[string]interface{}
	var planSubgrupo map[string]interface{}
	json.Unmarshal(datos, &parametros)
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
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/", "POST", &respuestaPost, plan); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaPost, &planSubgrupo)
			padre := planSubgrupo["_id"].(string)
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &respuestaHijos); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
				formulacionhelper.ClonarHijos(hijos, padre)
			}

		} else {
			return nil, errors.New("error del servicio CrearPlan: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		return planSubgrupo, nil
	} else {
		return nil, errors.New("error del servicio CrearPlan: La solicitud de consultar datos Plan Formato contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func GetPlanId(id string) (interface{}, error) {
	var res map[string]interface{}
	getProyect := make(map[string]interface{})
	var infoProyect map[string]interface{}
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &infoProyect)
		getProyect["nombre_proyecto"] = infoProyect["nombre"]
		getProyect["codigo_proyecto"] = infoProyect["descripcion"]
		getProyect["fecha_creacion"] = infoProyect["fecha_creacion"]
		return getProyect, nil
	} else {
		return nil, errors.New("error del servicio GetProyectoId: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func GuardarMeta(id string, datos []byte) (interface{}, error) {
	var body map[string]interface{}
	var res map[string]interface{}
	var entrada map[string]interface{}
	var resPlan map[string]interface{}
	var plan map[string]interface{}
	var respuestaGuardado map[string]interface{}

	json.Unmarshal(datos, &body)
	entrada = body["entrada"].(map[string]interface{})
	idSubDetalleProI := body["idSubDetalle"]
	indexMetaSubProI := body["indexMetaSubPro"] //para efectos de actividades será indexMeta
	maxIndex := formulacionhelper.GetIndexActividad(entrada)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, &resPlan); err != nil {
		return nil, errors.New("error del servicio GuardarMeta: La solicitud de obtener Plan \"id\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
	request.LimpiezaRespuestaRefactor(resPlan, &plan)
	if plan["estado_plan_id"] != "614d3ad301c7a200482fabfd" {
		var res map[string]interface{}
		plan["estado_plan_id"] = "614d3ad301c7a200482fabfd"
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/plan/"+id, "PUT", &res, plan); err != nil {
			return nil, errors.New("error del servicio GuardarMeta: La solicitud de actualizar estado \"id\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	}

	for key, element := range entrada {

		var respuesta map[string]interface{}
		var respuestaLimpia []map[string]interface{}
		var subgrupo_detalle map[string]interface{}
		dato_plan := make(map[string]interface{})

		if element != "" {
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+key, &respuesta); err != nil {
				return nil, errors.New("error del servicio GuardarMeta: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			actividad := make(map[string]interface{})

			if subgrupo_detalle["dato_plan"] == nil {
				actividad["index"] = 1
				actividad["dato"] = element
				actividad["activo"] = true
				i := strconv.Itoa(actividad["index"].(int))
				dato_plan[i] = actividad

				b, _ := json.Marshal(dato_plan)
				str := string(b)
				subgrupo_detalle["dato_plan"] = str
				fmt.Sprintln(subgrupo_detalle["dato_plan"], "dato plan")
				armonizacion_dato := make(map[string]interface{})
				aux := make(map[string]interface{})
				aux["idSubDetalleProI"] = idSubDetalleProI
				aux["indexMetaSubProI"] = indexMetaSubProI
				aux["indexMetaPlan"] = 1
				armonizacion_dato[i] = aux
				c, _ := json.Marshal(armonizacion_dato)
				strArmonizacion := string(c)
				subgrupo_detalle["armonizacion_dato"] = strArmonizacion
			} else {
				dato_plan_str := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_str), &dato_plan)

				actividad["index"] = maxIndex + 1
				actividad["dato"] = element
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
					aux := make(map[string]interface{})
					aux["idSubDetalleProI"] = idSubDetalleProI
					aux["indexMetaSubProI"] = indexMetaSubProI
					aux["indexMetaPlan"] = i
					armonizacion_dato[i] = aux
					c, _ := json.Marshal(armonizacion_dato)
					strArmonizacion := string(c)
					subgrupo_detalle["armonizacion_dato"] = strArmonizacion
				}
			}
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
				return nil, errors.New("error del servicio GuardarMeta: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
			request.LimpiezaRespuestaRefactor(res, &respuestaGuardado)
		}
	}
	return respuestaGuardado, nil
}

func ObtenerPlan(id string) (interface{}, error) {
	var subgrupo []map[string]interface{}
	var res map[string]interface{}
	var id_subgrupoDetalle string
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	armo_dato := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Armonizar,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		id_subgrupoDetalle = subgrupo[0]["_id"].(string)
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+id_subgrupoDetalle, &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			armonizacion_dato_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(armonizacion_dato_str), &armo_dato)

		} else {
			return nil, errors.New("error del servicio GetPlan: La solicitud de obtener subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		return armo_dato, nil
	} else {
		return nil, errors.New("error del servicio GetPlan: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func ArmonizarInversion(id string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var body map[string]interface{}
	var subgrupo []map[string]interface{}
	var id_subgrupoDetalle string
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var armonizacionUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	json.Unmarshal(datos, &body)
	armonizacion_data, _ := json.Marshal(body)
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Armonizar,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		subgrupoPost := make(map[string]interface{})
		subDetallePost := make(map[string]interface{})

		if len(subgrupo) != 0 {
			id_subgrupoDetalle = subgrupo[0]["_id"].(string)
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+id_subgrupoDetalle, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
				if len(respuestaLimpia) > 0 {
					subgrupo_detalle = respuestaLimpia[0]
					subDetallePost["subgrupo_id"] = id_subgrupoDetalle
					subDetallePost["fecha_creacion"] = subgrupo_detalle["fecha_creacion"]
					subDetallePost["nombre"] = "Detalle Información Armonización"
					subDetallePost["descripcion"] = "Armonizar"
					subDetallePost["dato"] = " "
					subDetallePost["activo"] = true
					subDetallePost["armonizacion_dato"] = string(armonizacion_data)
					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &armonizacionUpdate)
					} else {
						return nil, errors.New("error del servicio ArmonizarInversion: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
					}
				} else {
					subDetallePost["subgrupo_id"] = id_subgrupoDetalle
					subDetallePost["nombre"] = "Detalle Información Armonización"
					subDetallePost["descripcion"] = "Armonizar"
					subDetallePost["dato"] = " "
					subDetallePost["activo"] = true
					subDetallePost["armonizacion_dato"] = string(armonizacion_data)

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &armonizacionUpdate)
					} else {
						return nil, errors.New("error del servicio ArmonizarInversion: La solicitud de registrar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
					}
				}

			} else {
				return nil, errors.New("error del servicio ArmonizarInversion: La solicitud de obtener subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
		} else {
			subgrupoPost["nombre"] = "Armonización Plan Inversión"
			subgrupoPost["descripcion"] = "Armonizar"
			subgrupoPost["padre"] = id
			subgrupoPost["activo"] = true
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/", "POST", &respuesta, subgrupoPost); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
				subgrupo_detalle = respuestaLimpia[0]
				subDetallePost["subgrupo_id"] = subgrupo_detalle["_id"]
				subDetallePost["nombre"] = "Detalle Información Armonización"
				subDetallePost["descripcion"] = "Armonizar"
				subDetallePost["dato"] = " "
				subDetallePost["activo"] = true
				subDetallePost["armonizacion_dato"] = string(armonizacion_data)

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
					request.LimpiezaRespuestaRefactor(res, &armonizacionUpdate)
				} else {
					return nil, errors.New("error del servicio ArmonizarInversion: La solicitud de registrar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}

			} else {
				return nil, errors.New("error del servicio ArmonizarInversion: La solicitud de registrar subgrupo contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
		}
		return res, nil
	} else {
		return nil, errors.New("error del servicio ArmonizarInversion: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func ActualizarMetaPlan(id string, index string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}

	_ = id
	json.Unmarshal(datos, &body)
	entrada = body["entrada"].(map[string]interface{})
	indexMetaSubProI := body["indexMetaSubPro"]
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
					return nil, errors.New("error del servicio ActualizarMetaPlan: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					return nil, errors.New("error del servicio ActualizarMetaPlan: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			return nil, errors.New("error del servicio ActualizarMetaPlan: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

		subgrupo_detalle = respuestaLimpia[0]
		if subgrupo_detalle["armonizacion_dato"] != nil {
			dato_armonizacion_str := subgrupo_detalle["armonizacion_dato"].(string)
			json.Unmarshal([]byte(dato_armonizacion_str), &armonizacion_dato)
			if armonizacion_dato[index] != nil {
				aux_armonizacion := armonizacion_dato[index].(map[string]interface{})
				aux := make(map[string]interface{})
				aux["idSubDetalleProI"] = aux_armonizacion["idSubDetalleProI"]
				aux["indexMetaSubProI"] = indexMetaSubProI
				aux["presupuesto_programado"] = aux_armonizacion["presupuesto_programado"]
				armonizacion_dato[index] = aux
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
			return nil, errors.New("error del servicio ActualizarMetaPlan: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	}
	return entrada, nil
}

func AllMetasPlan(id string) (interface{}, error) {
	var res map[string]interface{}
	var hijos []map[string]interface{}
	var tabla map[string]interface{}
	var metas []map[string]interface{}
	var auxHijos []interface{}
	var data_source []map[string]interface{}
	inversionhelper.Limpia()
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &hijos)
		for i := 0; i < len(hijos); i++ {
			auxHijos = append(auxHijos, hijos[i]["_id"])
		}
		tabla = inversionhelper.GetTabla(auxHijos)
		metas = tabla["data_source"].([]map[string]interface{})
		for indexMeta := range metas {
			if metas[indexMeta]["activo"] == true {
				data_source = append(data_source, metas[indexMeta])
			}
		}
		return data_source, nil
	} else {
		return nil, errors.New("error del servicio AllMetasPlan: La solicitud de consultar metas del plan \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func InactivarMeta(id string, index string) (interface{}, error) {
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
		inversionhelper.GetSons(auxHijos, index)
		return tabla, nil
	} else {
		return nil, errors.New("error del servicio InactivarMeta: La solicitud de consultar metas del plan \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func ProgMagnitudesPlan(id string, index string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var body map[string]interface{}
	var subgrupo []map[string]interface{}
	var id_subgrupoDetalle string
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var magnitudesUpdate []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})

	json.Unmarshal(datos, &body)
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &subgrupo)
		subgrupoPost := make(map[string]interface{})
		subDetallePost := make(map[string]interface{})

		if len(subgrupo) > 0 {
			id_subgrupoDetalle = subgrupo[0]["_id"].(string)
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+id_subgrupoDetalle, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				if len(respuestaLimpia) > 0 {
					subgrupo_detalle = respuestaLimpia[0]
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

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &magnitudesUpdate)
					} else {
						return nil, errors.New("error del servicio ProgMagnitudesPlan: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
					}
				} else {
					subgrupo_detalle := make(map[string]interface{})
					magnitud := make(map[string]interface{})
					magnitud["index"] = index
					magnitud["dato"] = body
					magnitud["activo"] = true
					dato[index] = magnitud
					b, _ := json.Marshal(dato)
					str := string(b)
					subgrupo_detalle["dato"] = str

					subDetallePost["subgrupo_id"] = id_subgrupoDetalle
					subDetallePost["nombre"] = "Detalle Información Programación de Magnitudes y Presupuesto"
					subDetallePost["descripcion"] = "Magnitudes"
					subDetallePost["dato"] = subgrupo_detalle["dato"]
					subDetallePost["activo"] = true

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &magnitudesUpdate)
					} else {
						return nil, errors.New("error del servicio ProgMagnitudesPlan: La solicitud de registrar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
					}
				}

			} else {
				return nil, errors.New("error del servicio ProgMagnitudesPlan: La solicitud de obtener subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
		} else {
			subgrupoPost["nombre"] = "Programación Magnitudes y Prespuesto Plan de Inversión"
			subgrupoPost["descripcion"] = "Magnitudes"
			subgrupoPost["padre"] = id
			subgrupoPost["activo"] = true
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/", "POST", &respuesta, subgrupoPost); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
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

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
					request.LimpiezaRespuestaRefactor(res, &magnitudesUpdate)
				} else {
					return nil, errors.New("error del servicio ProgMagnitudesPlan: La solicitud de registrar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}

			} else {
				return nil, errors.New("error del servicio ProgMagnitudesPlan: La solicitud de registrar subgrupo contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
		}
		return res, nil
	} else {
		return nil, errors.New("error del servicio ProgMagnitudesPlan: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func MagnitudesProgramadas(id string, index string) (interface{}, error) {
	var res map[string]interface{}
	var subgrupo map[string]interface{}
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})
	var magnitud map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuestaLimpia)
		subgrupo = respuestaLimpia[0]
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+subgrupo["_id"].(string), &respuesta); err != nil {
			return nil, errors.New("error del servicio MagnitudesProgramadas: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
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
		return magnitud, nil
	} else {
		return nil, errors.New("error del servicio MagnitudesProgramadas: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func CrearGrupoMeta(datos []byte) (interface{}, error) {
	var respuesta map[string]interface{}
	var planFormato map[string]interface{}
	var parametros map[string]interface{}
	var respuestaPost map[string]interface{}
	var planSubgrupo map[string]interface{}
	json.Unmarshal(datos, &parametros)
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
			} else {
				return nil, errors.New("error del servicio CrearGrupoMeta: La solicitud de obtener subgrupo contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}

		} else {
			return nil, errors.New("error del servicio CrearGrupoMeta: La solicitud de crear un plan contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}

		return planSubgrupo, nil
	} else {
		return nil, errors.New("error del servicio CrearGrupoMeta: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func ActualizarActividadInv(id string, index string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}
	var resSubgrupo map[string]interface{}
	var subgrupo map[string]interface{}

	_ = id
	json.Unmarshal(datos, &body)
	entrada = body["entrada"].(map[string]interface{})
	indexMetaSubProI := body["indexMetaSubPro"]
	idDetalleFuentesPro := body["idDetalleFuentesPro"].(string)
	fuentesActividad := body["fuentesActividad"]
	ponderacionH := body["ponderacionH"]
	var dato_fuente []map[string]interface{}

	errGet := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+idDetalleFuentesPro, &resSubgrupo)
	if errGet == nil {
		request.LimpiezaRespuestaRefactor(resSubgrupo, &subgrupo)
		if subgrupo["dato"] != nil {

			dato_str := subgrupo["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato_fuente)
			for key := range dato_fuente {
				fuenteActividad := body["fuentesActividad"].([]interface{})
				for key2 := range fuenteActividad {
					if dato_fuente[key]["_id"] == fuenteActividad[key2].(map[string]interface{})["id"] {
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

			b, _ := json.Marshal(dato_fuente)
			str := string(b)
			subgrupo["dato"] = str
		}
		var resDetalle map[string]interface{}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+idDetalleFuentesPro, "PUT", &resDetalle, subgrupo); err != nil {
			return nil, errors.New("error del servicio ActualizarActividadInv: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}

	} else {
		return nil, errors.New("error del servicio ActualizarActividadInv: La solicitud de obtener subgrupo-detalle contiene un tipo de dato incorrecto o un parámetro inválido " + errGet.Error())
	}

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
					return nil, errors.New("error del servicio ActualizarActividadInv: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					return nil, errors.New("error del servicio ActualizarActividadInv: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			return nil, errors.New("error del servicio ActualizarActividadInv: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
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
			return nil, errors.New("error del servicio ActualizarActividadInv: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	}
	return entrada, nil
}

func ActualizarTablaActividad(id string, index string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}

	_ = id
	json.Unmarshal(datos, &body)
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
					return nil, errors.New("error del servicio ActualizarTablaActividad: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					return nil, errors.New("error del servicio ActualizarTablaActividad: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			return nil, errors.New("error del servicio ActualizarTablaActividad: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
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
			return nil, errors.New("error del servicio ActualizarTablaActividad: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	}
	return entrada, nil
}

func ActualizarPresupuestoMeta(id string, index string, datos []byte) (interface{}, error) {
	var res map[string]interface{}
	var entrada map[string]interface{}
	var body map[string]interface{}

	_ = id
	json.Unmarshal(datos, &body)
	entrada = body["entrada"].(map[string]interface{})
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
					return nil, errors.New("error del servicio ActualizarPresupuestoMeta: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)

				subgrupo_detalle = respuestaLimpia[0]
				if subgrupo_detalle["dato_plan"] != nil {
					meta := make(map[string]interface{})
					dato_plan_str := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_str), &dato_plan)
					for index_actividad := range dato_plan {
						if index_actividad == index {
							aux_actividad := dato_plan[index_actividad].(map[string]interface{})
							meta["index"] = index_actividad
							meta["dato"] = aux_actividad["dato"]
							meta["activo"] = aux_actividad["activo"]
							meta["observacion"] = element
							dato_plan[index_actividad] = meta
						}
					}
					b, _ := json.Marshal(dato_plan)
					str := string(b)
					subgrupo_detalle["dato_plan"] = str
				}

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
					return nil, errors.New("error del servicio ActualizarPresupuestoMeta: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
				}

			}
			continue
		}
		id_subgrupoDetalle = key
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+id_subgrupoDetalle, &respuesta); err != nil {
			return nil, errors.New("error del servicio ActualizarPresupuestoMeta: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
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
			return nil, errors.New("error del servicio ActualizarPresupuestoMeta: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	}
	return entrada, nil
}

func VerificarMagnitudesProgramadas(id string) (interface{}, error) {
	var res map[string]interface{}
	var subgrupo map[string]interface{}
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato := make(map[string]interface{})
	var magnitud int

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=descripcion:Magnitudes,activo:true,padre:"+id, &res); err == nil {
		request.LimpiezaRespuestaRefactor(res, &respuestaLimpia)
		subgrupo = respuestaLimpia[0]
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+subgrupo["_id"].(string), &respuesta); err != nil {
			return nil, errors.New("error del servicio VerificarMagnitudesProgramadas: La solicitud de obtener subgrupo-detalle \"key\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["dato"] != nil {

			dato_str := subgrupo_detalle["dato"].(string)
			json.Unmarshal([]byte(dato_str), &dato)
			magnitud = len(dato)
		}
		return magnitud, nil
	} else {
		return nil, errors.New("error del servicio VerificarMagnitudesProgramadas: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}
}

func VersionarPlanInv(id string) (interface{}, error) {
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
			return nil, errors.New("error del servicio VersionarPlanInv: La solicitud de versionar plan \"plan[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
		planVersionado = respuestaPost["Data"].(map[string]interface{})

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
						return nil, errors.New("error del servicio VersionarPlanInv: La solicitud de versionar actividades \"actividadPadre[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
					}

					actividadVersionada = respuestaPostActividad["Data"].(map[string]interface{})

					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+actividadPadre["_id"].(string), &respuestaHijos); err == nil {
						request.LimpiezaRespuestaRefactor(respuestaHijos, &hijos)
						formulacionhelper.VersionarHijos(hijos, actividadVersionada["_id"].(string))
					}

				}
			}
		} else {
			return nil, errors.New("error del servicio VersionarPlanInv: La solicitud de obtener plan contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
		}
	} else {
		return nil, errors.New("error del servicio VersionarPlanInv: La solicitud contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
	}

	return respuestaPost, nil
}
