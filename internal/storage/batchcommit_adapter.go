package storage

func (s *BatchCommitStore) Snapshot() []string {
	out := make([]string, len(s.committed))
	copy(out, s.committed)
	return out
}
func (s *BatchCommitStore) Count() int { return len(s.committed) }
