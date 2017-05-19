
type CAM struct{
	header ItsPduHeader 
	cam    CoopAwareness
}

type CoopAwareness struct{
	generationDeltaTime GenerationDeltaTime
	camParameters       CamParameters
}

type CamParameters struct{
	basicContainer          BasicContainer 
	highFrequencyContainer  HighFrequencyContainer 
	lowFrequencyContainer   LowFrequencyContainer OPTIONAL 
	specialVehicleContainer SpecialVehicleContainer OPTIONAL
    }

type HighFrequencyContainer struct{
	basicVehicleContainerHighFrequency BasicVehicleContainerHighFrequency 
	rsuContainerHighFrequency          RSUContainerHighFrequency
}

type LowFrequencyContainer {
	basicVehicleContainerLowFrequency BasicVehicleContainerLowFrequency 
    }

type SpecialVehicleContainer {
	publicTransportContainer  PublicTransportContainer 
	specialTransportContainer SpecialTransportContainer 
	dangerousGoodsContainer   DangerousGoodsContainer 
	roadWorksContainerBasic   RoadWorksContainerBasic 
	rescueContainer           RescueContainer 
	emergencyContainer        EmergencyContainer 
	safetyCarContainer        SafetyCarContainer 
}

type BasicContainer {
	stationType       StationType 
	referencePosition ReferencePosition
    }

type BasicVehicleContainerHighFrequency  {
	heading                  Heading 
	speed                    Speed 
	driveDirection           DriveDirection 
	vehicleLength            VehicleLength 
	vehicleWidth             VehicleWidth 
	longitudinalAcceleration LongitudinalAcceleration 
	curvature                Curvature 
	curvatureCalculationMode CurvatureCalculationMode 
	yawRate                  YawRate 
	accelerationControl      AccelerationControl OPTIONAL 
	lanePosition             LanePosition OPTIONAL 
	steeringWheelAngle       SteeringWheelAngle OPTIONAL 
	lateralAcceleration      LateralAcceleration OPTIONAL 
	verticalAcceleration     VerticalAcceleration OPTIONAL 
	performanceClass         PerformanceClass OPTIONAL 
	cenDsrcTollingZone       CenDsrcTollingZone OPTIONAL
    } 

type BasicVehicleContainerLowFrequency {
	vehicleRole    VehicleRole,
	exteriorLights ExteriorLights,
	pathHistory    PathHistory
    }

 type PublicTransportContainer  struct {
	embarkationStatus EmbarkationStatus,
	ptActivation      PtActivation OPTIONAL
    }

type SpecialTransportContainer  struct {
	specialTransportType SpecialTransportType,
	lightBarSirenInUse   LightBarSirenInUse
    }
type DangerousGoodsContainer  struct {
	dangerousGoodsBasic DangerousGoodsBasic
    }

   type RoadWorksContainerBasic  struct {
	roadworksSubCauseCode RoadworksSubCauseCode OPTIONAL,
	lightBarSirenInUse    LightBarSirenInUse,
	closedLanes           ClosedLanes OPTIONAL
    }
    
   type RescueContainer  struct {
	lightBarSirenInUse LightBarSirenInUse
    }
    
  type EmergencyContainer  struct {
	lightBarSirenInUse LightBarSirenInUse,
	incidentIndication CauseCode OPTIONAL,
	emergencyPriority  EmergencyPriority OPTIONAL
    }
    
   type SafetyCarContainer  struct {
	lightBarSirenInUse LightBarSirenInUse,
	incidentIndication CauseCode OPTIONAL,
	trafficRule        TrafficRule OPTIONAL,
	speedLimit         SpeedLimit OPTIONAL
    }
    
   type RSUContainerHighFrequency  struct {
	protectedCommunicationZonesRSU ProtectedCommunicationZonesRSU OPTIONAL,
    }

    type ItsPduHeader type {    //SEQUENCE
	protocolVersion int  
	messageID       int 
	 /*   denm(1),
	    cam(2),
	    poi(3),
	    spat(4),
	    map(5),
	    ivi(6),
	    ev-rsr(7)*/
	stationID       StationID
    }
    
    type StationID int //(0..4294967295)
    
    type ReferencePosition struct { // SEQUENCE 
	latitude                  Latitude,
	longitude                 Longitude,
	positionConfidenceEllipse PosConfidenceEllipse,
	altitude                  Altitude
    }
    
    type DeltaReferencePosition struct {
	deltaLatitude  DeltaLatitude,
	deltaLongitude DeltaLongitude,
	deltaAltitude  DeltaAltitude
    }
    
    type Longitude int 
    /*{
	oneMicrodegreeEast(10),
	oneMicrodegreeWest(-10),
	unavailable(1800000001)
    } (-1800000000..1800000001)*/
    
    type Latitude int
    /*{
	oneMicrodegreeNorth(10),
	oneMicrodegreeSouth(-10),
	unavailable(900000001)
    } (-900000000..900000001)*/
    
    type Altitude struct {
	altitudeValue      AltitudeValue,
	altitudeConfidence AltitudeConfidence
    }
    
    type AltitudeValue int
    /*{
	referenceEllipsoidSurface(0),
	oneCentimeter(1),
	unavailable(800001)
    } (-100000..800001)
*/

   type AltitudeConfidence Enumerated
   /* {
	alt-000-01(0),
	alt-000-02(1),
	alt-000-05(2),
	alt-000-10(3),
	alt-000-20(4),
	alt-000-50(5),
	alt-001-00(6),
	alt-002-00(7),
	alt-005-00(8),
	alt-010-00(9),
	alt-020-00(10),
	alt-050-00(11),
	alt-100-00(12),
	alt-200-00(13),
	outOfRange(14),
	unavailable(15)
    }*/
    
    type DeltaLongitude int
    /* {
	oneMicrodegreeEast(10),
	oneMicrodegreeWest(-10),
	unavailable(131072)
    } (-131071..131072)
    */


    type DeltaLatitude int
    /* {
	oneMicrodegreeNorth(10),
	oneMicrodegreeSouth(-10),
	unavailable(131072)
    } (-131071..131072)
    */

    type DeltaAltitude int
    /*{
	oneCentimeterUp(1),
	oneCentimeterDown(-1),
	unavailable(12800)
    } (-12700..12800)*/
    

    type PosConfidenceEllipse struct {
	semiMajorConfidence  SemiAxisLength
	semiMinorConfidence  SemiAxisLength
	semiMajorOrientation HeadingValue
    }
    
    type PathPoint struct {
	pathPosition  DeltaReferencePosition,
	pathDeltaTime PathDeltaTime //OPTIONAL
    }
    
    type PathDeltaTime int {
	tenMilliSecondsInPast(1)
    } (1..65535, ...)
    
    type PtActivation ::= SEQUENCE {
	ptActivationType PtActivationType,
	ptActivationData PtActivationData
    }
    
    type PtActivationType ::= INTEGER {
	undefinedCodingType(0),
	r09-16CodingType(1),
	vdv-50149CodingType(2)
    } (0..255)
    
    type PtActivationData ::= OCTET STRING (SIZE (1..20))
    
    type AccelerationControl ::= BIT STRING {
	brakePedalEngaged(0),
	gasPedalEngaged(1),
	emergencyBrakeEngaged(2),
	collisionWarningEngaged(3),
	accEngaged(4),
	cruiseControlEngaged(5),
	speedLimiterEngaged(6)
    } (SIZE (7))
    
    type SemiAxisLength ::= INTEGER {
	oneCentimeter(1),
	outOfRange(4094),
	unavailable(4095)
    } (0..4095)
    
    type CauseCode ::= SEQUENCE {
	causeCode    CauseCodeType,
	subCauseCode SubCauseCodeType
    }
    
    type CauseCodeType ::= INTEGER {
	reserved(0),
	trafficCondition(1),
	accident(2),
	roadworks(3),
	adverseWeatherCondition-Adhesion(6),
	hazardousLocation-SurfaceCondition(9),
	hazardousLocation-ObstacleOnTheRoad(10),
	hazardousLocation-AnimalOnTheRoad(11),
	humanPresenceOnTheRoad(12),
	wrongWayDriving(14),
	rescueAndRecoveryWorkInProgress(15),
	adverseWeatherCondition-ExtremeWeatherCondition(17),
	adverseWeatherCondition-Visibility(18),
	adverseWeatherCondition-Precipitation(19),
	slowVehicle(26),
	dangerousEndOfQueue(27),
	vehicleBreakdown(91),
	postCrash(92),
	humanProblem(93),
	stationaryVehicle(94),
	emergencyVehicleApproaching(95),
	hazardousLocation-DangerousCurve(96),
	collisionRisk(97),
	signalViolation(98),
	dangerousSituation(99)
    } (0..255)
    
    type SubCauseCodeType ::= INTEGER (0..255)
    
    Ttype rafficConditionSubCauseCode ::= INTEGER {
	unavailable(0),
	increasedVolumeOfTraffic(1),
	trafficJamSlowlyIncreasing(2),
	trafficJamIncreasing(3),
	trafficJamStronglyIncreasing(4),
	trafficStationary(5),
	trafficJamSlightlyDecreasing(6),
	trafficJamDecreasing(7),
	trafficJamStronglyDecreasing(8)
    } (0..255)
    
    type AccidentSubCauseCode ::= INTEGER {
	unavailable(0),
	multiVehicleAccident(1),
	heavyAccident(2),
	accidentInvolvingLorry(3),
	accidentInvolvingBus(4),
	accidentInvolvingHazardousMaterials(5),
	accidentOnOppositeLane(6),
	unsecuredAccident(7),
	assistanceRequested(8)
    } (0..255)
    
    type RoadworksSubCauseCode ::= INTEGER {
	unavailable(0),
	majorRoadworks(1),
	roadMarkingWork(2),
	slowMovingRoadMaintenance(3),
	shortTermStationaryRoadworks(4),
	streetCleaning(5),
	winterService(6)
    } (0..255)
    
    type HumanPresenceOnTheRoadSubCauseCode ::= INTEGER {
	unavailable(0),
	childrenOnRoadway(1),
	cyclistOnRoadway(2),
	motorcyclistOnRoadway(3)
    } (0..255)
    
    type WrongWayDrivingSubCauseCode ::= INTEGER {
	unavailable(0),
	wrongLane(1),
	wrongDirection(2)
    } (0..255)
    
    type AdverseWeatherCondition-ExtremeWeatherConditionSubCauseCode ::= INTEGER {
	unavailable(0),
	strongWinds(1),
	damagingHail(2),
	hurricane(3),
	thunderstorm(4),
	tornado(5),
	blizzard(6)
    } (0..255)
    
    type AdverseWeatherCondition-AdhesionSubCauseCode ::= INTEGER {
	unavailable(0),
	heavyFrostOnRoad(1),
	fuelOnRoad(2),
	mudOnRoad(3),
	snowOnRoad(4),
	iceOnRoad(5),
	blackIceOnRoad(6),
	oilOnRoad(7),
	looseChippings(8),
	instantBlackIce(9),
	roadsSalted(10)
    } (0..255)
    
    type AdverseWeatherCondition-VisibilitySubCauseCode ::= INTEGER {
	unavailable(0),
	fog(1),
	smoke(2),
	heavySnowfall(3),
	heavyRain(4),
	heavyHail(5),
	lowSunGlare(6),
	sandstorms(7),
	swarmsOfInsects(8)
    } (0..255)
    
    type AdverseWeatherCondition-PrecipitationSubCauseCode ::= INTEGER {
	unavailable(0),
	heavyRain(1),
	heavySnowfall(2),
	softHail(3)
    } (0..255)
    
    type SlowVehicleSubCauseCode ::= INTEGER {
	unavailable(0),
	maintenanceVehicle(1),
	vehiclesSlowingToLookAtAccident(2),
	abnormalLoad(3),
	abnormalWideLoad(4),
	convoy(5),
	snowplough(6),
	deicing(7),
	saltingVehicles(8)
    } (0..255)
    
    type StationaryVehicleSubCauseCode ::= INTEGER {
	unavailable(0),
	humanProblem(1),
	vehicleBreakdown(2),
	postCrash(3),
	publicTransportStop(4),
	carryingDangerousGoods(5)
    } (0..255)
    
    type HumanProblemSubCauseCode ::= INTEGER {
	unavailable(0),
	glycemiaProblem(1),
	heartProblem(2)
    } (0..255)
    
    type EmergencyVehicleApproachingSubCauseCode ::= INTEGER {
	unavailable(0),
	emergencyVehicleApproaching(1),
	prioritizedVehicleApproaching(2)
    } (0..255)
    
    type HazardousLocation-DangerousCurveSubCauseCode ::= INTEGER {
	unavailable(0),
	dangerousLeftTurnCurve(1),
	dangerousRightTurnCurve(2),
	multipleCurvesStartingWithUnknownTurningDirection(3),
	multipleCurvesStartingWithLeftTurn(4),
	multipleCurvesStartingWithRightTurn(5)
    } (0..255)
    
    type HazardousLocation-SurfaceConditionSubCauseCode ::= INTEGER {
	unavailable(0),
	rockfalls(1),
	earthquakeDamage(2),
	sewerCollapse(3),
	subsidence(4),
	snowDrifts(5),
	stormDamage(6),
	burstPipe(7),
	volcanoEruption(8),
	fallingIce(9)
    } (0..255)
    
    type HazardousLocation-ObstacleOnTheRoadSubCauseCode ::= INTEGER {
	unavailable(0),
	shedLoad(1),
	partsOfVehicles(2),
	partsOfTyres(3),
	bigObjects(4),
	fallenTrees(5),
	hubCaps(6),
	waitingVehicles(7)
    } (0..255)
    
    type HazardousLocation-AnimalOnTheRoadSubCauseCode ::= INTEGER {
	unavailable(0),
	wildAnimals(1),
	herdOfAnimals(2),
	smallAnimals(3),
	largeAnimals(4)
    } (0..255)
    
    type CollisionRiskSubCauseCode ::= INTEGER {
	unavailable(0),
	longitudinalCollisionRisk(1),
	crossingCollisionRisk(2),
	lateralCollisionRisk(3),
	vulnerableRoadUser(4)
    } (0..255)
    
    type SignalViolationSubCauseCode ::= INTEGER {
	unavailable(0),
	stopSignViolation(1),
	trafficLightViolation(2),
	turningRegulationViolation(3)
    } (0..255)
    
    type RescueAndRecoveryWorkInProgressSubCauseCode ::= INTEGER {
	unavailable(0),
	emergencyVehicles(1),
	rescueHelicopterLanding(2),
	policeActivityOngoing(3),
	medicalEmergencyOngoing(4),
	childAbductionInProgress(5)
    } (0..255)
    
    type DangerousEndOfQueueSubCauseCode ::= INTEGER {
	unavailable(0),
	suddenEndOfQueue(1),
	queueOverHill(2),
	queueAroundBend(3),
	queueInTunnel(4)
    } (0..255)
    
    type DangerousSituationSubCauseCode ::= INTEGER {
	unavailable(0),
	emergencyElectronicBrakeEngaged(1),
	preCrashSystemEngaged(2),
	espEngaged(3),
	absEngaged(4),
	aebEngaged(5),
	brakeWarningEngaged(6),
	collisionRiskWarningEngaged(7)
    } (0..255)
    
    type VehicleBreakdownSubCauseCode ::= INTEGER {
	unavailable(0),
	lackOfFuel(1),
	lackOfBatteryPower(2),
	engineProblem(3),
	transmissionProblem(4),
	engineCoolingProblem(5),
	brakingSystemProblem(6),
	steeringProblem(7),
	tyrePuncture(8)
    } (0..255)
    
    type PostCrashSubCauseCode ::= INTEGER {
	unavailable(0),
	accidentWithoutECallTriggered(1),
	accidentWithECallManuallyTriggered(2),
	accidentWithECallAutomaticallyTriggered(3),
	accidentWithECallTriggeredWithoutAccessToCellularNetwork(4)
    } (0..255)
    
    type Curvature ::= SEQUENCE {
	curvatureValue      CurvatureValue,
	curvatureConfidence CurvatureConfidence
    }
    
    type CurvatureValue ::= INTEGER {
	straight(0),
	reciprocalOf1MeterRadiusToRight(-30000),
	reciprocalOf1MeterRadiusToLeft(30000),
	unavailable(30001)
    } (-30000..30001)
    
    type CurvatureConfidence ::= ENUMERATED {
	onePerMeter-0-00002(0),
	onePerMeter-0-0001(1),
	onePerMeter-0-0005(2),
	onePerMeter-0-002(3),
	onePerMeter-0-01(4),
	onePerMeter-0-1(5),
	outOfRange(6),
	unavailable(7)
    }
    
    type CurvatureCalculationMode ::= ENUMERATED {
	yawRateUsed(0),
	yawRateNotUsed(1),
	unavailable(2),
	...
    }
    
    type Heading ::= SEQUENCE {
	headingValue      HeadingValue,
	headingConfidence HeadingConfidence
    }
    
    type HeadingValue ::= INTEGER {
	wgs84North(0),
	wgs84East(900),
	wgs84South(1800),
	wgs84West(2700),
	unavailable(3601)
    } (0..3601)
    
    type HeadingConfidence ::= INTEGER {
	equalOrWithinZeroPointOneDegree(1),
	equalOrWithinOneDegree(10),
	outOfRange(126),
	unavailable(127)
    } (1..127)
    
    type LanePosition ::= INTEGER {
	offTheRoad(-1),
	hardShoulder(0),
	outermostDrivingLane(1),
	secondLaneFromOutside(2)
    } (-1..14)
    
    type ClosedLanes ::= SEQUENCE {
	hardShoulderStatus HardShoulderStatus OPTIONAL,
	drivingLaneStatus  DrivingLaneStatus,
	                   ...
    }
    
    type HardShoulderStatus ::= ENUMERATED {
	availableForStopping(0),
	closed(1),
	availableForDriving(2)
    }
    
    type DrivingLaneStatus ::= BIT STRING {
	outermostLaneClosed(1),
	secondLaneFromOutsideClosed(2)
    } (SIZE (1..14))
    
    type PerformanceClass ::= INTEGER {
	unavailable(0),
	performanceClassA(1),
	performanceClassB(2)
    } (0..7)
    
    type SpeedValue ::= INTEGER {
	standstill(0),
	oneCentimeterPerSec(1),
	unavailable(16383)
    } (0..16383)
    
    type SpeedConfidence ::= INTEGER {
	equalOrWithinOneCentimeterPerSec(1),
	equalOrWithinOneMeterPerSec(100),
	outOfRange(126),
	unavailable(127)
    } (1..127)
    
    type VehicleMass ::= INTEGER {
	hundredKg(1),
	unavailable(1024)
    } (1..1024)
    
   type Speed ::= SEQUENCE {
	speedValue      SpeedValue,
	speedConfidence SpeedConfidence
    }
    
    type DriveDirection ::= ENUMERATED {
	forward(0),
	backward(1),
	unavailable(2)
    }
    
    type EmbarkationStatus ::= BOOLEAN
    
    type LongitudinalAcceleration ::= SEQUENCE {
	longitudinalAccelerationValue      LongitudinalAccelerationValue,
	longitudinalAccelerationConfidence AccelerationConfidence
    }
    
    type LongitudinalAccelerationValue ::= INTEGER {
	pointOneMeterPerSecSquaredForward(1),
	pointOneMeterPerSecSquaredBackward(-1),
	unavailable(161)
    } (-160..161)
    
    type AccelerationConfidence ::= INTEGER {
	pointOneMeterPerSecSquared(1),
	outOfRange(101),
	unavailable(102)
    } (0..102)
    
    type LateralAcceleration ::= SEQUENCE {
	lateralAccelerationValue      LateralAccelerationValue,
	lateralAccelerationConfidence AccelerationConfidence
    }
    
    type LateralAccelerationValue ::= INTEGER {
	pointOneMeterPerSecSquaredToRight(-1),
	pointOneMeterPerSecSquaredToLeft(1),
	unavailable(161)
    } (-160..161)
    
    type VerticalAcceleration ::= SEQUENCE {
	verticalAccelerationValue      VerticalAccelerationValue,
	verticalAccelerationConfidence AccelerationConfidence
    }
    
    type VerticalAccelerationValue ::= INTEGER {
	pointOneMeterPerSecSquaredUp(1),
	pointOneMeterPerSecSquaredDown(-1),
	unavailable(161)
    } (-160..161)
    
    type StationType ::= INTEGER {
	unknown(0),
	pedestrian(1),
	cyclist(2),
	moped(3),
	motorcycle(4),
	passengerCar(5),
	bus(6),
	lightTruck(7),
	heavyTruck(8),
	trailer(9),
	specialVehicles(10),
	tram(11),
	roadSideUnit(15)
    } (0..255)
    
    type ExteriorLights ::= BIT STRING {
	lowBeamHeadlightsOn(0),
	highBeamHeadlightsOn(1),
	leftTurnSignalOn(2),
	rightTurnSignalOn(3),
	daytimeRunningLightsOn(4),
	reverseLightOn(5),
	fogLightOn(6),
	parkingLightsOn(7)
    } (SIZE (8))
    
    type DangerousGoodsBasic ::= ENUMERATED {
	explosives1(0),
	explosives2(1),
	explosives3(2),
	explosives4(3),
	explosives5(4),
	explosives6(5),
	flammableGases(6),
	nonFlammableGases(7),
	toxicGases(8),
	flammableLiquids(9),
	flammableSolids(10),
	substancesLiableToSpontaneousCombustion(11),
	substancesEmittingFlammableGasesUponContactWithWater(12),
	oxidizingSubstances(13),
	organicPeroxides(14),
	toxicSubstances(15),
	infectiousSubstances(16),
	radioactiveMaterial(17),
	corrosiveSubstances(18),
	miscellaneousDangerousSubstances(19)
    }
    
    type DangerousGoodsExtended ::= SEQUENCE {
	dangerousGoodsType  DangerousGoodsBasic,
	unNumber            INTEGER (0..9999),
	elevatedTemperature BOOLEAN,
	tunnelsRestricted   BOOLEAN,
	limitedQuantity     BOOLEAN,
	emergencyActionCode IA5String (SIZE (1..24)) OPTIONAL,
	phoneNumber         IA5String (SIZE (1..24)) OPTIONAL,
	companyName         UTF8String (SIZE (1..24)) OPTIONAL
    }
    
    type SpecialTransportType ::= BIT STRING {
	heavyLoad(0),
	excessWidth(1),
	excessLength(2),
	excessHeight(3)
    } (SIZE (4))
    
    type LightBarSirenInUse ::= BIT STRING {
	lightBarActivated(0),
	sirenActivated(1)
    } (SIZE (2))
    
    type HeightLonCarr ::= INTEGER {
	oneCentimeter(1),
	unavailable(100)
    } (1..100)
    
    type PosLonCarr ::= INTEGER {
	oneCentimeter(1),
	unavailable(127)
    } (1..127)
    
    type PosPillar ::= INTEGER {
	tenCentimeters(1),
	unavailable(30)
    } (1..30)
    
    type PosCentMass ::= INTEGER {
	tenCentimeters(1),
	unavailable(63)
    } (1..63)
    
    type RequestResponseIndication ::= ENUMERATED {
	request(0),
	response(1)
    }
    
    type SpeedLimit ::= INTEGER {
	oneKmPerHour(1)
    } (1..255)
    
    type StationarySince ::= ENUMERATED {
	lessThan1Minute(0),
	lessThan2Minutes(1),
	lessThan15Minutes(2),
	equalOrGreater15Minutes(3)
    }
    
    type Temperature ::= INTEGER {
	equalOrSmallerThanMinus60Deg(-60),
	oneDegreeCelsius(1),
	equalOrGreaterThan67Deg(67)
    } (-60..67)
    
    type TrafficRule ::= ENUMERATED {
	noPassing(0),
	noPassingForTrucks(1),
	passToRight(2),
	passToLeft(3),
	...
    }
    
    type WheelBaseVehicle ::= INTEGER {
	tenCentimeters(1),
	unavailable(127)
    } (1..127)
    
    type TurningRadius ::= INTEGER {
	point4Meters(1),
	unavailable(255)
    } (1..255)
    
    type PosFrontAx ::= INTEGER {
	tenCentimeters(1),
	unavailable(20)
    } (1..20)
    
    type PositionOfOccupants ::= BIT STRING {
	row1LeftOccupied(0),
	row1RightOccupied(1),
	row1MidOccupied(2),
	row1NotDetectable(3),
	row1NotPresent(4),
	row2LeftOccupied(5),
	row2RightOccupied(6),
	row2MidOccupied(7),
	row2NotDetectable(8),
	row2NotPresent(9),
	row3LeftOccupied(10),
	row3RightOccupied(11),
	row3MidOccupied(12),
	row3NotDetectable(13),
	row3NotPresent(14),
	row4LeftOccupied(15),
	row4RightOccupied(16),
	row4MidOccupied(17),
	row4NotDetectable(18),
	row4NotPresent(19)
    } (SIZE (20))
    
    type PositioningSolutionType ::= ENUMERATED {
	noPositioningSolution(0),
	sGNSS(1),
	dGNSS(2),
	sGNSSplusDR(3),
	dGNSSplusDR(4),
	dR(5),
	...
    }
    
    type VehicleIdentification ::= SEQUENCE {
	wMInumber WMInumber OPTIONAL,
	vDS       VDS OPTIONAL,
	          ...
    }
    
    type WMInumber ::= IA5String (SIZE (1..3))
    
    type VDS ::= IA5String (SIZE (6))
    
    type EnergyStorageType ::= BIT STRING {
	hydrogenStorage(0),
	electricEnergyStorage(1),
	liquidPropaneGas(2),
	compressedNaturalGas(3),
	diesel(4),
	gasoline(5),
	ammonia(6)
    } (SIZE (7))
    
    type VehicleLength ::= SEQUENCE {
	vehicleLengthValue                VehicleLengthValue,
	vehicleLengthConfidenceIndication VehicleLengthConfidenceIndication
    }
    
    type VehicleLengthValue ::= INTEGER {
	tenCentimeters(1),
	outOfRange(1022),
	unavailable(1023)
    } (1..1023)
    
    type VehicleLengthConfidenceIndication ::= ENUMERATED {
	noTrailerPresent(0),
	trailerPresentWithKnownLength(1),
	trailerPresentWithUnknownLength(2),
	trailerPresenceIsUnknown(3),
	unavailable(4)
    }
    
    type VehicleWidth ::= INTEGER {
	tenCentimeters(1),
	outOfRange(61),
	unavailable(62)
    } (1..62)
    
    type PathHistory ::= SEQUENCE SIZE (0..40) OF PathPoint
    
    type EmergencyPriority ::= BIT STRING {
	requestForRightOfWay(0),
	requestForFreeCrossingAtATrafficLight(1)
    } (SIZE (2))
    
    type InformationQuality ::= INTEGER {
	unavailable(0),
	lowest(1),
	highest(7)
    } (0..7)
    
    type RoadType ::= ENUMERATED {
	urban-NoStructuralSeparationToOppositeLanes(0),
	urban-WithStructuralSeparationToOppositeLanes(1),
	nonUrban-NoStructuralSeparationToOppositeLanes(2),
	nonUrban-WithStructuralSeparationToOppositeLanes(3)
    }
    
    type SteeringWheelAngle ::= SEQUENCE {
	steeringWheelAngleValue      SteeringWheelAngleValue,
	steeringWheelAngleConfidence SteeringWheelAngleConfidence
    }
    
    type SteeringWheelAngleValue ::= INTEGER {
	straight(0),
	onePointFiveDegreesToRight(-1),
	onePointFiveDegreesToLeft(1),
	unavailable(512)
    } (-511..512)
    
    type SteeringWheelAngleConfidence ::= INTEGER {
	equalOrWithinOnePointFiveDegree(1),
	outOfRange(126),
	unavailable(127)
    } (1..127)
    
    type TimestampIts ::= INTEGER {
	utcStartOf2004(0),
	oneMillisecAfterUTCStartOf2004(1)
    } (0..4398046511103)
    
    type VehicleRole ::= ENUMERATED {
	default(0),
	publicTransport(1),
	specialTransport(2),
	dangerousGoods(3),
	roadWork(4),
	rescue(5),
	emergency(6),
	safetyCar(7),
	agriculture(8),
	commercial(9),
	military(10),
	roadOperator(11),
	taxi(12),
	reserved1(13),
	reserved2(14),
	reserved3(15)
    }
    
    type YawRate ::= SEQUENCE {
	yawRateValue      YawRateValue,
	yawRateConfidence YawRateConfidence
    }
    
    type YawRateValue ::= INTEGER {
	straight(0),
	degSec-000-01ToRight(-1),
	degSec-000-01ToLeft(1),
	unavailable(32767)
    } (-32766..32767)
    
    type YawRateConfidence ::= ENUMERATED {
	degSec-000-01(0),
	degSec-000-05(1),
	degSec-000-10(2),
	degSec-001-00(3),
	degSec-005-00(4),
	degSec-010-00(5),
	degSec-100-00(6),
	outOfRange(7),
	unavailable(8)
    }
    
    type ProtectedZoneType ::= ENUMERATED {
	cenDsrcTolling(0),
	...
    }
    
    type RelevanceDistance ::= ENUMERATED {
	lessThan50m(0),
	lessThan100m(1),
	lessThan200m(2),
	lessThan500m(3),
	lessThan1000m(4),
	lessThan5km(5),
	lessThan10km(6),
	over10km(7)
    }
    
    type RelevanceTrafficDirection ::= ENUMERATED {
	allTrafficDirections(0),
	upstreamTraffic(1),
	downstreamTraffic(2),
	oppositeTraffic(3)
    }
    
    type TransmissionInterval ::= INTEGER {
	oneMilliSecond(1),
	tenSeconds(10000)
    } (1..10000)
    
    type ValidityDuration ::= INTEGER {
	timeOfDetection(0),
	oneSecondAfterDetection(1)
    } (0..86400)
    
    type ActionID ::= SEQUENCE {
	originatingStationID StationID,
	sequenceNumber       SequenceNumber
    }
    
    type ItineraryPath ::= SEQUENCE SIZE (1..40) OF ReferencePosition
    
    type ProtectedCommunicationZone ::= SEQUENCE {
	protectedZoneType      ProtectedZoneType,
	expiryTime             TimestampIts OPTIONAL,
	protectedZoneLatitude  Latitude,
	protectedZoneLongitude Longitude,
	protectedZoneRadius    ProtectedZoneRadius OPTIONAL,
	protectedZoneID        ProtectedZoneID OPTIONAL
    }
    
    type Traces ::= SEQUENCE SIZE (1..7) OF PathHistory
    
    type NumberOfOccupants int
    /*
	oneOccupant(1),
	unavailable(127)
    } (0..127)*/
    
    type SequenceNumber int
    // (0..65535)
    
    type PositionOfPillars ::= SEQUENCE SIZE (1..3, ...) OF PosPillar
    
    type RestrictedTypes ::= SEQUENCE SIZE (1..3, ...) OF StationType
    
    type EventHistory ::= SEQUENCE SIZE (1..23) OF EventPoint
    
    type EventPoint ::= SEQUENCE {
	eventPosition      DeltaReferencePosition,
	eventDeltaTime     PathDeltaTime OPTIONAL,
	informationQuality InformationQuality
    }
    
    type ProtectedCommunicationZonesRSU ProtectedCommunicationZone[16]
    //::= SEQUENCE SIZE (1..16) OF
    
    type CenDsrcTollingZone ::= SEQUENCE {
	protectedZoneLatitude  Latitude,
	protectedZoneLongitude Longitude,
	cenDsrcTollingZoneID   CenDsrcTollingZoneID OPTIONAL
    }
    
    type ProtectedZoneRadius ::= INTEGER {
	oneMeter(1)
    } (1..255, ...)
    
    type ProtectedZoneID ::= INTEGER (0..134217727)
    
    type CenDsrcTollingZoneID ::= ProtectedZoneID