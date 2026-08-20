package storage

import "errors"

var ErrBatchCommitRejected = errors.New("batchcommit batch rejected")

type BatchCommitStore struct{ committed []string }

func (s *BatchCommitStore) Apply(values []string) error {
	pending := make([]string, 0, len(values))
	for _, value := range values {
		if value == "bad" {
			return ErrBatchCommitRejected
		}
		pending = append(pending, value)
	}
	s.committed = append(s.committed, pending...)
	return nil
}
