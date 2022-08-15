package domains

type DomainItem[K comparable] struct {
	Id    K
	Descr string
}

type Domain[K comparable] struct {
	s []DomainItem[K]
	m map[K]DomainItem[K]
}

func (d *Domain[K]) First() K {
	return d.s[0].Id
}

func (d *Domain[K]) Last() K {
	return d.s[len(d.s)-1].Id
}

func (d *Domain[K]) Slice() []DomainItem[K] {
	return d.s
}

func (d *Domain[K]) ById(Id K) DomainItem[K] {
	return d.m[Id]
}

func (d *Domain[K]) AsMap() map[K]DomainItem[K] {
	return d.m
}

func (d *Domain[K]) init(newSlice []DomainItem[K]) {
	d.s = newSlice
	d.m = make(map[K]DomainItem[K], len(d.s))
	for _, v := range d.s {
		d.m[v.Id] = v
	}
}
