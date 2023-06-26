// Start 3 workers….create an employee some employees and update the salaries of those employees using those 3 workers.
// Read on worker pattern for go routines for this.
package main

import (
	"fmt"
	"sync"
	"time"
)

// Creating Employee struct
type Employee struct {
	Id     int
	Name   string
	Salary int
}

func updateSalary(emp *Employee, wg *sync.WaitGroup, m *sync.Mutex) {
	//Each goroutines calls 'Done()' when it completes.
	defer wg.Done()
	m.Lock()
	//Increment of the salary
	emp.Salary += 500
	m.Unlock()
	fmt.Printf("The updated salary for Employee %s is $%d\n", emp.Name, emp.Salary)

}

func main() {
	//Creating 3 go routines which is worker
	numWorkers := 3
	//Wait for a group of goroutines to complete before proceeding.
	var wg sync.WaitGroup
	var m sync.Mutex
	wg.Add(numWorkers) //Adding goroutine to wait group

	//Adding values to the employee struct
	employees := []*Employee{
		{Id: 1, Name: "Michelle", Salary: 50000},
		{Id: 2, Name: "Tharun", Salary: 60000},
		{Id: 3, Name: "Ram", Salary: 45000},
	}

	//Loop through worker i.e. start
	// for i := 0; i < numWorkers; i++ {
	// 	//wg.Add(1)
	// 	go func() {
	// 		for _, value := range employees {
	// 			// &wg pointer to sync.WaitGroup and wait for collection of go routine to finish executing.
	// 			updateSalary(value, &wg, &m)

	// 		}
	// 	}()
	// }
	for i := 0; i < numWorkers; i++ {
		go func() {
			for _, emp := range employees {
				// Create a copy of the employee struct
				employeeCopy := *emp
				updateSalary(&employeeCopy, &wg, &m)
			}
		}()
	}

	//Allow some time for the goroutines to complete thier work
	time.Sleep(time.Minute)

	//Wait for all goroutines to finish
	wg.Wait()
}
