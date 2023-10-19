package main

func Merge(a []string, b []string) []string {
  return append(a, b...)
}

func Remove(a []string, index int) []string {
  // withoutThirdElement := append(names[:2], names[3:]...)
  withoutThirdElement := append(a[:(index)], a[index+1:]...)
  return withoutThirdElement

}

func RemoveLast(a []string) []string {

  return a[:len(a)-1]
}
