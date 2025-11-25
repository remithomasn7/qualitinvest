package services

import (
	"log"
	"time"
)

// DataScheduler gère la planification des tâches de collecte de données
type DataScheduler struct {
	dataCollectionService *DataCollectionService
	running               bool
}

// NewDataScheduler crée un nouveau scheduler
func NewDataScheduler(dataSvc *DataCollectionService) *DataScheduler {
	return &DataScheduler{
		dataCollectionService: dataSvc,
		running:               false,
	}
}

// Start démarre le scheduler
func (s *DataScheduler) Start() {
	if s.running {
		log.Println("Data scheduler already running")
		return
	}

	s.running = true
	log.Println("Starting data collection scheduler")

	// Mise à jour quotidienne des overviews à 2h du matin
	go s.scheduleDailyOverviewUpdate()

	// Mise à jour complète des données financières le dimanche à 3h
	go s.scheduleWeeklyFinancialUpdate()
}

// Stop arrête le scheduler
func (s *DataScheduler) Stop() {
	s.running = false
	log.Println("Data collection scheduler stopped")
}

// scheduleDailyOverviewUpdate planifie la mise à jour quotidienne des overviews
func (s *DataScheduler) scheduleDailyOverviewUpdate() {
	for s.running {
		// Calculer le temps jusqu'à 2h du matin
		now := time.Now()
		nextRun := time.Date(now.Year(), now.Month(), now.Day(), 2, 0, 0, 0, now.Location())

		// Si on a déjà passé 2h aujourd'hui, programmer pour demain
		if now.After(nextRun) {
			nextRun = nextRun.Add(24 * time.Hour)
		}

		waitDuration := nextRun.Sub(now)
		log.Printf("Next daily overview update scheduled in %v", waitDuration)

		time.Sleep(waitDuration)

		if s.running {
			log.Println("Starting daily overview update")
			if err := s.dataCollectionService.BatchUpdateOverviews(); err != nil {
				log.Printf("Daily overview update failed: %v", err)
			} else {
				log.Println("Daily overview update completed")
			}
		}
	}
}

// scheduleWeeklyFinancialUpdate planifie la mise à jour hebdomadaire des données financières
func (s *DataScheduler) scheduleWeeklyFinancialUpdate() {
	for s.running {
		// Calculer le temps jusqu'au prochain dimanche à 3h
		now := time.Now()
		daysUntilSunday := (7 - int(now.Weekday())) % 7
		if daysUntilSunday == 0 && now.Hour() >= 3 {
			daysUntilSunday = 7 // Si on est dimanche après 3h, attendre la semaine prochaine
		}

		nextRun := time.Date(now.Year(), now.Month(), now.Day(), 3, 0, 0, 0, now.Location())
		nextRun = nextRun.AddDate(0, 0, daysUntilSunday)

		waitDuration := nextRun.Sub(now)
		log.Printf("Next weekly financial update scheduled in %v", waitDuration)

		time.Sleep(waitDuration)

		if s.running {
			log.Println("Starting weekly financial data update")
			// TODO: Implémenter la logique de mise à jour financière hebdomadaire
			// Cela pourrait inclure la collecte de nouvelles données pour les entreprises existantes
			log.Println("Weekly financial data update completed (placeholder)")
		}
	}
}

// TriggerManualUpdate déclenche une mise à jour manuelle
func (s *DataScheduler) TriggerManualUpdate() error {
	log.Println("Starting manual data update")
	return s.dataCollectionService.BatchUpdateOverviews()
}
