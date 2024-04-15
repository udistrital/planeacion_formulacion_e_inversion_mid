package controllers

import (
	"bytes"
	"net/http"
	"testing"
)

func TestConsultarPlan(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/plan/613acf8edf020f82a056eb2b/0"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarPlan Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarPlan Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarPlan:", err.Error())
		t.Fail()
	}
}

func TestConsultarTodasActividades(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/actividad/6139894fdf020f41fc56e5af/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarTodasActividades Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarTodasActividades Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarTodasActividades:", err.Error())
		t.Fail()
	}
}

func TestConsultarIdentificaciones(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/identificacion/618dfa66f6fc976f4627d9d6/617b6630f6fc97b776279afa"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarIdentificaciones Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarIdentificaciones Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarIdentificaciones:", err.Error())
		t.Fail()
	}
}

func TestConsultarPlanVersiones(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/plan/versiones/pruebas/2020/Plan de acción Inversión 20202"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarPlanVersiones Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarPlanVersiones Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarPlanVersiones:", err.Error())
		t.Fail()
	}
}

func TestPonderacionActividades(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/actividad/ponderacion/64368f5aa280496519a44efc"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestPonderacionActividades Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestCPonderacionActividades Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestPonderacionActividades:", err.Error())
		t.Fail()
	}
}

// SE DEMORA MUCHO
/*func TestConsultarRubros(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/rubros"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarRubros Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarRubros Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarRubros:", err.Error())
		t.Fail()
	}
}*/

func TestConsultarUnidades(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/unidades"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarUnidades Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarUnidades Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarUnidades:", err.Error())
		t.Fail()
	}
}

func TestVinculacionTercero(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/tercero/59769"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestVinculacionTercero Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestVinculacionTercero Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestVinculacionTercero:", err.Error())
		t.Fail()
	}
}

func TestPlanes(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/planes/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestPlanes Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestPlanes Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestPlanes:", err.Error())
		t.Fail()
	}
}

func TestVerificarIdentificaciones(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/identificacion/verificacion/616f6911a985e921bca12e96"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestVerificarIdentificaciones Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestVerificarIdentificaciones Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error VerificarIdentificaciones:", err.Error())
		t.Fail()
	}
}

func TestClonarFormato(t *testing.T) {
	body := []byte(`{
		"dependencia_id": "PRB",
		"vigencia": "2024"
	}`)

	if response, err := http.Post("http://localhost:8082/v1/formulacion/formato/611e4a2dd403481fb638b6e9/clonar", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestClonarFormato Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestClonarFormato Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestClonarFormato:", err.Error())
		t.Fail()
	}
}

// NO SE ENCUENTRA EL FORMATO y ya no manda success
func TestConsultarArbolArmonizacion(t *testing.T) {
	body := []byte(`{}`)

	if response, err := http.Post("http://localhost:8082/v1/formulacion/arbol/armonizacion/611db9b4d403482fec38b637", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestConsultarArbolArmonizacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestConsultarArbolArmonizacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestConsultarArbolArmonizacion:", err.Error())
		t.Fail()
	}
}

func TestVersionarPlan(t *testing.T) {
	body := []byte(`{}`)

	if response, err := http.Post("http://localhost:8082/v1/formulacion/plan/61398379df020f786256e5a7/versionar", "application/json", bytes.NewBuffer(body)); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestVersionarPlan Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestVersionarPlan Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestVersionarPlan:", err.Error())
		t.Fail()
	}
}

// SE NECESITA EL JSON
func TestGuardarActividad(t *testing.T) {
	body := []byte(`{{
		"armo": "613991d2df020fd81056e5c8",
		"armoPI": "613991d2df020fd81056e5c8",
		"entrada": {"1":{"dato":"Prueba Meta segplan","index":1},"10":{"activo":false,"dato":"Prueba Meta segplan","index":"10"},"11":{"activo":true,"dato":"Prueba Meta segplan","index":11},"12":{"activo":true,"dato":"Prueba Meta segplan","index":12},"13":{"activo":true,"dato":"Prueba Meta segplan","index":13},"14":{"activo":false,"dato":"Prueba Meta segplan","index":"14"},"2":{"dato":"Prueba Meta segplan","index":"2"},"3":{"dato":"Prueba Meta segplan","index":"3"},"4":{"dato":"Prueba Meta segplan","index":4},"5":{"dato":"Prueba Meta segplan","index":5},"6":{"dato":"Prueba Meta segplan","index":6},"7":{"dato":"Prueba Meta segplan","index":7},"8":{"dato":"Prueba Meta segplan","index":8},"9":{"activo":false,"dato":"Prueba Meta segplan","index":"9"}}
	}}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:8082/v1/formulacion/actividad/", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestGuardarActividad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestGuardarActividad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

// SE NECESITA EL JSON
func TestActualizarActividad(t *testing.T) {
	body := []byte(`{{
		"armo": "613991d2df020fd81056e5c8",
		"armoPI": "613991d2df020fd81056e5c8",
		"entrada": {"1":{"dato":"Prueba Meta segplan","index":1},"10":{"activo":false,"dato":"Prueba Meta segplan","index":"10"},"11":{"activo":true,"dato":"Prueba Meta segplan","index":11},"12":{"activo":true,"dato":"Prueba Meta segplan","index":12},"13":{"activo":true,"dato":"Prueba Meta segplan","index":13},"14":{"activo":false,"dato":"Prueba Meta segplan","index":"14"},"2":{"dato":"Prueba Meta segplan","index":"2"},"3":{"dato":"Prueba Meta segplan","index":"3"},"4":{"dato":"Prueba Meta segplan","index":4},"5":{"dato":"Prueba Meta segplan","index":5},"6":{"dato":"Prueba Meta segplan","index":6},"7":{"dato":"Prueba Meta segplan","index":7},"8":{"dato":"Prueba Meta segplan","index":8},"9":{"activo":false,"dato":"Prueba Meta segplan","index":"9"}}
	}}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:8082/v1/formulacion/actividad/613991d2df020fd81056e5c8/1", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestActualizarActividad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestActualizarActividad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDesactivarActividad(t *testing.T) {
	body := []byte(`{}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:8082/v1/formulacion/actividad/613991d2df020fd81056e5c8/1/desactivar", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestDesactivarActividad Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestDesactivarActividad Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestGuardarIdentificacion(t *testing.T) {
	body := []byte(`{
		"nombre": "Identificación de Contratistas Plan de Acción de Funcionamiento 2022",
		"descripcion": "Identificación de Contratistas Plan de Acción de Funcionamiento 2022 OFICINA ASESORA DE ASUNTOS DISCIPLINARIOS",
		"plan_id": "616f6911a985e921bca12e96",
		"dato": "{}",
		"tipo_identificacion_id": "6184b3e6f6fc97850127bb68",
		"activo": true
	  }`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:8082/v1/formulacion/identificacion/616f6911a985e921bca12e96/6184b3e6f6fc97850127bb68", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestGuardarIdentificacion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestGuardarIdentificacion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestDesactivarIdentificacion(t *testing.T) {
	body := []byte(`{}`)

	if request, err := http.NewRequest(http.MethodPut, "http://localhost:8082/v1/formulacion/identificacion/616f6911a985e921bca12e96/6184b3e6f6fc97850127bb68/0", bytes.NewBuffer(body)); err == nil {
		client := &http.Client{}
		if response, err := client.Do(request); err == nil {
			if response.StatusCode != 200 {
				t.Error("Error TestDesactivarIdentificacion Se esperaba 200 y se obtuvo", response.StatusCode)
				t.Fail()
			} else {
				t.Log("TestDesactivarIdentificacion Finalizado Correctamente (OK)")
			}
		}
	} else {
		t.Error("Error al crear la solicitud PUT: ", err.Error())
		t.Fail()
	}
}

func TestPlanesEnFormulacion(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/planes_formulacion/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error TestPlanesEnFormulacion Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("TestPlanesEnFormulacion Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error TestPlanesEnFormulacion:", err.Error())
		t.Fail()
	}
}

func GetPlanesUnidadesComun(t *testing.T) {
	if response, err := http.Get("http://localhost:8082/v1/formulacion/planes_formulacion/"); err == nil {
		if response.StatusCode != 200 {
			t.Error("Error GetPlanesUnidadesComun Se esperaba 200 y se obtuvo", response.StatusCode)
			t.Fail()
		} else {
			t.Log("GetPlanesUnidadesComun Finalizado Correctamente (OK)")
		}
	} else {
		t.Error("Error GetPlanesUnidadesComun:", err.Error())
		t.Fail()
	}
}
