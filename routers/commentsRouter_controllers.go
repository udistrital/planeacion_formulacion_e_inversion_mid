package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ActualizarActividad",
            Router: "/actualizar_actividad/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ClonarFormato",
            Router: "/clonar_formato/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarArbolArmonizacion",
            Router: "/consultar_arbol_armonizacion/:id/",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarIdentificaciones",
            Router: "/consultar_identificaciones/:id/:idTipo",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarPlan",
            Router: "/consultar_plan/:id/:index",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarPlanVersiones",
            Router: "/consultar_plan_versiones/:unidad/:vigencia/:nombre",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarRubros",
            Router: "/consultar_rubros",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarTodasActividades",
            Router: "/consultar_todas_actividades/:id/",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "ConsultarUnidades",
            Router: "/consultar_unidades",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "DesactivarActividad",
            Router: "/desactivar_actividad/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "DesactivarIdentificacion",
            Router: "/desactivar_identificacion/:id/:idTipo/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "GuardarActividad",
            Router: "/guardar_actividad/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "GuardarIdentificacion",
            Router: "/guardar_identificacion/:id/:idTipo",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "Planes",
            Router: "/planes",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "PonderacionActividades",
            Router: "/ponderacion_actividades/:plan",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "VerificarIdentificaciones",
            Router: "/verificar_identificaciones/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "VersionarPlan",
            Router: "/versionar_plan/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:FormulacionController"],
        beego.ControllerComments{
            Method: "VinculacionTercero",
            Router: "/vinculacion_tercero/:tercero_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ActualizarActividad",
            Router: "/actualizar_actividad/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ActualizarMetaPlan",
            Router: "/actualizar_meta/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ActualizarPresupuestoMeta",
            Router: "/actualizar_presupuesto_meta/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ActualizarProyectoGeneral",
            Router: "/actualizar_proyecto/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ActualizarSubgrupoDetalle",
            Router: "/actualizar_subgrupo_detalle/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ActualizarTablaActividad",
            Router: "/actualizar_tabla_actividad/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarTodasMetasPlan",
            Router: "/consulta_todas_metas/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarPlan",
            Router: "/consultar_informacion_plan/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarMetasProyecto",
            Router: "/consultar_metas_proyecto/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarPlanIdentificador",
            Router: "/consultar_plan/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "CrearGrupoMeta",
            Router: "/crear_grupo_meta",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "CrearPlan",
            Router: "/crearplan",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "GuardarDocumentos",
            Router: "/guardar_documentos",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "GuardarMeta",
            Router: "/guardar_meta/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "InactivarMeta",
            Router: "/inactivar_meta/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ProgramarMagnitudesPlan",
            Router: "/magnitudes/:id/:index",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarMagnitudesProgramadas",
            Router: "/magnitudes/:id/:indexMeta",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "AgregarProyecto",
            Router: "/proyecto",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarProyectoId",
            Router: "/proyecto/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "EditarProyecto",
            Router: "/proyecto/:id",
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "ConsultarTodosProyectos",
            Router: "/proyectos/:aplicativo_id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "VerificarMagnitudesProgramadas",
            Router: "/verificar_magnitudes/:id",
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"] = append(beego.GlobalControllerRouter["github.com/udistrital/planeacion_formulacion_mid/controllers:InversionController"],
        beego.ControllerComments{
            Method: "VersionarPlan",
            Router: "/versionar_plan/:id",
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
