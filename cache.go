package cache

import "time"

type Cache struct {
	items map[string]struct {
		key      string
		value    string
		deadline time.Time
	}
}

func NewCache() Cache {
	i := make(map[string]struct {
		key      string
		value    string
		deadline time.Time
	})
	return Cache{items: i}
}

func (a *Cache) Get(key string) (string, bool) {
	var b string
	var ok bool
	i, found := a.items[key]
	if !found {
		return "", false
	}
	if i.deadline.IsZero() {
		b = i.value
		ok = true
	} else {
		if i.deadline.Before(time.Now()) {
			delete(a.items, key)
			b = ""
			ok = false
		} else {
			b = i.value
			ok = true
		}
	}
	return b, ok
}

func (a *Cache) Put(key, value string) {
	type item struct {
		key, value string
		deadline   time.Time
	}
	a.items[key] = item{
		value: value,
	}
}

func (a *Cache) Keys() []string {
	var v []string
	for i := range a.items {
		_, ok := a.Get(i)
		if !ok {
			continue
		}
		v = append(v, i)
	}
	return v
}

func (a *Cache) PutTill(key, value string, deadline time.Time) {
	type item struct {
		key, value string
		deadline   time.Time
	}
	a.items[key] = item{
		value:    value,
		deadline: deadline,
	}
}
