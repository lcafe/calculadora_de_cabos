package low_voltage

import (
	"github.com/lcafe/calculadora_de_cabos/internal/electrical/objects"
	"github.com/lcafe/calculadora_de_cabos/internal/electrical/objects/low_voltage"
)

var kTable = map[low_voltage.KTableKey]float64{
	{
		Material:       objects.ConductorMaterialCopper,
		Insulation:     objects.ConductorInsulationPVC,
		SectionOver300: false,
	}: 115,

	{
		Material:       objects.ConductorMaterialCopper,
		Insulation:     objects.ConductorInsulationPVC,
		SectionOver300: true,
	}: 103,

	{
		Material:       objects.ConductorMaterialCopper,
		Insulation:     objects.ConductorInsulationXLPE,
		SectionOver300: false,
	}: 143,
	{
		Material:       objects.ConductorMaterialCopper,
		Insulation:     objects.ConductorInsulationXLPE,
		SectionOver300: true,
	}: 143,
	{
		Material:       objects.ConductorMaterialCopper,
		Insulation:     objects.ConductorInsulationEPR,
		SectionOver300: false,
	}: 143,
	{
		Material:       objects.ConductorMaterialCopper,
		Insulation:     objects.ConductorInsulationEPR,
		SectionOver300: true,
	}: 143,

	{
		Material:       objects.ConductorMaterialAluminum,
		Insulation:     objects.ConductorInsulationPVC,
		SectionOver300: false,
	}: 76,
	{
		Material:       objects.ConductorMaterialAluminum,
		Insulation:     objects.ConductorInsulationPVC,
		SectionOver300: true,
	}: 68,

	{
		Material:       objects.ConductorMaterialAluminum,
		Insulation:     objects.ConductorInsulationXLPE,
		SectionOver300: false,
	}: 94,
	{
		Material:       objects.ConductorMaterialAluminum,
		Insulation:     objects.ConductorInsulationXLPE,
		SectionOver300: true,
	}: 94,
	{
		Material:       objects.ConductorMaterialAluminum,
		Insulation:     objects.ConductorInsulationEPR,
		SectionOver300: false,
	}: 94,
	{
		Material:       objects.ConductorMaterialAluminum,
		Insulation:     objects.ConductorInsulationEPR,
		SectionOver300: true,
	}: 94,
}
