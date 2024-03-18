package aggregations

import "testing"

func TestGetAggregationsInformationSuccessfully(t *testing.T) {

}

func TestGetAggsByHypertableSucessfully(t *testing.T) {}

func TestGetAggsByHypertableInexistentHypertable(t *testing.T) {
	err := GetAggsByHypertable("inexistent_hypertable")
}
