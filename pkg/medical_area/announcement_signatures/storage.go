package announcement_signatures

import (
	"github.com/jmoiron/sqlx"
)

const (
	Postgresql = "postgres"
	SqlServer  = "sqlserver"
	Oracle     = "oci8"
)

type ServicesAnnouncementSignaturesRepository interface {
	create(m *AnnouncementSignatures) error
	update(m *AnnouncementSignatures) error
	delete(id string) error
	getByID(id string) (*AnnouncementSignatures, error)
	getByIDPatient(id string) (*AnnouncementSignatures, error)
	getByDNIPatient(dni string) (*AnnouncementSignatures, error)
	getAll() ([]*AnnouncementSignatures, error)
	GetAnnouncement() (*Announcement, error)
}

func FactoryStorage(db *sqlx.DB, txID string) ServicesAnnouncementSignaturesRepository {
	return newAnnouncementSignaturesSqlServerRepository(db, txID)
}
