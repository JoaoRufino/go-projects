package main

import (
	"encoding/asn1"
	"fmt"
	"log"
	"unsafe"

	"github.com/JoaoRufino/canopus"
)

//#cgo LDFLAGS:-L/home/jplr/gitrepos/Golang-projects/asn/Final/libs -lit2s-asn-cam
//#cgo CFLAGS:-I/home/jplr/gitrepos/Golang-projects/asn/Final/libs
/*#include <stdlib.h>
#include <stdio.h>
#include <unistd.h>
#include <stdint.h>
#include <errno.h>
#include <syslog.h>
#include <CAM.h>
#include <INTEGER.h>
#include <asn_application.h>

#define syslog_emerg(msg, ...) syslog(LOG_EMERG, "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)
#define syslog_err(msg, ...)   syslog(LOG_ERR  , "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)

#ifndef NDEBUG
#define syslog_debug(msg, ...) syslog(LOG_DEBUG, "%s(%s:%d) [" msg "]", __FILE__, __func__, __LINE__, ##__VA_ARGS__)
#else
#define syslog_debug(msg, ...)
#endif

static CAM_t *cam;

int decode(uint8_t *buffer,uint8_t *der_buffer,int file_size, int der_size){
	asn_dec_rval_t dec;
	asn_enc_rval_t er;
	asn_codec_ctx_t *opt_codec_ctx = 0;
	cam = calloc(1, sizeof(CAM_t));
	if(cam == NULL) {
		syslog_emerg("calloc() failed: %m");
	}
	dec =  uper_decode_complete(opt_codec_ctx,&asn_DEF_CAM, (void **) &cam, buffer, file_size);
	switch(dec.code) {
		case RC_OK:
			er =der_encode_to_buffer(&asn_DEF_CAM,cam,der_buffer,der_size);
			assert(er.encoded>0);
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return er.encoded;
		case RC_FAIL:
			syslog_debug("Error decoding: RC_FAIL");
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return 0;
		case RC_WMORE:
			syslog_debug("ERROR decoding: RC_WMORE");
			xer_fprint(stdout, &asn_DEF_CAM, cam);
			return 0;
	}

	printf("DONE!!!!\n");


}


uint32_t get_lat(void) {
	return (cam->cam.camParameters.basicContainer.referencePosition.latitude);
}

uint32_t get_lon(void) {
	return (cam->cam.camParameters.basicContainer.referencePosition.longitude);
}

*/
import "C"

func main() {

	/*	var err error
		var buffer []byte
		buffer, err = ioutil.ReadFile("./it2s-CAM_Tx.code")*/

	server := canopus.NewServer()
	//var xml string , C.CString(xml)
	server.Get("/cam/json", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewPlainTextPayload("Acknowledged"))
		res := canopus.NewResponse(msg, nil)
		return res
	})
	server.Get("/denm/json", func(req canopus.Request) canopus.Response {
		msg := canopus.NewMessageOfType(canopus.MessageAcknowledgment, req.GetMessage().GetMessageId(), canopus.NewPlainTextPayload("Acknowledged"))
		res := canopus.NewResponse(msg, nil)
		return res
	})

	server.Post("/cam", func(req canopus.Request) canopus.Response {
		log.Println("cam")

		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("DONE! ")
		res := canopus.NewResponse(msg, nil)

		// Save to file
		payload := req.GetMessage().GetPayload().GetBytes()
		log.Println("len", len(payload))
		buffer := make([]byte, 256)
		C.decode((*C.uint8_t)(unsafe.Pointer(&payload[0])), (*C.uint8_t)(unsafe.Pointer(&buffer[0])), (C.int)(len(payload)), C.int(len(buffer))) //BEWARE VERY UNSAFE SIFILIS!!! AIDS!!!!
		var cam = new(CAM)
		_, er := asn1.Unmarshal(buffer, cam)
		check(er)
		fmt.Println(cam.Header.MessageID)

		changeVal := fmt.Sprint((uint32)(C.get_lat())) + "|" + fmt.Sprint((uint32)(C.get_lon()))
		server.NotifyChange("/cam/pos", changeVal, false)
		return res

	})
	/*	server.Post("/denm", func(req canopus.Request) canopus.Response {
		log.Println("denm")
		var bxml *[]byte
		msg := canopus.ContentMessage(req.GetMessage().GetMessageId(), canopus.MessageAcknowledgment)
		msg.SetStringPayload("DONE! ")
		res := canopus.NewResponse(msg, nil)

		// Save to file
		payload := req.GetMessage().GetPayload().GetBytes()
		log.Println("len", len(payload))
		C.DENMdecode((*C.uint8_t)(unsafe.Pointer(&payload[0])), (C.int)(len(payload)), (*C.uint8_t)(unsafe.Pointer(&bxml))) //BEWARE VERY UNSAFE SIFILIS!!! AIDS!!!!
		log.Println(&bxml)

		changeVal := fmt.Sprint((uint32)(C.get_lat())) + "|" + fmt.Sprint((uint32)(C.get_lon()))
		server.NotifyChange("/cam/pos", changeVal, false)
		return res

	})*/

	server.OnMessage(func(msg canopus.Message, inbound bool) {
		//canopus.PrintMessage(msg)
	})

	server.OnObserve(func(resource string, msg canopus.Message) {
		fmt.Println("Observe Requested for " + resource)
	})
	server.OnBlockMessage(func(msg canopus.Message, inbound bool) {
		canopus.PrintMessage(msg)
	})

	server.ListenAndServe(":5683")
	<-make(chan struct{})
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

//BitString ASN1 Data Type
type BitString struct {
	Bytes     []byte // bits packed into bytes.
	BitLength int    // length in bits.
}

//Enumerated ASN1 Data Type
type Enumerated int

//CAM Message
type CAM struct {
	Header ItsPduHeader  `asn1:"explicit,tag:16"`
	Cam    CoopAwareness `asn1:"explicit,tag:16"`
}

type CoopAwareness struct {
	GenerationDeltaTime GenerationDeltaTime `asn1:"explicit,tag:16"`
	CamParameters       CamParameters       `asn1:"explicit,tag:16"`
}

type CamParameters struct {
	BasicContainer          BasicContainer          `asn1:"explicit,tag:16"`
	HighFrequencyContainer  HighFrequencyContainer  `asn1:"explicit,tag:16"`
	LowFrequencyContainer   LowFrequencyContainer   `asn1:"optional,tag:16"` //OPTIONAL
	specialVehicleContainer SpecialVehicleContainer `asn1:"optional,tag:16"` //OPTIONAL
}

type HighFrequencyContainer struct {
	BasicVehicleContainerHighFrequency BasicVehicleContainerHighFrequency `asn1:"explicit,tag:16"`
	rsuContainerHighFrequency          RSUContainerHighFrequency          `asn1:"explicit,tag:16"`
}

type LowFrequencyContainer struct {
	BasicVehicleContainerLowFrequency BasicVehicleContainerLowFrequency `asn1:"explicit,tag:16"`
}

type SpecialVehicleContainer struct {
	PublicTransportContainer  PublicTransportContainer
	SpecialTransportContainer SpecialTransportContainer
	DangerousGoodsContainer   DangerousGoodsContainer
	RoadWorksContainerBasic   RoadWorksContainerBasic
	RescueContainer           RescueContainer
	EmergencyContainer        EmergencyContainer
	SafetyCarContainer        SafetyCarContainer
}

type BasicContainer struct {
	StationType       StationType
	ReferencePosition ReferencePosition
}

type BasicVehicleContainerHighFrequency struct {
	Heading                  Heading
	Speed                    Speed
	DriveDirection           DriveDirection
	VehicleLength            VehicleLength
	VehicleWidth             VehicleWidth
	LongitudinalAcceleration LongitudinalAcceleration
	Curvature                Curvature
	CurvatureCalculationMode CurvatureCalculationMode
	YawRate                  YawRate
	AccelerationControl      AccelerationControl  //OPTIONAL
	LanePosition             LanePosition         //OPTIONAL
	SteeringWheelAngle       SteeringWheelAngle   //OPTIONAL
	LateralAcceleration      LateralAcceleration  //OPTIONAL
	verticalAcceleration     VerticalAcceleration //OPTIONAL
	PerformanceClass         PerformanceClass     //OPTIONAL
	CenDsrcTollingZone       CenDsrcTollingZone   //OPTIONAL
}

type BasicVehicleContainerLowFrequency struct {
	VehicleRole    VehicleRole
	ExteriorLights ExteriorLights
	PathHistory    PathHistory
}

type PublicTransportContainer struct {
	EmbarkationStatus EmbarkationStatus
	PtActivation      PtActivation //OPTIONAL
}

type SpecialTransportContainer struct {
	SpecialTransportType SpecialTransportType
	LightBarSirenInUse   LightBarSirenInUse
}
type DangerousGoodsContainer struct {
	dangerousGoodsBasic DangerousGoodsBasic
}

type RoadWorksContainerBasic struct {
	RoadworksSubCauseCode RoadworksSubCauseCode //OPTIONAL
	LightBarSirenInUse    LightBarSirenInUse
	ClosedLanes           ClosedLanes //OPTIONAL
}

type RescueContainer struct {
	LightBarSirenInUse LightBarSirenInUse
}

type EmergencyContainer struct {
	LightBarSirenInUse LightBarSirenInUse
	IncidentIndication CauseCode         //OPTIONAL
	EmergencyPriority  EmergencyPriority //OPTIONAL
}

type SafetyCarContainer struct {
	LightBarSirenInUse LightBarSirenInUse
	IncidentIndication CauseCode   //OPTIONAL
	TrafficRule        TrafficRule //OPTIONAL
	speedLimit         SpeedLimit  //OPTIONAL
}

type RSUContainerHighFrequency struct {
	ProtectedCommunicationZonesRSU ProtectedCommunicationZonesRSU //OPTIONAL
}

type GenerationDeltaTime int

type ItsPduHeader struct { //SEQUENCE
	ProtocolVersion int
	MessageID       int
	/*   denm(1)
	cam(2)
	poi(3)
	spat(4)
	map(5)
	ivi(6)
	ev-rsr(7)*/
	StationID StationID
}

type StationID int //(0..4294967295)

type ReferencePosition struct { // SEQUENCE
	Latitude                  Latitude
	Longitude                 Longitude
	PositionConfidenceEllipse PosConfidenceEllipse
	Altitude                  Altitude
}

type DeltaReferencePosition struct {
	DeltaLatitude  DeltaLatitude
	DeltaLongitude DeltaLongitude
	DeltaAltitude  DeltaAltitude
}

type Longitude int

/*{
	oneMicrodegreeEast(10)
	oneMicrodegreeWest(-10)
	unavailable(1800000001)
    } (-1800000000..1800000001)*/

type Latitude int

/*{
	oneMicrodegreeNorth(10)
	oneMicrodegreeSouth(-10)
	unavailable(900000001)
    } (-900000000..900000001)*/

type Altitude struct {
	AltitudeValue      AltitudeValue
	AltitudeConfidence AltitudeConfidence
}

type AltitudeValue int

/*{
	referenceEllipsoidSurface(0)
	oneCentimeter(1)
	unavailable(800001)
    } (-100000..800001)
*/

type AltitudeConfidence Enumerated

/* {
	alt-000-01(0)
	alt-000-02(1)
	alt-000-05(2)
	alt-000-10(3)
	alt-000-20(4)
	alt-000-50(5)
	alt-001-00(6)
	alt-002-00(7)
	alt-005-00(8)
	alt-010-00(9)
	alt-020-00(10)
	alt-050-00(11)
	alt-100-00(12)
	alt-200-00(13)
	outOfRange(14)
	unavailable(15)
    }*/

type DeltaLongitude int

/* {
	oneMicrodegreeEast(10)
	oneMicrodegreeWest(-10)
	unavailable(131072)
    } (-131071..131072)
*/

type DeltaLatitude int

/* {
	oneMicrodegreeNorth(10)
	oneMicrodegreeSouth(-10)
	unavailable(131072)
    } (-131071..131072)
*/

type DeltaAltitude int

/*{
	oneCentimeterUp(1)
	oneCentimeterDown(-1)
	unavailable(12800)
    } (-12700..12800)*/

type PosConfidenceEllipse struct {
	SemiMajorConfidence  SemiAxisLength
	SemiMinorConfidence  SemiAxisLength
	SemiMajorOrientation HeadingValue
}

type PathPoint struct {
	PathPosition  DeltaReferencePosition
	PathDeltaTime PathDeltaTime //OPTIONAL
}

type PathDeltaTime int

/*{
	tenMilliSecondsInPast(1)
    } (1..65535  ...)*/

type PtActivation struct {
	PtActivationType PtActivationType
	PtActivationData PtActivationData
}

type PtActivationType int

/*{
	undefinedCodingType(0)
	r09-16CodingType(1)
	vdv-50149CodingType(2)
    } (0..255)*/

type PtActivationData []byte

//::= OCTET STRING (SIZE (1..20))

type AccelerationControl BitString

/*::= BIT STRING {
	brakePedalEngaged(0)
	gasPedalEngaged(1)
	emergencyBrakeEngaged(2)
	collisionWarningEngaged(3)
	accEngaged(4)
	cruiseControlEngaged(5)
	speedLimiterEngaged(6)
    } (SIZE (7))*/

type SemiAxisLength int

/* {
	oneCentimeter(1)
	outOfRange(4094)
	unavailable(4095)
    } (0..4095)*/

type CauseCode struct {
	CauseCode    CauseCodeType
	SubCauseCode SubCauseCodeType
}

type CauseCodeType int

/*{
	reserved(0)
	trafficCondition(1)
	accident(2)
	roadworks(3)
	adverseWeatherCondition-Adhesion(6)
	hazardousLocation-SurfaceCondition(9)
	hazardousLocation-ObstacleOnTheRoad(10)
	hazardousLocation-AnimalOnTheRoad(11)
	humanPresenceOnTheRoad(12)
	wrongWayDriving(14)
	rescueAndRecoveryWorkInProgress(15)
	adverseWeatherCondition-ExtremeWeatherCondition(17)
	adverseWeatherCondition-Visibility(18)
	adverseWeatherCondition-Precipitation(19)
	slowVehicle(26)
	dangerousEndOfQueue(27)
	vehicleBreakdown(91)
	postCrash(92)
	humanProblem(93)
	stationaryVehicle(94)
	emergencyVehicleApproaching(95)
	hazardousLocation-DangerousCurve(96)
	collisionRisk(97)
	signalViolation(98)
	dangerousSituation(99)
    } (0..255)*/

type SubCauseCodeType int

// (0..255)

type TrafficConditionSubCauseCode int

/*{
	unavailable(0)
	increasedVolumeOfTraffic(1)
	trafficJamSlowlyIncreasing(2)
	trafficJamIncreasing(3)
	trafficJamStronglyIncreasing(4)
	trafficStationary(5)
	trafficJamSlightlyDecreasing(6)
	trafficJamDecreasing(7)
	trafficJamStronglyDecreasing(8)
    } (0..255)*/

type AccidentSubCauseCode int

/*{
	unavailable(0)
	multiVehicleAccident(1)
	heavyAccident(2)
	accidentInvolvingLorry(3)
	accidentInvolvingBus(4)
	accidentInvolvingHazardousMaterials(5)
	accidentOnOppositeLane(6)
	unsecuredAccident(7)
	assistanceRequested(8)
    } (0..255)*/

type RoadworksSubCauseCode int

/*{
	unavailable(0)
	majorRoadworks(1)
	roadMarkingWork(2)
	slowMovingRoadMaintenance(3)
	shortTermStationaryRoadworks(4)
	streetCleaning(5)
	winterService(6)
    } (0..255)*/

type HumanPresenceOnTheRoadSubCauseCode int

/*{
	unavailable(0)
	childrenOnRoadway(1)
	cyclistOnRoadway(2)
	motorcyclistOnRoadway(3)
    } (0..255)*/

type WrongWayDrivingSubCauseCode int

/*{
	unavailable(0)
	wrongLane(1)
	wrongDirection(2)
    } (0..255)*/

type AdverseWeatherCondition_ExtremeWeatherConditionSubCauseCode int

/*{
	unavailable(0)
	strongWinds(1)
	damagingHail(2)
	hurricane(3)
	thunderstorm(4)
	tornado(5)
	blizzard(6)
    } (0..255)*/

type AdverseWeatherCondition_AdhesionSubCauseCode int

/*{
	unavailable(0)
	heavyFrostOnRoad(1)
	fuelOnRoad(2)
	mudOnRoad(3)
	snowOnRoad(4)
	iceOnRoad(5)
	blackIceOnRoad(6)
	oilOnRoad(7)
	looseChippings(8)
	instantBlackIce(9)
	roadsSalted(10)
    } (0..255)*/

type AdverseWeatherCondition_VisibilitySubCauseCode int

/*{
	unavailable(0)
	fog(1)
	smoke(2)
	heavySnowfall(3)
	heavyRain(4)
	heavyHail(5)
	lowSunGlare(6)
	sandstorms(7)
	swarmsOfInsects(8)
    } (0..255)*/

type AdverseWeatherCondition_PrecipitationSubCauseCode int

/*{
	unavailable(0)
	heavyRain(1)
	heavySnowfall(2)
	softHail(3)
    } (0..255)*/

type SlowVehicleSubCauseCode int

/*{
	unavailable(0)
	maintenanceVehicle(1)
	vehiclesSlowingToLookAtAccident(2)
	abnormalLoad(3)
	abnormalWideLoad(4)
	convoy(5)
	snowplough(6)
	deicing(7)
	saltingVehicles(8)
    } (0..255)*/

type StationaryVehicleSubCauseCode int

/*{
	unavailable(0)
	humanProblem(1)
	vehicleBreakdown(2)
	postCrash(3)
	publicTransportStop(4)
	carryingDangerousGoods(5)
    } (0..255)*/

type HumanProblemSubCauseCode int

/*{
	unavailable(0)
	glycemiaProblem(1)
	heartProblem(2)
    } (0..255)*/

type EmergencyVehicleApproachingSubCauseCode int

/*{
	unavailable(0)
	emergencyVehicleApproaching(1)
	prioritizedVehicleApproaching(2)
    } (0..255)*/

type HazardousLocation_DangerousCurveSubCauseCode int

/*{
	unavailable(0)
	dangerousLeftTurnCurve(1)
	dangerousRightTurnCurve(2)
	multipleCurvesStartingWithUnknownTurningDirection(3)
	multipleCurvesStartingWithLeftTurn(4)
	multipleCurvesStartingWithRightTurn(5)
    } (0..255)*/

type HazardousLocation_SurfaceConditionSubCauseCode int

/*{
	unavailable(0)
	rockfalls(1)
	earthquakeDamage(2)
	sewerCollapse(3)
	subsidence(4)
	snowDrifts(5)
	stormDamage(6)
	burstPipe(7)
	volcanoEruption(8)
	fallingIce(9)
    } (0..255)*/

type HazardousLocation_ObstacleOnTheRoadSubCauseCode int

/*{
	unavailable(0)
	shedLoad(1)
	partsOfVehicles(2)
	partsOfTyres(3)
	bigObjects(4)
	fallenTrees(5)
	hubCaps(6)
	waitingVehicles(7)
    } (0..255)*/

type HazardousLocation_AnimalOnTheRoadSubCauseCode int

/*{
	unavailable(0)
	wildAnimals(1)
	herdOfAnimals(2)
	smallAnimals(3)
	largeAnimals(4)
    } (0..255)*/

type CollisionRiskSubCauseCode int

/*{
	unavailable(0)
	longitudinalCollisionRisk(1)
	crossingCollisionRisk(2)
	lateralCollisionRisk(3)
	vulnerableRoadUser(4)
    } (0..255)*/

type SignalViolationSubCauseCode int

/*{
	unavailable(0)
	stopSignViolation(1)
	trafficLightViolation(2)
	turningRegulationViolation(3)
    } (0..255)*/

type RescueAndRecoveryWorkInProgressSubCauseCode int

/*{
	unavailable(0)
	emergencyVehicles(1)
	rescueHelicopterLanding(2)
	policeActivityOngoing(3)
	medicalEmergencyOngoing(4)
	childAbductionInProgress(5)
    } (0..255)*/

type DangerousEndOfQueueSubCauseCode int

/*{
	unavailable(0)
	suddenEndOfQueue(1)
	queueOverHill(2)
	queueAroundBend(3)
	queueInTunnel(4)
    } (0..255)*/

type DangerousSituationSubCauseCode int

/*{
	unavailable(0)
	emergencyElectronicBrakeEngaged(1)
	preCrashSystemEngaged(2)
	espEngaged(3)
	absEngaged(4)
	aebEngaged(5)
	brakeWarningEngaged(6)
	collisionRiskWarningEngaged(7)
    } (0..255)*/

type VehicleBreakdownSubCauseCode int

/*{
	unavailable(0)
	lackOfFuel(1)
	lackOfBatteryPower(2)
	engineProblem(3)
	transmissionProblem(4)
	engineCoolingProblem(5)
	brakingSystemProblem(6)
	steeringProblem(7)
	tyrePuncture(8)
    } (0..255)*/

type PostCrashSubCauseCode int

/*{
	unavailable(0)
	accidentWithoutECallTriggered(1)
	accidentWithECallManuallyTriggered(2)
	accidentWithECallAutomaticallyTriggered(3)
	accidentWithECallTriggeredWithoutAccessToCellularNetwork(4)
    } (0..255)*/

type Curvature struct {
	CurvatureValue      CurvatureValue
	CurvatureConfidence CurvatureConfidence
}

type CurvatureValue int

/*
    {
	straight(0)
	reciprocalOf1MeterRadiusToRight(-30000)
	reciprocalOf1MeterRadiusToLeft(30000)
	unavailable(30001)
    } (-30000..30001)*/

type CurvatureConfidence Enumerated

/*{
	onePerMeter-0-00002(0)
	onePerMeter-0-0001(1)
	onePerMeter-0-0005(2)
	onePerMeter-0-002(3)
	onePerMeter-0-01(4)
	onePerMeter-0-1(5)
	outOfRange(6)
	unavailable(7)
    }*/

type CurvatureCalculationMode Enumerated

/*{
	yawRateUsed(0)
	yawRateNotUsed(1)
	unavailable(2)
	...
    }*/

type Heading struct {
	HeadingValue      HeadingValue
	HeadingConfidence HeadingConfidence
}

type HeadingValue int

/* {
	wgs84North(0)
	wgs84East(900)
	wgs84South(1800)
	wgs84West(2700)
	unavailable(3601)
    } (0..3601)*/

type HeadingConfidence int

/*{
	equalOrWithinZeroPointOneDegree(1)
	equalOrWithinOneDegree(10)
	outOfRange(126)
	unavailable(127)
    } (1..127)*/

type LanePosition int

/*{
	offTheRoad(-1)
	hardShoulder(0)
	outermostDrivingLane(1)
	secondLaneFromOutside(2)
    } (-1..14)*/

type ClosedLanes struct {
	HardShoulderStatus HardShoulderStatus //OPTIONAL
	DrivingLaneStatus  DrivingLaneStatus
}

type HardShoulderStatus Enumerated

/* {
	availableForStopping(0)
	closed(1)
	availableForDriving(2)
    }*/

type DrivingLaneStatus BitString

/*{
	outermostLaneClosed(1)
	secondLaneFromOutsideClosed(2)
    } (SIZE (1..14))*/

type PerformanceClass int

/*{
	unavailable(0)
	performanceClassA(1)
	performanceClassB(2)
    } (0..7)*/

type SpeedValue int

/*{
	standstill(0)
	oneCentimeterPerSec(1)
	unavailable(16383)
    } (0..16383)*/

type SpeedConfidence int

/*{
	equalOrWithinOneCentimeterPerSec(1)
	equalOrWithinOneMeterPerSec(100)
	outOfRange(126)
	unavailable(127)
    } (1..127)*/

type VehicleMass int

/*{
	hundredKg(1)
	unavailable(1024)
    } (1..1024)*/

type Speed struct {
	SpeedValue      SpeedValue
	SpeedConfidence SpeedConfidence
}

type DriveDirection Enumerated

/*{
	forward(0)
	backward(1)
	unavailable(2)
    }*/

type EmbarkationStatus bool

type LongitudinalAcceleration struct {
	LongitudinalAccelerationValue      LongitudinalAccelerationValue
	LongitudinalAccelerationConfidence AccelerationConfidence
}

type LongitudinalAccelerationValue int

/*{
	pointOneMeterPerSecSquaredForward(1)
	pointOneMeterPerSecSquaredBackward(-1)
	unavailable(161)
    } (-160..161)*/

type AccelerationConfidence int

/*{
	pointOneMeterPerSecSquared(1)
	outOfRange(101)
	unavailable(102)
    } (0..102)*/

type LateralAcceleration struct {
	LateralAccelerationValue      LateralAccelerationValue
	LateralAccelerationConfidence AccelerationConfidence
}

type LateralAccelerationValue int

/*{
	pointOneMeterPerSecSquaredToRight(-1)
	pointOneMeterPerSecSquaredToLeft(1)
	unavailable(161)
    } (-160..161)*/

type VerticalAcceleration struct {
	VerticalAccelerationValue      VerticalAccelerationValue
	VerticalAccelerationConfidence AccelerationConfidence
}

type VerticalAccelerationValue int

/*{
	pointOneMeterPerSecSquaredUp(1)
	pointOneMeterPerSecSquaredDown(-1)
	unavailable(161)
    } (-160..161)*/

type StationType int

/*{
	unknown(0)
	pedestrian(1)
	cyclist(2)
	moped(3)
	motorcycle(4)
	passengerCar(5)
	bus(6)
	lightTruck(7)
	heavyTruck(8)
	trailer(9)
	specialVehicles(10)
	tram(11)
	roadSideUnit(15)
    } (0..255)*/

type ExteriorLights BitString

/* {
	lowBeamHeadlightsOn(0)
	highBeamHeadlightsOn(1)
	leftTurnSignalOn(2)
	rightTurnSignalOn(3)
	daytimeRunningLightsOn(4)
	reverseLightOn(5)
	fogLightOn(6)
	parkingLightsOn(7)
    } (SIZE (8))*/

type DangerousGoodsBasic Enumerated

/*{
	explosives1(0)
	explosives2(1)
	explosives3(2)
	explosives4(3)
	explosives5(4)
	explosives6(5)
	flammableGases(6)
	nonFlammableGases(7)
	toxicGases(8)
	flammableLiquids(9)
	flammableSolids(10)
	substancesLiableToSpontaneousCombustion(11)
	substancesEmittingFlammableGasesUponContactWithWater(12)
	oxidizingSubstances(13)
	organicPeroxides(14)
	toxicSubstances(15)
	infectiousSubstances(16)
	radioactiveMaterial(17)
	corrosiveSubstances(18)
	miscellaneousDangerousSubstances(19)
    }*/

type DangerousGoodsExtended struct {
	DangerousGoodsType  DangerousGoodsBasic
	UnNumber            int // (0..9999)
	ElevatedTemperature bool
	TunnelsRestricted   bool
	LimitedQuantity     bool
	EmergencyActionCode string //(SIZE (1..24))  //OPTIONAL
	PhoneNumber         string // (SIZE (1..24))  //OPTIONAL
	CompanyName         string // (SIZE (1..24))  //OPTIONAL
}

type SpecialTransportType BitString

/* {
	heavyLoad(0)
	excessWidth(1)
	excessLength(2)
	excessHeight(3)
    } (SIZE (4))*/

type LightBarSirenInUse BitString

/*{
	lightBarActivated(0)
	sirenActivated(1)
    } (SIZE (2))*/

type HeightLonCarr int

/*{
	oneCentimeter(1)
	unavailable(100)
    } (1..100)*/

type PosLonCarr int

/*{
	oneCentimeter(1)
	unavailable(127)
    } (1..127)*/

type PosPillar int

/*{
	tenCentimeters(1)
	unavailable(30)
    } (1..30)*/

type PosCentMass int

/*{
	tenCentimeters(1)
	unavailable(63)
    } (1..63)*/

type RequestResponseIndication Enumerated

/* {
	request(0)
	response(1)
    }*/

type SpeedLimit int

/*{
	oneKmPerHour(1)
    } (1..255)*/

type StationarySince Enumerated

/*{
	lessThan1Minute(0)
	lessThan2Minutes(1)
	lessThan15Minutes(2)
	equalOrGreater15Minutes(3)
    }*/

type Temperature int

/*{
	equalOrSmallerThanMinus60Deg(-60)
	oneDegreeCelsius(1)
	equalOrGreaterThan67Deg(67)
    } (-60..67)*/

type TrafficRule Enumerated

/*{
	noPassing(0)
	noPassingForTrucks(1)
	passToRight(2)
	passToLeft(3)
	...
    }*/

type WheelBaseVehicle int

/*{
	tenCentimeters(1)
	unavailable(127)
    } (1..127)*/

type TurningRadius int

/*{
	point4Meters(1)
	unavailable(255)
    } (1..255)*/

type PosFrontAx int

/*{
	tenCentimeters(1)
	unavailable(20)
    } (1..20)*/

type PositionOfOccupants BitString

/* {
	row1LeftOccupied(0)
	row1RightOccupied(1)
	row1MidOccupied(2)
	row1NotDetectable(3)
	row1NotPresent(4)
	row2LeftOccupied(5)
	row2RightOccupied(6)
	row2MidOccupied(7)
	row2NotDetectable(8)
	row2NotPresent(9)
	row3LeftOccupied(10)
	row3RightOccupied(11)
	row3MidOccupied(12)
	row3NotDetectable(13)
	row3NotPresent(14)
	row4LeftOccupied(15)
	row4RightOccupied(16)
	row4MidOccupied(17)
	row4NotDetectable(18)
	row4NotPresent(19)
    } (SIZE (20))*/

type PositioningSolutionType Enumerated

/*{
	noPositioningSolution(0)
	sGNSS(1)
	dGNSS(2)
	sGNSSplusDR(3)
	dGNSSplusDR(4)
	dR(5)
	...
    }*/

type VehicleIdentification struct {
	WMInumber WMInumber //OPTIONAL
	VDS       VDS       //OPTIONAL
}

type WMInumber string // (SIZE (1..3))

type VDS string //(SIZE (6))

type EnergyStorageType BitString

/* BIT STRING {
	hydrogenStorage(0)
	electricEnergyStorage(1)
	liquidPropaneGas(2)
	compressedNaturalGas(3)
	diesel(4)
	gasoline(5)
	ammonia(6)
    } (SIZE (7))*/

type VehicleLength struct {
	VehicleLengthValue                VehicleLengthValue
	VehicleLengthConfidenceIndication VehicleLengthConfidenceIndication
}

type VehicleLengthValue int

/*{
	tenCentimeters(1)
	outOfRange(1022)
	unavailable(1023)
    } (1..1023)*/

type VehicleLengthConfidenceIndication Enumerated

/*{
	noTrailerPresent(0)
	trailerPresentWithKnownLength(1)
	trailerPresentWithUnknownLength(2)
	trailerPresenceIsUnknown(3)
	unavailable(4)
    }*/

type VehicleWidth int

/*{
	tenCentimeters(1)
	outOfRange(61)
	unavailable(62)
    } (1..62)*/

type PathHistory []PathPoint

//::= SEQUENCE SIZE (0..40) OF

type EmergencyPriority BitString

/*{
	requestForRightOfWay(0)
	requestForFreeCrossingAtATrafficLight(1)
    } (SIZE (2))*/

type InformationQuality int

/*{
	unavailable(0)
	lowest(1)
	highest(7)
    } (0..7)*/

type RoadType Enumerated

/*{
	urban-NoStructuralSeparationToOppositeLanes(0)
	urban-WithStructuralSeparationToOppositeLanes(1)
	nonUrban-NoStructuralSeparationToOppositeLanes(2)
	nonUrban-WithStructuralSeparationToOppositeLanes(3)
    }*/

type SteeringWheelAngle struct {
	SteeringWheelAngleValue      SteeringWheelAngleValue
	SteeringWheelAngleConfidence SteeringWheelAngleConfidence
}

type SteeringWheelAngleValue int

/*{
	straight(0)
	onePointFiveDegreesToRight(-1)
	onePointFiveDegreesToLeft(1)
	unavailable(512)
    } (-511..512)*/

type SteeringWheelAngleConfidence int

/*{
	equalOrWithinOnePointFiveDegree(1)
	outOfRange(126)
	unavailable(127)
    } (1..127)*/

type TimestampIts int

/*{
	utcStartOf2004(0)
	oneMillisecAfterUTCStartOf2004(1)
    } (0..4398046511103)*/

type VehicleRole Enumerated

/* {
	default(0)
	publicTransport(1)
	specialTransport(2)
	dangerousGoods(3)
	roadWork(4)
	rescue(5)
	emergency(6)
	safetyCar(7)
	agriculture(8)
	commercial(9)
	military(10)
	roadOperator(11)
	taxi(12)
	reserved1(13)
	reserved2(14)
	reserved3(15)
    }*/

type YawRate struct {
	YawRateValue      YawRateValue
	YawRateConfidence YawRateConfidence
}

type YawRateValue int

/*{
	straight(0)
	degSec-000-01ToRight(-1)
	degSec-000-01ToLeft(1)
	unavailable(32767)
    } (-32766..32767)*/

type YawRateConfidence Enumerated

/* {
	degSec-000-01(0)
	degSec-000-05(1)
	degSec-000-10(2)
	degSec-001-00(3)
	degSec-005-00(4)
	degSec-010-00(5)
	degSec-100-00(6)
	outOfRange(7)
	unavailable(8)
    }*/

type ProtectedZoneType Enumerated

/*{
	cenDsrcTolling(0)
	...
    }*/

type RelevanceDistance Enumerated

/*{
	lessThan50m(0)
	lessThan100m(1)
	lessThan200m(2)
	lessThan500m(3)
	lessThan1000m(4)
	lessThan5km(5)
	lessThan10km(6)
	over10km(7)
    }*/

type RelevanceTrafficDirection Enumerated

/* {
	allTrafficDirections(0)
	upstreamTraffic(1)
	downstreamTraffic(2)
	oppositeTraffic(3)
    }*/

type TransmissionInterval int

/* {
	oneMilliSecond(1)
	tenSeconds(10000)
    } (1..10000)*/

type ValidityDuration int

/* {
	timeOfDetection(0)
	oneSecondAfterDetection(1)
    } (0..86400)*/

type ActionID struct {
	OriginatingStationID StationID
	SequenceNumber       SequenceNumber
}

type ItineraryPath []ReferencePosition

//::= SEQUENCE SIZE (1..40) OF

type ProtectedCommunicationZone struct {
	ProtectedZoneType      ProtectedZoneType
	ExpiryTime             TimestampIts //OPTIONAL
	ProtectedZoneLatitude  Latitude
	ProtectedZoneLongitude Longitude
	ProtectedZoneRadius    ProtectedZoneRadius //OPTIONAL
	ProtectedZoneID        ProtectedZoneID     //OPTIONAL
}

type Traces []PathHistory

//::= SEQUENCE SIZE (1..7) OF

type NumberOfOccupants int

/*
	oneOccupant(1)
	unavailable(127)
    } (0..127)*/

type SequenceNumber int

// (0..65535)

type PositionOfPillars []PosPillar

//::= SEQUENCE SIZE (1..3  ...) OF

type RestrictedTypes []StationType

//::= SEQUENCE SIZE (1..3  ...) OF

type EventHistory []EventPoint

//::= SEQUENCE SIZE (1..23) OF

type EventPoint struct {
	EventPosition      DeltaReferencePosition
	EventDeltaTime     PathDeltaTime //OPTIONAL
	InformationQuality InformationQuality
}

type ProtectedCommunicationZonesRSU []ProtectedCommunicationZone

//::= SEQUENCE SIZE (1..16) OF

type CenDsrcTollingZone struct {
	ProtectedZoneLatitude  Latitude
	ProtectedZoneLongitude Longitude
	SenDsrcTollingZoneID   CenDsrcTollingZoneID //OPTIONAL
}

type ProtectedZoneRadius int

//oneMeter(1) } (1..255  ...)

type ProtectedZoneID int

//int (0..134217727)

type CenDsrcTollingZoneID ProtectedZoneID
