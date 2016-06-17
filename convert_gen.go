//
// This file is generated. To change the content of this file, please do not
// apply the change to this file because it will get overwritten. Instead,
// change xenapi.go and execute 'go generate'.
//

package xenAPI

import (
	"fmt"
	"github.com/amfranz/go-xmlrpc-client"
	"reflect"
	"strconv"
	"time"
)

var _ = fmt.Errorf
var _ = xmlrpc.NewClient
var _ = reflect.TypeOf
var _ = strconv.Atoi
var _ = time.UTC

func convertBondRefToBondRecordMapToGo(context string, input interface{}) (goMap map[BondRef]BondRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[BondRef]BondRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertBondRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertBondRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertDRTaskRefToDRTaskRecordMapToGo(context string, input interface{}) (goMap map[DRTaskRef]DRTaskRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[DRTaskRef]DRTaskRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertDRTaskRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertDRTaskRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertGPUGroupRefToGPUGroupRecordMapToGo(context string, input interface{}) (goMap map[GPUGroupRef]GPUGroupRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[GPUGroupRef]GPUGroupRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertGPUGroupRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertGPUGroupRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPBDRefToPBDRecordMapToGo(context string, input interface{}) (goMap map[PBDRef]PBDRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PBDRef]PBDRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPBDRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPBDRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPCIRefToPCIRecordMapToGo(context string, input interface{}) (goMap map[PCIRef]PCIRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PCIRef]PCIRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPCIRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPCIRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPGPURefToPGPURecordMapToGo(context string, input interface{}) (goMap map[PGPURef]PGPURecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PGPURef]PGPURecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPGPURefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPGPURecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPIFRefToPIFRecordMapToGo(context string, input interface{}) (goMap map[PIFRef]PIFRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PIFRef]PIFRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPIFRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPIFRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPIFMetricsRefToPIFMetricsRecordMapToGo(context string, input interface{}) (goMap map[PIFMetricsRef]PIFMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PIFMetricsRef]PIFMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPIFMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPIFMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSMRefToSMRecordMapToGo(context string, input interface{}) (goMap map[SMRef]SMRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SMRef]SMRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSMRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSRRefToSRRecordMapToGo(context string, input interface{}) (goMap map[SRRef]SRRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SRRef]SRRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSRRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSRRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVBDRefToVBDRecordMapToGo(context string, input interface{}) (goMap map[VBDRef]VBDRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VBDRef]VBDRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVBDRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVBDRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVBDMetricsRefToVBDMetricsRecordMapToGo(context string, input interface{}) (goMap map[VBDMetricsRef]VBDMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VBDMetricsRef]VBDMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVBDMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVBDMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVDIRefToSRRefMapToXen(context string, goMap map[VDIRef]SRRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVDIRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertSRRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVDIRefToVDIRecordMapToGo(context string, input interface{}) (goMap map[VDIRef]VDIRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VDIRef]VDIRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVDIRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVDIRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVGPURefToVGPURecordMapToGo(context string, input interface{}) (goMap map[VGPURef]VGPURecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VGPURef]VGPURecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVGPURefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVGPURecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVGPUTypeRefToVGPUTypeRecordMapToGo(context string, input interface{}) (goMap map[VGPUTypeRef]VGPUTypeRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VGPUTypeRef]VGPUTypeRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVGPUTypeRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVGPUTypeRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVGPUTypeRefToIntMapToGo(context string, input interface{}) (goMap map[VGPUTypeRef]int, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VGPUTypeRef]int, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVGPUTypeRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertIntToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVIFRefToVIFRecordMapToGo(context string, input interface{}) (goMap map[VIFRef]VIFRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VIFRef]VIFRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVIFRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVIFRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVIFRefToNetworkRefMapToXen(context string, goMap map[VIFRef]NetworkRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVIFRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertNetworkRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVIFRefToStringMapToGo(context string, input interface{}) (goMap map[VIFRef]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VIFRef]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVIFRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVIFRefToStringMapToXen(context string, goMap map[VIFRef]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVIFRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVIFMetricsRefToVIFMetricsRecordMapToGo(context string, input interface{}) (goMap map[VIFMetricsRef]VIFMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VIFMetricsRef]VIFMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVIFMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVIFMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVLANRefToVLANRecordMapToGo(context string, input interface{}) (goMap map[VLANRef]VLANRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VLANRef]VLANRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVLANRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVLANRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToStringToStringMapMapToGo(context string, input interface{}) (goMap map[VMRef]map[string]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMRef]map[string]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToStringMapToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToVMRecordMapToGo(context string, input interface{}) (goMap map[VMRef]VMRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMRef]VMRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToStringSetMapToGo(context string, input interface{}) (goMap map[VMRef][]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMRef][]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringSetToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMRefToStringMapToXen(context string, goMap map[VMRef]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertVMRefToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertVMPPRefToVMPPRecordMapToGo(context string, input interface{}) (goMap map[VMPPRef]VMPPRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMPPRef]VMPPRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMPPRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMPPRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMApplianceRefToVMApplianceRecordMapToGo(context string, input interface{}) (goMap map[VMApplianceRef]VMApplianceRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMApplianceRef]VMApplianceRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMApplianceRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMApplianceRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMGuestMetricsRefToVMGuestMetricsRecordMapToGo(context string, input interface{}) (goMap map[VMGuestMetricsRef]VMGuestMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMGuestMetricsRef]VMGuestMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMGuestMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMGuestMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertVMMetricsRefToVMMetricsRecordMapToGo(context string, input interface{}) (goMap map[VMMetricsRef]VMMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMMetricsRef]VMMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertVMMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertVMMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertBlobRefToBlobRecordMapToGo(context string, input interface{}) (goMap map[BlobRef]BlobRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[BlobRef]BlobRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertBlobRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertBlobRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertConsoleRefToConsoleRecordMapToGo(context string, input interface{}) (goMap map[ConsoleRef]ConsoleRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[ConsoleRef]ConsoleRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertConsoleRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertConsoleRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertCrashdumpRefToCrashdumpRecordMapToGo(context string, input interface{}) (goMap map[CrashdumpRef]CrashdumpRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[CrashdumpRef]CrashdumpRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertCrashdumpRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertCrashdumpRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertEnumVMOperationsToStringMapToGo(context string, input interface{}) (goMap map[VMOperations]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[VMOperations]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertEnumVMOperationsToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertEnumVMOperationsToStringMapToXen(context string, goMap map[VMOperations]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertEnumVMOperationsToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertHostRefToHostRecordMapToGo(context string, input interface{}) (goMap map[HostRef]HostRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostRef]HostRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostRefToStringSetMapToGo(context string, input interface{}) (goMap map[HostRef][]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostRef][]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringSetToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostCPURefToHostCPURecordMapToGo(context string, input interface{}) (goMap map[HostCPURef]HostCPURecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostCPURef]HostCPURecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostCPURefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostCPURecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostCrashdumpRefToHostCrashdumpRecordMapToGo(context string, input interface{}) (goMap map[HostCrashdumpRef]HostCrashdumpRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostCrashdumpRef]HostCrashdumpRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostCrashdumpRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostCrashdumpRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostMetricsRefToHostMetricsRecordMapToGo(context string, input interface{}) (goMap map[HostMetricsRef]HostMetricsRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostMetricsRef]HostMetricsRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostMetricsRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostMetricsRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertHostPatchRefToHostPatchRecordMapToGo(context string, input interface{}) (goMap map[HostPatchRef]HostPatchRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[HostPatchRef]HostPatchRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertHostPatchRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertHostPatchRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertIntToFloatMapToGo(context string, input interface{}) (goMap map[int]float64, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[int]float64, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertIntToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertFloatToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertIntToIntMapToGo(context string, input interface{}) (goMap map[int]int, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[int]int, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertIntToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertIntToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertIntToStringSetMapToGo(context string, input interface{}) (goMap map[int][]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[int][]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertIntToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringSetToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertMessageRefToMessageRecordMapToGo(context string, input interface{}) (goMap map[MessageRef]MessageRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[MessageRef]MessageRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertMessageRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertMessageRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertNetworkRefToNetworkRecordMapToGo(context string, input interface{}) (goMap map[NetworkRef]NetworkRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[NetworkRef]NetworkRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertNetworkRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertNetworkRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPoolRefToPoolRecordMapToGo(context string, input interface{}) (goMap map[PoolRef]PoolRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PoolRef]PoolRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPoolRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPoolRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertPoolPatchRefToPoolPatchRecordMapToGo(context string, input interface{}) (goMap map[PoolPatchRef]PoolPatchRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[PoolPatchRef]PoolPatchRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertPoolPatchRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertPoolPatchRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertRoleRefToRoleRecordMapToGo(context string, input interface{}) (goMap map[RoleRef]RoleRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[RoleRef]RoleRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertRoleRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertRoleRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertSecretRefToSecretRecordMapToGo(context string, input interface{}) (goMap map[SecretRef]SecretRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SecretRef]SecretRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSecretRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSecretRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToBlobRefMapToGo(context string, input interface{}) (goMap map[string]BlobRef, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]BlobRef, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertBlobRefToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToBlobRefMapToXen(context string, goMap map[string]BlobRef) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertBlobRefToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumHostAllowedOperationsMapToGo(context string, input interface{}) (goMap map[string]HostAllowedOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]HostAllowedOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumHostAllowedOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumNetworkOperationsMapToGo(context string, input interface{}) (goMap map[string]NetworkOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]NetworkOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumNetworkOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumNetworkOperationsMapToXen(context string, goMap map[string]NetworkOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumNetworkOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumPoolAllowedOperationsMapToGo(context string, input interface{}) (goMap map[string]PoolAllowedOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]PoolAllowedOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumPoolAllowedOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumStorageOperationsMapToGo(context string, input interface{}) (goMap map[string]StorageOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]StorageOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumStorageOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumTaskAllowedOperationsMapToGo(context string, input interface{}) (goMap map[string]TaskAllowedOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]TaskAllowedOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumTaskAllowedOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVbdOperationsMapToGo(context string, input interface{}) (goMap map[string]VbdOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VbdOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVbdOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVbdOperationsMapToXen(context string, goMap map[string]VbdOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVbdOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVdiOperationsMapToGo(context string, input interface{}) (goMap map[string]VdiOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VdiOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVdiOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVdiOperationsMapToXen(context string, goMap map[string]VdiOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVdiOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVifOperationsMapToGo(context string, input interface{}) (goMap map[string]VifOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VifOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVifOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVifOperationsMapToXen(context string, goMap map[string]VifOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVifOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVMApplianceOperationMapToGo(context string, input interface{}) (goMap map[string]VMApplianceOperation, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VMApplianceOperation, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVMApplianceOperationToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVMApplianceOperationMapToXen(context string, goMap map[string]VMApplianceOperation) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVMApplianceOperationToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToEnumVMOperationsMapToGo(context string, input interface{}) (goMap map[string]VMOperations, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]VMOperations, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertEnumVMOperationsToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToEnumVMOperationsMapToXen(context string, goMap map[string]VMOperations) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertEnumVMOperationsToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertStringToIntMapToGo(context string, input interface{}) (goMap map[string]int, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]int, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertIntToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToStringMapToGo(context string, input interface{}) (goMap map[string]string, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[string]string, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertStringToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertStringToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertStringToStringMapToXen(context string, goMap map[string]string) (xenMap xmlrpc.Struct, err error) {
	xenMap = make(xmlrpc.Struct)
	for goKey, goValue := range goMap {
		keyContext := fmt.Sprintf("%s[%s]", context, goKey)
		xenKey, err := convertStringToXen(keyContext, goKey)
		if err != nil {
			return xenMap, err
		}
		xenValue, err := convertStringToXen(keyContext, goValue)
		if err != nil {
			return xenMap, err
		}
		xenMap[xenKey] = xenValue
	}
	return
}

func convertSubjectRefToSubjectRecordMapToGo(context string, input interface{}) (goMap map[SubjectRef]SubjectRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[SubjectRef]SubjectRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertSubjectRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertSubjectRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertTaskRefToTaskRecordMapToGo(context string, input interface{}) (goMap map[TaskRef]TaskRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[TaskRef]TaskRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertTaskRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertTaskRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertTunnelRefToTunnelRecordMapToGo(context string, input interface{}) (goMap map[TunnelRef]TunnelRecord, err error) {
	xenMap, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
	goMap = make(map[TunnelRef]TunnelRecord, len(xenMap))
	for xenKey, xenValue := range xenMap {
		keyContext := fmt.Sprintf("%s[%s]", context, xenKey)
		goKey, err := convertTunnelRefToGo(keyContext, xenKey)
		if err != nil {
			return goMap, err
		}
		goValue, err := convertTunnelRecordToGo(keyContext, xenValue)
		if err != nil {
			return goMap, err
		}
		goMap[goKey] = goValue
	}
	return
}

func convertBondRecordToGo(context string, input interface{}) (record BondRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  masterValue, ok := rpcStruct["master"]
	if ok {
  	record.Master, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "master"), masterValue)
		if err != nil {
			return
		}
	}
  slavesValue, ok := rpcStruct["slaves"]
	if ok {
  	record.Slaves, err = convertPIFRefSetToGo(fmt.Sprintf("%s.%s", context, "slaves"), slavesValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  primarySlaveValue, ok := rpcStruct["primary_slave"]
	if ok {
  	record.PrimarySlave, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "primary_slave"), primarySlaveValue)
		if err != nil {
			return
		}
	}
  modeValue, ok := rpcStruct["mode"]
	if ok {
  	record.Mode, err = convertEnumBondModeToGo(fmt.Sprintf("%s.%s", context, "mode"), modeValue)
		if err != nil {
			return
		}
	}
  propertiesValue, ok := rpcStruct["properties"]
	if ok {
  	record.Properties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "properties"), propertiesValue)
		if err != nil {
			return
		}
	}
  linksUpValue, ok := rpcStruct["links_up"]
	if ok {
  	record.LinksUp, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "links_up"), linksUpValue)
		if err != nil {
			return
		}
	}
	return
}

func convertBondRefSetToGo(context string, input interface{}) (slice []BondRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]BondRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertBondRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertBondRefToGo(context string, input interface{}) (ref BondRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = BondRef(value)
	}
	return
}

func convertBondRefToXen(context string, ref BondRef) (string, error) {
	return string(ref), nil
}

func convertDRTaskRecordToGo(context string, input interface{}) (record DRTaskRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  introducedSRsValue, ok := rpcStruct["introduced_SRs"]
	if ok {
  	record.IntroducedSRs, err = convertSRRefSetToGo(fmt.Sprintf("%s.%s", context, "introduced_SRs"), introducedSRsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertDRTaskRefSetToGo(context string, input interface{}) (slice []DRTaskRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]DRTaskRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertDRTaskRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertDRTaskRefToGo(context string, input interface{}) (ref DRTaskRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = DRTaskRef(value)
	}
	return
}

func convertDRTaskRefToXen(context string, ref DRTaskRef) (string, error) {
	return string(ref), nil
}

func convertGPUGroupRecordToGo(context string, input interface{}) (record GPUGroupRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  pgpusValue, ok := rpcStruct["PGPUs"]
	if ok {
  	record.PGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "PGPUs"), pgpusValue)
		if err != nil {
			return
		}
	}
  vgpusValue, ok := rpcStruct["VGPUs"]
	if ok {
  	record.VGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "VGPUs"), vgpusValue)
		if err != nil {
			return
		}
	}
  gpuTypesValue, ok := rpcStruct["GPU_types"]
	if ok {
  	record.GPUTypes, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "GPU_types"), gpuTypesValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  allocationAlgorithmValue, ok := rpcStruct["allocation_algorithm"]
	if ok {
  	record.AllocationAlgorithm, err = convertEnumAllocationAlgorithmToGo(fmt.Sprintf("%s.%s", context, "allocation_algorithm"), allocationAlgorithmValue)
		if err != nil {
			return
		}
	}
  supportedVGPUTypesValue, ok := rpcStruct["supported_VGPU_types"]
	if ok {
  	record.SupportedVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "supported_VGPU_types"), supportedVGPUTypesValue)
		if err != nil {
			return
		}
	}
  enabledVGPUTypesValue, ok := rpcStruct["enabled_VGPU_types"]
	if ok {
  	record.EnabledVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_VGPU_types"), enabledVGPUTypesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertGPUGroupRefSetToGo(context string, input interface{}) (slice []GPUGroupRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]GPUGroupRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertGPUGroupRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertGPUGroupRefToGo(context string, input interface{}) (ref GPUGroupRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = GPUGroupRef(value)
	}
	return
}

func convertGPUGroupRefToXen(context string, ref GPUGroupRef) (string, error) {
	return string(ref), nil
}

func convertPBDRecordToGo(context string, input interface{}) (record PBDRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  srValue, ok := rpcStruct["SR"]
	if ok {
  	record.SR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "SR"), srValue)
		if err != nil {
			return
		}
	}
  deviceConfigValue, ok := rpcStruct["device_config"]
	if ok {
  	record.DeviceConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "device_config"), deviceConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPBDRecordToXen(context string, record PBDRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["host"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "host"), record.Host)
  if err != nil {
		return
	}
  rpcStruct["SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "SR"), record.SR)
  if err != nil {
		return
	}
  rpcStruct["device_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "device_config"), record.DeviceConfig)
  if err != nil {
		return
	}
  rpcStruct["currently_attached"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "currently_attached"), record.CurrentlyAttached)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertPBDRefSetToGo(context string, input interface{}) (slice []PBDRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PBDRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPBDRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPBDRefToGo(context string, input interface{}) (ref PBDRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PBDRef(value)
	}
	return
}

func convertPBDRefToXen(context string, ref PBDRef) (string, error) {
	return string(ref), nil
}

func convertPCIRecordToGo(context string, input interface{}) (record PCIRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  classNameValue, ok := rpcStruct["class_name"]
	if ok {
  	record.ClassName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "class_name"), classNameValue)
		if err != nil {
			return
		}
	}
  vendorNameValue, ok := rpcStruct["vendor_name"]
	if ok {
  	record.VendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_name"), vendorNameValue)
		if err != nil {
			return
		}
	}
  deviceNameValue, ok := rpcStruct["device_name"]
	if ok {
  	record.DeviceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device_name"), deviceNameValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  pciIDValue, ok := rpcStruct["pci_id"]
	if ok {
  	record.PciID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "pci_id"), pciIDValue)
		if err != nil {
			return
		}
	}
  dependenciesValue, ok := rpcStruct["dependencies"]
	if ok {
  	record.Dependencies, err = convertPCIRefSetToGo(fmt.Sprintf("%s.%s", context, "dependencies"), dependenciesValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  subsystemVendorNameValue, ok := rpcStruct["subsystem_vendor_name"]
	if ok {
  	record.SubsystemVendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subsystem_vendor_name"), subsystemVendorNameValue)
		if err != nil {
			return
		}
	}
  subsystemDeviceNameValue, ok := rpcStruct["subsystem_device_name"]
	if ok {
  	record.SubsystemDeviceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subsystem_device_name"), subsystemDeviceNameValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPCIRefSetToGo(context string, input interface{}) (slice []PCIRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PCIRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPCIRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPCIRefSetToXen(context string, slice []PCIRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPCIRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertPCIRefToGo(context string, input interface{}) (ref PCIRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PCIRef(value)
	}
	return
}

func convertPCIRefToXen(context string, ref PCIRef) (string, error) {
	return string(ref), nil
}

func convertPGPURecordToGo(context string, input interface{}) (record PGPURecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  pciValue, ok := rpcStruct["PCI"]
	if ok {
  	record.PCI, err = convertPCIRefToGo(fmt.Sprintf("%s.%s", context, "PCI"), pciValue)
		if err != nil {
			return
		}
	}
  gpuGroupValue, ok := rpcStruct["GPU_group"]
	if ok {
  	record.GPUGroup, err = convertGPUGroupRefToGo(fmt.Sprintf("%s.%s", context, "GPU_group"), gpuGroupValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  supportedVGPUTypesValue, ok := rpcStruct["supported_VGPU_types"]
	if ok {
  	record.SupportedVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "supported_VGPU_types"), supportedVGPUTypesValue)
		if err != nil {
			return
		}
	}
  enabledVGPUTypesValue, ok := rpcStruct["enabled_VGPU_types"]
	if ok {
  	record.EnabledVGPUTypes, err = convertVGPUTypeRefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_VGPU_types"), enabledVGPUTypesValue)
		if err != nil {
			return
		}
	}
  residentVGPUsValue, ok := rpcStruct["resident_VGPUs"]
	if ok {
  	record.ResidentVGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "resident_VGPUs"), residentVGPUsValue)
		if err != nil {
			return
		}
	}
  supportedVGPUMaxCapacitiesValue, ok := rpcStruct["supported_VGPU_max_capacities"]
	if ok {
  	record.SupportedVGPUMaxCapacities, err = convertVGPUTypeRefToIntMapToGo(fmt.Sprintf("%s.%s", context, "supported_VGPU_max_capacities"), supportedVGPUMaxCapacitiesValue)
		if err != nil {
			return
		}
	}
  dom0AccessValue, ok := rpcStruct["dom0_access"]
	if ok {
  	record.Dom0Access, err = convertEnumPgpuDom0AccessToGo(fmt.Sprintf("%s.%s", context, "dom0_access"), dom0AccessValue)
		if err != nil {
			return
		}
	}
  isSystemDisplayDeviceValue, ok := rpcStruct["is_system_display_device"]
	if ok {
  	record.IsSystemDisplayDevice, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_system_display_device"), isSystemDisplayDeviceValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPGPURefSetToGo(context string, input interface{}) (slice []PGPURef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PGPURef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPGPURefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPGPURefToGo(context string, input interface{}) (ref PGPURef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PGPURef(value)
	}
	return
}

func convertPGPURefToXen(context string, ref PGPURef) (string, error) {
	return string(ref), nil
}

func convertPIFRecordToGo(context string, input interface{}) (record PIFRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  networkValue, ok := rpcStruct["network"]
	if ok {
  	record.Network, err = convertNetworkRefToGo(fmt.Sprintf("%s.%s", context, "network"), networkValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  macValue, ok := rpcStruct["MAC"]
	if ok {
  	record.MAC, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "MAC"), macValue)
		if err != nil {
			return
		}
	}
  mtuValue, ok := rpcStruct["MTU"]
	if ok {
  	record.MTU, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "MTU"), mtuValue)
		if err != nil {
			return
		}
	}
  vlanValue, ok := rpcStruct["VLAN"]
	if ok {
  	record.VLAN, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VLAN"), vlanValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok {
  	record.Metrics, err = convertPIFMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  physicalValue, ok := rpcStruct["physical"]
	if ok {
  	record.Physical, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "physical"), physicalValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  ipConfigurationModeValue, ok := rpcStruct["ip_configuration_mode"]
	if ok {
  	record.IPConfigurationMode, err = convertEnumIPConfigurationModeToGo(fmt.Sprintf("%s.%s", context, "ip_configuration_mode"), ipConfigurationModeValue)
		if err != nil {
			return
		}
	}
  ipValue, ok := rpcStruct["IP"]
	if ok {
  	record.IP, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "IP"), ipValue)
		if err != nil {
			return
		}
	}
  netmaskValue, ok := rpcStruct["netmask"]
	if ok {
  	record.Netmask, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "netmask"), netmaskValue)
		if err != nil {
			return
		}
	}
  gatewayValue, ok := rpcStruct["gateway"]
	if ok {
  	record.Gateway, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "gateway"), gatewayValue)
		if err != nil {
			return
		}
	}
  dnsValue, ok := rpcStruct["DNS"]
	if ok {
  	record.DNS, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "DNS"), dnsValue)
		if err != nil {
			return
		}
	}
  bondSlaveOfValue, ok := rpcStruct["bond_slave_of"]
	if ok {
  	record.BondSlaveOf, err = convertBondRefToGo(fmt.Sprintf("%s.%s", context, "bond_slave_of"), bondSlaveOfValue)
		if err != nil {
			return
		}
	}
  bondMasterOfValue, ok := rpcStruct["bond_master_of"]
	if ok {
  	record.BondMasterOf, err = convertBondRefSetToGo(fmt.Sprintf("%s.%s", context, "bond_master_of"), bondMasterOfValue)
		if err != nil {
			return
		}
	}
  vlanMasterOfValue, ok := rpcStruct["VLAN_master_of"]
	if ok {
  	record.VLANMasterOf, err = convertVLANRefToGo(fmt.Sprintf("%s.%s", context, "VLAN_master_of"), vlanMasterOfValue)
		if err != nil {
			return
		}
	}
  vlanSlaveOfValue, ok := rpcStruct["VLAN_slave_of"]
	if ok {
  	record.VLANSlaveOf, err = convertVLANRefSetToGo(fmt.Sprintf("%s.%s", context, "VLAN_slave_of"), vlanSlaveOfValue)
		if err != nil {
			return
		}
	}
  managementValue, ok := rpcStruct["management"]
	if ok {
  	record.Management, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "management"), managementValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  disallowUnplugValue, ok := rpcStruct["disallow_unplug"]
	if ok {
  	record.DisallowUnplug, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "disallow_unplug"), disallowUnplugValue)
		if err != nil {
			return
		}
	}
  tunnelAccessPIFOfValue, ok := rpcStruct["tunnel_access_PIF_of"]
	if ok {
  	record.TunnelAccessPIFOf, err = convertTunnelRefSetToGo(fmt.Sprintf("%s.%s", context, "tunnel_access_PIF_of"), tunnelAccessPIFOfValue)
		if err != nil {
			return
		}
	}
  tunnelTransportPIFOfValue, ok := rpcStruct["tunnel_transport_PIF_of"]
	if ok {
  	record.TunnelTransportPIFOf, err = convertTunnelRefSetToGo(fmt.Sprintf("%s.%s", context, "tunnel_transport_PIF_of"), tunnelTransportPIFOfValue)
		if err != nil {
			return
		}
	}
  ipv6ConfigurationModeValue, ok := rpcStruct["ipv6_configuration_mode"]
	if ok {
  	record.Ipv6ConfigurationMode, err = convertEnumIpv6ConfigurationModeToGo(fmt.Sprintf("%s.%s", context, "ipv6_configuration_mode"), ipv6ConfigurationModeValue)
		if err != nil {
			return
		}
	}
  ipv6Value, ok := rpcStruct["IPv6"]
	if ok {
  	record.IPv6, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "IPv6"), ipv6Value)
		if err != nil {
			return
		}
	}
  ipv6GatewayValue, ok := rpcStruct["ipv6_gateway"]
	if ok {
  	record.Ipv6Gateway, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ipv6_gateway"), ipv6GatewayValue)
		if err != nil {
			return
		}
	}
  primaryAddressTypeValue, ok := rpcStruct["primary_address_type"]
	if ok {
  	record.PrimaryAddressType, err = convertEnumPrimaryAddressTypeToGo(fmt.Sprintf("%s.%s", context, "primary_address_type"), primaryAddressTypeValue)
		if err != nil {
			return
		}
	}
  managedValue, ok := rpcStruct["managed"]
	if ok {
  	record.Managed, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "managed"), managedValue)
		if err != nil {
			return
		}
	}
  propertiesValue, ok := rpcStruct["properties"]
	if ok {
  	record.Properties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "properties"), propertiesValue)
		if err != nil {
			return
		}
	}
  capabilitiesValue, ok := rpcStruct["capabilities"]
	if ok {
  	record.Capabilities, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "capabilities"), capabilitiesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPIFRefSetToGo(context string, input interface{}) (slice []PIFRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PIFRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPIFRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPIFRefSetToXen(context string, slice []PIFRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPIFRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertPIFRefToGo(context string, input interface{}) (ref PIFRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PIFRef(value)
	}
	return
}

func convertPIFRefToXen(context string, ref PIFRef) (string, error) {
	return string(ref), nil
}

func convertPIFMetricsRecordToGo(context string, input interface{}) (record PIFMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  ioReadKbsValue, ok := rpcStruct["io_read_kbs"]
	if ok {
  	record.IoReadKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_read_kbs"), ioReadKbsValue)
		if err != nil {
			return
		}
	}
  ioWriteKbsValue, ok := rpcStruct["io_write_kbs"]
	if ok {
  	record.IoWriteKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_write_kbs"), ioWriteKbsValue)
		if err != nil {
			return
		}
	}
  carrierValue, ok := rpcStruct["carrier"]
	if ok {
  	record.Carrier, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "carrier"), carrierValue)
		if err != nil {
			return
		}
	}
  vendorIDValue, ok := rpcStruct["vendor_id"]
	if ok {
  	record.VendorID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_id"), vendorIDValue)
		if err != nil {
			return
		}
	}
  vendorNameValue, ok := rpcStruct["vendor_name"]
	if ok {
  	record.VendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_name"), vendorNameValue)
		if err != nil {
			return
		}
	}
  deviceIDValue, ok := rpcStruct["device_id"]
	if ok {
  	record.DeviceID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device_id"), deviceIDValue)
		if err != nil {
			return
		}
	}
  deviceNameValue, ok := rpcStruct["device_name"]
	if ok {
  	record.DeviceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device_name"), deviceNameValue)
		if err != nil {
			return
		}
	}
  speedValue, ok := rpcStruct["speed"]
	if ok {
  	record.Speed, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "speed"), speedValue)
		if err != nil {
			return
		}
	}
  duplexValue, ok := rpcStruct["duplex"]
	if ok {
  	record.Duplex, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "duplex"), duplexValue)
		if err != nil {
			return
		}
	}
  pciBusPathValue, ok := rpcStruct["pci_bus_path"]
	if ok {
  	record.PciBusPath, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "pci_bus_path"), pciBusPathValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPIFMetricsRefSetToGo(context string, input interface{}) (slice []PIFMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PIFMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPIFMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPIFMetricsRefToGo(context string, input interface{}) (ref PIFMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PIFMetricsRef(value)
	}
	return
}

func convertPIFMetricsRefToXen(context string, ref PIFMetricsRef) (string, error) {
	return string(ref), nil
}

func convertSMRecordToGo(context string, input interface{}) (record SMRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok {
  	record.Type, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  vendorValue, ok := rpcStruct["vendor"]
	if ok {
  	record.Vendor, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor"), vendorValue)
		if err != nil {
			return
		}
	}
  copyrightValue, ok := rpcStruct["copyright"]
	if ok {
  	record.Copyright, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "copyright"), copyrightValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  requiredAPIVersionValue, ok := rpcStruct["required_api_version"]
	if ok {
  	record.RequiredAPIVersion, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "required_api_version"), requiredAPIVersionValue)
		if err != nil {
			return
		}
	}
  configurationValue, ok := rpcStruct["configuration"]
	if ok {
  	record.Configuration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "configuration"), configurationValue)
		if err != nil {
			return
		}
	}
  capabilitiesValue, ok := rpcStruct["capabilities"]
	if ok {
  	record.Capabilities, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "capabilities"), capabilitiesValue)
		if err != nil {
			return
		}
	}
  featuresValue, ok := rpcStruct["features"]
	if ok {
  	record.Features, err = convertStringToIntMapToGo(fmt.Sprintf("%s.%s", context, "features"), featuresValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  driverFilenameValue, ok := rpcStruct["driver_filename"]
	if ok {
  	record.DriverFilename, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "driver_filename"), driverFilenameValue)
		if err != nil {
			return
		}
	}
  requiredClusterStackValue, ok := rpcStruct["required_cluster_stack"]
	if ok {
  	record.RequiredClusterStack, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "required_cluster_stack"), requiredClusterStackValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSMRefSetToGo(context string, input interface{}) (slice []SMRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SMRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSMRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSMRefToGo(context string, input interface{}) (ref SMRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SMRef(value)
	}
	return
}

func convertSMRefToXen(context string, ref SMRef) (string, error) {
	return string(ref), nil
}

func convertSRRecordToGo(context string, input interface{}) (record SRRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumStorageOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumStorageOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vdisValue, ok := rpcStruct["VDIs"]
	if ok {
  	record.VDIs, err = convertVDIRefSetToGo(fmt.Sprintf("%s.%s", context, "VDIs"), vdisValue)
		if err != nil {
			return
		}
	}
  pbdsValue, ok := rpcStruct["PBDs"]
	if ok {
  	record.PBDs, err = convertPBDRefSetToGo(fmt.Sprintf("%s.%s", context, "PBDs"), pbdsValue)
		if err != nil {
			return
		}
	}
  virtualAllocationValue, ok := rpcStruct["virtual_allocation"]
	if ok {
  	record.VirtualAllocation, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "virtual_allocation"), virtualAllocationValue)
		if err != nil {
			return
		}
	}
  physicalUtilisationValue, ok := rpcStruct["physical_utilisation"]
	if ok {
  	record.PhysicalUtilisation, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "physical_utilisation"), physicalUtilisationValue)
		if err != nil {
			return
		}
	}
  physicalSizeValue, ok := rpcStruct["physical_size"]
	if ok {
  	record.PhysicalSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "physical_size"), physicalSizeValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok {
  	record.Type, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  contentTypeValue, ok := rpcStruct["content_type"]
	if ok {
  	record.ContentType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "content_type"), contentTypeValue)
		if err != nil {
			return
		}
	}
  sharedValue, ok := rpcStruct["shared"]
	if ok {
  	record.Shared, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "shared"), sharedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  smConfigValue, ok := rpcStruct["sm_config"]
	if ok {
  	record.SmConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "sm_config"), smConfigValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  localCacheEnabledValue, ok := rpcStruct["local_cache_enabled"]
	if ok {
  	record.LocalCacheEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "local_cache_enabled"), localCacheEnabledValue)
		if err != nil {
			return
		}
	}
  introducedByValue, ok := rpcStruct["introduced_by"]
	if ok {
  	record.IntroducedBy, err = convertDRTaskRefToGo(fmt.Sprintf("%s.%s", context, "introduced_by"), introducedByValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSRRefSetToGo(context string, input interface{}) (slice []SRRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SRRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSRRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSRRefSetToXen(context string, slice []SRRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSRRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertSRRefToGo(context string, input interface{}) (ref SRRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SRRef(value)
	}
	return
}

func convertSRRefToXen(context string, ref SRRef) (string, error) {
	return string(ref), nil
}

func convertVBDRecordToGo(context string, input interface{}) (record VBDRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumVbdOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumVbdOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  vdiValue, ok := rpcStruct["VDI"]
	if ok {
  	record.VDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "VDI"), vdiValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  userdeviceValue, ok := rpcStruct["userdevice"]
	if ok {
  	record.Userdevice, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "userdevice"), userdeviceValue)
		if err != nil {
			return
		}
	}
  bootableValue, ok := rpcStruct["bootable"]
	if ok {
  	record.Bootable, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "bootable"), bootableValue)
		if err != nil {
			return
		}
	}
  modeValue, ok := rpcStruct["mode"]
	if ok {
  	record.Mode, err = convertEnumVbdModeToGo(fmt.Sprintf("%s.%s", context, "mode"), modeValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok {
  	record.Type, err = convertEnumVbdTypeToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  unpluggableValue, ok := rpcStruct["unpluggable"]
	if ok {
  	record.Unpluggable, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "unpluggable"), unpluggableValue)
		if err != nil {
			return
		}
	}
  storageLockValue, ok := rpcStruct["storage_lock"]
	if ok {
  	record.StorageLock, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "storage_lock"), storageLockValue)
		if err != nil {
			return
		}
	}
  emptyValue, ok := rpcStruct["empty"]
	if ok {
  	record.Empty, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "empty"), emptyValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  statusCodeValue, ok := rpcStruct["status_code"]
	if ok {
  	record.StatusCode, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "status_code"), statusCodeValue)
		if err != nil {
			return
		}
	}
  statusDetailValue, ok := rpcStruct["status_detail"]
	if ok {
  	record.StatusDetail, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "status_detail"), statusDetailValue)
		if err != nil {
			return
		}
	}
  runtimePropertiesValue, ok := rpcStruct["runtime_properties"]
	if ok {
  	record.RuntimeProperties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "runtime_properties"), runtimePropertiesValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmTypeValue, ok := rpcStruct["qos_algorithm_type"]
	if ok {
  	record.QosAlgorithmType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), qosAlgorithmTypeValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmParamsValue, ok := rpcStruct["qos_algorithm_params"]
	if ok {
  	record.QosAlgorithmParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), qosAlgorithmParamsValue)
		if err != nil {
			return
		}
	}
  qosSupportedAlgorithmsValue, ok := rpcStruct["qos_supported_algorithms"]
	if ok {
  	record.QosSupportedAlgorithms, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), qosSupportedAlgorithmsValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok {
  	record.Metrics, err = convertVBDMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVBDRecordToXen(context string, record VBDRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVbdOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVbdOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["VDI"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "VDI"), record.VDI)
  if err != nil {
		return
	}
  rpcStruct["device"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "device"), record.Device)
  if err != nil {
		return
	}
  rpcStruct["userdevice"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "userdevice"), record.Userdevice)
  if err != nil {
		return
	}
  rpcStruct["bootable"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "bootable"), record.Bootable)
  if err != nil {
		return
	}
  rpcStruct["mode"], err = convertEnumVbdModeToXen(fmt.Sprintf("%s.%s", context, "mode"), record.Mode)
  if err != nil {
		return
	}
  rpcStruct["type"], err = convertEnumVbdTypeToXen(fmt.Sprintf("%s.%s", context, "type"), record.Type)
  if err != nil {
		return
	}
  rpcStruct["unpluggable"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "unpluggable"), record.Unpluggable)
  if err != nil {
		return
	}
  rpcStruct["storage_lock"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "storage_lock"), record.StorageLock)
  if err != nil {
		return
	}
  rpcStruct["empty"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "empty"), record.Empty)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["currently_attached"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "currently_attached"), record.CurrentlyAttached)
  if err != nil {
		return
	}
  rpcStruct["status_code"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "status_code"), record.StatusCode)
  if err != nil {
		return
	}
  rpcStruct["status_detail"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "status_detail"), record.StatusDetail)
  if err != nil {
		return
	}
  rpcStruct["runtime_properties"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "runtime_properties"), record.RuntimeProperties)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_type"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), record.QosAlgorithmType)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), record.QosAlgorithmParams)
  if err != nil {
		return
	}
  rpcStruct["qos_supported_algorithms"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), record.QosSupportedAlgorithms)
  if err != nil {
		return
	}
  rpcStruct["metrics"], err = convertVBDMetricsRefToXen(fmt.Sprintf("%s.%s", context, "metrics"), record.Metrics)
  if err != nil {
		return
	}
	return
}

func convertVBDRefSetToGo(context string, input interface{}) (slice []VBDRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VBDRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVBDRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVBDRefSetToXen(context string, slice []VBDRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVBDRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVBDRefToGo(context string, input interface{}) (ref VBDRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VBDRef(value)
	}
	return
}

func convertVBDRefToXen(context string, ref VBDRef) (string, error) {
	return string(ref), nil
}

func convertVBDMetricsRecordToGo(context string, input interface{}) (record VBDMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  ioReadKbsValue, ok := rpcStruct["io_read_kbs"]
	if ok {
  	record.IoReadKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_read_kbs"), ioReadKbsValue)
		if err != nil {
			return
		}
	}
  ioWriteKbsValue, ok := rpcStruct["io_write_kbs"]
	if ok {
  	record.IoWriteKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_write_kbs"), ioWriteKbsValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVBDMetricsRefSetToGo(context string, input interface{}) (slice []VBDMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VBDMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVBDMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVBDMetricsRefToGo(context string, input interface{}) (ref VBDMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VBDMetricsRef(value)
	}
	return
}

func convertVBDMetricsRefToXen(context string, ref VBDMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVDIRecordToGo(context string, input interface{}) (record VDIRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumVdiOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumVdiOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  srValue, ok := rpcStruct["SR"]
	if ok {
  	record.SR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "SR"), srValue)
		if err != nil {
			return
		}
	}
  vbdsValue, ok := rpcStruct["VBDs"]
	if ok {
  	record.VBDs, err = convertVBDRefSetToGo(fmt.Sprintf("%s.%s", context, "VBDs"), vbdsValue)
		if err != nil {
			return
		}
	}
  crashDumpsValue, ok := rpcStruct["crash_dumps"]
	if ok {
  	record.CrashDumps, err = convertCrashdumpRefSetToGo(fmt.Sprintf("%s.%s", context, "crash_dumps"), crashDumpsValue)
		if err != nil {
			return
		}
	}
  virtualSizeValue, ok := rpcStruct["virtual_size"]
	if ok {
  	record.VirtualSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "virtual_size"), virtualSizeValue)
		if err != nil {
			return
		}
	}
  physicalUtilisationValue, ok := rpcStruct["physical_utilisation"]
	if ok {
  	record.PhysicalUtilisation, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "physical_utilisation"), physicalUtilisationValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok {
  	record.Type, err = convertEnumVdiTypeToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  sharableValue, ok := rpcStruct["sharable"]
	if ok {
  	record.Sharable, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "sharable"), sharableValue)
		if err != nil {
			return
		}
	}
  readOnlyValue, ok := rpcStruct["read_only"]
	if ok {
  	record.ReadOnly, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "read_only"), readOnlyValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  storageLockValue, ok := rpcStruct["storage_lock"]
	if ok {
  	record.StorageLock, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "storage_lock"), storageLockValue)
		if err != nil {
			return
		}
	}
  locationValue, ok := rpcStruct["location"]
	if ok {
  	record.Location, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "location"), locationValue)
		if err != nil {
			return
		}
	}
  managedValue, ok := rpcStruct["managed"]
	if ok {
  	record.Managed, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "managed"), managedValue)
		if err != nil {
			return
		}
	}
  missingValue, ok := rpcStruct["missing"]
	if ok {
  	record.Missing, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "missing"), missingValue)
		if err != nil {
			return
		}
	}
  parentValue, ok := rpcStruct["parent"]
	if ok {
  	record.Parent, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "parent"), parentValue)
		if err != nil {
			return
		}
	}
  xenstoreDataValue, ok := rpcStruct["xenstore_data"]
	if ok {
  	record.XenstoreData, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "xenstore_data"), xenstoreDataValue)
		if err != nil {
			return
		}
	}
  smConfigValue, ok := rpcStruct["sm_config"]
	if ok {
  	record.SmConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "sm_config"), smConfigValue)
		if err != nil {
			return
		}
	}
  isASnapshotValue, ok := rpcStruct["is_a_snapshot"]
	if ok {
  	record.IsASnapshot, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), isASnapshotValue)
		if err != nil {
			return
		}
	}
  snapshotOfValue, ok := rpcStruct["snapshot_of"]
	if ok {
  	record.SnapshotOf, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "snapshot_of"), snapshotOfValue)
		if err != nil {
			return
		}
	}
  snapshotsValue, ok := rpcStruct["snapshots"]
	if ok {
  	record.Snapshots, err = convertVDIRefSetToGo(fmt.Sprintf("%s.%s", context, "snapshots"), snapshotsValue)
		if err != nil {
			return
		}
	}
  snapshotTimeValue, ok := rpcStruct["snapshot_time"]
	if ok {
  	record.SnapshotTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "snapshot_time"), snapshotTimeValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  allowCachingValue, ok := rpcStruct["allow_caching"]
	if ok {
  	record.AllowCaching, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "allow_caching"), allowCachingValue)
		if err != nil {
			return
		}
	}
  onBootValue, ok := rpcStruct["on_boot"]
	if ok {
  	record.OnBoot, err = convertEnumOnBootToGo(fmt.Sprintf("%s.%s", context, "on_boot"), onBootValue)
		if err != nil {
			return
		}
	}
  metadataOfPoolValue, ok := rpcStruct["metadata_of_pool"]
	if ok {
  	record.MetadataOfPool, err = convertPoolRefToGo(fmt.Sprintf("%s.%s", context, "metadata_of_pool"), metadataOfPoolValue)
		if err != nil {
			return
		}
	}
  metadataLatestValue, ok := rpcStruct["metadata_latest"]
	if ok {
  	record.MetadataLatest, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "metadata_latest"), metadataLatestValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVDIRecordToXen(context string, record VDIRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVdiOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVdiOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "SR"), record.SR)
  if err != nil {
		return
	}
  rpcStruct["VBDs"], err = convertVBDRefSetToXen(fmt.Sprintf("%s.%s", context, "VBDs"), record.VBDs)
  if err != nil {
		return
	}
  rpcStruct["crash_dumps"], err = convertCrashdumpRefSetToXen(fmt.Sprintf("%s.%s", context, "crash_dumps"), record.CrashDumps)
  if err != nil {
		return
	}
  rpcStruct["virtual_size"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "virtual_size"), record.VirtualSize)
  if err != nil {
		return
	}
  rpcStruct["physical_utilisation"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "physical_utilisation"), record.PhysicalUtilisation)
  if err != nil {
		return
	}
  rpcStruct["type"], err = convertEnumVdiTypeToXen(fmt.Sprintf("%s.%s", context, "type"), record.Type)
  if err != nil {
		return
	}
  rpcStruct["sharable"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "sharable"), record.Sharable)
  if err != nil {
		return
	}
  rpcStruct["read_only"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "read_only"), record.ReadOnly)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["storage_lock"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "storage_lock"), record.StorageLock)
  if err != nil {
		return
	}
  rpcStruct["location"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "location"), record.Location)
  if err != nil {
		return
	}
  rpcStruct["managed"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "managed"), record.Managed)
  if err != nil {
		return
	}
  rpcStruct["missing"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "missing"), record.Missing)
  if err != nil {
		return
	}
  rpcStruct["parent"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "parent"), record.Parent)
  if err != nil {
		return
	}
  rpcStruct["xenstore_data"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "xenstore_data"), record.XenstoreData)
  if err != nil {
		return
	}
  rpcStruct["sm_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "sm_config"), record.SmConfig)
  if err != nil {
		return
	}
  rpcStruct["is_a_snapshot"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), record.IsASnapshot)
  if err != nil {
		return
	}
  rpcStruct["snapshot_of"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "snapshot_of"), record.SnapshotOf)
  if err != nil {
		return
	}
  rpcStruct["snapshots"], err = convertVDIRefSetToXen(fmt.Sprintf("%s.%s", context, "snapshots"), record.Snapshots)
  if err != nil {
		return
	}
  rpcStruct["snapshot_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "snapshot_time"), record.SnapshotTime)
  if err != nil {
		return
	}
  rpcStruct["tags"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "tags"), record.Tags)
  if err != nil {
		return
	}
  rpcStruct["allow_caching"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "allow_caching"), record.AllowCaching)
  if err != nil {
		return
	}
  rpcStruct["on_boot"], err = convertEnumOnBootToXen(fmt.Sprintf("%s.%s", context, "on_boot"), record.OnBoot)
  if err != nil {
		return
	}
  rpcStruct["metadata_of_pool"], err = convertPoolRefToXen(fmt.Sprintf("%s.%s", context, "metadata_of_pool"), record.MetadataOfPool)
  if err != nil {
		return
	}
  rpcStruct["metadata_latest"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "metadata_latest"), record.MetadataLatest)
  if err != nil {
		return
	}
	return
}

func convertVDIRefSetToGo(context string, input interface{}) (slice []VDIRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VDIRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVDIRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVDIRefSetToXen(context string, slice []VDIRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVDIRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVDIRefToGo(context string, input interface{}) (ref VDIRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VDIRef(value)
	}
	return
}

func convertVDIRefToXen(context string, ref VDIRef) (string, error) {
	return string(ref), nil
}

func convertVGPURecordToGo(context string, input interface{}) (record VGPURecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  gpuGroupValue, ok := rpcStruct["GPU_group"]
	if ok {
  	record.GPUGroup, err = convertGPUGroupRefToGo(fmt.Sprintf("%s.%s", context, "GPU_group"), gpuGroupValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok {
  	record.Type, err = convertVGPUTypeRefToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  residentOnValue, ok := rpcStruct["resident_on"]
	if ok {
  	record.ResidentOn, err = convertPGPURefToGo(fmt.Sprintf("%s.%s", context, "resident_on"), residentOnValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVGPURefSetToGo(context string, input interface{}) (slice []VGPURef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VGPURef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPURefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVGPURefSetToXen(context string, slice []VGPURef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPURefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVGPURefToGo(context string, input interface{}) (ref VGPURef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VGPURef(value)
	}
	return
}

func convertVGPURefToXen(context string, ref VGPURef) (string, error) {
	return string(ref), nil
}

func convertVGPUTypeRecordToGo(context string, input interface{}) (record VGPUTypeRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vendorNameValue, ok := rpcStruct["vendor_name"]
	if ok {
  	record.VendorName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor_name"), vendorNameValue)
		if err != nil {
			return
		}
	}
  modelNameValue, ok := rpcStruct["model_name"]
	if ok {
  	record.ModelName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "model_name"), modelNameValue)
		if err != nil {
			return
		}
	}
  framebufferSizeValue, ok := rpcStruct["framebuffer_size"]
	if ok {
  	record.FramebufferSize, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "framebuffer_size"), framebufferSizeValue)
		if err != nil {
			return
		}
	}
  maxHeadsValue, ok := rpcStruct["max_heads"]
	if ok {
  	record.MaxHeads, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "max_heads"), maxHeadsValue)
		if err != nil {
			return
		}
	}
  maxResolutionXValue, ok := rpcStruct["max_resolution_x"]
	if ok {
  	record.MaxResolutionX, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "max_resolution_x"), maxResolutionXValue)
		if err != nil {
			return
		}
	}
  maxResolutionYValue, ok := rpcStruct["max_resolution_y"]
	if ok {
  	record.MaxResolutionY, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "max_resolution_y"), maxResolutionYValue)
		if err != nil {
			return
		}
	}
  supportedOnPGPUsValue, ok := rpcStruct["supported_on_PGPUs"]
	if ok {
  	record.SupportedOnPGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "supported_on_PGPUs"), supportedOnPGPUsValue)
		if err != nil {
			return
		}
	}
  enabledOnPGPUsValue, ok := rpcStruct["enabled_on_PGPUs"]
	if ok {
  	record.EnabledOnPGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_on_PGPUs"), enabledOnPGPUsValue)
		if err != nil {
			return
		}
	}
  vgpusValue, ok := rpcStruct["VGPUs"]
	if ok {
  	record.VGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "VGPUs"), vgpusValue)
		if err != nil {
			return
		}
	}
  supportedOnGPUGroupsValue, ok := rpcStruct["supported_on_GPU_groups"]
	if ok {
  	record.SupportedOnGPUGroups, err = convertGPUGroupRefSetToGo(fmt.Sprintf("%s.%s", context, "supported_on_GPU_groups"), supportedOnGPUGroupsValue)
		if err != nil {
			return
		}
	}
  enabledOnGPUGroupsValue, ok := rpcStruct["enabled_on_GPU_groups"]
	if ok {
  	record.EnabledOnGPUGroups, err = convertGPUGroupRefSetToGo(fmt.Sprintf("%s.%s", context, "enabled_on_GPU_groups"), enabledOnGPUGroupsValue)
		if err != nil {
			return
		}
	}
  implementationValue, ok := rpcStruct["implementation"]
	if ok {
  	record.Implementation, err = convertEnumVgpuTypeImplementationToGo(fmt.Sprintf("%s.%s", context, "implementation"), implementationValue)
		if err != nil {
			return
		}
	}
  identifierValue, ok := rpcStruct["identifier"]
	if ok {
  	record.Identifier, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "identifier"), identifierValue)
		if err != nil {
			return
		}
	}
  experimentalValue, ok := rpcStruct["experimental"]
	if ok {
  	record.Experimental, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "experimental"), experimentalValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVGPUTypeRefSetToGo(context string, input interface{}) (slice []VGPUTypeRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VGPUTypeRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPUTypeRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVGPUTypeRefSetToXen(context string, slice []VGPUTypeRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVGPUTypeRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVGPUTypeRefToGo(context string, input interface{}) (ref VGPUTypeRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VGPUTypeRef(value)
	}
	return
}

func convertVGPUTypeRefToXen(context string, ref VGPUTypeRef) (string, error) {
	return string(ref), nil
}

func convertVIFRecordToGo(context string, input interface{}) (record VIFRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumVifOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumVifOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  deviceValue, ok := rpcStruct["device"]
	if ok {
  	record.Device, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "device"), deviceValue)
		if err != nil {
			return
		}
	}
  networkValue, ok := rpcStruct["network"]
	if ok {
  	record.Network, err = convertNetworkRefToGo(fmt.Sprintf("%s.%s", context, "network"), networkValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  macValue, ok := rpcStruct["MAC"]
	if ok {
  	record.MAC, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "MAC"), macValue)
		if err != nil {
			return
		}
	}
  mtuValue, ok := rpcStruct["MTU"]
	if ok {
  	record.MTU, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "MTU"), mtuValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  currentlyAttachedValue, ok := rpcStruct["currently_attached"]
	if ok {
  	record.CurrentlyAttached, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "currently_attached"), currentlyAttachedValue)
		if err != nil {
			return
		}
	}
  statusCodeValue, ok := rpcStruct["status_code"]
	if ok {
  	record.StatusCode, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "status_code"), statusCodeValue)
		if err != nil {
			return
		}
	}
  statusDetailValue, ok := rpcStruct["status_detail"]
	if ok {
  	record.StatusDetail, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "status_detail"), statusDetailValue)
		if err != nil {
			return
		}
	}
  runtimePropertiesValue, ok := rpcStruct["runtime_properties"]
	if ok {
  	record.RuntimeProperties, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "runtime_properties"), runtimePropertiesValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmTypeValue, ok := rpcStruct["qos_algorithm_type"]
	if ok {
  	record.QosAlgorithmType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), qosAlgorithmTypeValue)
		if err != nil {
			return
		}
	}
  qosAlgorithmParamsValue, ok := rpcStruct["qos_algorithm_params"]
	if ok {
  	record.QosAlgorithmParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), qosAlgorithmParamsValue)
		if err != nil {
			return
		}
	}
  qosSupportedAlgorithmsValue, ok := rpcStruct["qos_supported_algorithms"]
	if ok {
  	record.QosSupportedAlgorithms, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), qosSupportedAlgorithmsValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok {
  	record.Metrics, err = convertVIFMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  macAutogeneratedValue, ok := rpcStruct["MAC_autogenerated"]
	if ok {
  	record.MACAutogenerated, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "MAC_autogenerated"), macAutogeneratedValue)
		if err != nil {
			return
		}
	}
  lockingModeValue, ok := rpcStruct["locking_mode"]
	if ok {
  	record.LockingMode, err = convertEnumVifLockingModeToGo(fmt.Sprintf("%s.%s", context, "locking_mode"), lockingModeValue)
		if err != nil {
			return
		}
	}
  ipv4AllowedValue, ok := rpcStruct["ipv4_allowed"]
	if ok {
  	record.Ipv4Allowed, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ipv4_allowed"), ipv4AllowedValue)
		if err != nil {
			return
		}
	}
  ipv6AllowedValue, ok := rpcStruct["ipv6_allowed"]
	if ok {
  	record.Ipv6Allowed, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ipv6_allowed"), ipv6AllowedValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVIFRecordToXen(context string, record VIFRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVifOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVifOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["device"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "device"), record.Device)
  if err != nil {
		return
	}
  rpcStruct["network"], err = convertNetworkRefToXen(fmt.Sprintf("%s.%s", context, "network"), record.Network)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["MAC"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "MAC"), record.MAC)
  if err != nil {
		return
	}
  rpcStruct["MTU"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "MTU"), record.MTU)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["currently_attached"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "currently_attached"), record.CurrentlyAttached)
  if err != nil {
		return
	}
  rpcStruct["status_code"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "status_code"), record.StatusCode)
  if err != nil {
		return
	}
  rpcStruct["status_detail"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "status_detail"), record.StatusDetail)
  if err != nil {
		return
	}
  rpcStruct["runtime_properties"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "runtime_properties"), record.RuntimeProperties)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_type"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_type"), record.QosAlgorithmType)
  if err != nil {
		return
	}
  rpcStruct["qos_algorithm_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "qos_algorithm_params"), record.QosAlgorithmParams)
  if err != nil {
		return
	}
  rpcStruct["qos_supported_algorithms"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "qos_supported_algorithms"), record.QosSupportedAlgorithms)
  if err != nil {
		return
	}
  rpcStruct["metrics"], err = convertVIFMetricsRefToXen(fmt.Sprintf("%s.%s", context, "metrics"), record.Metrics)
  if err != nil {
		return
	}
  rpcStruct["MAC_autogenerated"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "MAC_autogenerated"), record.MACAutogenerated)
  if err != nil {
		return
	}
  rpcStruct["locking_mode"], err = convertEnumVifLockingModeToXen(fmt.Sprintf("%s.%s", context, "locking_mode"), record.LockingMode)
  if err != nil {
		return
	}
  rpcStruct["ipv4_allowed"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "ipv4_allowed"), record.Ipv4Allowed)
  if err != nil {
		return
	}
  rpcStruct["ipv6_allowed"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "ipv6_allowed"), record.Ipv6Allowed)
  if err != nil {
		return
	}
	return
}

func convertVIFRefSetToGo(context string, input interface{}) (slice []VIFRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VIFRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVIFRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVIFRefSetToXen(context string, slice []VIFRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVIFRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVIFRefToGo(context string, input interface{}) (ref VIFRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VIFRef(value)
	}
	return
}

func convertVIFRefToXen(context string, ref VIFRef) (string, error) {
	return string(ref), nil
}

func convertVIFMetricsRecordToGo(context string, input interface{}) (record VIFMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  ioReadKbsValue, ok := rpcStruct["io_read_kbs"]
	if ok {
  	record.IoReadKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_read_kbs"), ioReadKbsValue)
		if err != nil {
			return
		}
	}
  ioWriteKbsValue, ok := rpcStruct["io_write_kbs"]
	if ok {
  	record.IoWriteKbs, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "io_write_kbs"), ioWriteKbsValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVIFMetricsRefSetToGo(context string, input interface{}) (slice []VIFMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VIFMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVIFMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVIFMetricsRefToGo(context string, input interface{}) (ref VIFMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VIFMetricsRef(value)
	}
	return
}

func convertVIFMetricsRefToXen(context string, ref VIFMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVLANRecordToGo(context string, input interface{}) (record VLANRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  taggedPIFValue, ok := rpcStruct["tagged_PIF"]
	if ok {
  	record.TaggedPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "tagged_PIF"), taggedPIFValue)
		if err != nil {
			return
		}
	}
  untaggedPIFValue, ok := rpcStruct["untagged_PIF"]
	if ok {
  	record.UntaggedPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "untagged_PIF"), untaggedPIFValue)
		if err != nil {
			return
		}
	}
  tagValue, ok := rpcStruct["tag"]
	if ok {
  	record.Tag, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "tag"), tagValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVLANRefSetToGo(context string, input interface{}) (slice []VLANRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VLANRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVLANRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVLANRefToGo(context string, input interface{}) (ref VLANRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VLANRef(value)
	}
	return
}

func convertVLANRefToXen(context string, ref VLANRef) (string, error) {
	return string(ref), nil
}

func convertVMRecordToGo(context string, input interface{}) (record VMRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumVMOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumVMOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  powerStateValue, ok := rpcStruct["power_state"]
	if ok {
  	record.PowerState, err = convertEnumVMPowerStateToGo(fmt.Sprintf("%s.%s", context, "power_state"), powerStateValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  userVersionValue, ok := rpcStruct["user_version"]
	if ok {
  	record.UserVersion, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "user_version"), userVersionValue)
		if err != nil {
			return
		}
	}
  isATemplateValue, ok := rpcStruct["is_a_template"]
	if ok {
  	record.IsATemplate, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_a_template"), isATemplateValue)
		if err != nil {
			return
		}
	}
  suspendVDIValue, ok := rpcStruct["suspend_VDI"]
	if ok {
  	record.SuspendVDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "suspend_VDI"), suspendVDIValue)
		if err != nil {
			return
		}
	}
  residentOnValue, ok := rpcStruct["resident_on"]
	if ok {
  	record.ResidentOn, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "resident_on"), residentOnValue)
		if err != nil {
			return
		}
	}
  affinityValue, ok := rpcStruct["affinity"]
	if ok {
  	record.Affinity, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "affinity"), affinityValue)
		if err != nil {
			return
		}
	}
  memoryOverheadValue, ok := rpcStruct["memory_overhead"]
	if ok {
  	record.MemoryOverhead, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_overhead"), memoryOverheadValue)
		if err != nil {
			return
		}
	}
  memoryTargetValue, ok := rpcStruct["memory_target"]
	if ok {
  	record.MemoryTarget, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_target"), memoryTargetValue)
		if err != nil {
			return
		}
	}
  memoryStaticMaxValue, ok := rpcStruct["memory_static_max"]
	if ok {
  	record.MemoryStaticMax, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_static_max"), memoryStaticMaxValue)
		if err != nil {
			return
		}
	}
  memoryDynamicMaxValue, ok := rpcStruct["memory_dynamic_max"]
	if ok {
  	record.MemoryDynamicMax, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_dynamic_max"), memoryDynamicMaxValue)
		if err != nil {
			return
		}
	}
  memoryDynamicMinValue, ok := rpcStruct["memory_dynamic_min"]
	if ok {
  	record.MemoryDynamicMin, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_dynamic_min"), memoryDynamicMinValue)
		if err != nil {
			return
		}
	}
  memoryStaticMinValue, ok := rpcStruct["memory_static_min"]
	if ok {
  	record.MemoryStaticMin, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_static_min"), memoryStaticMinValue)
		if err != nil {
			return
		}
	}
  vcpusParamsValue, ok := rpcStruct["VCPUs_params"]
	if ok {
  	record.VCPUsParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_params"), vcpusParamsValue)
		if err != nil {
			return
		}
	}
  vcpusMaxValue, ok := rpcStruct["VCPUs_max"]
	if ok {
  	record.VCPUsMax, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VCPUs_max"), vcpusMaxValue)
		if err != nil {
			return
		}
	}
  vcpusAtStartupValue, ok := rpcStruct["VCPUs_at_startup"]
	if ok {
  	record.VCPUsAtStartup, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VCPUs_at_startup"), vcpusAtStartupValue)
		if err != nil {
			return
		}
	}
  actionsAfterShutdownValue, ok := rpcStruct["actions_after_shutdown"]
	if ok {
  	record.ActionsAfterShutdown, err = convertEnumOnNormalExitToGo(fmt.Sprintf("%s.%s", context, "actions_after_shutdown"), actionsAfterShutdownValue)
		if err != nil {
			return
		}
	}
  actionsAfterRebootValue, ok := rpcStruct["actions_after_reboot"]
	if ok {
  	record.ActionsAfterReboot, err = convertEnumOnNormalExitToGo(fmt.Sprintf("%s.%s", context, "actions_after_reboot"), actionsAfterRebootValue)
		if err != nil {
			return
		}
	}
  actionsAfterCrashValue, ok := rpcStruct["actions_after_crash"]
	if ok {
  	record.ActionsAfterCrash, err = convertEnumOnCrashBehaviourToGo(fmt.Sprintf("%s.%s", context, "actions_after_crash"), actionsAfterCrashValue)
		if err != nil {
			return
		}
	}
  consolesValue, ok := rpcStruct["consoles"]
	if ok {
  	record.Consoles, err = convertConsoleRefSetToGo(fmt.Sprintf("%s.%s", context, "consoles"), consolesValue)
		if err != nil {
			return
		}
	}
  vifsValue, ok := rpcStruct["VIFs"]
	if ok {
  	record.VIFs, err = convertVIFRefSetToGo(fmt.Sprintf("%s.%s", context, "VIFs"), vifsValue)
		if err != nil {
			return
		}
	}
  vbdsValue, ok := rpcStruct["VBDs"]
	if ok {
  	record.VBDs, err = convertVBDRefSetToGo(fmt.Sprintf("%s.%s", context, "VBDs"), vbdsValue)
		if err != nil {
			return
		}
	}
  crashDumpsValue, ok := rpcStruct["crash_dumps"]
	if ok {
  	record.CrashDumps, err = convertCrashdumpRefSetToGo(fmt.Sprintf("%s.%s", context, "crash_dumps"), crashDumpsValue)
		if err != nil {
			return
		}
	}
  vtpmsValue, ok := rpcStruct["VTPMs"]
	if ok {
  	record.VTPMs, err = convertVTPMRefSetToGo(fmt.Sprintf("%s.%s", context, "VTPMs"), vtpmsValue)
		if err != nil {
			return
		}
	}
  pvBootloaderValue, ok := rpcStruct["PV_bootloader"]
	if ok {
  	record.PVBootloader, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_bootloader"), pvBootloaderValue)
		if err != nil {
			return
		}
	}
  pvKernelValue, ok := rpcStruct["PV_kernel"]
	if ok {
  	record.PVKernel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_kernel"), pvKernelValue)
		if err != nil {
			return
		}
	}
  pvRamdiskValue, ok := rpcStruct["PV_ramdisk"]
	if ok {
  	record.PVRamdisk, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_ramdisk"), pvRamdiskValue)
		if err != nil {
			return
		}
	}
  pvArgsValue, ok := rpcStruct["PV_args"]
	if ok {
  	record.PVArgs, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_args"), pvArgsValue)
		if err != nil {
			return
		}
	}
  pvBootloaderArgsValue, ok := rpcStruct["PV_bootloader_args"]
	if ok {
  	record.PVBootloaderArgs, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_bootloader_args"), pvBootloaderArgsValue)
		if err != nil {
			return
		}
	}
  pvLegacyArgsValue, ok := rpcStruct["PV_legacy_args"]
	if ok {
  	record.PVLegacyArgs, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PV_legacy_args"), pvLegacyArgsValue)
		if err != nil {
			return
		}
	}
  hvmBootPolicyValue, ok := rpcStruct["HVM_boot_policy"]
	if ok {
  	record.HVMBootPolicy, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "HVM_boot_policy"), hvmBootPolicyValue)
		if err != nil {
			return
		}
	}
  hvmBootParamsValue, ok := rpcStruct["HVM_boot_params"]
	if ok {
  	record.HVMBootParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "HVM_boot_params"), hvmBootParamsValue)
		if err != nil {
			return
		}
	}
  hvmShadowMultiplierValue, ok := rpcStruct["HVM_shadow_multiplier"]
	if ok {
  	record.HVMShadowMultiplier, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "HVM_shadow_multiplier"), hvmShadowMultiplierValue)
		if err != nil {
			return
		}
	}
  platformValue, ok := rpcStruct["platform"]
	if ok {
  	record.Platform, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "platform"), platformValue)
		if err != nil {
			return
		}
	}
  pciBusValue, ok := rpcStruct["PCI_bus"]
	if ok {
  	record.PCIBus, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "PCI_bus"), pciBusValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  domidValue, ok := rpcStruct["domid"]
	if ok {
  	record.Domid, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "domid"), domidValue)
		if err != nil {
			return
		}
	}
  domarchValue, ok := rpcStruct["domarch"]
	if ok {
  	record.Domarch, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "domarch"), domarchValue)
		if err != nil {
			return
		}
	}
  lastBootCPUFlagsValue, ok := rpcStruct["last_boot_CPU_flags"]
	if ok {
  	record.LastBootCPUFlags, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "last_boot_CPU_flags"), lastBootCPUFlagsValue)
		if err != nil {
			return
		}
	}
  isControlDomainValue, ok := rpcStruct["is_control_domain"]
	if ok {
  	record.IsControlDomain, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_control_domain"), isControlDomainValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok {
  	record.Metrics, err = convertVMMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  guestMetricsValue, ok := rpcStruct["guest_metrics"]
	if ok {
  	record.GuestMetrics, err = convertVMGuestMetricsRefToGo(fmt.Sprintf("%s.%s", context, "guest_metrics"), guestMetricsValue)
		if err != nil {
			return
		}
	}
  lastBootedRecordValue, ok := rpcStruct["last_booted_record"]
	if ok {
  	record.LastBootedRecord, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "last_booted_record"), lastBootedRecordValue)
		if err != nil {
			return
		}
	}
  recommendationsValue, ok := rpcStruct["recommendations"]
	if ok {
  	record.Recommendations, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "recommendations"), recommendationsValue)
		if err != nil {
			return
		}
	}
  xenstoreDataValue, ok := rpcStruct["xenstore_data"]
	if ok {
  	record.XenstoreData, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "xenstore_data"), xenstoreDataValue)
		if err != nil {
			return
		}
	}
  haAlwaysRunValue, ok := rpcStruct["ha_always_run"]
	if ok {
  	record.HaAlwaysRun, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_always_run"), haAlwaysRunValue)
		if err != nil {
			return
		}
	}
  haRestartPriorityValue, ok := rpcStruct["ha_restart_priority"]
	if ok {
  	record.HaRestartPriority, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ha_restart_priority"), haRestartPriorityValue)
		if err != nil {
			return
		}
	}
  isASnapshotValue, ok := rpcStruct["is_a_snapshot"]
	if ok {
  	record.IsASnapshot, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), isASnapshotValue)
		if err != nil {
			return
		}
	}
  snapshotOfValue, ok := rpcStruct["snapshot_of"]
	if ok {
  	record.SnapshotOf, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "snapshot_of"), snapshotOfValue)
		if err != nil {
			return
		}
	}
  snapshotsValue, ok := rpcStruct["snapshots"]
	if ok {
  	record.Snapshots, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "snapshots"), snapshotsValue)
		if err != nil {
			return
		}
	}
  snapshotTimeValue, ok := rpcStruct["snapshot_time"]
	if ok {
  	record.SnapshotTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "snapshot_time"), snapshotTimeValue)
		if err != nil {
			return
		}
	}
  transportableSnapshotIDValue, ok := rpcStruct["transportable_snapshot_id"]
	if ok {
  	record.TransportableSnapshotID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "transportable_snapshot_id"), transportableSnapshotIDValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  blockedOperationsValue, ok := rpcStruct["blocked_operations"]
	if ok {
  	record.BlockedOperations, err = convertEnumVMOperationsToStringMapToGo(fmt.Sprintf("%s.%s", context, "blocked_operations"), blockedOperationsValue)
		if err != nil {
			return
		}
	}
  snapshotInfoValue, ok := rpcStruct["snapshot_info"]
	if ok {
  	record.SnapshotInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "snapshot_info"), snapshotInfoValue)
		if err != nil {
			return
		}
	}
  snapshotMetadataValue, ok := rpcStruct["snapshot_metadata"]
	if ok {
  	record.SnapshotMetadata, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "snapshot_metadata"), snapshotMetadataValue)
		if err != nil {
			return
		}
	}
  parentValue, ok := rpcStruct["parent"]
	if ok {
  	record.Parent, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "parent"), parentValue)
		if err != nil {
			return
		}
	}
  childrenValue, ok := rpcStruct["children"]
	if ok {
  	record.Children, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "children"), childrenValue)
		if err != nil {
			return
		}
	}
  biosStringsValue, ok := rpcStruct["bios_strings"]
	if ok {
  	record.BiosStrings, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "bios_strings"), biosStringsValue)
		if err != nil {
			return
		}
	}
  protectionPolicyValue, ok := rpcStruct["protection_policy"]
	if ok {
  	record.ProtectionPolicy, err = convertVMPPRefToGo(fmt.Sprintf("%s.%s", context, "protection_policy"), protectionPolicyValue)
		if err != nil {
			return
		}
	}
  isSnapshotFromVmppValue, ok := rpcStruct["is_snapshot_from_vmpp"]
	if ok {
  	record.IsSnapshotFromVmpp, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_snapshot_from_vmpp"), isSnapshotFromVmppValue)
		if err != nil {
			return
		}
	}
  applianceValue, ok := rpcStruct["appliance"]
	if ok {
  	record.Appliance, err = convertVMApplianceRefToGo(fmt.Sprintf("%s.%s", context, "appliance"), applianceValue)
		if err != nil {
			return
		}
	}
  startDelayValue, ok := rpcStruct["start_delay"]
	if ok {
  	record.StartDelay, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "start_delay"), startDelayValue)
		if err != nil {
			return
		}
	}
  shutdownDelayValue, ok := rpcStruct["shutdown_delay"]
	if ok {
  	record.ShutdownDelay, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "shutdown_delay"), shutdownDelayValue)
		if err != nil {
			return
		}
	}
  orderValue, ok := rpcStruct["order"]
	if ok {
  	record.Order, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "order"), orderValue)
		if err != nil {
			return
		}
	}
  vgpusValue, ok := rpcStruct["VGPUs"]
	if ok {
  	record.VGPUs, err = convertVGPURefSetToGo(fmt.Sprintf("%s.%s", context, "VGPUs"), vgpusValue)
		if err != nil {
			return
		}
	}
  attachedPCIsValue, ok := rpcStruct["attached_PCIs"]
	if ok {
  	record.AttachedPCIs, err = convertPCIRefSetToGo(fmt.Sprintf("%s.%s", context, "attached_PCIs"), attachedPCIsValue)
		if err != nil {
			return
		}
	}
  suspendSRValue, ok := rpcStruct["suspend_SR"]
	if ok {
  	record.SuspendSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "suspend_SR"), suspendSRValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok {
  	record.Version, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  generationIDValue, ok := rpcStruct["generation_id"]
	if ok {
  	record.GenerationID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "generation_id"), generationIDValue)
		if err != nil {
			return
		}
	}
  hardwarePlatformVersionValue, ok := rpcStruct["hardware_platform_version"]
	if ok {
  	record.HardwarePlatformVersion, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "hardware_platform_version"), hardwarePlatformVersionValue)
		if err != nil {
			return
		}
	}
  autoUpdateDriversValue, ok := rpcStruct["auto_update_drivers"]
	if ok {
  	record.AutoUpdateDrivers, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "auto_update_drivers"), autoUpdateDriversValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMRecordToXen(context string, record VMRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVMOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVMOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["power_state"], err = convertEnumVMPowerStateToXen(fmt.Sprintf("%s.%s", context, "power_state"), record.PowerState)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["user_version"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "user_version"), record.UserVersion)
  if err != nil {
		return
	}
  rpcStruct["is_a_template"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_a_template"), record.IsATemplate)
  if err != nil {
		return
	}
  rpcStruct["suspend_VDI"], err = convertVDIRefToXen(fmt.Sprintf("%s.%s", context, "suspend_VDI"), record.SuspendVDI)
  if err != nil {
		return
	}
  rpcStruct["resident_on"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "resident_on"), record.ResidentOn)
  if err != nil {
		return
	}
  rpcStruct["affinity"], err = convertHostRefToXen(fmt.Sprintf("%s.%s", context, "affinity"), record.Affinity)
  if err != nil {
		return
	}
  rpcStruct["memory_overhead"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_overhead"), record.MemoryOverhead)
  if err != nil {
		return
	}
  rpcStruct["memory_target"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_target"), record.MemoryTarget)
  if err != nil {
		return
	}
  rpcStruct["memory_static_max"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_static_max"), record.MemoryStaticMax)
  if err != nil {
		return
	}
  rpcStruct["memory_dynamic_max"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_dynamic_max"), record.MemoryDynamicMax)
  if err != nil {
		return
	}
  rpcStruct["memory_dynamic_min"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_dynamic_min"), record.MemoryDynamicMin)
  if err != nil {
		return
	}
  rpcStruct["memory_static_min"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "memory_static_min"), record.MemoryStaticMin)
  if err != nil {
		return
	}
  rpcStruct["VCPUs_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "VCPUs_params"), record.VCPUsParams)
  if err != nil {
		return
	}
  rpcStruct["VCPUs_max"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "VCPUs_max"), record.VCPUsMax)
  if err != nil {
		return
	}
  rpcStruct["VCPUs_at_startup"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "VCPUs_at_startup"), record.VCPUsAtStartup)
  if err != nil {
		return
	}
  rpcStruct["actions_after_shutdown"], err = convertEnumOnNormalExitToXen(fmt.Sprintf("%s.%s", context, "actions_after_shutdown"), record.ActionsAfterShutdown)
  if err != nil {
		return
	}
  rpcStruct["actions_after_reboot"], err = convertEnumOnNormalExitToXen(fmt.Sprintf("%s.%s", context, "actions_after_reboot"), record.ActionsAfterReboot)
  if err != nil {
		return
	}
  rpcStruct["actions_after_crash"], err = convertEnumOnCrashBehaviourToXen(fmt.Sprintf("%s.%s", context, "actions_after_crash"), record.ActionsAfterCrash)
  if err != nil {
		return
	}
  rpcStruct["consoles"], err = convertConsoleRefSetToXen(fmt.Sprintf("%s.%s", context, "consoles"), record.Consoles)
  if err != nil {
		return
	}
  rpcStruct["VIFs"], err = convertVIFRefSetToXen(fmt.Sprintf("%s.%s", context, "VIFs"), record.VIFs)
  if err != nil {
		return
	}
  rpcStruct["VBDs"], err = convertVBDRefSetToXen(fmt.Sprintf("%s.%s", context, "VBDs"), record.VBDs)
  if err != nil {
		return
	}
  rpcStruct["crash_dumps"], err = convertCrashdumpRefSetToXen(fmt.Sprintf("%s.%s", context, "crash_dumps"), record.CrashDumps)
  if err != nil {
		return
	}
  rpcStruct["VTPMs"], err = convertVTPMRefSetToXen(fmt.Sprintf("%s.%s", context, "VTPMs"), record.VTPMs)
  if err != nil {
		return
	}
  rpcStruct["PV_bootloader"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_bootloader"), record.PVBootloader)
  if err != nil {
		return
	}
  rpcStruct["PV_kernel"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_kernel"), record.PVKernel)
  if err != nil {
		return
	}
  rpcStruct["PV_ramdisk"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_ramdisk"), record.PVRamdisk)
  if err != nil {
		return
	}
  rpcStruct["PV_args"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_args"), record.PVArgs)
  if err != nil {
		return
	}
  rpcStruct["PV_bootloader_args"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_bootloader_args"), record.PVBootloaderArgs)
  if err != nil {
		return
	}
  rpcStruct["PV_legacy_args"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PV_legacy_args"), record.PVLegacyArgs)
  if err != nil {
		return
	}
  rpcStruct["HVM_boot_policy"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "HVM_boot_policy"), record.HVMBootPolicy)
  if err != nil {
		return
	}
  rpcStruct["HVM_boot_params"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "HVM_boot_params"), record.HVMBootParams)
  if err != nil {
		return
	}
  rpcStruct["HVM_shadow_multiplier"], err = convertFloatToXen(fmt.Sprintf("%s.%s", context, "HVM_shadow_multiplier"), record.HVMShadowMultiplier)
  if err != nil {
		return
	}
  rpcStruct["platform"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "platform"), record.Platform)
  if err != nil {
		return
	}
  rpcStruct["PCI_bus"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "PCI_bus"), record.PCIBus)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["domid"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "domid"), record.Domid)
  if err != nil {
		return
	}
  rpcStruct["domarch"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "domarch"), record.Domarch)
  if err != nil {
		return
	}
  rpcStruct["last_boot_CPU_flags"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "last_boot_CPU_flags"), record.LastBootCPUFlags)
  if err != nil {
		return
	}
  rpcStruct["is_control_domain"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_control_domain"), record.IsControlDomain)
  if err != nil {
		return
	}
  rpcStruct["metrics"], err = convertVMMetricsRefToXen(fmt.Sprintf("%s.%s", context, "metrics"), record.Metrics)
  if err != nil {
		return
	}
  rpcStruct["guest_metrics"], err = convertVMGuestMetricsRefToXen(fmt.Sprintf("%s.%s", context, "guest_metrics"), record.GuestMetrics)
  if err != nil {
		return
	}
  rpcStruct["last_booted_record"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "last_booted_record"), record.LastBootedRecord)
  if err != nil {
		return
	}
  rpcStruct["recommendations"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "recommendations"), record.Recommendations)
  if err != nil {
		return
	}
  rpcStruct["xenstore_data"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "xenstore_data"), record.XenstoreData)
  if err != nil {
		return
	}
  rpcStruct["ha_always_run"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "ha_always_run"), record.HaAlwaysRun)
  if err != nil {
		return
	}
  rpcStruct["ha_restart_priority"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "ha_restart_priority"), record.HaRestartPriority)
  if err != nil {
		return
	}
  rpcStruct["is_a_snapshot"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_a_snapshot"), record.IsASnapshot)
  if err != nil {
		return
	}
  rpcStruct["snapshot_of"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "snapshot_of"), record.SnapshotOf)
  if err != nil {
		return
	}
  rpcStruct["snapshots"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "snapshots"), record.Snapshots)
  if err != nil {
		return
	}
  rpcStruct["snapshot_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "snapshot_time"), record.SnapshotTime)
  if err != nil {
		return
	}
  rpcStruct["transportable_snapshot_id"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "transportable_snapshot_id"), record.TransportableSnapshotID)
  if err != nil {
		return
	}
  rpcStruct["blobs"], err = convertStringToBlobRefMapToXen(fmt.Sprintf("%s.%s", context, "blobs"), record.Blobs)
  if err != nil {
		return
	}
  rpcStruct["tags"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "tags"), record.Tags)
  if err != nil {
		return
	}
  rpcStruct["blocked_operations"], err = convertEnumVMOperationsToStringMapToXen(fmt.Sprintf("%s.%s", context, "blocked_operations"), record.BlockedOperations)
  if err != nil {
		return
	}
  rpcStruct["snapshot_info"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "snapshot_info"), record.SnapshotInfo)
  if err != nil {
		return
	}
  rpcStruct["snapshot_metadata"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "snapshot_metadata"), record.SnapshotMetadata)
  if err != nil {
		return
	}
  rpcStruct["parent"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "parent"), record.Parent)
  if err != nil {
		return
	}
  rpcStruct["children"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "children"), record.Children)
  if err != nil {
		return
	}
  rpcStruct["bios_strings"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "bios_strings"), record.BiosStrings)
  if err != nil {
		return
	}
  rpcStruct["protection_policy"], err = convertVMPPRefToXen(fmt.Sprintf("%s.%s", context, "protection_policy"), record.ProtectionPolicy)
  if err != nil {
		return
	}
  rpcStruct["is_snapshot_from_vmpp"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_snapshot_from_vmpp"), record.IsSnapshotFromVmpp)
  if err != nil {
		return
	}
  rpcStruct["appliance"], err = convertVMApplianceRefToXen(fmt.Sprintf("%s.%s", context, "appliance"), record.Appliance)
  if err != nil {
		return
	}
  rpcStruct["start_delay"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "start_delay"), record.StartDelay)
  if err != nil {
		return
	}
  rpcStruct["shutdown_delay"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "shutdown_delay"), record.ShutdownDelay)
  if err != nil {
		return
	}
  rpcStruct["order"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "order"), record.Order)
  if err != nil {
		return
	}
  rpcStruct["VGPUs"], err = convertVGPURefSetToXen(fmt.Sprintf("%s.%s", context, "VGPUs"), record.VGPUs)
  if err != nil {
		return
	}
  rpcStruct["attached_PCIs"], err = convertPCIRefSetToXen(fmt.Sprintf("%s.%s", context, "attached_PCIs"), record.AttachedPCIs)
  if err != nil {
		return
	}
  rpcStruct["suspend_SR"], err = convertSRRefToXen(fmt.Sprintf("%s.%s", context, "suspend_SR"), record.SuspendSR)
  if err != nil {
		return
	}
  rpcStruct["version"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "version"), record.Version)
  if err != nil {
		return
	}
  rpcStruct["generation_id"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "generation_id"), record.GenerationID)
  if err != nil {
		return
	}
  rpcStruct["hardware_platform_version"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "hardware_platform_version"), record.HardwarePlatformVersion)
  if err != nil {
		return
	}
  rpcStruct["auto_update_drivers"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "auto_update_drivers"), record.AutoUpdateDrivers)
  if err != nil {
		return
	}
	return
}

func convertVMRefSetToGo(context string, input interface{}) (slice []VMRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMRefSetToXen(context string, slice []VMRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVMRefToGo(context string, input interface{}) (ref VMRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMRef(value)
	}
	return
}

func convertVMRefToXen(context string, ref VMRef) (string, error) {
	return string(ref), nil
}

func convertVMPPRecordToGo(context string, input interface{}) (record VMPPRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  isPolicyEnabledValue, ok := rpcStruct["is_policy_enabled"]
	if ok {
  	record.IsPolicyEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_policy_enabled"), isPolicyEnabledValue)
		if err != nil {
			return
		}
	}
  backupTypeValue, ok := rpcStruct["backup_type"]
	if ok {
  	record.BackupType, err = convertEnumVmppBackupTypeToGo(fmt.Sprintf("%s.%s", context, "backup_type"), backupTypeValue)
		if err != nil {
			return
		}
	}
  backupRetentionValueValue, ok := rpcStruct["backup_retention_value"]
	if ok {
  	record.BackupRetentionValue, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "backup_retention_value"), backupRetentionValueValue)
		if err != nil {
			return
		}
	}
  backupFrequencyValue, ok := rpcStruct["backup_frequency"]
	if ok {
  	record.BackupFrequency, err = convertEnumVmppBackupFrequencyToGo(fmt.Sprintf("%s.%s", context, "backup_frequency"), backupFrequencyValue)
		if err != nil {
			return
		}
	}
  backupScheduleValue, ok := rpcStruct["backup_schedule"]
	if ok {
  	record.BackupSchedule, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "backup_schedule"), backupScheduleValue)
		if err != nil {
			return
		}
	}
  isBackupRunningValue, ok := rpcStruct["is_backup_running"]
	if ok {
  	record.IsBackupRunning, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_backup_running"), isBackupRunningValue)
		if err != nil {
			return
		}
	}
  backupLastRunTimeValue, ok := rpcStruct["backup_last_run_time"]
	if ok {
  	record.BackupLastRunTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "backup_last_run_time"), backupLastRunTimeValue)
		if err != nil {
			return
		}
	}
  archiveTargetTypeValue, ok := rpcStruct["archive_target_type"]
	if ok {
  	record.ArchiveTargetType, err = convertEnumVmppArchiveTargetTypeToGo(fmt.Sprintf("%s.%s", context, "archive_target_type"), archiveTargetTypeValue)
		if err != nil {
			return
		}
	}
  archiveTargetConfigValue, ok := rpcStruct["archive_target_config"]
	if ok {
  	record.ArchiveTargetConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "archive_target_config"), archiveTargetConfigValue)
		if err != nil {
			return
		}
	}
  archiveFrequencyValue, ok := rpcStruct["archive_frequency"]
	if ok {
  	record.ArchiveFrequency, err = convertEnumVmppArchiveFrequencyToGo(fmt.Sprintf("%s.%s", context, "archive_frequency"), archiveFrequencyValue)
		if err != nil {
			return
		}
	}
  archiveScheduleValue, ok := rpcStruct["archive_schedule"]
	if ok {
  	record.ArchiveSchedule, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "archive_schedule"), archiveScheduleValue)
		if err != nil {
			return
		}
	}
  isArchiveRunningValue, ok := rpcStruct["is_archive_running"]
	if ok {
  	record.IsArchiveRunning, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_archive_running"), isArchiveRunningValue)
		if err != nil {
			return
		}
	}
  archiveLastRunTimeValue, ok := rpcStruct["archive_last_run_time"]
	if ok {
  	record.ArchiveLastRunTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "archive_last_run_time"), archiveLastRunTimeValue)
		if err != nil {
			return
		}
	}
  vmsValue, ok := rpcStruct["VMs"]
	if ok {
  	record.VMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "VMs"), vmsValue)
		if err != nil {
			return
		}
	}
  isAlarmEnabledValue, ok := rpcStruct["is_alarm_enabled"]
	if ok {
  	record.IsAlarmEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_alarm_enabled"), isAlarmEnabledValue)
		if err != nil {
			return
		}
	}
  alarmConfigValue, ok := rpcStruct["alarm_config"]
	if ok {
  	record.AlarmConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "alarm_config"), alarmConfigValue)
		if err != nil {
			return
		}
	}
  recentAlertsValue, ok := rpcStruct["recent_alerts"]
	if ok {
  	record.RecentAlerts, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "recent_alerts"), recentAlertsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMPPRecordToXen(context string, record VMPPRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["is_policy_enabled"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_policy_enabled"), record.IsPolicyEnabled)
  if err != nil {
		return
	}
  rpcStruct["backup_type"], err = convertEnumVmppBackupTypeToXen(fmt.Sprintf("%s.%s", context, "backup_type"), record.BackupType)
  if err != nil {
		return
	}
  rpcStruct["backup_retention_value"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "backup_retention_value"), record.BackupRetentionValue)
  if err != nil {
		return
	}
  rpcStruct["backup_frequency"], err = convertEnumVmppBackupFrequencyToXen(fmt.Sprintf("%s.%s", context, "backup_frequency"), record.BackupFrequency)
  if err != nil {
		return
	}
  rpcStruct["backup_schedule"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "backup_schedule"), record.BackupSchedule)
  if err != nil {
		return
	}
  rpcStruct["is_backup_running"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_backup_running"), record.IsBackupRunning)
  if err != nil {
		return
	}
  rpcStruct["backup_last_run_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "backup_last_run_time"), record.BackupLastRunTime)
  if err != nil {
		return
	}
  rpcStruct["archive_target_type"], err = convertEnumVmppArchiveTargetTypeToXen(fmt.Sprintf("%s.%s", context, "archive_target_type"), record.ArchiveTargetType)
  if err != nil {
		return
	}
  rpcStruct["archive_target_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "archive_target_config"), record.ArchiveTargetConfig)
  if err != nil {
		return
	}
  rpcStruct["archive_frequency"], err = convertEnumVmppArchiveFrequencyToXen(fmt.Sprintf("%s.%s", context, "archive_frequency"), record.ArchiveFrequency)
  if err != nil {
		return
	}
  rpcStruct["archive_schedule"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "archive_schedule"), record.ArchiveSchedule)
  if err != nil {
		return
	}
  rpcStruct["is_archive_running"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_archive_running"), record.IsArchiveRunning)
  if err != nil {
		return
	}
  rpcStruct["archive_last_run_time"], err = convertTimeToXen(fmt.Sprintf("%s.%s", context, "archive_last_run_time"), record.ArchiveLastRunTime)
  if err != nil {
		return
	}
  rpcStruct["VMs"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "VMs"), record.VMs)
  if err != nil {
		return
	}
  rpcStruct["is_alarm_enabled"], err = convertBoolToXen(fmt.Sprintf("%s.%s", context, "is_alarm_enabled"), record.IsAlarmEnabled)
  if err != nil {
		return
	}
  rpcStruct["alarm_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "alarm_config"), record.AlarmConfig)
  if err != nil {
		return
	}
  rpcStruct["recent_alerts"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "recent_alerts"), record.RecentAlerts)
  if err != nil {
		return
	}
	return
}

func convertVMPPRefSetToGo(context string, input interface{}) (slice []VMPPRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMPPRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMPPRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMPPRefToGo(context string, input interface{}) (ref VMPPRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMPPRef(value)
	}
	return
}

func convertVMPPRefToXen(context string, ref VMPPRef) (string, error) {
	return string(ref), nil
}

func convertVMApplianceRecordToGo(context string, input interface{}) (record VMApplianceRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumVMApplianceOperationSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumVMApplianceOperationMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vmsValue, ok := rpcStruct["VMs"]
	if ok {
  	record.VMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "VMs"), vmsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMApplianceRecordToXen(context string, record VMApplianceRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumVMApplianceOperationSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumVMApplianceOperationMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["VMs"], err = convertVMRefSetToXen(fmt.Sprintf("%s.%s", context, "VMs"), record.VMs)
  if err != nil {
		return
	}
	return
}

func convertVMApplianceRefSetToGo(context string, input interface{}) (slice []VMApplianceRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMApplianceRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMApplianceRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMApplianceRefToGo(context string, input interface{}) (ref VMApplianceRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMApplianceRef(value)
	}
	return
}

func convertVMApplianceRefToXen(context string, ref VMApplianceRef) (string, error) {
	return string(ref), nil
}

func convertVMGuestMetricsRecordToGo(context string, input interface{}) (record VMGuestMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  osVersionValue, ok := rpcStruct["os_version"]
	if ok {
  	record.OsVersion, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "os_version"), osVersionValue)
		if err != nil {
			return
		}
	}
  pvDriversVersionValue, ok := rpcStruct["PV_drivers_version"]
	if ok {
  	record.PVDriversVersion, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "PV_drivers_version"), pvDriversVersionValue)
		if err != nil {
			return
		}
	}
  pvDriversUpToDateValue, ok := rpcStruct["PV_drivers_up_to_date"]
	if ok {
  	record.PVDriversUpToDate, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "PV_drivers_up_to_date"), pvDriversUpToDateValue)
		if err != nil {
			return
		}
	}
  networkPathsOptimizedValue, ok := rpcStruct["network_paths_optimized"]
	if ok {
  	record.NetworkPathsOptimized, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "network_paths_optimized"), networkPathsOptimizedValue)
		if err != nil {
			return
		}
	}
  storagePathsOptimizedValue, ok := rpcStruct["storage_paths_optimized"]
	if ok {
  	record.StoragePathsOptimized, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "storage_paths_optimized"), storagePathsOptimizedValue)
		if err != nil {
			return
		}
	}
  memoryValue, ok := rpcStruct["memory"]
	if ok {
  	record.Memory, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "memory"), memoryValue)
		if err != nil {
			return
		}
	}
  disksValue, ok := rpcStruct["disks"]
	if ok {
  	record.Disks, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "disks"), disksValue)
		if err != nil {
			return
		}
	}
  networksValue, ok := rpcStruct["networks"]
	if ok {
  	record.Networks, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "networks"), networksValue)
		if err != nil {
			return
		}
	}
  otherValue, ok := rpcStruct["other"]
	if ok {
  	record.Other, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other"), otherValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  liveValue, ok := rpcStruct["live"]
	if ok {
  	record.Live, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "live"), liveValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMGuestMetricsRefSetToGo(context string, input interface{}) (slice []VMGuestMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMGuestMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMGuestMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMGuestMetricsRefToGo(context string, input interface{}) (ref VMGuestMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMGuestMetricsRef(value)
	}
	return
}

func convertVMGuestMetricsRefToXen(context string, ref VMGuestMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVMMetricsRecordToGo(context string, input interface{}) (record VMMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  memoryActualValue, ok := rpcStruct["memory_actual"]
	if ok {
  	record.MemoryActual, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_actual"), memoryActualValue)
		if err != nil {
			return
		}
	}
  vcpusNumberValue, ok := rpcStruct["VCPUs_number"]
	if ok {
  	record.VCPUsNumber, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "VCPUs_number"), vcpusNumberValue)
		if err != nil {
			return
		}
	}
  vcpusUtilisationValue, ok := rpcStruct["VCPUs_utilisation"]
	if ok {
  	record.VCPUsUtilisation, err = convertIntToFloatMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_utilisation"), vcpusUtilisationValue)
		if err != nil {
			return
		}
	}
  vcpusCPUValue, ok := rpcStruct["VCPUs_CPU"]
	if ok {
  	record.VCPUsCPU, err = convertIntToIntMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_CPU"), vcpusCPUValue)
		if err != nil {
			return
		}
	}
  vcpusParamsValue, ok := rpcStruct["VCPUs_params"]
	if ok {
  	record.VCPUsParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_params"), vcpusParamsValue)
		if err != nil {
			return
		}
	}
  vcpusFlagsValue, ok := rpcStruct["VCPUs_flags"]
	if ok {
  	record.VCPUsFlags, err = convertIntToStringSetMapToGo(fmt.Sprintf("%s.%s", context, "VCPUs_flags"), vcpusFlagsValue)
		if err != nil {
			return
		}
	}
  stateValue, ok := rpcStruct["state"]
	if ok {
  	record.State, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "state"), stateValue)
		if err != nil {
			return
		}
	}
  startTimeValue, ok := rpcStruct["start_time"]
	if ok {
  	record.StartTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "start_time"), startTimeValue)
		if err != nil {
			return
		}
	}
  installTimeValue, ok := rpcStruct["install_time"]
	if ok {
  	record.InstallTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "install_time"), installTimeValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVMMetricsRefSetToGo(context string, input interface{}) (slice []VMMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVMMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVMMetricsRefToGo(context string, input interface{}) (ref VMMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VMMetricsRef(value)
	}
	return
}

func convertVMMetricsRefToXen(context string, ref VMMetricsRef) (string, error) {
	return string(ref), nil
}

func convertVTPMRecordToGo(context string, input interface{}) (record VTPMRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  backendValue, ok := rpcStruct["backend"]
	if ok {
  	record.Backend, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "backend"), backendValue)
		if err != nil {
			return
		}
	}
	return
}

func convertVTPMRecordToXen(context string, record VTPMRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["backend"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "backend"), record.Backend)
  if err != nil {
		return
	}
	return
}

func convertVTPMRefSetToGo(context string, input interface{}) (slice []VTPMRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VTPMRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVTPMRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertVTPMRefSetToXen(context string, slice []VTPMRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertVTPMRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertVTPMRefToGo(context string, input interface{}) (ref VTPMRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = VTPMRef(value)
	}
	return
}

func convertVTPMRefToXen(context string, ref VTPMRef) (string, error) {
	return string(ref), nil
}

func convertBlobRecordToGo(context string, input interface{}) (record BlobRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  publicValue, ok := rpcStruct["public"]
	if ok {
  	record.Public, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "public"), publicValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  mimeTypeValue, ok := rpcStruct["mime_type"]
	if ok {
  	record.MimeType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "mime_type"), mimeTypeValue)
		if err != nil {
			return
		}
	}
	return
}

func convertBlobRefSetToGo(context string, input interface{}) (slice []BlobRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]BlobRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertBlobRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertBlobRefToGo(context string, input interface{}) (ref BlobRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = BlobRef(value)
	}
	return
}

func convertBlobRefToXen(context string, ref BlobRef) (string, error) {
	return string(ref), nil
}

func convertBoolToGo(context string, input interface{}) (value bool, err error) {
	if input == nil {
		return
	}
	value, ok := input.(bool)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "bool", context, reflect.TypeOf(input), input)
	}
	return
}

func convertBoolToXen(context string, value bool) (bool, error) {
	return value, nil
}

func convertConsoleRecordToGo(context string, input interface{}) (record ConsoleRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  protocolValue, ok := rpcStruct["protocol"]
	if ok {
  	record.Protocol, err = convertEnumConsoleProtocolToGo(fmt.Sprintf("%s.%s", context, "protocol"), protocolValue)
		if err != nil {
			return
		}
	}
  locationValue, ok := rpcStruct["location"]
	if ok {
  	record.Location, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "location"), locationValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertConsoleRecordToXen(context string, record ConsoleRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["protocol"], err = convertEnumConsoleProtocolToXen(fmt.Sprintf("%s.%s", context, "protocol"), record.Protocol)
  if err != nil {
		return
	}
  rpcStruct["location"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "location"), record.Location)
  if err != nil {
		return
	}
  rpcStruct["VM"], err = convertVMRefToXen(fmt.Sprintf("%s.%s", context, "VM"), record.VM)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertConsoleRefSetToGo(context string, input interface{}) (slice []ConsoleRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]ConsoleRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertConsoleRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertConsoleRefSetToXen(context string, slice []ConsoleRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertConsoleRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertConsoleRefToGo(context string, input interface{}) (ref ConsoleRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = ConsoleRef(value)
	}
	return
}

func convertConsoleRefToXen(context string, ref ConsoleRef) (string, error) {
	return string(ref), nil
}

func convertCrashdumpRecordToGo(context string, input interface{}) (record CrashdumpRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  vmValue, ok := rpcStruct["VM"]
	if ok {
  	record.VM, err = convertVMRefToGo(fmt.Sprintf("%s.%s", context, "VM"), vmValue)
		if err != nil {
			return
		}
	}
  vdiValue, ok := rpcStruct["VDI"]
	if ok {
  	record.VDI, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "VDI"), vdiValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertCrashdumpRefSetToGo(context string, input interface{}) (slice []CrashdumpRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]CrashdumpRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertCrashdumpRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertCrashdumpRefSetToXen(context string, slice []CrashdumpRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertCrashdumpRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertCrashdumpRefToGo(context string, input interface{}) (ref CrashdumpRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = CrashdumpRef(value)
	}
	return
}

func convertCrashdumpRefToXen(context string, ref CrashdumpRef) (string, error) {
	return string(ref), nil
}

func convertDataSourceRecordSetToGo(context string, input interface{}) (slice []DataSourceRecord, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]DataSourceRecord, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertDataSourceRecordToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertDataSourceRecordToGo(context string, input interface{}) (record DataSourceRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  enabledValue, ok := rpcStruct["enabled"]
	if ok {
  	record.Enabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enabled"), enabledValue)
		if err != nil {
			return
		}
	}
  standardValue, ok := rpcStruct["standard"]
	if ok {
  	record.Standard, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "standard"), standardValue)
		if err != nil {
			return
		}
	}
  unitsValue, ok := rpcStruct["units"]
	if ok {
  	record.Units, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "units"), unitsValue)
		if err != nil {
			return
		}
	}
  minValue, ok := rpcStruct["min"]
	if ok {
  	record.Min, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "min"), minValue)
		if err != nil {
			return
		}
	}
  maxValue, ok := rpcStruct["max"]
	if ok {
  	record.Max, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "max"), maxValue)
		if err != nil {
			return
		}
	}
  valueValue, ok := rpcStruct["value"]
	if ok {
  	record.Value, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "value"), valueValue)
		if err != nil {
			return
		}
	}
	return
}

func convertTimeToGo(context string, input interface{}) (value time.Time, err error) {
	if input == nil {
		return
	}
	value, ok := input.(time.Time)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "time.Time", context, reflect.TypeOf(input), input)
	}
	return
}

func convertTimeToXen(context string, value time.Time) (time.Time, error) {
	return value, nil
}

func convertEnumAfterApplyGuidanceSetToGo(context string, input interface{}) (slice []AfterApplyGuidance, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]AfterApplyGuidance, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumAfterApplyGuidanceToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumAfterApplyGuidanceToGo(context string, input interface{}) (value AfterApplyGuidance, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "restartHVM":
      value = AfterApplyGuidanceRestartHVM
    case "restartPV":
      value = AfterApplyGuidanceRestartPV
    case "restartHost":
      value = AfterApplyGuidanceRestartHost
    case "restartXAPI":
      value = AfterApplyGuidanceRestartXAPI
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "AfterApplyGuidance", context)
	}
	return
}

func convertEnumAllocationAlgorithmToGo(context string, input interface{}) (value AllocationAlgorithm, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "breadth_first":
      value = AllocationAlgorithmBreadthFirst
    case "depth_first":
      value = AllocationAlgorithmDepthFirst
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "AllocationAlgorithm", context)
	}
	return
}

func convertEnumAllocationAlgorithmToXen(context string, value AllocationAlgorithm) (string, error) {
	return string(value), nil
}

func convertEnumBondModeToGo(context string, input interface{}) (value BondMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "balance-slb":
      value = BondModeBalanceSlb
    case "active-backup":
      value = BondModeActiveBackup
    case "lacp":
      value = BondModeLacp
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "BondMode", context)
	}
	return
}

func convertEnumBondModeToXen(context string, value BondMode) (string, error) {
	return string(value), nil
}

func convertEnumClsToGo(context string, input interface{}) (value Cls, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "VM":
      value = ClsVM
    case "Host":
      value = ClsHost
    case "SR":
      value = ClsSR
    case "Pool":
      value = ClsPool
    case "VMPP":
      value = ClsVMPP
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "Cls", context)
	}
	return
}

func convertEnumClsToXen(context string, value Cls) (string, error) {
	return string(value), nil
}

func convertEnumConsoleProtocolToGo(context string, input interface{}) (value ConsoleProtocol, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "vt100":
      value = ConsoleProtocolVt100
    case "rfb":
      value = ConsoleProtocolRfb
    case "rdp":
      value = ConsoleProtocolRdp
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "ConsoleProtocol", context)
	}
	return
}

func convertEnumConsoleProtocolToXen(context string, value ConsoleProtocol) (string, error) {
	return string(value), nil
}

func convertEnumEventOperationToGo(context string, input interface{}) (value EventOperation, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "add":
      value = EventOperationAdd
    case "del":
      value = EventOperationDel
    case "mod":
      value = EventOperationMod
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "EventOperation", context)
	}
	return
}

func convertEnumHostAllowedOperationsSetToGo(context string, input interface{}) (slice []HostAllowedOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostAllowedOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumHostAllowedOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumHostAllowedOperationsToGo(context string, input interface{}) (value HostAllowedOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "provision":
      value = HostAllowedOperationsProvision
    case "evacuate":
      value = HostAllowedOperationsEvacuate
    case "shutdown":
      value = HostAllowedOperationsShutdown
    case "reboot":
      value = HostAllowedOperationsReboot
    case "power_on":
      value = HostAllowedOperationsPowerOn
    case "vm_start":
      value = HostAllowedOperationsVMStart
    case "vm_resume":
      value = HostAllowedOperationsVMResume
    case "vm_migrate":
      value = HostAllowedOperationsVMMigrate
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "HostAllowedOperations", context)
	}
	return
}

func convertEnumHostDisplayToGo(context string, input interface{}) (value HostDisplay, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "enabled":
      value = HostDisplayEnabled
    case "disable_on_reboot":
      value = HostDisplayDisableOnReboot
    case "disabled":
      value = HostDisplayDisabled
    case "enable_on_reboot":
      value = HostDisplayEnableOnReboot
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "HostDisplay", context)
	}
	return
}

func convertEnumHostDisplayToXen(context string, value HostDisplay) (string, error) {
	return string(value), nil
}

func convertEnumIPConfigurationModeToGo(context string, input interface{}) (value IPConfigurationMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "None":
      value = IPConfigurationModeNone
    case "DHCP":
      value = IPConfigurationModeDHCP
    case "Static":
      value = IPConfigurationModeStatic
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "IPConfigurationMode", context)
	}
	return
}

func convertEnumIPConfigurationModeToXen(context string, value IPConfigurationMode) (string, error) {
	return string(value), nil
}

func convertEnumIpv6ConfigurationModeToGo(context string, input interface{}) (value Ipv6ConfigurationMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "None":
      value = Ipv6ConfigurationModeNone
    case "DHCP":
      value = Ipv6ConfigurationModeDHCP
    case "Static":
      value = Ipv6ConfigurationModeStatic
    case "Autoconf":
      value = Ipv6ConfigurationModeAutoconf
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "Ipv6ConfigurationMode", context)
	}
	return
}

func convertEnumIpv6ConfigurationModeToXen(context string, value Ipv6ConfigurationMode) (string, error) {
	return string(value), nil
}

func convertEnumNetworkDefaultLockingModeToGo(context string, input interface{}) (value NetworkDefaultLockingMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "unlocked":
      value = NetworkDefaultLockingModeUnlocked
    case "disabled":
      value = NetworkDefaultLockingModeDisabled
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "NetworkDefaultLockingMode", context)
	}
	return
}

func convertEnumNetworkDefaultLockingModeToXen(context string, value NetworkDefaultLockingMode) (string, error) {
	return string(value), nil
}

func convertEnumNetworkOperationsSetToGo(context string, input interface{}) (slice []NetworkOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]NetworkOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumNetworkOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumNetworkOperationsSetToXen(context string, slice []NetworkOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumNetworkOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumNetworkOperationsToGo(context string, input interface{}) (value NetworkOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attaching":
      value = NetworkOperationsAttaching
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "NetworkOperations", context)
	}
	return
}

func convertEnumNetworkOperationsToXen(context string, value NetworkOperations) (string, error) {
	return string(value), nil
}

func convertEnumOnBootToGo(context string, input interface{}) (value OnBoot, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "reset":
      value = OnBootReset
    case "persist":
      value = OnBootPersist
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "OnBoot", context)
	}
	return
}

func convertEnumOnBootToXen(context string, value OnBoot) (string, error) {
	return string(value), nil
}

func convertEnumOnCrashBehaviourToGo(context string, input interface{}) (value OnCrashBehaviour, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "destroy":
      value = OnCrashBehaviourDestroy
    case "coredump_and_destroy":
      value = OnCrashBehaviourCoredumpAndDestroy
    case "restart":
      value = OnCrashBehaviourRestart
    case "coredump_and_restart":
      value = OnCrashBehaviourCoredumpAndRestart
    case "preserve":
      value = OnCrashBehaviourPreserve
    case "rename_restart":
      value = OnCrashBehaviourRenameRestart
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "OnCrashBehaviour", context)
	}
	return
}

func convertEnumOnCrashBehaviourToXen(context string, value OnCrashBehaviour) (string, error) {
	return string(value), nil
}

func convertEnumOnNormalExitToGo(context string, input interface{}) (value OnNormalExit, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "destroy":
      value = OnNormalExitDestroy
    case "restart":
      value = OnNormalExitRestart
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "OnNormalExit", context)
	}
	return
}

func convertEnumOnNormalExitToXen(context string, value OnNormalExit) (string, error) {
	return string(value), nil
}

func convertEnumPgpuDom0AccessToGo(context string, input interface{}) (value PgpuDom0Access, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "enabled":
      value = PgpuDom0AccessEnabled
    case "disable_on_reboot":
      value = PgpuDom0AccessDisableOnReboot
    case "disabled":
      value = PgpuDom0AccessDisabled
    case "enable_on_reboot":
      value = PgpuDom0AccessEnableOnReboot
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PgpuDom0Access", context)
	}
	return
}

func convertEnumPoolAllowedOperationsSetToGo(context string, input interface{}) (slice []PoolAllowedOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolAllowedOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumPoolAllowedOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumPoolAllowedOperationsToGo(context string, input interface{}) (value PoolAllowedOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "ha_enable":
      value = PoolAllowedOperationsHaEnable
    case "ha_disable":
      value = PoolAllowedOperationsHaDisable
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PoolAllowedOperations", context)
	}
	return
}

func convertEnumPrimaryAddressTypeToGo(context string, input interface{}) (value PrimaryAddressType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "IPv4":
      value = PrimaryAddressTypeIPv4
    case "IPv6":
      value = PrimaryAddressTypeIPv6
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "PrimaryAddressType", context)
	}
	return
}

func convertEnumPrimaryAddressTypeToXen(context string, value PrimaryAddressType) (string, error) {
	return string(value), nil
}

func convertEnumStorageOperationsSetToGo(context string, input interface{}) (slice []StorageOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]StorageOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumStorageOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumStorageOperationsToGo(context string, input interface{}) (value StorageOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "scan":
      value = StorageOperationsScan
    case "destroy":
      value = StorageOperationsDestroy
    case "forget":
      value = StorageOperationsForget
    case "plug":
      value = StorageOperationsPlug
    case "unplug":
      value = StorageOperationsUnplug
    case "update":
      value = StorageOperationsUpdate
    case "vdi_create":
      value = StorageOperationsVdiCreate
    case "vdi_introduce":
      value = StorageOperationsVdiIntroduce
    case "vdi_destroy":
      value = StorageOperationsVdiDestroy
    case "vdi_resize":
      value = StorageOperationsVdiResize
    case "vdi_clone":
      value = StorageOperationsVdiClone
    case "vdi_snapshot":
      value = StorageOperationsVdiSnapshot
    case "pbd_create":
      value = StorageOperationsPbdCreate
    case "pbd_destroy":
      value = StorageOperationsPbdDestroy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "StorageOperations", context)
	}
	return
}

func convertEnumTaskAllowedOperationsSetToGo(context string, input interface{}) (slice []TaskAllowedOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]TaskAllowedOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumTaskAllowedOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumTaskAllowedOperationsToGo(context string, input interface{}) (value TaskAllowedOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "cancel":
      value = TaskAllowedOperationsCancel
    case "destroy":
      value = TaskAllowedOperationsDestroy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "TaskAllowedOperations", context)
	}
	return
}

func convertEnumTaskStatusTypeToGo(context string, input interface{}) (value TaskStatusType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "pending":
      value = TaskStatusTypePending
    case "success":
      value = TaskStatusTypeSuccess
    case "failure":
      value = TaskStatusTypeFailure
    case "cancelling":
      value = TaskStatusTypeCancelling
    case "cancelled":
      value = TaskStatusTypeCancelled
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "TaskStatusType", context)
	}
	return
}

func convertEnumVbdModeToGo(context string, input interface{}) (value VbdMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "RO":
      value = VbdModeRO
    case "RW":
      value = VbdModeRW
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VbdMode", context)
	}
	return
}

func convertEnumVbdModeToXen(context string, value VbdMode) (string, error) {
	return string(value), nil
}

func convertEnumVbdOperationsSetToGo(context string, input interface{}) (slice []VbdOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VbdOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVbdOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVbdOperationsSetToXen(context string, slice []VbdOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVbdOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVbdOperationsToGo(context string, input interface{}) (value VbdOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attach":
      value = VbdOperationsAttach
    case "eject":
      value = VbdOperationsEject
    case "insert":
      value = VbdOperationsInsert
    case "plug":
      value = VbdOperationsPlug
    case "unplug":
      value = VbdOperationsUnplug
    case "unplug_force":
      value = VbdOperationsUnplugForce
    case "pause":
      value = VbdOperationsPause
    case "unpause":
      value = VbdOperationsUnpause
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VbdOperations", context)
	}
	return
}

func convertEnumVbdOperationsToXen(context string, value VbdOperations) (string, error) {
	return string(value), nil
}

func convertEnumVbdTypeToGo(context string, input interface{}) (value VbdType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "CD":
      value = VbdTypeCD
    case "Disk":
      value = VbdTypeDisk
    case "Floppy":
      value = VbdTypeFloppy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VbdType", context)
	}
	return
}

func convertEnumVbdTypeToXen(context string, value VbdType) (string, error) {
	return string(value), nil
}

func convertEnumVdiOperationsSetToGo(context string, input interface{}) (slice []VdiOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VdiOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVdiOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVdiOperationsSetToXen(context string, slice []VdiOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVdiOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVdiOperationsToGo(context string, input interface{}) (value VdiOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "scan":
      value = VdiOperationsScan
    case "clone":
      value = VdiOperationsClone
    case "copy":
      value = VdiOperationsCopy
    case "resize":
      value = VdiOperationsResize
    case "resize_online":
      value = VdiOperationsResizeOnline
    case "snapshot":
      value = VdiOperationsSnapshot
    case "destroy":
      value = VdiOperationsDestroy
    case "forget":
      value = VdiOperationsForget
    case "update":
      value = VdiOperationsUpdate
    case "force_unlock":
      value = VdiOperationsForceUnlock
    case "generate_config":
      value = VdiOperationsGenerateConfig
    case "blocked":
      value = VdiOperationsBlocked
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VdiOperations", context)
	}
	return
}

func convertEnumVdiOperationsToXen(context string, value VdiOperations) (string, error) {
	return string(value), nil
}

func convertEnumVdiTypeToGo(context string, input interface{}) (value VdiType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "system":
      value = VdiTypeSystem
    case "user":
      value = VdiTypeUser
    case "ephemeral":
      value = VdiTypeEphemeral
    case "suspend":
      value = VdiTypeSuspend
    case "crashdump":
      value = VdiTypeCrashdump
    case "ha_statefile":
      value = VdiTypeHaStatefile
    case "metadata":
      value = VdiTypeMetadata
    case "redo_log":
      value = VdiTypeRedoLog
    case "rrd":
      value = VdiTypeRrd
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VdiType", context)
	}
	return
}

func convertEnumVdiTypeToXen(context string, value VdiType) (string, error) {
	return string(value), nil
}

func convertEnumVgpuTypeImplementationToGo(context string, input interface{}) (value VgpuTypeImplementation, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "passthrough":
      value = VgpuTypeImplementationPassthrough
    case "nvidia":
      value = VgpuTypeImplementationNvidia
    case "gvt_g":
      value = VgpuTypeImplementationGvtG
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VgpuTypeImplementation", context)
	}
	return
}

func convertEnumVifLockingModeToGo(context string, input interface{}) (value VifLockingMode, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "network_default":
      value = VifLockingModeNetworkDefault
    case "locked":
      value = VifLockingModeLocked
    case "unlocked":
      value = VifLockingModeUnlocked
    case "disabled":
      value = VifLockingModeDisabled
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VifLockingMode", context)
	}
	return
}

func convertEnumVifLockingModeToXen(context string, value VifLockingMode) (string, error) {
	return string(value), nil
}

func convertEnumVifOperationsSetToGo(context string, input interface{}) (slice []VifOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VifOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVifOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVifOperationsSetToXen(context string, slice []VifOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVifOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVifOperationsToGo(context string, input interface{}) (value VifOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "attach":
      value = VifOperationsAttach
    case "plug":
      value = VifOperationsPlug
    case "unplug":
      value = VifOperationsUnplug
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VifOperations", context)
	}
	return
}

func convertEnumVifOperationsToXen(context string, value VifOperations) (string, error) {
	return string(value), nil
}

func convertEnumVMApplianceOperationSetToGo(context string, input interface{}) (slice []VMApplianceOperation, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMApplianceOperation, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMApplianceOperationToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVMApplianceOperationSetToXen(context string, slice []VMApplianceOperation) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMApplianceOperationToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVMApplianceOperationToGo(context string, input interface{}) (value VMApplianceOperation, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "start":
      value = VMApplianceOperationStart
    case "clean_shutdown":
      value = VMApplianceOperationCleanShutdown
    case "hard_shutdown":
      value = VMApplianceOperationHardShutdown
    case "shutdown":
      value = VMApplianceOperationShutdown
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VMApplianceOperation", context)
	}
	return
}

func convertEnumVMApplianceOperationToXen(context string, value VMApplianceOperation) (string, error) {
	return string(value), nil
}

func convertEnumVMOperationsSetToGo(context string, input interface{}) (slice []VMOperations, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]VMOperations, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMOperationsToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEnumVMOperationsSetToXen(context string, slice []VMOperations) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEnumVMOperationsToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertEnumVMOperationsToGo(context string, input interface{}) (value VMOperations, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "snapshot":
      value = VMOperationsSnapshot
    case "clone":
      value = VMOperationsClone
    case "copy":
      value = VMOperationsCopy
    case "create_template":
      value = VMOperationsCreateTemplate
    case "revert":
      value = VMOperationsRevert
    case "checkpoint":
      value = VMOperationsCheckpoint
    case "snapshot_with_quiesce":
      value = VMOperationsSnapshotWithQuiesce
    case "provision":
      value = VMOperationsProvision
    case "start":
      value = VMOperationsStart
    case "start_on":
      value = VMOperationsStartOn
    case "pause":
      value = VMOperationsPause
    case "unpause":
      value = VMOperationsUnpause
    case "clean_shutdown":
      value = VMOperationsCleanShutdown
    case "clean_reboot":
      value = VMOperationsCleanReboot
    case "hard_shutdown":
      value = VMOperationsHardShutdown
    case "power_state_reset":
      value = VMOperationsPowerStateReset
    case "hard_reboot":
      value = VMOperationsHardReboot
    case "suspend":
      value = VMOperationsSuspend
    case "csvm":
      value = VMOperationsCsvm
    case "resume":
      value = VMOperationsResume
    case "resume_on":
      value = VMOperationsResumeOn
    case "pool_migrate":
      value = VMOperationsPoolMigrate
    case "migrate_send":
      value = VMOperationsMigrateSend
    case "get_boot_record":
      value = VMOperationsGetBootRecord
    case "send_sysrq":
      value = VMOperationsSendSysrq
    case "send_trigger":
      value = VMOperationsSendTrigger
    case "query_services":
      value = VMOperationsQueryServices
    case "shutdown":
      value = VMOperationsShutdown
    case "call_plugin":
      value = VMOperationsCallPlugin
    case "changing_memory_live":
      value = VMOperationsChangingMemoryLive
    case "awaiting_memory_live":
      value = VMOperationsAwaitingMemoryLive
    case "changing_dynamic_range":
      value = VMOperationsChangingDynamicRange
    case "changing_static_range":
      value = VMOperationsChangingStaticRange
    case "changing_memory_limits":
      value = VMOperationsChangingMemoryLimits
    case "changing_shadow_memory":
      value = VMOperationsChangingShadowMemory
    case "changing_shadow_memory_live":
      value = VMOperationsChangingShadowMemoryLive
    case "changing_VCPUs":
      value = VMOperationsChangingVCPUs
    case "changing_VCPUs_live":
      value = VMOperationsChangingVCPUsLive
    case "assert_operation_valid":
      value = VMOperationsAssertOperationValid
    case "data_source_op":
      value = VMOperationsDataSourceOp
    case "update_allowed_operations":
      value = VMOperationsUpdateAllowedOperations
    case "make_into_template":
      value = VMOperationsMakeIntoTemplate
    case "import":
      value = VMOperationsImport
    case "export":
      value = VMOperationsExport
    case "metadata_export":
      value = VMOperationsMetadataExport
    case "reverting":
      value = VMOperationsReverting
    case "destroy":
      value = VMOperationsDestroy
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VMOperations", context)
	}
	return
}

func convertEnumVMOperationsToXen(context string, value VMOperations) (string, error) {
	return string(value), nil
}

func convertEnumVMPowerStateToGo(context string, input interface{}) (value VMPowerState, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "Halted":
      value = VMPowerStateHalted
    case "Paused":
      value = VMPowerStatePaused
    case "Running":
      value = VMPowerStateRunning
    case "Suspended":
      value = VMPowerStateSuspended
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VMPowerState", context)
	}
	return
}

func convertEnumVMPowerStateToXen(context string, value VMPowerState) (string, error) {
	return string(value), nil
}

func convertEnumVmppArchiveFrequencyToGo(context string, input interface{}) (value VmppArchiveFrequency, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "never":
      value = VmppArchiveFrequencyNever
    case "always_after_backup":
      value = VmppArchiveFrequencyAlwaysAfterBackup
    case "daily":
      value = VmppArchiveFrequencyDaily
    case "weekly":
      value = VmppArchiveFrequencyWeekly
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppArchiveFrequency", context)
	}
	return
}

func convertEnumVmppArchiveFrequencyToXen(context string, value VmppArchiveFrequency) (string, error) {
	return string(value), nil
}

func convertEnumVmppArchiveTargetTypeToGo(context string, input interface{}) (value VmppArchiveTargetType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "none":
      value = VmppArchiveTargetTypeNone
    case "cifs":
      value = VmppArchiveTargetTypeCifs
    case "nfs":
      value = VmppArchiveTargetTypeNfs
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppArchiveTargetType", context)
	}
	return
}

func convertEnumVmppArchiveTargetTypeToXen(context string, value VmppArchiveTargetType) (string, error) {
	return string(value), nil
}

func convertEnumVmppBackupFrequencyToGo(context string, input interface{}) (value VmppBackupFrequency, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "hourly":
      value = VmppBackupFrequencyHourly
    case "daily":
      value = VmppBackupFrequencyDaily
    case "weekly":
      value = VmppBackupFrequencyWeekly
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppBackupFrequency", context)
	}
	return
}

func convertEnumVmppBackupFrequencyToXen(context string, value VmppBackupFrequency) (string, error) {
	return string(value), nil
}

func convertEnumVmppBackupTypeToGo(context string, input interface{}) (value VmppBackupType, err error) {
	strValue, err := convertStringToGo(context, input)
	if err != nil {
		return
	}
  switch strValue {
    case "snapshot":
      value = VmppBackupTypeSnapshot
    case "checkpoint":
      value = VmppBackupTypeCheckpoint
    default:
      err = fmt.Errorf("Unable to parse XenAPI response: got value %q for enum %s at %s, but this is not any of the known values", strValue, "VmppBackupType", context)
	}
	return
}

func convertEnumVmppBackupTypeToXen(context string, value VmppBackupType) (string, error) {
	return string(value), nil
}

func convertEventRecordSetToGo(context string, input interface{}) (slice []EventRecord, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]EventRecord, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertEventRecordToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertEventRecordToGo(context string, input interface{}) (record EventRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  idValue, ok := rpcStruct["id"]
	if ok {
  	record.ID, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "id"), idValue)
		if err != nil {
			return
		}
	}
  timestampValue, ok := rpcStruct["timestamp"]
	if ok {
  	record.Timestamp, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp"), timestampValue)
		if err != nil {
			return
		}
	}
  classValue, ok := rpcStruct["class"]
	if ok {
  	record.Class, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "class"), classValue)
		if err != nil {
			return
		}
	}
  operationValue, ok := rpcStruct["operation"]
	if ok {
  	record.Operation, err = convertEnumEventOperationToGo(fmt.Sprintf("%s.%s", context, "operation"), operationValue)
		if err != nil {
			return
		}
	}
  refValue, ok := rpcStruct["ref"]
	if ok {
  	record.Ref, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ref"), refValue)
		if err != nil {
			return
		}
	}
  objUUIDValue, ok := rpcStruct["obj_uuid"]
	if ok {
  	record.ObjUUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "obj_uuid"), objUUIDValue)
		if err != nil {
			return
		}
	}
	return
}

func convertFloatToGo(context string, input interface{}) (value float64, err error) {
	if input == nil {
		return
	}
	value, ok := input.(float64)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "float64", context, reflect.TypeOf(input), input)
	}
	return
}

func convertFloatToXen(context string, value float64) (float64, error) {
	return value, nil
}

func convertHostRecordToGo(context string, input interface{}) (record HostRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  memoryOverheadValue, ok := rpcStruct["memory_overhead"]
	if ok {
  	record.MemoryOverhead, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_overhead"), memoryOverheadValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumHostAllowedOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumHostAllowedOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  apiVersionMajorValue, ok := rpcStruct["API_version_major"]
	if ok {
  	record.APIVersionMajor, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "API_version_major"), apiVersionMajorValue)
		if err != nil {
			return
		}
	}
  apiVersionMinorValue, ok := rpcStruct["API_version_minor"]
	if ok {
  	record.APIVersionMinor, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "API_version_minor"), apiVersionMinorValue)
		if err != nil {
			return
		}
	}
  apiVersionVendorValue, ok := rpcStruct["API_version_vendor"]
	if ok {
  	record.APIVersionVendor, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "API_version_vendor"), apiVersionVendorValue)
		if err != nil {
			return
		}
	}
  apiVersionVendorImplementationValue, ok := rpcStruct["API_version_vendor_implementation"]
	if ok {
  	record.APIVersionVendorImplementation, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "API_version_vendor_implementation"), apiVersionVendorImplementationValue)
		if err != nil {
			return
		}
	}
  enabledValue, ok := rpcStruct["enabled"]
	if ok {
  	record.Enabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "enabled"), enabledValue)
		if err != nil {
			return
		}
	}
  softwareVersionValue, ok := rpcStruct["software_version"]
	if ok {
  	record.SoftwareVersion, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "software_version"), softwareVersionValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  capabilitiesValue, ok := rpcStruct["capabilities"]
	if ok {
  	record.Capabilities, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "capabilities"), capabilitiesValue)
		if err != nil {
			return
		}
	}
  cpuConfigurationValue, ok := rpcStruct["cpu_configuration"]
	if ok {
  	record.CPUConfiguration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "cpu_configuration"), cpuConfigurationValue)
		if err != nil {
			return
		}
	}
  schedPolicyValue, ok := rpcStruct["sched_policy"]
	if ok {
  	record.SchedPolicy, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "sched_policy"), schedPolicyValue)
		if err != nil {
			return
		}
	}
  supportedBootloadersValue, ok := rpcStruct["supported_bootloaders"]
	if ok {
  	record.SupportedBootloaders, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "supported_bootloaders"), supportedBootloadersValue)
		if err != nil {
			return
		}
	}
  residentVMsValue, ok := rpcStruct["resident_VMs"]
	if ok {
  	record.ResidentVMs, err = convertVMRefSetToGo(fmt.Sprintf("%s.%s", context, "resident_VMs"), residentVMsValue)
		if err != nil {
			return
		}
	}
  loggingValue, ok := rpcStruct["logging"]
	if ok {
  	record.Logging, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "logging"), loggingValue)
		if err != nil {
			return
		}
	}
  pifsValue, ok := rpcStruct["PIFs"]
	if ok {
  	record.PIFs, err = convertPIFRefSetToGo(fmt.Sprintf("%s.%s", context, "PIFs"), pifsValue)
		if err != nil {
			return
		}
	}
  suspendImageSrValue, ok := rpcStruct["suspend_image_sr"]
	if ok {
  	record.SuspendImageSr, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "suspend_image_sr"), suspendImageSrValue)
		if err != nil {
			return
		}
	}
  crashDumpSrValue, ok := rpcStruct["crash_dump_sr"]
	if ok {
  	record.CrashDumpSr, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "crash_dump_sr"), crashDumpSrValue)
		if err != nil {
			return
		}
	}
  crashdumpsValue, ok := rpcStruct["crashdumps"]
	if ok {
  	record.Crashdumps, err = convertHostCrashdumpRefSetToGo(fmt.Sprintf("%s.%s", context, "crashdumps"), crashdumpsValue)
		if err != nil {
			return
		}
	}
  patchesValue, ok := rpcStruct["patches"]
	if ok {
  	record.Patches, err = convertHostPatchRefSetToGo(fmt.Sprintf("%s.%s", context, "patches"), patchesValue)
		if err != nil {
			return
		}
	}
  pbdsValue, ok := rpcStruct["PBDs"]
	if ok {
  	record.PBDs, err = convertPBDRefSetToGo(fmt.Sprintf("%s.%s", context, "PBDs"), pbdsValue)
		if err != nil {
			return
		}
	}
  hostCPUsValue, ok := rpcStruct["host_CPUs"]
	if ok {
  	record.HostCPUs, err = convertHostCPURefSetToGo(fmt.Sprintf("%s.%s", context, "host_CPUs"), hostCPUsValue)
		if err != nil {
			return
		}
	}
  cpuInfoValue, ok := rpcStruct["cpu_info"]
	if ok {
  	record.CPUInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "cpu_info"), cpuInfoValue)
		if err != nil {
			return
		}
	}
  hostnameValue, ok := rpcStruct["hostname"]
	if ok {
  	record.Hostname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "hostname"), hostnameValue)
		if err != nil {
			return
		}
	}
  addressValue, ok := rpcStruct["address"]
	if ok {
  	record.Address, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "address"), addressValue)
		if err != nil {
			return
		}
	}
  metricsValue, ok := rpcStruct["metrics"]
	if ok {
  	record.Metrics, err = convertHostMetricsRefToGo(fmt.Sprintf("%s.%s", context, "metrics"), metricsValue)
		if err != nil {
			return
		}
	}
  licenseParamsValue, ok := rpcStruct["license_params"]
	if ok {
  	record.LicenseParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "license_params"), licenseParamsValue)
		if err != nil {
			return
		}
	}
  haStatefilesValue, ok := rpcStruct["ha_statefiles"]
	if ok {
  	record.HaStatefiles, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ha_statefiles"), haStatefilesValue)
		if err != nil {
			return
		}
	}
  haNetworkPeersValue, ok := rpcStruct["ha_network_peers"]
	if ok {
  	record.HaNetworkPeers, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ha_network_peers"), haNetworkPeersValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  externalAuthTypeValue, ok := rpcStruct["external_auth_type"]
	if ok {
  	record.ExternalAuthType, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "external_auth_type"), externalAuthTypeValue)
		if err != nil {
			return
		}
	}
  externalAuthServiceNameValue, ok := rpcStruct["external_auth_service_name"]
	if ok {
  	record.ExternalAuthServiceName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "external_auth_service_name"), externalAuthServiceNameValue)
		if err != nil {
			return
		}
	}
  externalAuthConfigurationValue, ok := rpcStruct["external_auth_configuration"]
	if ok {
  	record.ExternalAuthConfiguration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "external_auth_configuration"), externalAuthConfigurationValue)
		if err != nil {
			return
		}
	}
  editionValue, ok := rpcStruct["edition"]
	if ok {
  	record.Edition, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "edition"), editionValue)
		if err != nil {
			return
		}
	}
  licenseServerValue, ok := rpcStruct["license_server"]
	if ok {
  	record.LicenseServer, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "license_server"), licenseServerValue)
		if err != nil {
			return
		}
	}
  biosStringsValue, ok := rpcStruct["bios_strings"]
	if ok {
  	record.BiosStrings, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "bios_strings"), biosStringsValue)
		if err != nil {
			return
		}
	}
  powerOnModeValue, ok := rpcStruct["power_on_mode"]
	if ok {
  	record.PowerOnMode, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "power_on_mode"), powerOnModeValue)
		if err != nil {
			return
		}
	}
  powerOnConfigValue, ok := rpcStruct["power_on_config"]
	if ok {
  	record.PowerOnConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "power_on_config"), powerOnConfigValue)
		if err != nil {
			return
		}
	}
  localCacheSrValue, ok := rpcStruct["local_cache_sr"]
	if ok {
  	record.LocalCacheSr, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "local_cache_sr"), localCacheSrValue)
		if err != nil {
			return
		}
	}
  chipsetInfoValue, ok := rpcStruct["chipset_info"]
	if ok {
  	record.ChipsetInfo, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "chipset_info"), chipsetInfoValue)
		if err != nil {
			return
		}
	}
  pcisValue, ok := rpcStruct["PCIs"]
	if ok {
  	record.PCIs, err = convertPCIRefSetToGo(fmt.Sprintf("%s.%s", context, "PCIs"), pcisValue)
		if err != nil {
			return
		}
	}
  pgpusValue, ok := rpcStruct["PGPUs"]
	if ok {
  	record.PGPUs, err = convertPGPURefSetToGo(fmt.Sprintf("%s.%s", context, "PGPUs"), pgpusValue)
		if err != nil {
			return
		}
	}
  sslLegacyValue, ok := rpcStruct["ssl_legacy"]
	if ok {
  	record.SslLegacy, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ssl_legacy"), sslLegacyValue)
		if err != nil {
			return
		}
	}
  guestVCPUsParamsValue, ok := rpcStruct["guest_VCPUs_params"]
	if ok {
  	record.GuestVCPUsParams, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "guest_VCPUs_params"), guestVCPUsParamsValue)
		if err != nil {
			return
		}
	}
  displayValue, ok := rpcStruct["display"]
	if ok {
  	record.Display, err = convertEnumHostDisplayToGo(fmt.Sprintf("%s.%s", context, "display"), displayValue)
		if err != nil {
			return
		}
	}
  virtualHardwarePlatformVersionsValue, ok := rpcStruct["virtual_hardware_platform_versions"]
	if ok {
  	record.VirtualHardwarePlatformVersions, err = convertIntSetToGo(fmt.Sprintf("%s.%s", context, "virtual_hardware_platform_versions"), virtualHardwarePlatformVersionsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostRefSetToGo(context string, input interface{}) (slice []HostRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostRefSetToXen(context string, slice []HostRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertHostRefToGo(context string, input interface{}) (ref HostRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostRef(value)
	}
	return
}

func convertHostRefToXen(context string, ref HostRef) (string, error) {
	return string(ref), nil
}

func convertHostCPURecordToGo(context string, input interface{}) (record HostCPURecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  numberValue, ok := rpcStruct["number"]
	if ok {
  	record.Number, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "number"), numberValue)
		if err != nil {
			return
		}
	}
  vendorValue, ok := rpcStruct["vendor"]
	if ok {
  	record.Vendor, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vendor"), vendorValue)
		if err != nil {
			return
		}
	}
  speedValue, ok := rpcStruct["speed"]
	if ok {
  	record.Speed, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "speed"), speedValue)
		if err != nil {
			return
		}
	}
  modelnameValue, ok := rpcStruct["modelname"]
	if ok {
  	record.Modelname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "modelname"), modelnameValue)
		if err != nil {
			return
		}
	}
  familyValue, ok := rpcStruct["family"]
	if ok {
  	record.Family, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "family"), familyValue)
		if err != nil {
			return
		}
	}
  modelValue, ok := rpcStruct["model"]
	if ok {
  	record.Model, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "model"), modelValue)
		if err != nil {
			return
		}
	}
  steppingValue, ok := rpcStruct["stepping"]
	if ok {
  	record.Stepping, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "stepping"), steppingValue)
		if err != nil {
			return
		}
	}
  flagsValue, ok := rpcStruct["flags"]
	if ok {
  	record.Flags, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "flags"), flagsValue)
		if err != nil {
			return
		}
	}
  featuresValue, ok := rpcStruct["features"]
	if ok {
  	record.Features, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "features"), featuresValue)
		if err != nil {
			return
		}
	}
  utilisationValue, ok := rpcStruct["utilisation"]
	if ok {
  	record.Utilisation, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "utilisation"), utilisationValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostCPURefSetToGo(context string, input interface{}) (slice []HostCPURef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostCPURef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostCPURefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostCPURefToGo(context string, input interface{}) (ref HostCPURef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostCPURef(value)
	}
	return
}

func convertHostCPURefToXen(context string, ref HostCPURef) (string, error) {
	return string(ref), nil
}

func convertHostCrashdumpRecordToGo(context string, input interface{}) (record HostCrashdumpRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  timestampValue, ok := rpcStruct["timestamp"]
	if ok {
  	record.Timestamp, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp"), timestampValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostCrashdumpRefSetToGo(context string, input interface{}) (slice []HostCrashdumpRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostCrashdumpRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostCrashdumpRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostCrashdumpRefToGo(context string, input interface{}) (ref HostCrashdumpRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostCrashdumpRef(value)
	}
	return
}

func convertHostCrashdumpRefToXen(context string, ref HostCrashdumpRef) (string, error) {
	return string(ref), nil
}

func convertHostMetricsRecordToGo(context string, input interface{}) (record HostMetricsRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  memoryTotalValue, ok := rpcStruct["memory_total"]
	if ok {
  	record.MemoryTotal, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_total"), memoryTotalValue)
		if err != nil {
			return
		}
	}
  memoryFreeValue, ok := rpcStruct["memory_free"]
	if ok {
  	record.MemoryFree, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "memory_free"), memoryFreeValue)
		if err != nil {
			return
		}
	}
  liveValue, ok := rpcStruct["live"]
	if ok {
  	record.Live, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "live"), liveValue)
		if err != nil {
			return
		}
	}
  lastUpdatedValue, ok := rpcStruct["last_updated"]
	if ok {
  	record.LastUpdated, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_updated"), lastUpdatedValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostMetricsRefSetToGo(context string, input interface{}) (slice []HostMetricsRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostMetricsRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostMetricsRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostMetricsRefToGo(context string, input interface{}) (ref HostMetricsRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostMetricsRef(value)
	}
	return
}

func convertHostMetricsRefToXen(context string, ref HostMetricsRef) (string, error) {
	return string(ref), nil
}

func convertHostPatchRecordToGo(context string, input interface{}) (record HostPatchRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  hostValue, ok := rpcStruct["host"]
	if ok {
  	record.Host, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "host"), hostValue)
		if err != nil {
			return
		}
	}
  appliedValue, ok := rpcStruct["applied"]
	if ok {
  	record.Applied, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "applied"), appliedValue)
		if err != nil {
			return
		}
	}
  timestampAppliedValue, ok := rpcStruct["timestamp_applied"]
	if ok {
  	record.TimestampApplied, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp_applied"), timestampAppliedValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  poolPatchValue, ok := rpcStruct["pool_patch"]
	if ok {
  	record.PoolPatch, err = convertPoolPatchRefToGo(fmt.Sprintf("%s.%s", context, "pool_patch"), poolPatchValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertHostPatchRefSetToGo(context string, input interface{}) (slice []HostPatchRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]HostPatchRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertHostPatchRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertHostPatchRefToGo(context string, input interface{}) (ref HostPatchRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = HostPatchRef(value)
	}
	return
}

func convertHostPatchRefToXen(context string, ref HostPatchRef) (string, error) {
	return string(ref), nil
}

func convertIntSetToGo(context string, input interface{}) (slice []int, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]int, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertIntToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertIntToGo(context string, input interface{}) (value int, err error) {
	strValue, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
  	value, err = strconv.Atoi(strValue)
	}
	return
}

func convertIntToXen(context string, value int) (string, error) {
	return strconv.Itoa(value), nil
}

func convertMessageRecordToGo(context string, input interface{}) (record MessageRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameValue, ok := rpcStruct["name"]
	if ok {
  	record.Name, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name"), nameValue)
		if err != nil {
			return
		}
	}
  priorityValue, ok := rpcStruct["priority"]
	if ok {
  	record.Priority, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "priority"), priorityValue)
		if err != nil {
			return
		}
	}
  clsValue, ok := rpcStruct["cls"]
	if ok {
  	record.Cls, err = convertEnumClsToGo(fmt.Sprintf("%s.%s", context, "cls"), clsValue)
		if err != nil {
			return
		}
	}
  objUUIDValue, ok := rpcStruct["obj_uuid"]
	if ok {
  	record.ObjUUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "obj_uuid"), objUUIDValue)
		if err != nil {
			return
		}
	}
  timestampValue, ok := rpcStruct["timestamp"]
	if ok {
  	record.Timestamp, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "timestamp"), timestampValue)
		if err != nil {
			return
		}
	}
  bodyValue, ok := rpcStruct["body"]
	if ok {
  	record.Body, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "body"), bodyValue)
		if err != nil {
			return
		}
	}
	return
}

func convertMessageRefSetToGo(context string, input interface{}) (slice []MessageRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]MessageRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertMessageRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertMessageRefToGo(context string, input interface{}) (ref MessageRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = MessageRef(value)
	}
	return
}

func convertMessageRefToXen(context string, ref MessageRef) (string, error) {
	return string(ref), nil
}

func convertNetworkRecordToGo(context string, input interface{}) (record NetworkRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumNetworkOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumNetworkOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  vifsValue, ok := rpcStruct["VIFs"]
	if ok {
  	record.VIFs, err = convertVIFRefSetToGo(fmt.Sprintf("%s.%s", context, "VIFs"), vifsValue)
		if err != nil {
			return
		}
	}
  pifsValue, ok := rpcStruct["PIFs"]
	if ok {
  	record.PIFs, err = convertPIFRefSetToGo(fmt.Sprintf("%s.%s", context, "PIFs"), pifsValue)
		if err != nil {
			return
		}
	}
  mtuValue, ok := rpcStruct["MTU"]
	if ok {
  	record.MTU, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "MTU"), mtuValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  bridgeValue, ok := rpcStruct["bridge"]
	if ok {
  	record.Bridge, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "bridge"), bridgeValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  defaultLockingModeValue, ok := rpcStruct["default_locking_mode"]
	if ok {
  	record.DefaultLockingMode, err = convertEnumNetworkDefaultLockingModeToGo(fmt.Sprintf("%s.%s", context, "default_locking_mode"), defaultLockingModeValue)
		if err != nil {
			return
		}
	}
  assignedIpsValue, ok := rpcStruct["assigned_ips"]
	if ok {
  	record.AssignedIps, err = convertVIFRefToStringMapToGo(fmt.Sprintf("%s.%s", context, "assigned_ips"), assignedIpsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertNetworkRecordToXen(context string, record NetworkRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["name_label"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_label"), record.NameLabel)
  if err != nil {
		return
	}
  rpcStruct["name_description"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "name_description"), record.NameDescription)
  if err != nil {
		return
	}
  rpcStruct["allowed_operations"], err = convertEnumNetworkOperationsSetToXen(fmt.Sprintf("%s.%s", context, "allowed_operations"), record.AllowedOperations)
  if err != nil {
		return
	}
  rpcStruct["current_operations"], err = convertStringToEnumNetworkOperationsMapToXen(fmt.Sprintf("%s.%s", context, "current_operations"), record.CurrentOperations)
  if err != nil {
		return
	}
  rpcStruct["VIFs"], err = convertVIFRefSetToXen(fmt.Sprintf("%s.%s", context, "VIFs"), record.VIFs)
  if err != nil {
		return
	}
  rpcStruct["PIFs"], err = convertPIFRefSetToXen(fmt.Sprintf("%s.%s", context, "PIFs"), record.PIFs)
  if err != nil {
		return
	}
  rpcStruct["MTU"], err = convertIntToXen(fmt.Sprintf("%s.%s", context, "MTU"), record.MTU)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["bridge"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "bridge"), record.Bridge)
  if err != nil {
		return
	}
  rpcStruct["blobs"], err = convertStringToBlobRefMapToXen(fmt.Sprintf("%s.%s", context, "blobs"), record.Blobs)
  if err != nil {
		return
	}
  rpcStruct["tags"], err = convertStringSetToXen(fmt.Sprintf("%s.%s", context, "tags"), record.Tags)
  if err != nil {
		return
	}
  rpcStruct["default_locking_mode"], err = convertEnumNetworkDefaultLockingModeToXen(fmt.Sprintf("%s.%s", context, "default_locking_mode"), record.DefaultLockingMode)
  if err != nil {
		return
	}
  rpcStruct["assigned_ips"], err = convertVIFRefToStringMapToXen(fmt.Sprintf("%s.%s", context, "assigned_ips"), record.AssignedIps)
  if err != nil {
		return
	}
	return
}

func convertNetworkRefSetToGo(context string, input interface{}) (slice []NetworkRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]NetworkRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertNetworkRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertNetworkRefToGo(context string, input interface{}) (ref NetworkRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = NetworkRef(value)
	}
	return
}

func convertNetworkRefToXen(context string, ref NetworkRef) (string, error) {
	return string(ref), nil
}

func convertPoolRecordToGo(context string, input interface{}) (record PoolRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  masterValue, ok := rpcStruct["master"]
	if ok {
  	record.Master, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "master"), masterValue)
		if err != nil {
			return
		}
	}
  defaultSRValue, ok := rpcStruct["default_SR"]
	if ok {
  	record.DefaultSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "default_SR"), defaultSRValue)
		if err != nil {
			return
		}
	}
  suspendImageSRValue, ok := rpcStruct["suspend_image_SR"]
	if ok {
  	record.SuspendImageSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "suspend_image_SR"), suspendImageSRValue)
		if err != nil {
			return
		}
	}
  crashDumpSRValue, ok := rpcStruct["crash_dump_SR"]
	if ok {
  	record.CrashDumpSR, err = convertSRRefToGo(fmt.Sprintf("%s.%s", context, "crash_dump_SR"), crashDumpSRValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  haEnabledValue, ok := rpcStruct["ha_enabled"]
	if ok {
  	record.HaEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_enabled"), haEnabledValue)
		if err != nil {
			return
		}
	}
  haConfigurationValue, ok := rpcStruct["ha_configuration"]
	if ok {
  	record.HaConfiguration, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "ha_configuration"), haConfigurationValue)
		if err != nil {
			return
		}
	}
  haStatefilesValue, ok := rpcStruct["ha_statefiles"]
	if ok {
  	record.HaStatefiles, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "ha_statefiles"), haStatefilesValue)
		if err != nil {
			return
		}
	}
  haHostFailuresToTolerateValue, ok := rpcStruct["ha_host_failures_to_tolerate"]
	if ok {
  	record.HaHostFailuresToTolerate, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "ha_host_failures_to_tolerate"), haHostFailuresToTolerateValue)
		if err != nil {
			return
		}
	}
  haPlanExistsForValue, ok := rpcStruct["ha_plan_exists_for"]
	if ok {
  	record.HaPlanExistsFor, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "ha_plan_exists_for"), haPlanExistsForValue)
		if err != nil {
			return
		}
	}
  haAllowOvercommitValue, ok := rpcStruct["ha_allow_overcommit"]
	if ok {
  	record.HaAllowOvercommit, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_allow_overcommit"), haAllowOvercommitValue)
		if err != nil {
			return
		}
	}
  haOvercommittedValue, ok := rpcStruct["ha_overcommitted"]
	if ok {
  	record.HaOvercommitted, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "ha_overcommitted"), haOvercommittedValue)
		if err != nil {
			return
		}
	}
  blobsValue, ok := rpcStruct["blobs"]
	if ok {
  	record.Blobs, err = convertStringToBlobRefMapToGo(fmt.Sprintf("%s.%s", context, "blobs"), blobsValue)
		if err != nil {
			return
		}
	}
  tagsValue, ok := rpcStruct["tags"]
	if ok {
  	record.Tags, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "tags"), tagsValue)
		if err != nil {
			return
		}
	}
  guiConfigValue, ok := rpcStruct["gui_config"]
	if ok {
  	record.GuiConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "gui_config"), guiConfigValue)
		if err != nil {
			return
		}
	}
  healthCheckConfigValue, ok := rpcStruct["health_check_config"]
	if ok {
  	record.HealthCheckConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "health_check_config"), healthCheckConfigValue)
		if err != nil {
			return
		}
	}
  wlbURLValue, ok := rpcStruct["wlb_url"]
	if ok {
  	record.WlbURL, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "wlb_url"), wlbURLValue)
		if err != nil {
			return
		}
	}
  wlbUsernameValue, ok := rpcStruct["wlb_username"]
	if ok {
  	record.WlbUsername, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "wlb_username"), wlbUsernameValue)
		if err != nil {
			return
		}
	}
  wlbEnabledValue, ok := rpcStruct["wlb_enabled"]
	if ok {
  	record.WlbEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "wlb_enabled"), wlbEnabledValue)
		if err != nil {
			return
		}
	}
  wlbVerifyCertValue, ok := rpcStruct["wlb_verify_cert"]
	if ok {
  	record.WlbVerifyCert, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "wlb_verify_cert"), wlbVerifyCertValue)
		if err != nil {
			return
		}
	}
  redoLogEnabledValue, ok := rpcStruct["redo_log_enabled"]
	if ok {
  	record.RedoLogEnabled, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "redo_log_enabled"), redoLogEnabledValue)
		if err != nil {
			return
		}
	}
  redoLogVdiValue, ok := rpcStruct["redo_log_vdi"]
	if ok {
  	record.RedoLogVdi, err = convertVDIRefToGo(fmt.Sprintf("%s.%s", context, "redo_log_vdi"), redoLogVdiValue)
		if err != nil {
			return
		}
	}
  vswitchControllerValue, ok := rpcStruct["vswitch_controller"]
	if ok {
  	record.VswitchController, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "vswitch_controller"), vswitchControllerValue)
		if err != nil {
			return
		}
	}
  restrictionsValue, ok := rpcStruct["restrictions"]
	if ok {
  	record.Restrictions, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "restrictions"), restrictionsValue)
		if err != nil {
			return
		}
	}
  metadataVDIsValue, ok := rpcStruct["metadata_VDIs"]
	if ok {
  	record.MetadataVDIs, err = convertVDIRefSetToGo(fmt.Sprintf("%s.%s", context, "metadata_VDIs"), metadataVDIsValue)
		if err != nil {
			return
		}
	}
  haClusterStackValue, ok := rpcStruct["ha_cluster_stack"]
	if ok {
  	record.HaClusterStack, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "ha_cluster_stack"), haClusterStackValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumPoolAllowedOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumPoolAllowedOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPoolRefSetToGo(context string, input interface{}) (slice []PoolRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPoolRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPoolRefToGo(context string, input interface{}) (ref PoolRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PoolRef(value)
	}
	return
}

func convertPoolRefToXen(context string, ref PoolRef) (string, error) {
	return string(ref), nil
}

func convertPoolPatchRecordToGo(context string, input interface{}) (record PoolPatchRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  versionValue, ok := rpcStruct["version"]
	if ok {
  	record.Version, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "version"), versionValue)
		if err != nil {
			return
		}
	}
  sizeValue, ok := rpcStruct["size"]
	if ok {
  	record.Size, err = convertIntToGo(fmt.Sprintf("%s.%s", context, "size"), sizeValue)
		if err != nil {
			return
		}
	}
  poolAppliedValue, ok := rpcStruct["pool_applied"]
	if ok {
  	record.PoolApplied, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "pool_applied"), poolAppliedValue)
		if err != nil {
			return
		}
	}
  hostPatchesValue, ok := rpcStruct["host_patches"]
	if ok {
  	record.HostPatches, err = convertHostPatchRefSetToGo(fmt.Sprintf("%s.%s", context, "host_patches"), hostPatchesValue)
		if err != nil {
			return
		}
	}
  afterApplyGuidanceValue, ok := rpcStruct["after_apply_guidance"]
	if ok {
  	record.AfterApplyGuidance, err = convertEnumAfterApplyGuidanceSetToGo(fmt.Sprintf("%s.%s", context, "after_apply_guidance"), afterApplyGuidanceValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertPoolPatchRefSetToGo(context string, input interface{}) (slice []PoolPatchRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]PoolPatchRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertPoolPatchRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertPoolPatchRefToGo(context string, input interface{}) (ref PoolPatchRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = PoolPatchRef(value)
	}
	return
}

func convertPoolPatchRefToXen(context string, ref PoolPatchRef) (string, error) {
	return string(ref), nil
}

func convertRoleRecordToGo(context string, input interface{}) (record RoleRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  subrolesValue, ok := rpcStruct["subroles"]
	if ok {
  	record.Subroles, err = convertRoleRefSetToGo(fmt.Sprintf("%s.%s", context, "subroles"), subrolesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertRoleRefSetToGo(context string, input interface{}) (slice []RoleRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]RoleRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertRoleRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertRoleRefSetToXen(context string, slice []RoleRef) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertRoleRefToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertRoleRefToGo(context string, input interface{}) (ref RoleRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = RoleRef(value)
	}
	return
}

func convertRoleRefToXen(context string, ref RoleRef) (string, error) {
	return string(ref), nil
}

func convertSecretRecordToGo(context string, input interface{}) (record SecretRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  valueValue, ok := rpcStruct["value"]
	if ok {
  	record.Value, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "value"), valueValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSecretRecordToXen(context string, record SecretRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["value"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "value"), record.Value)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertSecretRefSetToGo(context string, input interface{}) (slice []SecretRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SecretRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSecretRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSecretRefToGo(context string, input interface{}) (ref SecretRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SecretRef(value)
	}
	return
}

func convertSecretRefToXen(context string, ref SecretRef) (string, error) {
	return string(ref), nil
}

func convertSessionRecordToGo(context string, input interface{}) (record SessionRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  thisHostValue, ok := rpcStruct["this_host"]
	if ok {
  	record.ThisHost, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "this_host"), thisHostValue)
		if err != nil {
			return
		}
	}
  thisUserValue, ok := rpcStruct["this_user"]
	if ok {
  	record.ThisUser, err = convertUserRefToGo(fmt.Sprintf("%s.%s", context, "this_user"), thisUserValue)
		if err != nil {
			return
		}
	}
  lastActiveValue, ok := rpcStruct["last_active"]
	if ok {
  	record.LastActive, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "last_active"), lastActiveValue)
		if err != nil {
			return
		}
	}
  poolValue, ok := rpcStruct["pool"]
	if ok {
  	record.Pool, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "pool"), poolValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  isLocalSuperuserValue, ok := rpcStruct["is_local_superuser"]
	if ok {
  	record.IsLocalSuperuser, err = convertBoolToGo(fmt.Sprintf("%s.%s", context, "is_local_superuser"), isLocalSuperuserValue)
		if err != nil {
			return
		}
	}
  subjectValue, ok := rpcStruct["subject"]
	if ok {
  	record.Subject, err = convertSubjectRefToGo(fmt.Sprintf("%s.%s", context, "subject"), subjectValue)
		if err != nil {
			return
		}
	}
  validationTimeValue, ok := rpcStruct["validation_time"]
	if ok {
  	record.ValidationTime, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "validation_time"), validationTimeValue)
		if err != nil {
			return
		}
	}
  authUserSidValue, ok := rpcStruct["auth_user_sid"]
	if ok {
  	record.AuthUserSid, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "auth_user_sid"), authUserSidValue)
		if err != nil {
			return
		}
	}
  authUserNameValue, ok := rpcStruct["auth_user_name"]
	if ok {
  	record.AuthUserName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "auth_user_name"), authUserNameValue)
		if err != nil {
			return
		}
	}
  rbacPermissionsValue, ok := rpcStruct["rbac_permissions"]
	if ok {
  	record.RbacPermissions, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "rbac_permissions"), rbacPermissionsValue)
		if err != nil {
			return
		}
	}
  tasksValue, ok := rpcStruct["tasks"]
	if ok {
  	record.Tasks, err = convertTaskRefSetToGo(fmt.Sprintf("%s.%s", context, "tasks"), tasksValue)
		if err != nil {
			return
		}
	}
  parentValue, ok := rpcStruct["parent"]
	if ok {
  	record.Parent, err = convertSessionRefToGo(fmt.Sprintf("%s.%s", context, "parent"), parentValue)
		if err != nil {
			return
		}
	}
  originatorValue, ok := rpcStruct["originator"]
	if ok {
  	record.Originator, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "originator"), originatorValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSessionRefToGo(context string, input interface{}) (ref SessionRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SessionRef(value)
	}
	return
}

func convertSessionRefToXen(context string, ref SessionRef) (string, error) {
	return string(ref), nil
}

func convertStringSetToGo(context string, input interface{}) (slice []string, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]string, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertStringToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertStringSetToXen(context string, slice []string) (set []interface{}, err error) {
	set = make([]interface{}, len(slice))
	for index, item := range slice {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertStringToXen(itemContext, item)
		if err != nil {
			return set, err
		}
		set[index] = itemValue
	}
	return
}

func convertStringToGo(context string, input interface{}) (value string, err error) {
	if input == nil {
		return
	}
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	}
	return
}

func convertStringToXen(context string, value string) (string, error) {
	return value, nil
}

func convertSubjectRecordToGo(context string, input interface{}) (record SubjectRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  subjectIdentifierValue, ok := rpcStruct["subject_identifier"]
	if ok {
  	record.SubjectIdentifier, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "subject_identifier"), subjectIdentifierValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  rolesValue, ok := rpcStruct["roles"]
	if ok {
  	record.Roles, err = convertRoleRefSetToGo(fmt.Sprintf("%s.%s", context, "roles"), rolesValue)
		if err != nil {
			return
		}
	}
	return
}

func convertSubjectRecordToXen(context string, record SubjectRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["subject_identifier"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "subject_identifier"), record.SubjectIdentifier)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
  rpcStruct["roles"], err = convertRoleRefSetToXen(fmt.Sprintf("%s.%s", context, "roles"), record.Roles)
  if err != nil {
		return
	}
	return
}

func convertSubjectRefSetToGo(context string, input interface{}) (slice []SubjectRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]SubjectRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertSubjectRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertSubjectRefToGo(context string, input interface{}) (ref SubjectRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = SubjectRef(value)
	}
	return
}

func convertSubjectRefToXen(context string, ref SubjectRef) (string, error) {
	return string(ref), nil
}

func convertTaskRecordToGo(context string, input interface{}) (record TaskRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  nameLabelValue, ok := rpcStruct["name_label"]
	if ok {
  	record.NameLabel, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_label"), nameLabelValue)
		if err != nil {
			return
		}
	}
  nameDescriptionValue, ok := rpcStruct["name_description"]
	if ok {
  	record.NameDescription, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "name_description"), nameDescriptionValue)
		if err != nil {
			return
		}
	}
  allowedOperationsValue, ok := rpcStruct["allowed_operations"]
	if ok {
  	record.AllowedOperations, err = convertEnumTaskAllowedOperationsSetToGo(fmt.Sprintf("%s.%s", context, "allowed_operations"), allowedOperationsValue)
		if err != nil {
			return
		}
	}
  currentOperationsValue, ok := rpcStruct["current_operations"]
	if ok {
  	record.CurrentOperations, err = convertStringToEnumTaskAllowedOperationsMapToGo(fmt.Sprintf("%s.%s", context, "current_operations"), currentOperationsValue)
		if err != nil {
			return
		}
	}
  createdValue, ok := rpcStruct["created"]
	if ok {
  	record.Created, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "created"), createdValue)
		if err != nil {
			return
		}
	}
  finishedValue, ok := rpcStruct["finished"]
	if ok {
  	record.Finished, err = convertTimeToGo(fmt.Sprintf("%s.%s", context, "finished"), finishedValue)
		if err != nil {
			return
		}
	}
  statusValue, ok := rpcStruct["status"]
	if ok {
  	record.Status, err = convertEnumTaskStatusTypeToGo(fmt.Sprintf("%s.%s", context, "status"), statusValue)
		if err != nil {
			return
		}
	}
  residentOnValue, ok := rpcStruct["resident_on"]
	if ok {
  	record.ResidentOn, err = convertHostRefToGo(fmt.Sprintf("%s.%s", context, "resident_on"), residentOnValue)
		if err != nil {
			return
		}
	}
  progressValue, ok := rpcStruct["progress"]
	if ok {
  	record.Progress, err = convertFloatToGo(fmt.Sprintf("%s.%s", context, "progress"), progressValue)
		if err != nil {
			return
		}
	}
  atypeValue, ok := rpcStruct["type"]
	if ok {
  	record.Type, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "type"), atypeValue)
		if err != nil {
			return
		}
	}
  resultValue, ok := rpcStruct["result"]
	if ok {
  	record.Result, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "result"), resultValue)
		if err != nil {
			return
		}
	}
  errorInfoValue, ok := rpcStruct["error_info"]
	if ok {
  	record.ErrorInfo, err = convertStringSetToGo(fmt.Sprintf("%s.%s", context, "error_info"), errorInfoValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
  subtaskOfValue, ok := rpcStruct["subtask_of"]
	if ok {
  	record.SubtaskOf, err = convertTaskRefToGo(fmt.Sprintf("%s.%s", context, "subtask_of"), subtaskOfValue)
		if err != nil {
			return
		}
	}
  subtasksValue, ok := rpcStruct["subtasks"]
	if ok {
  	record.Subtasks, err = convertTaskRefSetToGo(fmt.Sprintf("%s.%s", context, "subtasks"), subtasksValue)
		if err != nil {
			return
		}
	}
  backtraceValue, ok := rpcStruct["backtrace"]
	if ok {
  	record.Backtrace, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "backtrace"), backtraceValue)
		if err != nil {
			return
		}
	}
	return
}

func convertTaskRefSetToGo(context string, input interface{}) (slice []TaskRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]TaskRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertTaskRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertTaskRefToGo(context string, input interface{}) (ref TaskRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = TaskRef(value)
	}
	return
}

func convertTaskRefToXen(context string, ref TaskRef) (string, error) {
	return string(ref), nil
}

func convertTunnelRecordToGo(context string, input interface{}) (record TunnelRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  accessPIFValue, ok := rpcStruct["access_PIF"]
	if ok {
  	record.AccessPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "access_PIF"), accessPIFValue)
		if err != nil {
			return
		}
	}
  transportPIFValue, ok := rpcStruct["transport_PIF"]
	if ok {
  	record.TransportPIF, err = convertPIFRefToGo(fmt.Sprintf("%s.%s", context, "transport_PIF"), transportPIFValue)
		if err != nil {
			return
		}
	}
  statusValue, ok := rpcStruct["status"]
	if ok {
  	record.Status, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "status"), statusValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertTunnelRefSetToGo(context string, input interface{}) (slice []TunnelRef, err error) {
	set, ok := input.([]interface{})
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "[]interface{}", context, reflect.TypeOf(input), input)
		return
	}
	slice = make([]TunnelRef, len(set))
	for index, item := range set {
		itemContext := fmt.Sprintf("%s[%d]", context, index)
		itemValue, err := convertTunnelRefToGo(itemContext, item)
		if err != nil {
			return slice, err
		}
		slice[index] = itemValue
	}
	return
}

func convertTunnelRefToGo(context string, input interface{}) (ref TunnelRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = TunnelRef(value)
	}
	return
}

func convertTunnelRefToXen(context string, ref TunnelRef) (string, error) {
	return string(ref), nil
}

func convertUserRecordToGo(context string, input interface{}) (record UserRecord, err error) {
	rpcStruct, ok := input.(xmlrpc.Struct)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "xmlrpc.Struct", context, reflect.TypeOf(input), input)
		return
	}
  uuidValue, ok := rpcStruct["uuid"]
	if ok {
  	record.UUID, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "uuid"), uuidValue)
		if err != nil {
			return
		}
	}
  shortNameValue, ok := rpcStruct["short_name"]
	if ok {
  	record.ShortName, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "short_name"), shortNameValue)
		if err != nil {
			return
		}
	}
  fullnameValue, ok := rpcStruct["fullname"]
	if ok {
  	record.Fullname, err = convertStringToGo(fmt.Sprintf("%s.%s", context, "fullname"), fullnameValue)
		if err != nil {
			return
		}
	}
  otherConfigValue, ok := rpcStruct["other_config"]
	if ok {
  	record.OtherConfig, err = convertStringToStringMapToGo(fmt.Sprintf("%s.%s", context, "other_config"), otherConfigValue)
		if err != nil {
			return
		}
	}
	return
}

func convertUserRecordToXen(context string, record UserRecord) (rpcStruct xmlrpc.Struct, err error) {
  rpcStruct = make(xmlrpc.Struct)
  rpcStruct["uuid"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "uuid"), record.UUID)
  if err != nil {
		return
	}
  rpcStruct["short_name"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "short_name"), record.ShortName)
  if err != nil {
		return
	}
  rpcStruct["fullname"], err = convertStringToXen(fmt.Sprintf("%s.%s", context, "fullname"), record.Fullname)
  if err != nil {
		return
	}
  rpcStruct["other_config"], err = convertStringToStringMapToXen(fmt.Sprintf("%s.%s", context, "other_config"), record.OtherConfig)
  if err != nil {
		return
	}
	return
}

func convertUserRefToGo(context string, input interface{}) (ref UserRef, err error) {
	value, ok := input.(string)
	if !ok {
		err = fmt.Errorf("Failed to parse XenAPI response: expected Go type %s at %s but got Go type %s with value %v", "string", context, reflect.TypeOf(input), input)
	} else {
		ref = UserRef(value)
	}
	return
}

func convertUserRefToXen(context string, ref UserRef) (string, error) {
	return string(ref), nil
}
