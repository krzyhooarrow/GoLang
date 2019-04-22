package main

var version int
var EmployeeTaskPickupDelay = 3     // co ile pracownik pobiera zadanie
const SimulationTime = 10000			// czas symulacji /while(1) nieskonczonosc
const EmployersCounter = 10			// ilosc pracownikow
var NewTaskTime = 2					// co ile szef ma nowe zadanie
var TaskArgRange = 50				// zakres liczb przy tworzeniu nowego zadnia listy zadan
var taskListSize = 20				// rozmiar listy zadan
var productListSize = 100			// rozmiar listy produktow
var Clientinterval = 20				// co ile klient kupuje cos
var ClientCounter = 5				// ilosc klientow
var avgMachineWorkTime = 10			// sredni czas pracy maszyny
const machineCount = 10				// ilosc maszyn
const walkInterval  = 1 			// czas potrzebny na zmiane maszyny
