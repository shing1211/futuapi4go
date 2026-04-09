package examples

import (
	"os"
	"os/exec"
	"path/filepath"
	"testing"
)

// getExamplesDir returns the path to the examples directory
func getExamplesDir() string {
	// Try to find examples directory relative to module root
	wd, _ := os.Getwd()
	// Walk up to find cmd/examples
	for i := 0; i < 5; i++ {
		p := filepath.Join(wd, "cmd", "examples")
		if _, err := os.Stat(p); err == nil {
			return p
		}
		wd = filepath.Dir(wd)
	}
	return ""
}

func TestExamplesCompile(t *testing.T) {
	examplesDir := getExamplesDir()
	if examplesDir == "" {
		t.Skip("examples directory not found")
	}

	// List of example directories to test
	examples := []string{
		"qot_get_basic_qot",
		"qot_get_kl",
		"qot_get_order_book",
		"qot_get_ticker",
		"qot_get_rt",
		"qot_get_broker",
		"qot_get_capital_flow",
		"qot_get_static_info",
		"qot_get_trade_date",
		"qot_subscribe",
		"qot_stock_filter",
		"trd_get_acc_list",
		"trd_get_funds",
		"trd_get_position_list",
		"trd_unlock_trade",
		"trd_place_order",
		"trd_get_order_list",
		"trd_modify_order",
		"sys_get_global_state",
	}

	for _, ex := range examples {
		t.Run(ex, func(t *testing.T) {
			examplePath := filepath.Join(examplesDir, ex)
			if _, err := os.Stat(examplePath); os.IsNotExist(err) {
				t.Skipf("example %s not found", ex)
			}

			// Try to compile the example
			cmd := exec.Command("go", "build", "-o", os.DevNull, ".")
			cmd.Dir = examplePath
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("example %s failed to compile:\n%s\n%v", ex, string(output), err)
			}
		})
	}
}

func TestAlgoExamplesCompile(t *testing.T) {
	examplesDir := getExamplesDir()
	if examplesDir == "" {
		t.Skip("examples directory not found")
	}

	algoExamples := []string{
		"algo_sma_crossover",
		"algo_grid_trading",
		"algo_market_making",
		"algo_breakout_trading",
		"algo_vwap_execution",
	}

	for _, ex := range algoExamples {
		t.Run(ex, func(t *testing.T) {
			examplePath := filepath.Join(examplesDir, ex)
			if _, err := os.Stat(examplePath); os.IsNotExist(err) {
				t.Skipf("example %s not found", ex)
			}

			cmd := exec.Command("go", "build", "-o", os.DevNull, ".")
			cmd.Dir = examplePath
			output, err := cmd.CombinedOutput()
			if err != nil {
				t.Errorf("algo example %s failed to compile:\n%s\n%v", ex, string(output), err)
			}
		})
	}
}

func TestSimulatorCompiles(t *testing.T) {
	examplesDir := getExamplesDir()
	if examplesDir == "" {
		t.Skip("examples directory not found")
	}

	simPath := filepath.Join(examplesDir, "..", "simulator")
	if _, err := os.Stat(simPath); os.IsNotExist(err) {
		t.Skip("simulator directory not found")
	}

	cmd := exec.Command("go", "build", "-o", os.DevNull, ".")
	cmd.Dir = simPath
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Errorf("simulator failed to compile:\n%s\n%v", string(output), err)
	}
}

