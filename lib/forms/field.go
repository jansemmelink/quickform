package forms

import "sync"

//NewFields makes a new un-named root field where you can add sub fields to
func NewFields() IField {
	return &field{
		parent:   nil,
		fullname: "",
		name:     "",
		sub:      make(map[string]IField),
	}
}

//IField ...
type IField interface {
	FullName() string
	Name() string
	Add(name string) IField
	Sub(name string) IField
	Subs() map[string]IField
}

//field implements IField
type field struct {
	parent   IField
	fullname string
	name     string

	mutex sync.Mutex
	sub   map[string]IField
}

func (f *field) FullName() string {
	return f.fullname
}

func (f *field) Name() string {
	return f.name
}

func (f *field) Sub(name string) IField {
	if s, ok := f.sub[name]; ok {
		return s
	}
	return nil
}

func (f *field) Subs() map[string]IField {
	return f.sub
}

func (f *field) Add(name string) IField {
	f.mutex.Lock()
	f.mutex.Unlock()
	if name == "" {
		panic(log.Wrapf(nil, "field(%s).Add(\"\")", f.fullname))
	}
	if _, ok := f.sub[name]; ok {
		panic(log.Wrapf(nil, "field(%s).Add(%s) = duplicate field", f.fullname, name))
	}
	newField := &field{
		parent:   f,
		fullname: f.fullname + "." + name,
		name:     name,
		sub:      make(map[string]IField),
	}
	return newField
}
