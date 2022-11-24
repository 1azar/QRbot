package usecases

import (
	"github.com/1azar/QRChan/domain"
	"sync"
	"time"
)

// QRSettingsBuffer is safe map for buffering QESettings
// it has method to erase its cache.
// this struct reduce interactions with databases when user enters to options mode
type QRSettingsBuffer struct {
	sync.Mutex
	cache map[int64]domain.QRSettings //TODO make erasing timed for every users qr setting
}

func NewQrSettingsBuffer(ErasingMinutes time.Duration) *QRSettingsBuffer {
	qsb := QRSettingsBuffer{
		cache: make(map[int64]domain.QRSettings),
	}
	qsb.eraserStart(ErasingMinutes) // starts ticker for periodic erasing cache
	return &qsb
}

func (qsb *QRSettingsBuffer) Set(key int64, value domain.QRSettings) {
	qsb.Lock()
	defer qsb.Unlock()
	qsb.cache[key] = value
}

func (qsb *QRSettingsBuffer) Get(key int64) (domain.QRSettings, bool) {
	var val domain.QRSettings
	qsb.Lock()
	val = qsb.cache[key]
	qsb.Unlock()
	if val == (domain.QRSettings{}) {
		return val, false
	}
	return val, true
}

func (qsb *QRSettingsBuffer) eraserStart(Minutes time.Duration) {
	ticker := time.NewTicker(time.Minute * time.Duration(Minutes))
	go func() {
		for _ = range ticker.C {
			qsb.Lock()
			qsb.cache = make(map[int64]domain.QRSettings)
			qsb.Unlock()
		}
	}()
}
