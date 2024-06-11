package pacer

import "time"

type PacerCallback func(time.Time, time.Time) error

func ExecuteWithPace(start time.Time, end time.Time, pace time.Duration, callback PacerCallback) error {
	pointer := start

	for pointer.Before(end) {
		var err error

		if pointer.Add(pace).Before(end) {
			err = callback(pointer, pointer.Add(pace))
		} else {
			err = callback(pointer, end)
		}

		if err != nil {
			return err
		}

		pointer = pointer.Add(pace)
	}

	return nil
}
