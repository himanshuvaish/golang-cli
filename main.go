package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

type Driver struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Wins  int    `json:"wins"`
	Poles int    `json:"poles"`
}

var drivers = []Driver{
	{ID: 1, Name: "Lewis Hamilton", Wins: 95, Poles: 98},
	{ID: 2, Name: "Sebastian Vettel", Wins: 53, Poles: 57},
	{ID: 3, Name: "Ayrton Senna", Wins: 41, Poles: 65},
}

func main() {
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(driverCmd())
	rootCmd.Execute()
}

func driverCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "driver",
		Short: "Driver operations",
	}

	cmd.AddCommand(getDriversCmd())
	cmd.AddCommand(getDriverCmd())
	cmd.AddCommand(addDriverCmd())
	cmd.AddCommand(updateDriverCmd())
	cmd.AddCommand(deleteDriverCmd())

	return cmd
}

func getDriversCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get",
		Short: "Get all drivers",
		Run: func(cmd *cobra.Command, args []string) {
			for _, driver := range drivers {
				fmt.Printf("ID: %d, Name: %s, Wins: %d, Poles: %d\n", driver.ID, driver.Name, driver.Wins, driver.Poles)
			}
		},
	}
}

func getDriverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "get [id]",
		Short: "Get a driver by ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			for _, driver := range drivers {
				if driver.ID == id {
					fmt.Printf("ID: %d, Name: %s, Wins: %d, Poles: %d\n", driver.ID, driver.Name, driver.Wins, driver.Poles)
					return
				}
			}
			fmt.Println("Driver not found")
		},
	}
}

func addDriverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "add [id] [name] [wins] [poles]",
		Short: "Add a new driver",
		Args:  cobra.ExactArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			wins, _ := strconv.Atoi(args[2])
			poles, _ := strconv.Atoi(args[3])
			newDriver := Driver{ID: id, Name: args[1], Wins: wins, Poles: poles}
			drivers = append(drivers, newDriver)
			fmt.Printf("Added driver: ID: %d, Name: %s, Wins: %d, Poles: %d\n", newDriver.ID, newDriver.Name, newDriver.Wins, newDriver.Poles)
		},
	}
}

func updateDriverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "update [id] [wins] [poles]",
		Short: "Update a driver by ID",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			wins, _ := strconv.Atoi(args[1])
			poles, _ := strconv.Atoi(args[2])
			for i, driver := range drivers {
				if driver.ID == id {
					drivers[i].Wins += wins
					drivers[i].Poles += poles
					fmt.Printf("Updated driver: ID: %d, Name: %s, Wins: %d, Poles: %d\n", drivers[i].ID, drivers[i].Name, drivers[i].Wins, drivers[i].Poles)
					return
				}
			}
			fmt.Println("Driver not found")
		},
	}
}

func deleteDriverCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a driver by ID",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, _ := strconv.Atoi(args[0])
			for i, driver := range drivers {
				if driver.ID == id {
					drivers = append(drivers[:i], drivers[i+1:]...)
					fmt.Printf("Deleted driver with ID %d\n", id)
					return
				}
			}
			fmt.Println("Driver not found")
		},
	}
}
