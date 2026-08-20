package storage

import "errors"

var ErrBatchCommitRejected = errors.New("batchcommit batch rejected")

type BatchCommitStore struct{ committed []string }

func (s *BatchCommitStore) Apply(values []string) error {
	for _, value := range values {
		if value == "bad" {
			return ErrBatchCommitRejected
		}
		s.committed = append(s.committed, value)
	}
	return nil
}
