package main

import (
	"strings"
	"reflect"
)

type SimplePrimitive struct {
	Id		string	`json:"id"`
	Class		string  `json:"class"`
	Type		string	`json:"type"`
	Provider	string	`json:"provider,omitempty"`
	SimpleMeta	interface{} `json:"meta,omitempty"`
	SimpleParam     interface{} `json:"param,omitempty"`
	SimpleOp	[]interface{} `json:"op,omitempty"`
}

type SimpleGroup struct {
	Id		string	`json:"id"`
	SimpleMeta	interface{} `json:"meta,omitempty"`
	SimplePrimitive []SimplePrimitive `json:"primitives"`
}

type SimpleMaster struct {
	Id		string	`json:"id"`
	SimpleMeta	interface{} `json:"meta,omitempty"`
	SimplePrimitive SimplePrimitive `json:"primitive"`
	SimpleGroup	SimpleGroup `json:"group"`
}

type SimpleResource struct {
	Id	string	`json:"id"`
	Type	string	`json:"type"`
}

func isString(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.String:
		return true
	}
	return false
}

func isBlank(value reflect.Value) bool {
    switch value.Kind() {
    case reflect.String:
        return value.Len() == 0
    case reflect.Bool:
        return !value.Bool()
    case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
        return value.Int() == 0
    case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
        return value.Uint() == 0
    case reflect.Float32, reflect.Float64:
        return value.Float() == 0
    case reflect.Interface, reflect.Ptr:
        return value.IsNil()
    }
    return reflect.DeepEqual(value.Interface(), reflect.Zero(value.Type()).Interface())
}

func (s *SimplePrimitive)InstanceSimplePrimitive(item *Primitive) {
	s.Id = item.Id
	s.Class = item.Class
	s.Provider = item.Provider
	s.Type = item.Type

	nv := make(map[string]string)
	for _, item := range item.MetaAttributes {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	if len(nv) == 0 {
		s.SimpleMeta = nil
	} else {
		s.SimpleMeta = nv
	}

	nv = make(map[string]string)
	for _, item := range item.InstanceAttributes {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	if len(nv) == 0 {
		s.SimpleParam = nil
	} else {
		s.SimpleParam = nv
	}

	if item.Operations != nil {
		for _, item := range item.Operations.Op {
			nv := make(map[string]string)
			tt := reflect.TypeOf(*item)
			tv := reflect.ValueOf(*item)
			for i := 0; i < tt.NumField(); i++ {
				field := tt.Field(i)
				field_v := tv.Field(i)
				if field.Name != "Id" && isString(field_v) && !isBlank(field_v) {
					nv[strings.ToLower(field.Name)] = field_v.String()
				}
			}
			s.SimpleOp = append(s.SimpleOp, nv)
		}
	}
}

func (s *SimpleGroup)InstanceSimpleGroup(item *Group) {
	s.Id = item.Id

	nv := make(map[string]string)
	for _, item := range item.MetaAttributes {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	if len(nv) == 0 {
		s.SimpleMeta = nil
	} else {
		s.SimpleMeta = nv
	}

	primitives := make([]SimplePrimitive, 0)
	for _, item := range item.Primitive {
		simple_item := &SimplePrimitive{}
		simple_item.InstanceSimplePrimitive(item)
		primitives = append(primitives, *simple_item)
	}
	s.SimplePrimitive = primitives
}

func (s *SimpleMaster)InstanceSimpleMaster(item *Master) {
	s.Id = item.Id

	nv := make(map[string]string)
	for _, item := range item.MetaAttributes {
		for _, nv_item := range item.Nvpair {
			nv[nv_item.Name] = nv_item.Value
		}
	}
	if len(nv) == 0 {
		s.SimpleMeta = nil
	} else {
		s.SimpleMeta = nv
	}

	if item.Primitive != nil {
		simple_item := &SimplePrimitive{}
		simple_item.InstanceSimplePrimitive(item.Primitive)
		s.SimplePrimitive = *simple_item
	} else if item.Group != nil {
		simple_item := &SimpleGroup{}
		simple_item.InstanceSimpleGroup(item.Group)
		s.SimpleGroup = *simple_item
	}
}

func handleConfigResources(urllist []string, cib Cib) (bool, interface{}) {
	/*
	if len(urllist) == 4 {
		cib.Configuration.Resources.URLType = "all"
	} else {
		resId := urllist[4]

		mapIdType := make(map[string]TypeIndex)
		for pi, pitem := range cib.Configuration.Resources.Primitive {
			mapIdType[pitem.Id] = TypeIndex{"primitive", pi}
		}
		for gi, gitem := range cib.Configuration.Resources.Group {
			mapIdType[gitem.Id] = TypeIndex{"group", gi}
		}
		for ci, citem := range cib.Configuration.Resources.Clone {
			mapIdType[citem.Id] = TypeIndex{"clone", ci}
		}
		for mi, mitem := range cib.Configuration.Resources.Master {
			mapIdType[mitem.Id] = TypeIndex{"master", mi}
		}

		val, ok := mapIdType[resId]
		if ok {
			cib.Configuration.Resources.URLType = val.Type
			cib.Configuration.Resources.URLIndex = val.Index
		} else {
			return false
		}
	}
	*/
	if len(urllist) == 4 {
		resources := make([]SimpleResource, 0)
		for _, item := range cib.Configuration.Resources.Primitive {
			resources = append(resources, SimpleResource{Id: item.Id, Type: "primitive"})
		}
		for _, item := range cib.Configuration.Resources.Group {
			resources = append(resources, SimpleResource{Id: item.Id, Type: "group"})
		}
		for _, item := range cib.Configuration.Resources.Clone {
			resources = append(resources, SimpleResource{Id: item.Id, Type: "clone"})
		}
		for _, item := range cib.Configuration.Resources.Master {
			resources = append(resources, SimpleResource{Id: item.Id, Type: "master"})
		}
		for _, item := range cib.Configuration.Resources.Bundle {
			resources = append(resources, SimpleResource{Id: item.Id, Type: "bundle"})
		}
		return true, resources
	} else {
		return handleConfigPrimitives(urllist, cib)
	}
}

/*
func getSimpleNv(t []*interface{}) interface{} {
	tSlice, ok := t.([]*interface{})
	if !ok {
		fmt.Println("1111")
	}
	fmt.Println(tSlice)
	for _, v := range tSlice {
		fmt.Println(v)
	}
	return nil
}
*/
/*
func getSimpleMeta(metas []*MetaAttributes) interface{} {
       nv := make(map[string]string)
       for _, item := range metas {
               for _, nv_item := range item.Nvpair {
                       nv[nv_item.Name] = nv_item.Value
               }
       }
       return nv
}

func getSimpleParam(metas []*InstanceAttributes) interface{} {
       nv := make(map[string]string)
       for _, item := range metas {
               for _, nv_item := range item.Nvpair {
                       nv[nv_item.Name] = nv_item.Value
               }
       }
       return nv
}
*/
func handleConfigPrimitives(urllist []string, cib Cib) (bool, interface{}) {

	primitives_data := cib.Configuration.Resources.Primitive
	if primitives_data == nil {
		return true, nil
	}

	var resId string
	if len(urllist) == 4 {
		resId = ""
	} else if len(urllist) == 5 {
		resId = urllist[4]
	}

	primitives := make([]SimplePrimitive, 0)
	for _, item := range primitives_data {
		simple_item := &SimplePrimitive{}
		simple_item.InstanceSimplePrimitive(item)
		if resId != "" && item.Id == resId {
			primitives = append(primitives, *simple_item)
			return true, primitives[0]
		} else if resId == "" {
			primitives = append(primitives, *simple_item)
		}
	}
	return true, primitives
}

func handleConfigGroups(urllist []string, cib Cib) (bool, interface{}) {
	groups_data := cib.Configuration.Resources.Group
	if groups_data == nil {
		return true, nil
	}

	var resId string
	if len(urllist) == 4 {
		resId = ""
	} else if len(urllist) == 5 {
		resId = urllist[4]
	}

	groups := make([]SimpleGroup, 0)
	for _, item := range groups_data {
		simple_item := &SimpleGroup{}
		simple_item.InstanceSimpleGroup(item)
		if resId != "" && item.Id == resId {
			groups = append(groups, *simple_item)
			return true, groups[0]
		} else if resId == "" {
			groups = append(groups, *simple_item)
		}
	}
	return true, groups
}


func handleConfigMasters(urllist []string, cib Cib) (bool, interface{}) {
	masters_data := cib.Configuration.Resources.Master
	if masters_data == nil {
		return true, nil
	}

	var resId string
	if len(urllist) == 4 {
		resId = ""
	} else if len(urllist) == 5 {
		resId = urllist[4]
	}

	masters := make([]SimpleMaster, 0)
	for _, item := range masters_data {
		simple_item := &SimpleMaster{}
		simple_item.InstanceSimpleMaster(item)
		if resId != "" && item.Id == resId {
			masters = append(masters, *simple_item)
			return true, masters[0]
		} else if resId == "" {
			masters = append(masters, *simple_item)
		}
	}
	return true, masters
}


func handleStateResources(urllist []string, cib Cib) (bool, interface{}) {
	return true, nil
}
