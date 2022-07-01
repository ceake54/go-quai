package tests

import (
	"github.com/spruce-solutions/go-quai/core/types"
)

// Example test for a fork choice scenario shown in slide00 (not a real slide)
func ForkChoiceTest_Slide05() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{ // Zone1
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, 2, 2}, [3]string{"", "z00_1", "z00_1"}, ""}, //Fork at z00_1
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, 3, 5}, [3]string{"", "z01_1", ""}, ""}, //Fork at z01_1
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 4, 8}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 2, 1}, [3]string{"", "z00_1", ""}, "z01_1"},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide00 (not a real slide)
func ForkChoiceTest_Slide06() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{}, // ... Region1 omitted
		{ // Region2
			{ //Zone1
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, "z00_1"},
			},
			{}, //... Zone2 omitted
			{}, //... Zone3 omitted
		},
		{ //Region 3
			{ //Zone1
				&types.BlockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
			},
			{ //Zone2
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 1, 2}, [3]string{"z00_1", "", ""}, ""}, //Fork at z00_1
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
			},
			{ //Zone3
				&types.BlockGenSpec{[3]int{-1, 2, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 2}, [3]string{}, ""},
			},
		},
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide08
func ForkChoiceTest_Slide08() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 11}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{-1, 3, 12}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 12}, [3]string{"", "", "z00_1"}, ""}, //Fork at z00_1
				&types.BlockGenSpec{[3]int{-1, -1, 13}, [3]string{}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide11
func ForkChoiceTest_Slide11() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 11}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{-1, 3, 12}, [3]string{}, "z00_3"},
				&types.BlockGenSpec{[3]int{-1, -1, 12}, [3]string{"", "", "z00_1"}, "z00_2"}, //Fork at z00_1
				&types.BlockGenSpec{[3]int{-1, -1, 13}, [3]string{"", "", "z00_2"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 13}, [3]string{"", "", "z00_3"}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide13
func ForkChoiceTest_Slide13() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, "z01_1"},
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				//...
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{"", "z00_1", "z01_1"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide15
func ForkChoiceTest_Slide15() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{ // Zone1
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 7}, [3]string{}, ""},
			},
			{ // Zone2
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, "z12_1"},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				//...
				&types.BlockGenSpec{[3]int{3, 4, 7}, [3]string{"z31_1", "", ""}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{
			{}, // ... Zone1 omitted
			{ // Zone2
				&types.BlockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
			},
			{}, // ... Zone3 omitted
		},
		{
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 3}, [3]string{"z12_1", "", ""}, "z31_1"},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 7}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 4, 7}, [3]string{}, ""},
			},
		},
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide17
func ForkChoiceTest_Slide17() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{"", "z00_1", ""}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 7}, [3]string{"", "z02_1", ""}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 2, 1}, [3]string{"", "z00_1", ""}, "z02_1"},
			},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide18
func ForkChoiceTest_Slide18() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, "z00_1"},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{"", "z00_1", ""}, ""},
				&types.BlockGenSpec{[3]int{1, 2, 4}, [3]string{"", "z00_1", ""}, ""},
			},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide23
func ForkChoiceTest_Slide23() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, "z00_1"},
				//...
				&types.BlockGenSpec{[3]int{-1, 2, 5}, [3]string{"", "", "z00_1"}, ""},
				&types.BlockGenSpec{[3]int{1, 3, 5}, [3]string{"", "", "z00_1"}, ""},
			},
			{},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide25
func ForkChoiceTest_Slide25() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 1, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
			},
			{},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide26
func ForkChoiceTest_Slide26() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{1, 2, 5}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
			},
			{},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide27
func ForkChoiceTest_Slide27() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{2, 3, 4}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, 3, 4}, [3]string{"", "z00_1", "z00_1"}, ""}, //Fork at z00_1
				&types.BlockGenSpec{[3]int{-1, 13, 14}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{},
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide29
func ForkChoiceTest_Slide29() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 1, 3}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{"", "", "z00_1"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 60}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{
			{
				&types.BlockGenSpec{[3]int{2, 1, 1}, [3]string{}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide30
func ForkChoiceTest_Slide30() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 1, 3}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 60}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{3, 3, 4}, [3]string{"z10_1", "z00_1", "z00_1"}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{
			{
				&types.BlockGenSpec{[3]int{2, 1, 1}, [3]string{}, "z10_1"},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide31
func ForkChoiceTest_Slide31() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{"", "", "z00_1"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{3, 3, 6}, [3]string{"z12_1", "z00_1", ""}, ""},
			},
			{}, // ... Zone2 omitted
			{}, // ... Zone3 omitted
		},
		{
			{}, // ... Zone1 omitted
			{}, // ... Zone2 omitted
			{
				&types.BlockGenSpec{[3]int{2, 1, 1}, [3]string{}, "z12_1"},
			},
		},
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide32
func ForkChoiceTest_Slide32() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{3, 3, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{"", "", "z00_1"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 61}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{4, 3, 62}, [3]string{"", "z00_1", ""}, ""},
			},
			{},
			{},
		},
		{
			{},
			{},
			{
				&types.BlockGenSpec{[3]int{2, 1, 1}, [3]string{}, "z12_1"},
			},
		},
		{
			{
				&types.BlockGenSpec{[3]int{3, 1, 1}, [3]string{"z12_1", "", ""}, ""},
			},
			{},
			{},
		},
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide36
func ForkChoiceTest_Slide36() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{5, 5, 7}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 5}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 7, 11}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 12}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{6, 8, 13}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 4, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 6, 2}, [3]string{}, ""},
			},
		},
		{
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 2}, [3]string{}, ""},
			},
			{},
			{},
		},
		{}, // ... Region3 omitted
	}
	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide37
func ForkChoiceTest_Slide37() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 5, 7}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 5}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 7, 11}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 12}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{5, 8, 13}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 4, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 6, 2}, [3]string{}, ""},
			},
		},
		{
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 2}, [3]string{}, ""},
			},
			{},
			{},
		},
		{}, // ... Region3 omitted
	}
	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide38
func ForkChoiceTest_Slide38() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 5, 7}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 7, 11}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 12}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 8, 13}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 3, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 4, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 6, 3}, [3]string{}, ""},
			},
		},
		{},
		{}, // ... Region3 omitted
	}
	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide39
func ForkChoiceTest_Slide39() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 5, 7}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 5}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 7, 11}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 4, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 6, 2}, [3]string{}, ""},
			},
		},
		{},
		{}, // ... Region3 omitted
	}
	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide40
func ForkChoiceTest_Slide40() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 5, 9}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 10}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{4, 4, 7}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{3, 3, 1}, [3]string{}, ""},
			},
		},
		{
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 1, 2}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 2, 3}, [3]string{}, ""},
			},
			{}, //Zone3 omitted
		},
		{}, //Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide41
func ForkChoiceTest_Slide41() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{2, 2, 3}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, -1, 3}, [3]string{"", "", "z00_1"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 50}, [3]string{}, ""},
			},
			{},
			{},
		},
		{}, // ... Region2 omitted
		{}, // ... Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide42
func ForkChoiceTest_Slide42() {

	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 2, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{2, 3, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{3, 4, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 8, 12}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 5, 3}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{6, 7, 7}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{5, 6, 1}, [3]string{}, ""},
			},
		},
		{
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 1, 2}, [3]string{}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, -1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{4, 2, 3}, [3]string{}, ""},
			},
			{}, //Zone3 omitted
		},
		{}, //Region3 omitted
	}

	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}

// Example test for a fork choice scenario shown in slide43
func ForkChoiceTest_Slide43() {
	/**************************************************************************
	 * Define the test scenario:
	 *************************************************************************/
	// The graph defines blocks to be mined by each of the chains in the network, to create a chain fork test scenario.
	graph := [3][3][]*types.BlockGenSpec{
		{ // Region1
			{
				&types.BlockGenSpec{[3]int{-1, 1, 1}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 2}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 4}, [3]string{}, "z00_1"},
				&types.BlockGenSpec{[3]int{-1, 3, 5}, [3]string{}, ""},
				// ...
				&types.BlockGenSpec{[3]int{-1, -1, 5}, [3]string{"", "", "z00_1"}, ""},
				&types.BlockGenSpec{[3]int{-1, -1, 6}, [3]string{}, ""},
				&types.BlockGenSpec{[3]int{-1, 3, 8}, [3]string{"", "z01_1", ""}, ""},
			},
			{
				&types.BlockGenSpec{[3]int{-1, 2, 1}, [3]string{}, "z01_1"},
				&types.BlockGenSpec{[3]int{-1, 4, 2}, [3]string{}, ""},
			},
			{},
		},
		{},
		{}, // ... Region3 omitted
	}
	/**************************************************************************
	 * Generate blocks and run the test scenario simulation:
	 *************************************************************************/
	clients, blocks := GenAndRun(graph)

	/**************************************************************************
	 * PASS/FAIL criteria:
	 *************************************************************************/
	dutLoc := []int{0, 0}
	clients.CheckCriteria(dutLoc, blocks)
	// TEST PASS
}
