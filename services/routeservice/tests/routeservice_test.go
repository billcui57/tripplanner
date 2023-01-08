package routeServiceTest

import (
	"github.com/go-test/deep"
	routeService "github/billcui57/tripplanner/services/routeservice"
	"github/billcui57/tripplanner/types"
	"testing"
)

func TestGetSegmentStepsInToDaysDrives(t *testing.T) {

	steps := make([]types.IStep, 20)
	for i := range steps {
		steps[i] = *types.NewIStepBuilder().
			StartLocation(
				*types.NewIGeoCodeBuilder().
					Latitude(float64(i)).
					Longitude(float64(i)).
					Build(),
			).
			EndLocation(
				*types.NewIGeoCodeBuilder().
					Latitude(float64(i + 1)).
					Longitude(float64(i + 1)).
					Build(),
			).
			DurationInSeconds(1).
			DistanceInMeters(1).
			Build()
	}
	route := types.NewIRouteBuilder().Steps(steps).Build()

	daydrives, err := routeService.SegmentRouteInToDaysDrives(route, 10)

	if err != nil {
		t.Errorf("Got Error")
	}

	expectedDayDrives := make([]types.IDayDrive, 2)
	expectedDayDrives[0] = *types.NewIDayDriveBuilder().
		StartLocation(
			*types.NewIGeoCodeBuilder().
				Latitude(float64(0)).
				Longitude(float64(0)).
				Build(),
		).
		EndLocation(
			*types.NewIGeoCodeBuilder().
				Latitude(float64(10)).
				Longitude(float64(10)).
				Build(),
		).
		DurationInSeconds(10).
		DistanceInMeters(10).
		Build()
	expectedDayDrives[1] = *types.NewIDayDriveBuilder().
		StartLocation(
			*types.NewIGeoCodeBuilder().
				Latitude(float64(10)).
				Longitude(float64(10)).
				Build(),
		).
		EndLocation(
			*types.NewIGeoCodeBuilder().
				Latitude(float64(20)).
				Longitude(float64(20)).
				Build(),
		).
		DurationInSeconds(10).
		DistanceInMeters(10).
		Build()
	if diff := deep.Equal(daydrives, expectedDayDrives); diff != nil {
		t.Error(diff)
	}
}
