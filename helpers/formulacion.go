package helpers

import (
	"encoding/json"
	"fmt"
	"sort"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

var DatavalidaT = []string{}
var datos_fuente []map[string]interface{}
var columnas_visibles []string

func Limpia() {
	DatavalidaT = []string{}
}

func LimpiaTabla() {
	datos_fuente = nil
	columnas_visibles = nil
}

func ClonarHijos(hijos []map[string]interface{}, padre string) {
	for i := 0; i < len(hijos); i++ {
		hijo := make(map[string]interface{})
		hijo["nombre"] = hijos[i]["nombre"]
		hijo["descripcion"] = hijos[i]["descripcion"]
		hijo["activo"] = hijos[i]["activo"]
		hijo["padre"] = padre
		hijo["bandera_tabla"] = hijos[i]["bandera_tabla"]

		var respuestaPost map[string]interface{}
		var respuestaLimpia map[string]interface{}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/registrar_nodo/", "POST", &respuestaPost, &hijo); err == nil {
			respuestaLimpia = respuestaPost["Data"].(map[string]interface{})

			var respuestaHijos map[string]interface{}
			var respuestaHijosDetalle map[string]interface{}
			var subHijos []map[string]interface{}
			var subHijosDetalle []map[string]interface{}

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+hijos[i]["_id"].(string), &respuestaHijosDetalle); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijosDetalle, &subHijosDetalle)
				ClonarHijosDetalle(subHijosDetalle, respuestaLimpia["_id"].(string))
			}

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+hijos[i]["_id"].(string), &respuestaHijos); err == nil {
				request.LimpiezaRespuestaRefactor(respuestaHijos, &subHijos)
				ClonarHijos(subHijos, respuestaLimpia["_id"].(string))
			}
		}
	}
}

func ClonarHijosDetalle(subHijosDetalle []map[string]interface{}, subgrupo_id string) {
	for i := 0; i < len(subHijosDetalle); i++ {
		hijoDetalle := make(map[string]interface{})
		hijoDetalle["nombre"] = subHijosDetalle[i]["nombre"]
		hijoDetalle["descripcion"] = subHijosDetalle[i]["descripcion"]
		hijoDetalle["subgrupo_id"] = subgrupo_id
		hijoDetalle["activo"] = subHijosDetalle[i]["activo"]
		hijoDetalle["dato"] = subHijosDetalle[i]["dato"]

		var respuestaPost map[string]interface{}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/", "POST", &respuestaPost, &hijoDetalle); err != nil {
			panic(err)
		}
	}
}

func ConstruirArbolFormulacion(hijos []map[string]interface{}, indice string) [][]map[string]interface{} {
	var arbol []map[string]interface{}
	var requeridos []map[string]interface{}
	armonizacion := make([]map[string]interface{}, 1)
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var resultado [][]map[string]interface{}
	var nodo map[string]interface{}

	for i := 0; i < len(hijos); i++ {
		if hijos[i]["activo"] == true {
			forkData := make(map[string]interface{})
			var identificador string
			forkData["id"] = hijos[i]["_id"]
			forkData["nombre"] = hijos[i]["nombre"]
			jsonString, _ := json.Marshal(hijos[i]["_id"])
			json.Unmarshal(jsonString, &identificador)

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador, &respuesta); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
				nodo = respuestaLimpia[0]

				if len(nodo) != 0 {
					var detalle map[string]interface{}
					dato_string := nodo["dato"].(string)
					json.Unmarshal([]byte(dato_string), &detalle)
					if (detalle["type"] != nil) && (detalle["required"] != nil) && (detalle["options"] == nil) {
						forkData["type"] = detalle["type"]
						forkData["required"] = detalle["required"]
					} else if (detalle["type"] != nil) && (detalle["required"] != nil) && (detalle["options"] != nil) {
						forkData["type"] = detalle["type"]
						forkData["required"] = detalle["required"]
						forkData["options"] = detalle["options"]
					} else {
						forkData["type"] = " "
						forkData["required"] = " "
					}
				}
			} else {
				panic(err)
			}

			if len(hijos[i]["hijos"].([]interface{})) > 0 {
				forkData["sub"] = make([]map[string]interface{}, len(ConsultarHijos(hijos[i]["hijos"].([]interface{}))))
				forkData["sub"] = ConsultarHijos(hijos[i]["hijos"].([]interface{}))
			} else {
				forkData["sub"] = ""
			}
			arbol = append(arbol, forkData)
			agregar(identificador)
		}
	}
	requeridos, armonizacion[0] = Convertir(DatavalidaT, indice)
	resultado = append(resultado, arbol)
	resultado = append(resultado, requeridos)
	resultado = append(resultado, armonizacion)
	return resultado
}

func ConsultarHijos(hijos []interface{}) (hijosArbol []map[string]interface{}) {
	var respuesta1 map[string]interface{}
	var respuesta2 map[string]interface{}
	var nodo map[string]interface{}
	var detalle []map[string]interface{}

	for _, hijo := range hijos {
		forkData := make(map[string]interface{})
		var identificador string
		hijoString := fmt.Sprintf("%v", hijo)
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+hijoString, &respuesta1); err != nil {
			panic(err)
		}
		request.LimpiezaRespuestaRefactor(respuesta1, &nodo)
		if nodo["activo"] == true {
			forkData["id"] = nodo["_id"]
			forkData["nombre"] = nodo["nombre"]
			jsonString, _ := json.Marshal(nodo["_id"])
			json.Unmarshal(jsonString, &identificador)

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+identificador, &respuesta2); err == nil {
				request.LimpiezaRespuestaRefactor(respuesta2, &detalle)
				if len(detalle) != 0 {
					var datoDetalle map[string]interface{}
					dato_string := fmt.Sprintf("%v", detalle[0]["dato"])
					json.Unmarshal([]byte(dato_string), &datoDetalle)

					if (datoDetalle["type"] != nil) && (datoDetalle["required"] != nil) && (datoDetalle["options"] == nil) {
						forkData["type"] = datoDetalle["type"]
						forkData["required"] = datoDetalle["required"]
					} else if (datoDetalle["type"] != nil) && (datoDetalle["required"] != nil) && (datoDetalle["options"] != nil) {
						forkData["type"] = datoDetalle["type"]
						forkData["required"] = datoDetalle["required"]
						forkData["options"] = datoDetalle["options"]
					} else {
						forkData["type"] = " "
						forkData["required"] = " "
					}
				}
			}
			if len(nodo["hijos"].([]interface{})) > 0 {
				if len(ConsultarHijos(nodo["hijos"].([]interface{}))) == 0 {
					forkData["sub"] = ""
				} else {
					forkData["sub"] = ConsultarHijos(nodo["hijos"].([]interface{}))
				}
			}
			hijosArbol = append(hijosArbol, forkData)
		}
		agregar(identificador)
	}
	return
}

func agregar(identificador string) {
	if !request.Contains(DatavalidaT, identificador) && identificador != "" {
		DatavalidaT = append(DatavalidaT, identificador)
	}
}

func Convertir(valido []string, indice string) ([]map[string]interface{}, map[string]interface{}) {
	var validadores []map[string]interface{}
	var respuesta map[string]interface{}
	var dato_armonizacion map[string]interface{}
	armonizacion := make(map[string]interface{})
	forkData := make(map[string]interface{})

	for _, dato := range valido {
		var subgrupo_detalle []map[string]interface{}
		var actividad map[string]interface{}
		var dato_plan map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+dato, &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &subgrupo_detalle)
			if len(subgrupo_detalle) > 0 {
				if subgrupo_detalle[0]["armonizacion_dato"] != nil {
					dato_armonizacion_string := subgrupo_detalle[0]["armonizacion_dato"].(string)
					json.Unmarshal([]byte(dato_armonizacion_string), &dato_armonizacion)
					variable_auxiliar := dato_armonizacion[indice]

					if variable_auxiliar != nil {
						armonizacion["armo"] = dato_armonizacion[indice].(map[string]interface{})["armonizacionPED"]
						armonizacion["armoPI"] = dato_armonizacion[indice].(map[string]interface{})["armonizacionPI"]
						armonizacion["fuentesActividad"] = dato_armonizacion[indice].(map[string]interface{})["fuentesActividad"]
						armonizacion["indexMetaSubProI"] = dato_armonizacion[indice].(map[string]interface{})["indexMetaSubProI"]
						armonizacion["ponderacionH"] = dato_armonizacion[indice].(map[string]interface{})["ponderacionH"]
						armonizacion["object"] = variable_auxiliar
					}
				}
				if subgrupo_detalle[0]["dato_plan"] != nil {
					dato_plan_string := subgrupo_detalle[0]["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_string), &dato_plan)

					if dato_plan[indice] == nil {
						forkData[dato] = ""
					} else {
						actividad = dato_plan[indice].(map[string]interface{})

						if dato != "" {
							forkData[dato] = actividad["dato"]

							if actividad["observacion"] != nil {
								keyObservacion := dato + "_o"
								forkData[keyObservacion] = ConsultarObservacion(actividad)
							} else {
								keyObservacion := dato + "_o"
								forkData[keyObservacion] = "Sin observación"
							}
						}
					}
				} else {
					forkData[dato] = ""
				}
			} else {
				forkData[dato] = ""
			}
		}
	}
	validadores = append(validadores, forkData)
	return validadores, armonizacion
}

func ConsultarObservacion(actividad map[string]interface{}) string {
	if actividad["observacion"] == nil {
		return ""
	} else {
		str := fmt.Sprintf("%v", actividad["observacion"])
		return str
	}
}

func RecorrerHijos(hijos []map[string]interface{}, indice string) {
	var subgrupo map[string]interface{}
	var respuesta map[string]interface{}

	for i := 0; i < len(hijos); i++ {
		if len(hijos[i]["hijos"].([]interface{})) != 0 {
			hijosSubgrupo := hijos[i]["hijos"].([]interface{})
			for j := 0; j < len(hijosSubgrupo); j++ {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+hijosSubgrupo[j].(string), &respuesta); err == nil {
					request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)
					if len(subgrupo["hijos"].([]interface{})) != 0 {
						RecorrerSubgrupos(subgrupo["hijos"].([]interface{}), indice)
					} else {
						DesactivarActividad(subgrupo["_id"].(string), indice)
					}
				} else {
					panic(map[string]interface{}{"funcion": "RecorrerHijos", "err": "Error obteniendo subgrupo \"subgrupo[\"_id\"].(string)\"", "status": "400", "log": err})
				}
			}
		} else {
			DesactivarActividad(hijos[i]["_id"].(string), indice)
		}
	}
}

func RecorrerSubgrupos(hijos []interface{}, indice string) {
	var respuesta map[string]interface{}
	var subgrupo map[string]interface{}

	for i := 0; i < len(hijos); i++ {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+hijos[i].(string), &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)
			if len(subgrupo["hijos"].([]interface{})) != 0 {
				RecorrerSubgrupos(subgrupo["hijos"].([]interface{}), indice)
			} else {
				DesactivarActividad(subgrupo["_id"].(string), indice)
			}
		} else {
			panic(map[string]interface{}{"funcion": "RecorrerSubgrupos", "err": "Error obteniendo subgrupo \"subgrupo[\"_id\"].(string)\"", "status": "400", "log": err})
		}
	}
}

func DesactivarActividad(subgrupo_identificador string, indice string) {
	var respuesta1 map[string]interface{}
	var respuesta2 map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupoDetalle map[string]interface{}
	var dato_plan map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+subgrupo_identificador, &respuesta2); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta2, &respuestaLimpia)
		subgrupoDetalle = respuestaLimpia[0]

		if subgrupoDetalle["dato_plan"] != nil {
			actividad := make(map[string]interface{})
			dato_plan_string := subgrupoDetalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_string), &dato_plan)

			for indice_actividad := range dato_plan {
				if indice_actividad == indice {
					auxiliar_actividad := dato_plan[indice_actividad].(map[string]interface{})
					actividad["index"] = indice_actividad
					actividad["dato"] = auxiliar_actividad["dato"]

					if auxiliar_actividad["observacion"] != nil {
						actividad["observacion"] = auxiliar_actividad["observacion"]
					}

					actividad["activo"] = false
					dato_plan[indice_actividad] = actividad
				}
			}
			auxiliar, _ := json.Marshal(dato_plan)
			str := string(auxiliar)
			subgrupoDetalle["dato_plan"] = str
		}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupoDetalle["_id"].(string), "PUT", &respuesta1, subgrupoDetalle); err != nil {
			panic(map[string]interface{}{"funcion": "DesactivarActividad(Helper)", "err": "Error obteniendo subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
	} else {
		panic(map[string]interface{}{"funcion": "DesactivarActividad(Helper)", "err": "Error obteniendo subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
	}
}

func ConsultarTabla(hijos []interface{}) map[string]interface{} {
	tabla := make(map[string]interface{})
	var respuesta map[string]interface{}
	var subgrupo map[string]interface{}

	for i := 0; i < len(hijos); i++ {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+hijos[i].(string), &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)

			if subgrupo["bandera_tabla"] == true {
				columnas_visibles = append(columnas_visibles, subgrupo["nombre"].(string))
				ConsultarActividadTabla(subgrupo)
			}

			if len(subgrupo["hijos"].([]interface{})) != 0 {
				var respuestaHijos map[string]interface{}
				var subgrupoHijo map[string]interface{}

				for j := 0; j < len(subgrupo["hijos"].([]interface{})); j++ {
					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["hijos"].([]interface{})[j].(string), &respuestaHijos); err == nil {
						request.LimpiezaRespuestaRefactor(respuestaHijos, &subgrupoHijo)
						if subgrupoHijo["bandera_tabla"] == true {
							columnas_visibles = append(columnas_visibles, subgrupoHijo["nombre"].(string))
							ConsultarActividadTabla(subgrupoHijo)
						}
					} else {
						panic(err)
					}
				}
			}
		} else {
			panic(err)
		}
	}
	tabla["displayed_columns"] = columnas_visibles
	tabla["data_source"] = datos_fuente
	return tabla
}

func ConsultarActividadTabla(subgrupo map[string]interface{}) {
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	var dato_plan map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+subgrupo["_id"].(string), &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if datos_fuente == nil {
			if subgrupo_detalle["dato_plan"] != nil {
				dato_plan_string := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_string), &dato_plan)

				for llave := range dato_plan {
					actividad := make(map[string]interface{})
					elemento := dato_plan[llave].(map[string]interface{})
					actividad["index"] = llave
					actividad[subgrupo["nombre"].(string)] = elemento["dato"]
					actividad["activo"] = elemento["activo"]
					datos_fuente = append(datos_fuente, actividad)
				}
			}
		} else {
			for i := 0; i < len(datos_fuente); i++ {
				if subgrupo_detalle["dato_plan"] != nil {
					var datos = datos_fuente[i]
					dato_plan_string := subgrupo_detalle["dato_plan"].(string)
					json.Unmarshal([]byte(dato_plan_string), &dato_plan)

					if dato_plan[datos["index"].(string)] != nil {
						elemento := dato_plan[datos["index"].(string)].(map[string]interface{})
						datos[subgrupo["nombre"].(string)] = elemento["dato"]
					}
				}
			}
		}
	}
	sort.SliceStable(datos_fuente, func(i, j int) bool {
		a, _ := strconv.Atoi(datos_fuente[i]["index"].(string))
		b, _ := strconv.Atoi(datos_fuente[j]["index"].(string))
		return a < b
	})
}

func ConsultarArmonizacion(identificador string) map[string]interface{} {
	var armonizacion map[string]interface{}
	var respuesta map[string]interface{}
	var subgrupo map[string]interface{}
	var recorrido []map[string]interface{}

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+identificador, &respuesta); err == nil {
		request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)
		recorrido = append(recorrido, subgrupo)

		for subgrupo != nil || subgrupo["padre"] != nil {
			var auxiliarRespuesta map[string]interface{}
			var auxiliarSubgrupo map[string]interface{}

			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["padre"].(string), &auxiliarRespuesta); err == nil {
				request.LimpiezaRespuestaRefactor(auxiliarRespuesta, &auxiliarSubgrupo)
				recorrido = append(recorrido, auxiliarSubgrupo)
			} else {
				panic(err)
			}
			subgrupo = auxiliarSubgrupo
		}
	} else {
		panic(err)
	}
	recorrido = recorrido[:len(recorrido)-1]
	armonizacion = ConsultarRama(recorrido)
	return armonizacion
}

func ConsultarRama(recorrido []map[string]interface{}) map[string]interface{} {
	var armonizacion map[string]interface{}

	for i := 0; i < len(recorrido); i++ {
		forkData := make(map[string]interface{})
		forkData["_id"] = recorrido[i]["_id"]
		forkData["nombre"] = recorrido[i]["nombre"]
		forkData["descripcion"] = recorrido[i]["descripcion"]
		forkData["activo"] = recorrido[i]["activo"]

		if armonizacion != nil {
			forkData["children"] = armonizacion
		}
		armonizacion = forkData
	}
	return armonizacion
}

func ConsultarIndexActividad(entrada map[string]interface{}) int {
	var respuesta map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	dato_plan := make(map[string]interface{})
	var maximoIndice = 0

	for llave, elemento := range entrada {
		if elemento != "" {
			if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+llave, &respuesta); err != nil {
				panic(map[string]interface{}{"funcion": "GuardarPlan", "err": "Error get subgrupo-detalle \"key\"", "status": "400", "log": err})
			}
			request.LimpiezaRespuestaRefactor(respuesta, &respuestaLimpia)
			subgrupo_detalle = respuestaLimpia[0]

			if subgrupo_detalle["dato_plan"] == nil {
				maximoIndice = 0
			} else {
				dato_plan_texto := subgrupo_detalle["dato_plan"].(string)
				json.Unmarshal([]byte(dato_plan_texto), &dato_plan)

				for llave2 := range dato_plan {
					if indice, err := strconv.Atoi(llave2); err == nil {
						if indice > maximoIndice {
							maximoIndice = indice
						}
					} else {
						panic(err)
					}
				}
			}
		}
	}
	return maximoIndice
}

func OrdenarVersiones(versiones []map[string]interface{}) []map[string]interface{} {
	var versionesOrdenadas []map[string]interface{}

	for i := range versiones {
		if versiones[i]["padre_plan_id"] == nil {
			versionesOrdenadas = append(versionesOrdenadas, versiones[i])
		}
	}

	for len(versionesOrdenadas) < len(versiones) {
		versionesOrdenadas = append(versionesOrdenadas, ConsultarVersionHija(versionesOrdenadas[len(versionesOrdenadas)-1]["_id"], versiones))
	}
	return versionesOrdenadas
}

func VersionarHijos(hijos []map[string]interface{}, padre string) {
	var respuestaPost map[string]interface{}
	var subgrupoVersionado map[string]interface{}

	for i := 0; i < len(hijos); i++ {
		hijo := make(map[string]interface{})
		hijo["nombre"] = hijos[i]["nombre"]
		hijo["descripcion"] = hijos[i]["descripcion"]
		hijo["activo"] = hijos[i]["activo"]
		hijo["padre"] = padre
		hijo["bandera_tabla"] = hijos[i]["bandera_tabla"]

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/registrar_nodo", "POST", &respuestaPost, hijo); err != nil {
			panic(map[string]interface{}{"funcion": "VersionarHijos", "err": "Error versionando subgrupo \"hijo[\"_id\"].(string)\"", "status": "400", "log": err})
		}
		subgrupoVersionado = respuestaPost["Data"].(map[string]interface{})
		var respuestaHijos map[string]interface{}
		var respuestaHijosDetalle map[string]interface{}
		var subHijos []map[string]interface{}
		var subHijosDetalle []map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+hijos[i]["_id"].(string), &respuestaHijosDetalle); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijosDetalle, &subHijosDetalle)
			VersionarHijosDetalle(subHijosDetalle, subgrupoVersionado["_id"].(string))
		} else {
			panic(err)
		}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/hijos/"+hijos[i]["_id"].(string), &respuestaHijos); err == nil {
			request.LimpiezaRespuestaRefactor(respuestaHijos, &subHijos)
			VersionarHijos(subHijos, subgrupoVersionado["_id"].(string))
		} else {
			panic(err)
		}
	}
}

func VersionarHijosDetalle(subHijosDetalle []map[string]interface{}, subgrupo_id string) {
	for i := 0; i < len(subHijosDetalle); i++ {
		hijoDetalle := make(map[string]interface{})
		hijoDetalle["nombre"] = subHijosDetalle[i]["nombre"]
		hijoDetalle["descripcion"] = subHijosDetalle[i]["descripcion"]
		hijoDetalle["subgrupo_id"] = subgrupo_id
		hijoDetalle["activo"] = subHijosDetalle[i]["activo"]
		hijoDetalle["dato"] = subHijosDetalle[i]["dato"]
		hijoDetalle["dato_plan"] = subHijosDetalle[i]["dato_plan"]
		hijoDetalle["armonizacion_dato"] = subHijosDetalle[i]["armonizacion_dato"]
		var respuestaPost map[string]interface{}

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle", "POST", &respuestaPost, hijoDetalle); err != nil {
			panic(map[string]interface{}{"funcion": "VersionarHijosDetalle", "err": "Error versionando subgrupo_detalle ", "status": "400", "log": err})
		}
	}
}

func VersionarIdentificaciones(identificaciones []map[string]interface{}, id string) {
	for i := 0; i < len(identificaciones); i++ {
		var aux map[string]interface{} = identificaciones[i]
		identificacion := make(map[string]interface{})
		var respuestaPost map[string]interface{}

		identificacion["nombre"] = aux["nombre"]
		identificacion["descripcion"] = aux["descripcion"]
		identificacion["plan_id"] = id
		identificacion["dato"] = aux["dato"]
		identificacion["tipo_identificacion_id"] = aux["tipo_identificacion_id"]
		identificacion["activo"] = aux["activo"]

		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/identificacion", "POST", &respuestaPost, identificacion); err != nil {
			panic(map[string]interface{}{"funcion": "VersionaIdentificaciones", "err": "Error versionando identificaciones", "status": "400", "log": err})
		}
	}
}

func ConsultarVersionHija(id interface{}, versiones []map[string]interface{}) map[string]interface{} {
	for i := range versiones {
		if versiones[i]["padre_plan_id"] == id {
			return versiones[i]
		}
	}
	return nil
}

func ConsultarHijosRubro(entrada []interface{}) []map[string]interface{} {
	var hojas []map[string]interface{}
	var respuesta map[string]interface{}
	var respuestaLimpia interface{}

	for i := 0; i < len(entrada); i++ {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanCuentasService")+"/arbol_rubro/"+entrada[i].(string), &respuesta); err == nil {
			if respuesta["Body"] != nil {
				respuestaLimpia = respuesta["Body"].(map[string]interface{})

				if len(respuestaLimpia.(map[string]interface{})["Hijos"].([]interface{})) != 0 {
					var aux = respuestaLimpia.(map[string]interface{})["Hijos"]
					hojas = append(hojas, ConsultarHijosRubro(aux.([]interface{}))...)
				} else {
					hojas = append(hojas, respuestaLimpia.(map[string]interface{}))
				}
			}
		} else {
			panic(map[string]interface{}{"funcion": "ConsultarHijosRubros", "err": "Error arbol_rubros", "status": "400", "log": err})
		}
	}
	return hojas
}

func VerificarDataIdentificaciones(identificaciones []map[string]interface{}, tipoUnidad string) bool {
	var bandera bool

	if tipoUnidad == "facultad" {
		for i := 0; i < len(identificaciones); i++ {
			identificacion := identificaciones[i]
			if identificacion["tipo_identificacion_id"] == "61897518f6fc97091727c3c3" {
				if identificacion["dato"] == "{}" {
					bandera = false
					break
				} else {
					bandera = true
				}
			}
			if identificacion["tipo_identificacion_id"] == "6184b3e6f6fc97850127bb68" {
				if identificacion["dato"] == "{}" {
					bandera = false
					break
				} else {
					bandera = true
				}
			}
		}
	} else if tipoUnidad == "unidad" {
		for i := 0; i < len(identificaciones); i++ {
			identificacion := identificaciones[i]
			if identificacion["tipo_identificacion_id"] == "6184b3e6f6fc97850127bb68" {
				if identificacion["dato"] == "{}" {
					bandera = false
					break
				} else {
					bandera = true
				}
			}
		}
	}
	return bandera
}

func FiltrarIdentificaciones(dato map[string]interface{}) []map[string]interface{} {
	datos := []map[string]interface{}{}
	for llave := range dato {
		elemento := dato[llave].(map[string]interface{})
		if elemento["activo"] == true {
			datos = append(datos, elemento)
		}
	}
	return datos
}
