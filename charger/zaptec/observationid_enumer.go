// Code generated by "enumer -type ObservationID"; DO NOT EDIT.

//
package zaptec

import (
	"fmt"
)

const _ObservationIDName = "IsOcppConnectedIsOnlinePulseUnknownOfflineModeAuthenticationRequiredPaymentActivePaymentCurrencyPaymentSessionUnitPricePaymentEnergyUnitPricePaymentTimeUnitPriceCommunicationModePermanentCableLockProductCodeHmiBrightnessLockCableWhenConnectedSoftStartDisabledFirmwareApiHostMIDBlinkEnabledTemperatureInternal5TemperatureInternal6TemperatureInternalLimitTemperatureInternalMaxLimitHumidityVoltagePhase1VoltagePhase2VoltagePhase3CurrentPhase1CurrentPhase2CurrentPhase3ChargerMaxCurrentChargerMinCurrentActivePhasesTotalChargePowerRcdCurrentInternal12vCurrentPowerFactorSetPhasesMaxPhasesChargerOfflinePhaseChargerOfflineCurrentRcdCalibrationRcdCalibrationNoiseTotalChargePowerSessionSignedMeterValueSignedMeterValueIntervalSessionEnergyCountExportActiveSessionEnergyCountExportReactiveSessionEnergyCountImportActiveSessionEnergyCountImportReactiveSoftStartTimeChargeDurationChargeModeChargePilotLevelInstantChargePilotLevelAveragePilotVsProximityTimeChargeCurrentInstallationMaxLimitChargeCurrentSetChargerOperationModeIsEnabledIsStandAloneChargerCurrentUserUuidDeprecatedCableTypeNetworkTypeDetectedCarGridTestResultFinalStopActiveSessionIdentifierChargerCurrentUserUuidCompletedSessionNewChargeCardAuthenticationListVersionEnabledNfcTechnologiesLteRoamingDisabledInstallationIdRoutingIdNotificationsWarningsDiagnosticsModeInternalDiagnosticsLogDiagnosticsStringCommunicationSignalStrengthCloudConnectionStatusMcuResetSourceMcuRxErrorsMcuToVariscitePacketErrorsVarisciteToMcuPacketErrorsUptimeVarisciteUptimeMCUCarSessionLogCommunicationModeConfigurationInconsistencyRawPilotMonitorIT3PhaseDiagnosticsLogPilotTestResultsUnconditionalNfcDetectionIndicationEmcTestCounterProductionTestResultsPostProductionTestResultsSmartMainboardSoftwareApplicationVersionSmartMainboardSoftwareBootloaderVersionSmartComputerSoftwareApplicationVersionSmartComputerSoftwareBootloaderVersionSmartComputerHardwareVersionMacMainMacPlcModuleGridMacWiFiMacPlcModuleEvLteImsiLteMsisdnLteIccidLteImeiMIDCalibration"

var _ObservationIDMap = map[ObservationID]string{
	-3:  _ObservationIDName[0:15],
	-2:  _ObservationIDName[15:23],
	-1:  _ObservationIDName[23:28],
	0:   _ObservationIDName[28:35],
	1:   _ObservationIDName[35:46],
	120: _ObservationIDName[46:68],
	130: _ObservationIDName[68:81],
	131: _ObservationIDName[81:96],
	132: _ObservationIDName[96:119],
	133: _ObservationIDName[119:141],
	134: _ObservationIDName[141:161],
	150: _ObservationIDName[161:178],
	151: _ObservationIDName[178:196],
	152: _ObservationIDName[196:207],
	153: _ObservationIDName[207:220],
	154: _ObservationIDName[220:242],
	155: _ObservationIDName[242:259],
	156: _ObservationIDName[259:274],
	170: _ObservationIDName[274:289],
	201: _ObservationIDName[289:309],
	202: _ObservationIDName[309:329],
	203: _ObservationIDName[329:353],
	241: _ObservationIDName[353:380],
	270: _ObservationIDName[380:388],
	501: _ObservationIDName[388:401],
	502: _ObservationIDName[401:414],
	503: _ObservationIDName[414:427],
	507: _ObservationIDName[427:440],
	508: _ObservationIDName[440:453],
	509: _ObservationIDName[453:466],
	510: _ObservationIDName[466:483],
	511: _ObservationIDName[483:500],
	512: _ObservationIDName[500:512],
	513: _ObservationIDName[512:528],
	515: _ObservationIDName[528:538],
	517: _ObservationIDName[538:556],
	518: _ObservationIDName[556:567],
	519: _ObservationIDName[567:576],
	520: _ObservationIDName[576:585],
	522: _ObservationIDName[585:604],
	523: _ObservationIDName[604:625],
	540: _ObservationIDName[625:639],
	541: _ObservationIDName[639:658],
	553: _ObservationIDName[658:681],
	554: _ObservationIDName[681:697],
	555: _ObservationIDName[697:721],
	560: _ObservationIDName[721:751],
	561: _ObservationIDName[751:783],
	562: _ObservationIDName[783:813],
	563: _ObservationIDName[813:845],
	570: _ObservationIDName[845:858],
	701: _ObservationIDName[858:872],
	702: _ObservationIDName[872:882],
	703: _ObservationIDName[882:905],
	704: _ObservationIDName[905:928],
	706: _ObservationIDName[928:948],
	707: _ObservationIDName[948:981],
	708: _ObservationIDName[981:997],
	710: _ObservationIDName[997:1017],
	711: _ObservationIDName[1017:1026],
	712: _ObservationIDName[1026:1038],
	713: _ObservationIDName[1038:1070],
	714: _ObservationIDName[1070:1079],
	715: _ObservationIDName[1079:1090],
	716: _ObservationIDName[1090:1101],
	717: _ObservationIDName[1101:1115],
	718: _ObservationIDName[1115:1130],
	721: _ObservationIDName[1130:1147],
	722: _ObservationIDName[1147:1169],
	723: _ObservationIDName[1169:1185],
	750: _ObservationIDName[1185:1198],
	751: _ObservationIDName[1198:1223],
	752: _ObservationIDName[1223:1245],
	753: _ObservationIDName[1245:1263],
	800: _ObservationIDName[1263:1277],
	801: _ObservationIDName[1277:1286],
	803: _ObservationIDName[1286:1299],
	804: _ObservationIDName[1299:1307],
	805: _ObservationIDName[1307:1322],
	807: _ObservationIDName[1322:1344],
	808: _ObservationIDName[1344:1361],
	809: _ObservationIDName[1361:1388],
	810: _ObservationIDName[1388:1409],
	811: _ObservationIDName[1409:1423],
	812: _ObservationIDName[1423:1434],
	813: _ObservationIDName[1434:1460],
	814: _ObservationIDName[1460:1486],
	820: _ObservationIDName[1486:1501],
	821: _ObservationIDName[1501:1510],
	850: _ObservationIDName[1510:1523],
	851: _ObservationIDName[1523:1566],
	852: _ObservationIDName[1566:1581],
	853: _ObservationIDName[1581:1603],
	854: _ObservationIDName[1603:1619],
	855: _ObservationIDName[1619:1654],
	899: _ObservationIDName[1654:1668],
	900: _ObservationIDName[1668:1689],
	901: _ObservationIDName[1689:1714],
	908: _ObservationIDName[1714:1754],
	909: _ObservationIDName[1754:1793],
	911: _ObservationIDName[1793:1832],
	912: _ObservationIDName[1832:1870],
	913: _ObservationIDName[1870:1898],
	950: _ObservationIDName[1898:1905],
	951: _ObservationIDName[1905:1921],
	952: _ObservationIDName[1921:1928],
	953: _ObservationIDName[1928:1942],
	960: _ObservationIDName[1942:1949],
	961: _ObservationIDName[1949:1958],
	962: _ObservationIDName[1958:1966],
	963: _ObservationIDName[1966:1973],
	980: _ObservationIDName[1973:1987],
}

func (i ObservationID) String() string {
	if str, ok := _ObservationIDMap[i]; ok {
		return str
	}
	return fmt.Sprintf("ObservationID(%d)", i)
}

var _ObservationIDValues = []ObservationID{-3, -2, -1, 0, 1, 120, 130, 131, 132, 133, 134, 150, 151, 152, 153, 154, 155, 156, 170, 201, 202, 203, 241, 270, 501, 502, 503, 507, 508, 509, 510, 511, 512, 513, 515, 517, 518, 519, 520, 522, 523, 540, 541, 553, 554, 555, 560, 561, 562, 563, 570, 701, 702, 703, 704, 706, 707, 708, 710, 711, 712, 713, 714, 715, 716, 717, 718, 721, 722, 723, 750, 751, 752, 753, 800, 801, 803, 804, 805, 807, 808, 809, 810, 811, 812, 813, 814, 820, 821, 850, 851, 852, 853, 854, 855, 899, 900, 901, 908, 909, 911, 912, 913, 950, 951, 952, 953, 960, 961, 962, 963, 980}

var _ObservationIDNameToValueMap = map[string]ObservationID{
	_ObservationIDName[0:15]:      -3,
	_ObservationIDName[15:23]:     -2,
	_ObservationIDName[23:28]:     -1,
	_ObservationIDName[28:35]:     0,
	_ObservationIDName[35:46]:     1,
	_ObservationIDName[46:68]:     120,
	_ObservationIDName[68:81]:     130,
	_ObservationIDName[81:96]:     131,
	_ObservationIDName[96:119]:    132,
	_ObservationIDName[119:141]:   133,
	_ObservationIDName[141:161]:   134,
	_ObservationIDName[161:178]:   150,
	_ObservationIDName[178:196]:   151,
	_ObservationIDName[196:207]:   152,
	_ObservationIDName[207:220]:   153,
	_ObservationIDName[220:242]:   154,
	_ObservationIDName[242:259]:   155,
	_ObservationIDName[259:274]:   156,
	_ObservationIDName[274:289]:   170,
	_ObservationIDName[289:309]:   201,
	_ObservationIDName[309:329]:   202,
	_ObservationIDName[329:353]:   203,
	_ObservationIDName[353:380]:   241,
	_ObservationIDName[380:388]:   270,
	_ObservationIDName[388:401]:   501,
	_ObservationIDName[401:414]:   502,
	_ObservationIDName[414:427]:   503,
	_ObservationIDName[427:440]:   507,
	_ObservationIDName[440:453]:   508,
	_ObservationIDName[453:466]:   509,
	_ObservationIDName[466:483]:   510,
	_ObservationIDName[483:500]:   511,
	_ObservationIDName[500:512]:   512,
	_ObservationIDName[512:528]:   513,
	_ObservationIDName[528:538]:   515,
	_ObservationIDName[538:556]:   517,
	_ObservationIDName[556:567]:   518,
	_ObservationIDName[567:576]:   519,
	_ObservationIDName[576:585]:   520,
	_ObservationIDName[585:604]:   522,
	_ObservationIDName[604:625]:   523,
	_ObservationIDName[625:639]:   540,
	_ObservationIDName[639:658]:   541,
	_ObservationIDName[658:681]:   553,
	_ObservationIDName[681:697]:   554,
	_ObservationIDName[697:721]:   555,
	_ObservationIDName[721:751]:   560,
	_ObservationIDName[751:783]:   561,
	_ObservationIDName[783:813]:   562,
	_ObservationIDName[813:845]:   563,
	_ObservationIDName[845:858]:   570,
	_ObservationIDName[858:872]:   701,
	_ObservationIDName[872:882]:   702,
	_ObservationIDName[882:905]:   703,
	_ObservationIDName[905:928]:   704,
	_ObservationIDName[928:948]:   706,
	_ObservationIDName[948:981]:   707,
	_ObservationIDName[981:997]:   708,
	_ObservationIDName[997:1017]:  710,
	_ObservationIDName[1017:1026]: 711,
	_ObservationIDName[1026:1038]: 712,
	_ObservationIDName[1038:1070]: 713,
	_ObservationIDName[1070:1079]: 714,
	_ObservationIDName[1079:1090]: 715,
	_ObservationIDName[1090:1101]: 716,
	_ObservationIDName[1101:1115]: 717,
	_ObservationIDName[1115:1130]: 718,
	_ObservationIDName[1130:1147]: 721,
	_ObservationIDName[1147:1169]: 722,
	_ObservationIDName[1169:1185]: 723,
	_ObservationIDName[1185:1198]: 750,
	_ObservationIDName[1198:1223]: 751,
	_ObservationIDName[1223:1245]: 752,
	_ObservationIDName[1245:1263]: 753,
	_ObservationIDName[1263:1277]: 800,
	_ObservationIDName[1277:1286]: 801,
	_ObservationIDName[1286:1299]: 803,
	_ObservationIDName[1299:1307]: 804,
	_ObservationIDName[1307:1322]: 805,
	_ObservationIDName[1322:1344]: 807,
	_ObservationIDName[1344:1361]: 808,
	_ObservationIDName[1361:1388]: 809,
	_ObservationIDName[1388:1409]: 810,
	_ObservationIDName[1409:1423]: 811,
	_ObservationIDName[1423:1434]: 812,
	_ObservationIDName[1434:1460]: 813,
	_ObservationIDName[1460:1486]: 814,
	_ObservationIDName[1486:1501]: 820,
	_ObservationIDName[1501:1510]: 821,
	_ObservationIDName[1510:1523]: 850,
	_ObservationIDName[1523:1566]: 851,
	_ObservationIDName[1566:1581]: 852,
	_ObservationIDName[1581:1603]: 853,
	_ObservationIDName[1603:1619]: 854,
	_ObservationIDName[1619:1654]: 855,
	_ObservationIDName[1654:1668]: 899,
	_ObservationIDName[1668:1689]: 900,
	_ObservationIDName[1689:1714]: 901,
	_ObservationIDName[1714:1754]: 908,
	_ObservationIDName[1754:1793]: 909,
	_ObservationIDName[1793:1832]: 911,
	_ObservationIDName[1832:1870]: 912,
	_ObservationIDName[1870:1898]: 913,
	_ObservationIDName[1898:1905]: 950,
	_ObservationIDName[1905:1921]: 951,
	_ObservationIDName[1921:1928]: 952,
	_ObservationIDName[1928:1942]: 953,
	_ObservationIDName[1942:1949]: 960,
	_ObservationIDName[1949:1958]: 961,
	_ObservationIDName[1958:1966]: 962,
	_ObservationIDName[1966:1973]: 963,
	_ObservationIDName[1973:1987]: 980,
}

// ObservationIDString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func ObservationIDString(s string) (ObservationID, error) {
	if val, ok := _ObservationIDNameToValueMap[s]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to ObservationID values", s)
}

// ObservationIDValues returns all values of the enum
func ObservationIDValues() []ObservationID {
	return _ObservationIDValues
}

// IsAObservationID returns "true" if the value is listed in the enum definition. "false" otherwise
func (i ObservationID) IsAObservationID() bool {
	_, ok := _ObservationIDMap[i]
	return ok
}