package storage

func (s *BatchCommitStore) Snapshot() []string { return s.committed }
func (s *BatchCommitStore) Count() int         { return len(s.committed) }
