package usecases

import (
	"fmt"
	"github.com/1azar/QRChan/domain"
	"time"
)

type Logger interface {
	Info(message ...any)
	Error(message ...any)
}

type QRInteract struct {
	qrSettingsRepository domain.QRSettingsRepository
	QRSettingsBuffer     QRSettingsBuffer //QRSettingsBuffer
	logger               Logger
}

func NewQRInteract(
	qrSettingsRepository domain.QRSettingsRepository,
	logger Logger,
	qsBufferErasingPeriodMinutes time.Duration, //after every qsBufferErasingPeriodMinutes minutes qrSettingsRepository will be erased repeatedly
) *QRInteract {

	qs := QRInteract{
		qrSettingsRepository: qrSettingsRepository,
		QRSettingsBuffer:     *NewQrSettingsBuffer(qsBufferErasingPeriodMinutes),
		logger:               logger,
	}
	return &qs
}

// StoreQRSettings settings should be sent by web or bot interfaces
func (interactor *QRInteract) StoreQRSettings(qs domain.QRSettings) error {
	if err := interactor.qrSettingsRepository.Store(qs); err != nil {
		interactor.logger.Info(fmt.Sprintf("Could not sotre QR Settings in the Repository for user:%v\n%+v", qs.ID, qs))
		return err
	}
	interactor.logger.Info(fmt.Sprintf("Stored QR Settings in the Repository for user: %v", qs.ID))
	return nil
}

// FindQRSettings search for QR settings for user in the repository and returns it.
// If there is no QR settings for current user in the repository then it generates new default settings and returns it.
func (interactor *QRInteract) FindQRSettings(id int64) (domain.QRSettings, error) {
	qs, err := interactor.qrSettingsRepository.FindById(id)
	if err != nil {
		interactor.logger.Error(fmt.Sprintf("Error during searching QR Settings in the Repository for user: %v", id))
		return domain.QRSettings{}, err
	}
	if qs == (domain.QRSettings{}) {
		interactor.logger.Info(fmt.Sprintf("Could not find QR Settings in the repository for user: %v", id))
		return interactor.NewQRSettings(id), nil
	}
	interactor.logger.Info(fmt.Sprintf("Found QR Settings in the Repository for user: %v\n%+v", id, qs))
	return qs, nil
}

func (interactor *QRInteract) NewQRSettings(id int64) domain.QRSettings {
	interactor.logger.Info(fmt.Sprintf("Generating new QRSettings for user: %v", id))
	return domain.NewQRSettings(id)
}
