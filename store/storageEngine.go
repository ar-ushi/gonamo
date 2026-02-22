package storageEngine

import ("github.com/ar-ushi/gonamo/types")

type storageEngine interface {
    Put(key types.Key, value types.Value) error
    Get(key types.Key) ([]types.Value, bool, error)
}

type StorageEngine struct {
	lock sync.RWMutex
	data map[types.Key] [] types.Value
}

func newStorageEngine() *StorageEngine{
	return &StorageEngine{
		data: make(map[types.Key][]types.Value),
	}
}


func (s *StorageEngine) Get(key types.Key) ([]types.Value, bool){
	s.lock.Lock()
	defer s.lock.Unlock()

	existing, exists := s.data[key];
	copyStore := make([]types.Value, 0, len(existing))
	copy(copyStore, existing)
	return copyStore, exists
}

func (s *StorageEngine) Put(key types.Key, value types.Value){
	s.lock.Lock();
	defer s.lock.Unlock(); //TIL this helps deal with early returns
newValStore := make([]types.Value, 0, len(existing)+1)
existing, exists := s.data[key]; // returns data + clock
	if !exists {
		s.data[key] = []types.Value value
		return
	}
	for i, oldVal := range existing {
		/** New Write is Stale, data already synced*/
		if (value.Clock.Compare(oldVal.Clock,  vclock.Equal|vclock.Ancestor)){
			return;
		}
		/** Happy Logic - Dropping the previous value if clock is less than or equal to new clock */
		if (oldVal.Clock.Compare(value.Clock, vclock.Equal|vclock.Ancestor)){
			continue;
		} else {
			newValStore = append(newValStore, oldVal)
		}
		
	}
	s.data[key] = append(newValStore, value);
	
}