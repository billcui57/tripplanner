package routeServiceTest

import (
	"github/billcui57/tripplanner/types"
	"testing"
)

func TestGetSegmentRouteInToDaysDrives(t *testing.T) {

	daydrive := types.NewIDayDriveBuilder().
		EndLocation(
			//Ottawa
			*types.NewIGeoCodeBuilder().
				Longitude(45.4215).
				Latitude(75.6972).
				Build(),
		).
		StartLocation(
			//San Francisco
			*types.NewIGeoCodeBuilder().
				Longitude(37.7749).
				Latitude(122.4194).
				Build(),
		).Build()
	daydrive.StartLocation.Latitude = 2
}
