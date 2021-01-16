package db
type Record struct {
    Name string  `json:"name"`
    Time float32 `json:"time"`
    Pool_sz float32 `json:"pool_sz"`
}
func NewRecord(name string, time float32, pool_sz float32) Record {
    return Record{name, time, pool_sz}
}
type Database struct {
    contents []Record
}
func New() Database {
    contents := make([]Record, 0)
    return Database{contents}
}
func (database *Database) AddRecord(r Record) {
    database.contents = append(database.contents, r)
}
func (database *Database) GetRecords() []Record {
    return database.contents
}