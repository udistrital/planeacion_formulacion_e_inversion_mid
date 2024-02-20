package helpers

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/udistrital/utils_oas/request"
)

var validaDataT = []string{}

func LimpiaInversion() {
	datos_fuente = nil
	columnas_visibles = nil
	validaDataT = []string{}
}

func RegistrarInformacionComplementaria(identificadorProyecto string, informacionProyecto map[string]interface{}, nombreComplementario string) error {
	var respuestaSubgrupo map[string]interface{}
	informacionSubgrupo := map[string]interface{}{
		"activo":      true,
		"padre":       identificadorProyecto,
		"nombre":      nombreComplementario,
		"descripcion": informacionProyecto["codigo_proyecto"],
	}

	var err error
	if err = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo", "POST", &respuestaSubgrupo, informacionSubgrupo); err == nil {
		identificadorSubgrupo := respuestaSubgrupo["Data"].(map[string]interface{})["_id"].(string)
		detalle, _ := json.Marshal(informacionProyecto["data"])
		var respuestaDetalle map[string]interface{}
		subgrupoDetalle := map[string]interface{}{
			"activo":      true,
			"subgrupo_id": identificadorSubgrupo,
			"nombre":      nombreComplementario,
			"descripcion": informacionProyecto["codigo_proyecto"],
			"dato":        string(detalle),
		}

		if err = request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle", "POST", &respuestaDetalle, subgrupoDetalle); err != nil {
			return err
		}
	}
	return err
}

func ActualizarPresupuestoDisponible(informacionFuente []interface{}) {
	for _, fuente := range informacionFuente {
		var dataFuente map[string]interface{}

		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/fuentes-apropiacion/"+fuente.(map[string]interface{})["id"].(string), &dataFuente); err == nil {
			respuestaFuente := dataFuente["Data"].(map[string]interface{})
			var dataFuente map[string]interface{}
			respuestaFuente["presupuestoDisponible"] = fuente.(map[string]interface{})["presupuestoDisponible"]

			if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/fuentes-apropiacion/"+fuente.(map[string]interface{})["id"].(string), "PUT", &dataFuente, respuestaFuente); err != nil {
				panic(err)
			}
		} else {
			panic(err)
		}
	}
}

func ActualizarInformacionComplementariaDetalle(identificadorSubgrupo string, detalleData []interface{}) error {
	var respuestaSubgrupo map[string]interface{}
	var subgrupo map[string]interface{}

	var err error
	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+identificadorSubgrupo, &respuestaSubgrupo); err == nil {
		request.LimpiezaRespuestaRefactor(respuestaSubgrupo, &subgrupo)
		detalle, _ := json.Marshal(detalleData)
		subgrupo["dato"] = string(detalle)

		var respuestaDetalle map[string]interface{}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+identificadorSubgrupo, "PUT", &respuestaDetalle, subgrupo); err != nil {
			return err
		}
	}
	return err
}

func ConsultarDataProyectos(informacionProyecto map[string]interface{}) map[string]interface{} {
	consultaProyecto := make(map[string]interface{})
	var subgruposData map[string]interface{}
	var informacionSubgrupos []map[string]interface{}
	consultaProyecto["nombre_proyecto"] = informacionProyecto["nombre"]
	consultaProyecto["codigo_proyecto"] = informacionProyecto["descripcion"]
	consultaProyecto["fecha_creacion"] = informacionProyecto["fecha_creacion"]
	consultaProyecto["id"] = informacionProyecto["_id"]
	padreIdentificador := informacionProyecto["_id"].(string)

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo?query=padre:"+padreIdentificador, &subgruposData); err == nil {
		request.LimpiezaRespuestaRefactor(subgruposData, &informacionSubgrupos)

		for i := range informacionSubgrupos {
			var subgrupoDetalle map[string]interface{}
			var detalleSubgrupos []map[string]interface{}

			if strings.Contains(strings.ToLower(informacionSubgrupos[i]["nombre"].(string)), "fuentes") {
				if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle?query=subgrupo_id:"+informacionSubgrupos[i]["_id"].(string), &subgrupoDetalle); err == nil {
					request.LimpiezaRespuestaRefactor(subgrupoDetalle, &detalleSubgrupos)
					armonizacion_dato_str := detalleSubgrupos[0]["dato"].(string)
					var subgrupo_dato []map[string]interface{}
					json.Unmarshal([]byte(armonizacion_dato_str), &subgrupo_dato)
					consultaProyecto["subgrupo_id_fuentes"] = informacionSubgrupos[i]["_id"]
					consultaProyecto["fuentes"] = subgrupo_dato
					consultaProyecto["id_detalle_fuentes"] = detalleSubgrupos[0]["_id"]
				}
			}
		}
	}
	return consultaProyecto
}

func ConsultarHijosInversion(hijos []interface{}, indice string) {
	var respuesta map[string]interface{}
	var subgrupo map[string]interface{}

	for i := 0; i < len(hijos); i++ {
		if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+hijos[i].(string), &respuesta); err == nil {
			request.LimpiezaRespuestaRefactor(respuesta, &subgrupo)

			if subgrupo["bandera_tabla"] == true {
				BorrarMetas(subgrupo, indice)
			}

			if len(subgrupo["hijos"].([]interface{})) != 0 {
				var respuestaHijos map[string]interface{}
				var subgrupoHijo map[string]interface{}

				for j := 0; j < len(subgrupo["hijos"].([]interface{})); j++ {
					if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo/"+subgrupo["hijos"].([]interface{})[j].(string), &respuestaHijos); err == nil {
						request.LimpiezaRespuestaRefactor(respuestaHijos, &subgrupoHijo)

						if subgrupoHijo["bandera_tabla"] == true {
							columnas_visibles = append(columnas_visibles, subgrupoHijo["nombre"].(string))
							BorrarMetas(subgrupoHijo, indice)
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
}

func BorrarMetas(subgrupo map[string]interface{}, indice string) {
	var respuesta map[string]interface{}
	var respuestaJ map[string]interface{}
	var respuestaLimpia []map[string]interface{}
	var subgrupo_detalle map[string]interface{}
	var dato_plan map[string]interface{}
	actividad := make(map[string]interface{})

	if err := request.GetJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/detalle/"+subgrupo["_id"].(string), &respuestaJ); err == nil {
		request.LimpiezaRespuestaRefactor(respuestaJ, &respuestaLimpia)
		subgrupo_detalle = respuestaLimpia[0]

		if subgrupo_detalle["dato_plan"] != nil {
			dato_plan_str := subgrupo_detalle["dato_plan"].(string)
			json.Unmarshal([]byte(dato_plan_str), &dato_plan)

			for llave := range dato_plan {
				if llave == indice {
					elemento := dato_plan[llave].(map[string]interface{})
					actividad["index"] = elemento["index"]
					actividad["dato"] = elemento["dato"]
					actividad["activo"] = false
					if elemento["observacion"] != nil {
						actividad["observacion"] = elemento["observacion"]
					}
					dato_plan[llave] = actividad
				}
			}
			b, _ := json.Marshal(dato_plan)
			str := string(b)
			subgrupo_detalle["dato_plan"] = str
		}
		if err := request.SendJson("http://"+beego.AppConfig.String("PlanesService")+"/subgrupo-detalle/"+subgrupo_detalle["_id"].(string), "PUT", &respuesta, subgrupo_detalle); err != nil {
			panic(map[string]interface{}{"funcion": "BorrarMetas", "err": "Error actualizando subgrupo-detalle \"subgrupo_detalle[\"_id\"].(string)\"", "status": "400", "log": err})
		}
	}
}
