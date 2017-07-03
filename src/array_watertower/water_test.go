package main

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {

	aWater := []int{5, 3, 7, 2, 6, 4, 5, 9, 1, 2}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 14 {
		fmt.Printf("The amount of water is %d\n", nAmount)
		t.Fail()
	}
}

func TestSingleTower(t *testing.T) {

	aWater := []int{5}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 0 {
		t.Fail()
	}
}

func TestOpenEnd(t *testing.T) {

	aWater := []int{5, 3, 7, 2, 6, 4, 5, 6, 1, 2}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 10 {
		t.Fail()
	}
}

func TestCraigsProblematic(t *testing.T) {

	aWater := []int{10, 9, 8, 7, 6, 5, 6, 7, 8, 7, 6, 5, 6, 7, 3, 5}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 15 {
		t.Fail()
	}
}

func TestEqualTowers(t *testing.T) {

	aWater := []int{5, 5}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 0 {
		t.Fail()
	}
}

func TestUnEqualTowers(t *testing.T) {

	aWater := []int{4, 5}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 0 {
		t.Fail()
	}
}

func TestZeroHeightTowers(t *testing.T) {

	aWater := []int{0, 0}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 0 {
		t.Fail()
	}
}

func TestOddTowers(t *testing.T) {

	aWater := []int{0, 0, 1}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 0 {
		t.Fail()
	}

}

func TestOddTowersHaveWater(t *testing.T) {

	aWater := []int{5, 1, 5}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 4 {
		t.Fail()
	}

}

func TestEvenTowersHaveWater(t *testing.T) {

	aWater := []int{5, 1, 1, 5}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 8 {
		t.Fail()
	}
}

func TestEvenTowersNoWater(t *testing.T) {

	aWater := []int{1, 5, 5, 1}
	nAmount := aWaterTowerCompute(aWater)

	if nAmount != 0 {
		t.Fail()
	}

}
