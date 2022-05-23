package cache

import "time"

type item struct {
	value string
	timer time.Time
}

type Cache struct {
	storage map[string]item
}

func NewCache() Cache {
	storage := make(map[string]item)
	return Cache{storage: storage}
}

func (k Cache) Get(key string) (string, bool) {
	item, ok := k.storage[key]
	if !ok {
		return "", false
	}
	if item.timer == 0 || time.Now().Before(item.timer) {
		return item.value, true
	}
	delete(k.storage, key)
	return "", false
}

func (k Cache) Put(key, value string) {
	k.storage[key] = item{value: value, timer:0}
	//помещает значение со связанным ключом в кеш. Значение, введенное с помощью этого метода, никогда не истекает (имеет бесконечный срок). 
	//Ввод в существующий ключ должен перезаписать значение
}

func (k Cache) Keys() []string {
	var keys []string
	for key, item := range k.box {
		if item.timer == 0 || time.Now().Before(item.timer) {
			keys = append(keys, key)
		}
	}

	return keys
	//должна возвращать часть существующих (не просроченных ключей)
}

func (k Cache) PutTill(key, value string, deadline time.Time) {
	k.storage[key] = item{value: value, timer: deadline}
	//Должен делать то же самое, что и Put, но срок действия пары ключ/значение должен истечь к указанному крайнему сроку.
}
