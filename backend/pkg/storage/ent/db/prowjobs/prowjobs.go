// Code generated by entc, DO NOT EDIT.

package prowjobs

const (
	// Label holds the string label denoting the prowjobs type in the database.
	Label = "prow_jobs"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldJobID holds the string denoting the job_id field in the database.
	FieldJobID = "job_id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldDuration holds the string denoting the duration field in the database.
	FieldDuration = "duration"
	// FieldTestsCount holds the string denoting the tests_count field in the database.
	FieldTestsCount = "tests_count"
	// FieldFailedCount holds the string denoting the failed_count field in the database.
	FieldFailedCount = "failed_count"
	// FieldSkippedCount holds the string denoting the skipped_count field in the database.
	FieldSkippedCount = "skipped_count"
	// FieldJobType holds the string denoting the job_type field in the database.
	FieldJobType = "job_type"
	// EdgeProwJobs holds the string denoting the prow_jobs edge name in mutations.
	EdgeProwJobs = "prow_jobs"
	// RepositoryFieldID holds the string denoting the ID field of the Repository.
	RepositoryFieldID = "repo_id"
	// Table holds the table name of the prowjobs in the database.
	Table = "prow_jobs"
	// ProwJobsTable is the table that holds the prow_jobs relation/edge.
	ProwJobsTable = "prow_jobs"
	// ProwJobsInverseTable is the table name for the Repository entity.
	// It exists in this package in order to avoid circular dependency with the "repository" package.
	ProwJobsInverseTable = "repositories"
	// ProwJobsColumn is the table column denoting the prow_jobs relation/edge.
	ProwJobsColumn = "repository_prow_jobs"
)

// Columns holds all SQL columns for prowjobs fields.
var Columns = []string{
	FieldID,
	FieldJobID,
	FieldCreatedAt,
	FieldDuration,
	FieldTestsCount,
	FieldFailedCount,
	FieldSkippedCount,
	FieldJobType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prow_jobs"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"repository_prow_jobs",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
