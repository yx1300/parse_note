package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
)

/**
思路：
	1、设置正则表达式
	2、读取源代码绝对地址(open函数要求)，每行读取，并去掉前缀空格
	3、判断每行是否符合正则表达式，符合提取姓名存入map
	4、读取结束后对map进行排序
	5、根据要求k，返回前出现次数最多的前K个
*/
type BlamePeople struct {
	Name       string
	ProblemNum int
}

func parse(addr string) []BlamePeople {
	m := make(map[string]int)
	// 读取文件路径
	f, err := os.Open("D:\\codes\\go_codes\\sgg_goLearning\\test\\parse_note\\test.go")
	defer f.Close()
	if err != nil {
		fmt.Println("路径错误:", err)
		return nil
	}
	// 正则匹配
	reg := regexp.MustCompile(`\/\/\s+(todo|FIXME|fixme|TODO)`)
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			fmt.Println(err)
			break
		}
		if len(line) == 0 {
			break
		}
		if reg.MatchString(strings.TrimSpace(line)) {
			// TODO(tom): I'l do it later
			name := strings.Split(strings.Split(line, "(")[1], ")")[0]
			m[name]++
		}
	}

	res := sortMap(m, 100)
	return res

}

// 根据map里的value值进行排序 并返回前k位
func sortMap(m map[string]int, k int) []BlamePeople {
	var bps []BlamePeople
	for k, v := range m {
		bps = append(bps, BlamePeople{
			Name:       k,
			ProblemNum: v,
		})
	}
	sort.Slice(bps, func(i, j int) bool {
		return bps[i].ProblemNum > bps[j].ProblemNum
	})
	if k > len(bps) {
		return bps
	}
	return bps[:k]
}
func main() {
	res := parse("D:\\codes\\go_codes\\sgg_goLearning\\test\\parse_note\\test.go")
	fmt.Println(res)
}
