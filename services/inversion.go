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
			fmt.Println(subgruposData, "respuesta")
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
						fmt.Println(datoMeta_str, "datoMetas_str")
						json.Unmarshal([]byte(datoMeta_str), &dato_metas)
						fmt.Println(dato_metas, "datoMetas")
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
				fmt.Println(subgrupo_detalle["armonizacion_dato"], "armonización dato")
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
				fmt.Println(subgrupo_detalle["dato_plan"], "dato_plan 2")

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
					fmt.Println(subgrupo_detalle["armonizacion_dato"], "armonización dato 2")
				}
			}
			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subgrupo_detalle); err != nil {
				return nil, errors.New("error del servicio GuardarMeta: La solicitud de actualizar subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\" contiene un tipo de dato incorrecto o un parámetro inválido " + err.Error())
			}
			request.LimpiezaRespuestaRefactor(res, &respuestaGuardado)
			fmt.Println(respuestaGuardado, "actividad Guardada")
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
		fmt.Println(res, "subgrupo")
		id_subgrupoDetalle = subgrupo[0]["_id"].(string)
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+id_subgrupoDetalle, &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]
			fmt.Println(subgrupo_detalle, "subgrupo Detalle")
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
			fmt.Println(id_subgrupoDetalle, "id_subgrupoDetalle")
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=activo:true,subgrupo_id:"+id_subgrupoDetalle, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
				fmt.Println(respuesta, "respuesta")
				fmt.Println(respuestaLimpia, "respuestaLimpia")
				if len(respuestaLimpia) > 0 {
					fmt.Println(respuesta, "subgrupoDetalleb PUT")
					subgrupo_detalle = respuestaLimpia[0]
					subDetallePost["subgrupo_id"] = id_subgrupoDetalle
					subDetallePost["fecha_creacion"] = subgrupo_detalle["fecha_creacion"]
					subDetallePost["nombre"] = "Detalle Información Armonización"
					subDetallePost["descripcion"] = "Armonizar"
					subDetallePost["dato"] = " "
					subDetallePost["activo"] = true
					subDetallePost["armonizacion_dato"] = string(armonizacion_data)
					fmt.Println(subDetallePost, "dataJSON")
					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &armonizacionUpdate)
						fmt.Println(armonizacionUpdate, "update911")
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
					fmt.Println(subDetallePost["armonizacion_dato"], "dataJSON")

					if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
						request.LimpiezaRespuestaRefactor(res, &armonizacionUpdate)
						fmt.Println(res, "update926")
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
				fmt.Println(respuestaLimpia, "respuesta subgrupo POST")
				subgrupo_detalle = respuestaLimpia[0]
				subDetallePost["subgrupo_id"] = subgrupo_detalle["_id"]
				subDetallePost["nombre"] = "Detalle Información Armonización"
				subDetallePost["descripcion"] = "Armonizar"
				subDetallePost["dato"] = " "
				subDetallePost["activo"] = true
				subDetallePost["armonizacion_dato"] = string(armonizacion_data)
				fmt.Println(subDetallePost["armonizacion_dato"], "dataJSON")

				if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &res, subDetallePost); err == nil {
					request.LimpiezaRespuestaRefactor(res, &armonizacionUpdate)
					fmt.Println(armonizacionUpdate, "update954")
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
					fmt.Println(dato_plan, "dato_plan")
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
				fmt.Println(res, "res 1058")

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
				fmt.Println(armonizacion_dato, "armonizacion_dato")
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
		fmt.Println(res, "res 1121")
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

// func funcion(id string, index string, body []byte) (interface{}, error) {
//
// 	return result, nil
// }
