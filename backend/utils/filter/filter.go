package filter
//
//import "reflect"
//
//func filter(tp reflect.Type) func(ss []interface{}, test func(tp) bool) {
//	return func(ss []interface{}, test func(tp) bool)(ret []interface{}) {
//		for _, s := range ss {
//			if test(s) {
//				ret = append(ret, s)
//			}
//		}
//		return
//	}
//}
//
//func main () {
//	type Group struct {
//		name string
//		money int
//	}
//	sss := []*Group{
//		&Group{
//			name:  "abc",
//			money: 0,
//		},
//		&Group{
//			name:  "h",
//			money: 1,
//		},
//	}
//	filter(sss, func([]*Group) bool {
//
//	})
//}