package schemavalidator

import "github.com/jackc/pgx/v5/pgtype"

func (schemaValidator *SchemaValidator) ValidateUUID(data string) (*pgtype.UUID, error) {
	id := pgtype.UUID{}
	if err := schemaValidator.validate.Var(data, "required,uuid"); err != nil {
		return nil, err
	}
	if err := id.Scan(data); err != nil {
		return nil, err
	}

	return &id, nil
}
