package services

type RowMapped interface {
	Map(rawRow string) map[string]string
}
