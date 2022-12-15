package cli

import (
	"fmt"

	Service "github.com/oswaldom-code/affiliate-tracker/src/aplication/system_services"
	"github.com/spf13/cobra"
)

const (
	CONFIRMATION = "> Enter Yes to confirm or No to cancel (Y/n): "
)

func askConfirmation() bool {
	var answer string
	// ask for confirmation
	fmt.Print(CONFIRMATION)
	fmt.Scanln(&answer)
	if answer == "Y" || answer == "y" {
		return true
	}
	return false
}

func RunCliCmd(cmd *cobra.Command, args []string) error {
	healthService := Service.HealthService()
	function, err := cmd.Flags().GetString("function")
	if err != nil {
		fmt.Println("Error:", err)
		return err
	}
	switch function {
	case "test":
		err = healthService.TestDb()
		if err != nil {
			fmt.Printf("> ❌Test error: %s\n", err)
			return err
		}
		fmt.Println("> ✅ Test connection to database success ")
		return nil
	}
	return nil
}
