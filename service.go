package booking

import "time"

type Service interface {
	AvailableDays() ([]time.Time, error)
	Book(bookingForm) (bookingId, error)
}

type service struct {
	dataPath       string
	processorToken string
	reserver       reserver
}

func NewService(dataPath, processorToken string) service {
	s := service{}
	s.dataPath = dataPath
	s.processorToken = processorToken
	s.reserver = newCalendar()
	return s
}

// Returns list of available dates in the future
func (s service) AvailableDays() ([]time.Time, error) {
	days := []time.Time{
		time.Date(2015, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2015, 2, 1, 0, 0, 0, 0, time.UTC),
	}
	return days, nil
}

func (s service) Book(bookingForm) (bookingId, error) {
	return 0, nil
}
