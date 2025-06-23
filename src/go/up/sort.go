package main

import (
	"fmt"
	. "fmt"
	"os"
	"sort"
	"text/tabwriter"
	"time"
)

func main() {
	// sort.Sort()
	// 核心就三个接口
	// type Interface interface {
	//     Len() int
	//     Less(i, j int) bool // i, j are indices of sequence elements
	//     Swap(i, j int)
	// }

	var ko = StringSort{"2", "1", "adf", "o2002"}
	sort.Sort(ko)
	Println("ko:", ko)

	var ko2 = []string{"2", "1", "adf", "o2002"}
	ko3 := StringSort(ko2)
	sort.Sort(ko3)
	Println("ko2:", ko2) // 发生了变化
	Println("ko3:", ko3) // 发生了变化

	// 内部提供了一些函数, 方便我们操作
	values := []int{3, 1, 4, 1}
	Println(sort.IntsAreSorted(values)) // "false"
	sort.Ints(values)
	Println(values)                     // "[1 1 3 4]"
	Println(sort.IntsAreSorted(values)) // "true"
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	Println(values)                     // "[4 3 1 1]"
	Println(sort.IntsAreSorted(values)) // "false"

	// before
	Println("\n****Before***")
	printTracks(tracks)
	sort.Sort(byArtist(tracks))
	// after
	Println("\n****After***")
	printTracks(tracks)

	// 反序
	Println("\n****Reverse***")
	sort.Sort(sort.Reverse(byArtist(tracks)))
	printTracks(tracks)

	// 自定义排序
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	Println("\n****CustomSort Title->Year->Length***")
	printTracks(tracks)

	// 1. 定义结构
	// 2. 实现三个 len ,less, swap方法
	// 3. 实现自定义方法

}

type StringSort []string

func (p StringSort) Len() int {
	return len(p)
}

func (p StringSort) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p StringSort) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
}

func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}
func printTracks(tracks []*Track) {
	const format = "%v\t%v\t%v\t%v\t%v\t\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 2, ' ', 0)
	fmt.Fprintf(tw, format, "Title", "Artist", "Album", "Year", "Length")
	fmt.Fprintf(tw, format, "-----", "------", "-----", "----", "------")
	for _, t := range tracks {
		fmt.Fprintf(tw, format, t.Title, t.Artist, t.Album, t.Year, t.Length)
	}
	tw.Flush() // calculate column widths and print table
}

type byArtist []*Track

func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Year < x[j].Year }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }

type customSort struct {
	t    []*Track
	less func(i, j *Track) bool
}

func (x customSort) Len() int           { return len(x.t) }
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }
